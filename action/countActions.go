package action

import (
	Lib "flx/lib"
	Mod "flx/model"
	"sync"
	"time"
)

const POOLLIMIT = 40

func BeginCount() {
	sources, channels := Mod.RequestTokens()
	sourceNames := Mod.GatherTokens().Sources
	channelNames := Mod.GatherTokens().Channels

	toWriteStruct := Lib.InitVariantCountAll()

	stageOne, stageTwo := CountVariants(Mod.GET_INVENTORY_VARIANTS_PATH, Mod.GetInventoryVariant{Page: 0, PageSize: 100, IncludeLinkedProductVariants: true}, sources, sourceNames, "Inventory Variant Count:", true)

	stageThree := CountListingVariants(channelNames, channels, Mod.GetCountListingVariant{})

	for i, v := range sourceNames {
		toWriteStruct.InventoryVariant[v] = stageOne[i]
		toWriteStruct.ProductVariant[v] = stageTwo[i]
	}

	for _, v := range channelNames {
		toWriteStruct.ChannelVariant[v] = stageThree[v]
	}

	toWriteStruct.CreatedAt = time.Now()
	Lib.WriteJsonToFile("output.txt", toWriteStruct)
}

func CountVariants[T Mod.GetFamily](path string, query T, tokens []string, tokenNames []string, message string, extra bool) ([]int, []int) {
	var results []int
	var resultsExtra []int = nil

	for i := 0; i < len(tokens); i++ {
		var wg sync.WaitGroup
		ch := make(chan int, 1)
		chExtra := make(chan int, 1)
		wg.Add(POOLLIMIT)
		for j := 1; j <= POOLLIMIT; j++ {
			queryLocal := query.StepPage(j).(T)
			go ConcurrentCount(path, &wg, ch, chExtra, tokens[i], queryLocal, extra)
		}
		wg.Wait()
		val := <-ch
		valExtra := <-chExtra
		close(ch)
		close(chExtra)

		results = append(results, val)
		if extra == true {
			resultsExtra = append(resultsExtra, valExtra)
		}
	}
	return results, resultsExtra
}

func ConcurrentCount[T Mod.GetFamily](path string, wg *sync.WaitGroup, ch chan int, chExtra chan int, token string, query T, extra bool) {

	resp := Lib.GetDataList(path+Mod.QueryUrl(query), token)
	count := len(resp)
	for count > 0 {
		if len(ch) == 0 {
			ch <- count
			if extra {
				chExtra <- len(resp[0]["linkedProductVariants"].([]interface{}))
				for i := 1; i < count; i++ {
					x := <-chExtra
					chExtra <- x + len(resp[i]["linkedProductVariants"].([]interface{}))
				}
			}
		} else {
			x := <-ch
			ch <- x + count
			if extra {
				for i := 0; i < count; i++ {
					x := <-chExtra
					chExtra <- x + len(resp[i]["linkedProductVariants"].([]interface{}))
				}
			}
		}
		query = query.StepPage(POOLLIMIT).(T)
		resp = Lib.GetDataList(path+Mod.QueryUrl(query), token)
		count = len(resp)
	}
	wg.Done()
}

func CountListingVariants(channelNames []string, channelTokens []string, query Mod.GetCountListingVariant) map[string]int {
	return (map[string]int{channelNames[0]: int(Lib.GetDataJson(Mod.GET_LISTING_VARIANTS_PATH+Mod.COUNT_URL_EXT+Mod.QueryUrl(query), channelTokens[0])["count"].(float64)),
		channelNames[1]: int(Lib.GetDataJson(Mod.GET_LISTING_VARIANTS_PATH+Mod.COUNT_URL_EXT+Mod.QueryUrl(query), channelTokens[1])["count"].(float64))})
}

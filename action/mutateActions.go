package action

import (
	Lib "flx/lib"
	Mod "flx/model"
	"fmt"
	"sync"
	"time"
)

type logType interface {
	string | int
}

// TODO: Move logging outside and implement more accurate debugging - using data matching technique
type UpdateLog[T logType] struct {
	Log         []string
	ChangeOrQty T
	CreatedAt   time.Time
}

func GetVariants[T Mod.GetFamily](path string, token string, query T) []map[string]interface{} {
	return Lib.GetData[[]map[string]interface{}](path+Mod.QueryUrl(query), token)
}

func updateQtyBody(data *[]map[string]interface{}, qty int) string {
	output := ""
	for _, v := range *data {
		v["quantityOverwrite"] = make(map[string]bool, 2)
		v["quantityOverwrite"].(map[string]bool)["quantityOverwritten"] = true
		v["quantityOverwrite"].(map[string]bool)["isLockedByOrderVolumeProtection"] = false
		v["quantity"] = qty

		if qty > 0 {
			v["changeToListingStatus"] = map[string]string{"handle": "listed"}
		}
		output += fmt.Sprintf("%s's qty updated to: %d\n", v["sku"], qty)
	}
	return output
}

func updateStateBody(data *[]map[string]interface{}, state Mod.VariantState) string {
	output := ""
	for _, v := range *data {
		v["changeToListingStatus"] = map[string]string{"handle": string(state)}
		output += fmt.Sprintf("%s's status updating to: %s\n", v["sku"], state)
	}
	return output
}

func takeAction(toBeUpdated []map[string]interface{}, routine func() Mod.UpdateListingVariantQuery) {
	chunks := Lib.PartitionByN(toBeUpdated, Mod.MAXROUTINE)
	routineCount := len(chunks)
	var wg sync.WaitGroup
	wg.Add(routineCount)
	for i := 0; i < routineCount; i++ {
		x := chunks[i]
		go func() {
			for _, v := range x {
				Lib.PostData[map[string]interface{}](Mod.FLX_URL+Mod.LISTING_URL_EXT+Mod.PLURAL_VARIANT_URL_EXT+Mod.QueryUrl(routine()), v, Mod.RequestAccToken())
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func UpdateListingQty(allVariantFile string, regex []string, qty int) string {
	res := Lib.ReadAllLineAndFilter(allVariantFile, regex)
	toBeUpdated := []map[string]interface{}{}
	output := ""
	log := UpdateLog[int]{
		ChangeOrQty: qty,
		CreatedAt:   time.Now(),
	}
	for _, v := range res {
		resp := GetVariants(Mod.FLX_URL+Mod.LISTING_URL_EXT+Mod.PLURAL_VARIANT_URL_EXT, Mod.RequestAccToken(), Mod.GetListingVariant{Skus: v})
		log.Log = append(log.Log, v...)
		output = updateQtyBody(&resp, qty)
		toBeUpdated = append(toBeUpdated, resp...)
	}
	takeAction(toBeUpdated, Mod.QtyUpdateQuery)
	output += "\nBULK QTY UPDATE - COMPLETE"
	Lib.WriteJsonToFile("output/updateQtyLog.txt", log)
	return output
}

func UpdateListingState(allVariantFile string, regex []string, state Mod.VariantState) string {
	res := Lib.ReadAllLineAndFilter(allVariantFile, regex)
	toBeUpdated := []map[string]interface{}{}
	output := ""
	log := UpdateLog[string]{
		ChangeOrQty: string(state),
		CreatedAt:   time.Now(),
	}
	for _, v := range res {
		resp := GetVariants(Mod.FLX_URL+Mod.LISTING_URL_EXT+Mod.PLURAL_VARIANT_URL_EXT, Mod.RequestAccToken(), Mod.GetListingVariant{Skus: v})
		log.Log = append(log.Log, v...)
		output = updateStateBody(&resp, Mod.DELISTED)
		toBeUpdated = append(toBeUpdated, resp...)
	}
	takeAction(toBeUpdated, Mod.StatusUpdateQuery)
	output += "\nBULK DELIST - COMPLETE"
	Lib.WriteJsonToFile("output/updateDelistLog.txt", log)
	return output
}

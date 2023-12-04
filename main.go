package main

import (
	Lib "flx/lib"
	Mod "flx/model"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
)

const POOLLIMIT = 40

type Tokens struct {
	sources  []string
	channels []string
}

func gatherTokens() Tokens {
	return Tokens{
		sources:  []string{"AB", "SS", "SM"},
		channels: []string{"Amazon", "Walmart"},
	}
}

func main() {
	sources := []string{os.Getenv("FLX_AB_TOKEN"), os.Getenv("FLX_SS_TOKEN"), os.Getenv("FLX_SM_TOKEN")}
	channels := []string{os.Getenv("FLX_AZ_TOKEN"), os.Getenv("FLX_WALMART_TOKEN")}

	CountVariants(Mod.GET_INVENTORY_VARIANTS_PATH, Mod.GetInventoryVariant{Page: 0, PageSize: 100}, sources, gatherTokens().sources, "Inventory Variant Count:")
	CountVariants(Mod.GET_PRODUCT_VARIANTS_PATH, Mod.GetProductVariant{Page: 0, PageSize: 100}, sources, gatherTokens().sources, "Product Variant Count:")
	CountVariants(Mod.GET_LISTING_VARIANTS_PATH, Mod.GetListingVariant{Page: 0, PageSize: 100}, channels, gatherTokens().channels, "Listing Variant Count:")

}

func CountVariants[T Mod.GetFamily](path string, query T, tokens []string, tokenNames []string, message string) {
	for i := 0; i < len(tokens); i++ {
		var wg sync.WaitGroup
		var ops atomic.Uint64

		wg.Add(POOLLIMIT)

		for j := 1; j <= POOLLIMIT; j++ {
			query = query.StepPage(j).(T)
			go ConcurrentCount(path, &wg, &ops, tokens[i], query)
		}
		wg.Wait()
		fmt.Printf("%s "+message+" %d\n", tokenNames[i], ops.Load())
	}
}

func ConcurrentCount[T Mod.GetFamily](path string, wg *sync.WaitGroup, ops *atomic.Uint64, token string, query T) {

	for count := len(Lib.GetDataList(path+Mod.QueryUrl(query), token)); count != 0; count = len(Lib.GetDataList(path+Mod.QueryUrl(query), token)) {
		ops.Add(uint64(count))
		query = query.StepPage(POOLLIMIT).(T)
	}
	wg.Done()
}

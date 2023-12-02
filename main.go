package main

//sm -> 612515
//fol -> 570935
//sns -> 47537
//ab -> 48226
import (
	Lib "flx/lib"
	Mod "flx/model"
	"os"

	"github.com/kr/pretty"
)

func main() {

	// sources := []int{Mod.AB, Mod.SNS, Mod.SM, Mod.FOL}

	// //1 -> Source, 2 -> product, 3 -> channel
	// var sourceData map[int][]int = make(map[int][]int)

	// for i := 0; i < len(sources); i++ {
	// 	sourceData[sources[i]] = make([]int, 100)
	// }

	// for i := 0; i < len(sources); i++ {
	// 	sourceData[sources[i]][0] = len(Lib.GetDataList(Mod.GET_SEARCH_VARIANTS_PATH, model.GetSearchInventoryVariants{FilterSourceId: 48226}))
	// 	// sourceData[sources[i]][0] = len(Lib.GetDataList(Mod.GET_PRODUCT_VARIANTS_PATH, model.GetProductVariant{SourceId: sources[i]}))
	// 	// sourceData[sources[i]][0] = len(Lib.GetDataList(Mod.GET_INVENTORY_VARIANTS_PATH, model.GetInventoryVariant{SourceId: sources[i]}))
	// }

	extent := Mod.QueryUrl(Mod.GetInventoryVariant{Page: 1})
	pretty.Print(Lib.GetDataList(Mod.GET_INVENTORY_VARIANTS_PATH+extent, os.Getenv("FLX_AB_TOKEN")))
}

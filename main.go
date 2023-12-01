package main

import (
	Lib "flx/lib"
	Mod "flx/model"

	"github.com/kr/pretty"
)

func main() {

	res := Lib.GetDataList(Mod.GET_SOURCES_PATH, Mod.GET_SOURCES_PATH)

	var sourceIds []interface{}

	for i := 0; i < len(res); i++ {
		pretty.Println(res)
		break
	}

	//1 -> Source, 2 -> product, 3 -> channel
	// 	var sourceData map[int][]int = make(map[int][]int)

	// 	for i := 0; i < len(sources); i++ {
	// 		sourceData[sources[i]][0] = len(lib.GetData(Mod.GET_INVENTORY_VARIANTS_PATH, model.GetInventoryVariant{SourceId: sources[i]}))
	// 		sourceData[sources[i]][0] = len(lib.GetData(Mod.GET_PRODUCT_VARIANTS_PATH, model.GetProductVariant{SourceId: sources[i]}))
	// 		sourceData[sources[i]][0] = len(lib.GetData(Mod.GET_INVENTORY_VARIANTS_PATH, model.GetInventoryVariant{SourceId: sources[i]}))
	// 	}
	//

	// pretty.Print(Lib.GetData(Mod.GET_INVENTORY_VARIANTS_PATH, Mod.GetInventoryVariant{PageSize: 66, SourceId: Mod.SNS})[0])
	pretty.Print(sourceIds)
}

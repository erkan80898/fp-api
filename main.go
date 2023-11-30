package main

import (
	lib "flx/lib"
	"flx/model"
	Mod "flx/model"

	"github.com/kr/pretty"
)

func main() {

	result := lib.GetData(Mod.GET_LISTING_VARIANTS_PATH, model.GetListingVariant{Skus: []string{"F_M_BB_AST"}})
	pretty.Print(result)
	println("-------------------------------------------------------")

	result = lib.GetData(Mod.GET_INVENTORY_VARIANTS_PATH, model.GetInventoryVariant{SourceId: Mod.ALPHA})
	pretty.Print(result)
	println("-----------------------END--------------------------------")
}

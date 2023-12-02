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

	extent := Mod.QueryUrl(Mod.GetInventoryVariant{Page: 4})
	// pretty.Print(Lib.GetDataList(Mod.GET_INVENTORY_VARIANTS_PATH+extent, os.Getenv("FLX_AB_TOKEN")))
	// pretty.Print(Lib.GetDataList(Mod.GET_INVENTORY_VARIANTS_PATH+extent, os.Getenv("FLX_SS_TOKEN")))
	pretty.Print(Lib.GetDataList(Mod.GET_INVENTORY_VARIANTS_PATH+extent, os.Getenv("FLX_SM_TOKEN")))
}

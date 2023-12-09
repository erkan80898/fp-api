package action

import (
	Lib "flx/lib"
	Mod "flx/model"
	"fmt"
	"strings"
	"time"

	"github.com/kr/pretty"
)

type UpdateLog struct {
	Log       []string
	CreatedAt time.Time
}

func GetVariants[T Mod.GetFamily](path string, token string, query T) []map[string]interface{} {
	return Lib.GetDataList(path+Mod.QueryUrl(query), token)
}

func UpdateListingQty(allVariantFile string, regex []string, qty int) string {
	res := Lib.ReadAllLineAndFilter(allVariantFile, regex)
	toBeUpdated := []map[string]interface{}{}
	output := ""
	log := UpdateLog{
		CreatedAt: time.Now(),
	}

	for _, v := range res {

		resp := GetVariants(Mod.FLX_URL+Mod.LISTING_URL_EXT+Mod.PLURAL_VARIANT_URL_EXT, Mod.RequestAccToken(), Mod.GetListingVariant{Skus: v, IncludeOverwrites: true})
		log.Log = append(log.Log, v...)
		output = UpdateQtyBody(&resp, qty)
		toBeUpdated = append(toBeUpdated, resp...)
	}

	for _, v := range toBeUpdated {
		Lib.PostDataJson(Mod.FLX_URL+Mod.LISTING_URL_EXT+Mod.PLURAL_VARIANT_URL_EXT+Mod.QueryUrl(Mod.QtyUpdateOnlyQuery()), v, Mod.RequestAccToken())
	}

	output += "\nBULK QTY UPDATE - COMPLETE"

	pretty.Println(log)
	Lib.WriteJsonToFile("updateQtyLog.txt", log)
	return output
}

func UpdateQtyBody(data *[]map[string]interface{}, qty int) string {
	output := ""

	for _, v := range *data {
		v["quantity"] = qty
		output += fmt.Sprintf("%s's qty updated to: %d\n", v["sku"], qty)
	}
	return output
}

func Run(qty int, skusAsText string) string {

	parts := strings.Split(skusAsText, ",")

	for i, _ := range parts {
		parts[i] = strings.TrimSpace(parts[i])
		if len(parts[i]) <= 4 {
			parts[i] = "_" + parts[i] + "_"
		}
	}

	return UpdateListingQty("fruitListingVariant.csv", parts, qty)
}

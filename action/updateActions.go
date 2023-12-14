package action

import (
	Lib "flx/lib"
	Mod "flx/model"
	"fmt"
	"strings"
	"sync"
	"time"
)

type UpdateLog struct {
	Log       []string
	Qty       int
	CreatedAt time.Time
}

func GetVariants[T Mod.GetFamily](path string, token string, query T) []map[string]interface{} {
	return Lib.GetData[[]map[string]interface{}](path+Mod.QueryUrl(query), token)
}

func UpdateListingQty(allVariantFile string, regex []string, qty int) string {
	res := Lib.ReadAllLineAndFilter(allVariantFile, regex)
	toBeUpdated := []map[string]interface{}{}
	output := ""
	log := UpdateLog{
		Qty:       qty,
		CreatedAt: time.Now(),
	}

	for _, v := range res {
		resp := GetVariants(Mod.FLX_URL+Mod.LISTING_URL_EXT+Mod.PLURAL_VARIANT_URL_EXT, Mod.RequestAccToken(), Mod.GetListingVariant{Skus: v})
		log.Log = append(log.Log, v...)
		output = UpdateQtyBody(&resp, qty)
		toBeUpdated = append(toBeUpdated, resp...)
	}

	chunks := Lib.PartitionByN(toBeUpdated, Mod.MAXROUTINE)

	routineCount := len(chunks)
	var wg sync.WaitGroup
	wg.Add(routineCount)

	for i := 0; i < routineCount; i++ {
		x := chunks[i]
		go func() {
			for _, v := range x {
				Lib.PostData[map[string]interface{}](Mod.FLX_URL+Mod.LISTING_URL_EXT+Mod.PLURAL_VARIANT_URL_EXT+Mod.QueryUrl(Mod.QtyUpdateOnlyQuery()), v, Mod.RequestAccToken())
			}
			wg.Done()
		}()
	}

	wg.Wait()
	output += "\nBULK QTY UPDATE - COMPLETE"

	Lib.WriteJsonToFile("updateQtyLog.txt", log)
	return output
}

func UpdateQtyBody(data *[]map[string]interface{}, qty int) string {
	output := ""

	for _, v := range *data {
		v["quantityOverwrite"] = make(map[string]bool, 2)
		v["quantityOverwrite"].(map[string]bool)["quantityOverwritten"] = true
		v["quantityOverwrite"].(map[string]bool)["isLockedByOrderVolumeProtection"] = false
		v["quantity"] = qty
		output += fmt.Sprintf("%s's qty updated to: %d\n", v["sku"], qty)
	}
	return output
}

func Run(qty int, skusAsText string) string {

	parts := strings.Split(skusAsText, ",")
	res := []string{}
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
		if strings.Index(parts[i], "[") != -1 {
			letter := string(parts[i][0])
			rest := string(parts[i][2 : len(parts[i])-1])
			nums := strings.Split(rest, "+")
			for _, v := range nums {
				res = append(res, "_"+letter+v+"_")
			}
		} else if len(parts[i]) <= 4 {
			res = append(res, "_"+parts[i]+"_")
		} else {
			res = append(res, parts[i])
		}
	}

	return UpdateListingQty("fruitListingVariant.csv", res, qty)
}

package action

import (
	Lib "flx/lib"
	Mod "flx/model"
)

func GetVariants[T Mod.GetFamily](path string, token string, query T) []map[string]interface{} {
	return Lib.GetDataList(path+Mod.QueryUrl(query), token)
}

func UpdateListingQty(allVariantFile string, regex []string, qty int) error {
	res := Lib.ReadAllLineAndFilter(allVariantFile, regex)
	toBeUpdated := []map[string]interface{}{}

	for _, v := range res {

		resp := GetVariants(Mod.FLX_URL+Mod.LISTING_URL_EXT+Mod.PLURAL_VARIANT_URL_EXT, Mod.RequestAccToken(), Mod.GetListingVariant{Skus: v, IncludeOverwrites: true})
		Lib.UpdateQtyBody(&resp, qty)
		toBeUpdated = append(toBeUpdated, resp...)
	}

	for _, v := range toBeUpdated {
		Lib.PostDataJson(Mod.FLX_URL+Mod.LISTING_URL_EXT+Mod.PLURAL_VARIANT_URL_EXT+Mod.QueryUrl(Mod.QtyUpdateOnlyQuery()), v, Mod.RequestAccToken())
	}

	println("BULK QTY UPDATE - COMPLETE")
	return nil
}

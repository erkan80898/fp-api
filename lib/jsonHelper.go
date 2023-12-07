package lib

import (
	Mod "flx/model"
)

func UpdateQtyBody(data []interface{}, qty int) []Mod.UpdateListingVariantBody {
	result := []Mod.UpdateListingVariantBody{}

	for _, v := range data {
		entry := Mod.UpdateListingVariantBody{}

		v := v.(map[string]interface{})
		entry.Sku = v["sku"].(string)
		entry.ChannelId = int(v["channelId"].(float64))

		entry.ProductVariantId = 0
		entry.QuantityOverwrite = map[string]interface{}{"quantityOverwritten": true, "lastManualQuantity": 10}
		entry.Quantity = qty
		result = append(result, entry)
	}
	return result
}

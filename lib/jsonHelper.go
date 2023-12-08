package lib

import "fmt"

func UpdateQtyBody(data *[]map[string]interface{}, qty int) {

	for _, v := range *data {
		v["quantity"] = qty
		fmt.Printf("%s's qty updated to: %d\n", v["sku"], qty)
	}
}

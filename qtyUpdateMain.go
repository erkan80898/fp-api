package main

import (
	Action "flx/action"
)

func main() {
	Action.UpdateListingQty("fruitListingVariant.csv", []string{"_P20_", "8_.*P_6DCSHC1_", "_M41_", "_S1_", "_S4_", "_P2_", "_P14_"}, 0)
}

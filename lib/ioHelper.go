package lib

import (
	Mod "flx/model"
)

type VariantCountAll struct {
	InventoryVariant map[string]int
	ProductVariant   map[string]int
	ChannelVariant   map[string]int
}

func InitVariantCountAll() VariantCountAll {
	tokenNames := Mod.GatherTokens()

	invVar := make(map[string]int)
	productVar := make(map[string]int)
	channelVar := make(map[string]int)

	for _, v := range tokenNames.Sources {
		invVar[v] = 0
	}
	for _, v := range tokenNames.Sources {
		productVar[v] = 0
	}
	for _, v := range tokenNames.Channels {
		channelVar[v] = 0
	}

	return VariantCountAll{invVar, productVar, channelVar}
}

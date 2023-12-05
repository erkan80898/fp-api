package lib

import (
	"encoding/json"
	Mod "flx/model"
	"io/ioutil"
	"log"
	"os"
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

func WriteJsonToFile(fileName string, data interface{}) {
	jsonData, err := json.Marshal(data)

	if err != nil {
		log.Panic(err)
	}
	ioutil.WriteFile(fileName, jsonData, os.ModePerm)
}

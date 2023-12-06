package lib

import (
	"encoding/json"
	Mod "flx/model"
	"io"
	"log"
	"os"
	"time"
)

type VariantCountAll struct {
	InventoryVariant map[string]int
	ProductVariant   map[string]int
	ChannelVariant   map[string]int
	CreatedAt        time.Time
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

	return VariantCountAll{invVar, productVar, channelVar, time.Time{}}
}

func WriteJsonToFile(fileName string, data interface{}) {
	jsonData, err := json.Marshal(data)

	if err != nil {
		log.Panic(err)
	}

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Panic(err)
	}

	if _, err := io.WriteString(file, string(jsonData)+"\n"); err != nil {
		log.Panic(err)
	}
}

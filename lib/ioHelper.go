package lib

import (
	"bufio"
	"encoding/json"
	Mod "flx/model"
	"io"
	"log"
	"os"
	"regexp"
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

// Max sku allowance = 50, so this function will break things up once its over 50 for the first dimension
func ReadAllLineAndFilter(fileName string, rule string) [][]string {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	res := [][]string{}

	for i, j := 0, 0; fileScanner.Scan(); {
		if i == 50 {
			j++
			i = 0
		}

		line := fileScanner.Text()
		if hit, _ := regexp.MatchString(rule, line); hit {
			if j == len(res) {
				res = append(res, []string{})
			}
			res[j] = append(res[j], line)
			i++
		}
	}
	return res
}

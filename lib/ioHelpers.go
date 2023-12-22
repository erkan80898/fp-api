package lib

import (
	"bufio"
	Mod "flx/model"
	"github.com/kr/pretty"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
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
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Panic(err)
	}
	if _, err := io.WriteString(file, pretty.Sprint(data)+"\n"); err != nil {
		log.Panic(err)
	}
}

// max sku allowance = 50
// break things up once its over 50 for the first dimension
func ReadAllLineAndFilter(fileName string, rules []string) [][]string {
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
		line := strings.ToUpper(fileScanner.Text())
		for _, v := range rules {
			v = strings.ToUpper(v)
			if hit, _ := regexp.MatchString(v, line); hit {
				if j == len(res) {
					res = append(res, []string{})
				}
				res[j] = append(res[j], line)
				i++
			}
		}
	}
	return res
}

package model

import (
	"reflect"
	"strconv"
)

const (
	FLX_URL                = "https://api.flxpoint.com/"
	INVENTORY_URL_EXT      = "inventory/"
	PRODUCT_URL_EXT        = "product/"
	LISTING_URL_EXT        = "listing/"
	SOURCES_URL_EXT        = "sources/"
	CHANNEL_URL_EXT        = "channel/"
	PARENT_URL_EXT         = "parent/"
	VARIANT_URL_EXT        = "variant/"
	PLURAL_PARENT_URL_EXT  = "parents/"
	PLURAL_VARIANT_URL_EXT = "variants/"

	SEARCH_URL_EXT = "search/"
)

const (
	GET_INVENTORY_VARIANTS_PATH = FLX_URL + INVENTORY_URL_EXT + PLURAL_VARIANT_URL_EXT
	GET_PRODUCT_VARIANTS_PATH   = FLX_URL + PRODUCT_URL_EXT + PLURAL_VARIANT_URL_EXT
	GET_LISTING_VARIANTS_PATH   = FLX_URL + LISTING_URL_EXT + PLURAL_VARIANT_URL_EXT
	GET_SOURCES_PATH            = FLX_URL + SOURCES_URL_EXT
	GET_SEARCH_VARIANTS_PATH    = FLX_URL + INVENTORY_URL_EXT + SEARCH_URL_EXT + VARIANT_URL_EXT
)

const (
	AMAZON  int = 48226
	WALMART int = 47537
)

const (
	SNS int = 47537
	AB  int = 48226
	SM  int = 612515
	FOL int = 570935
)

type GetInventoryVariant struct {
	Page                    int      `json:"page"`
	PageSize                int      `json:"pageSize"`
	FilterArchived          bool     `json:"filterArchived"`
	FilterNeedsDeleting     bool     `json:"filterNeedsDeleting"`
	Ids                     []int    `json:"ids"`
	IncludeAttributes       bool     `json:"includeAttributes"`
	IncludeCategories       bool     `json:"includeCategories"`
	IncludeCustomAggregates bool     `json:"includeCustomAggregates"`
	IncludeCustomFields     bool     `json:"includeCustomFields"`
	IncludeImages           bool     `json:"includeImages"`
	IncludeOptions          bool     `json:"includeOptions"`
	IncludeParent           bool     `json:"includeParent"`
	Skus                    []string `json:"skus"`
	SourceId                int      `json:"sourceId"`
	UpdatedAfter            string   `json:"updatedAfter"`
}

type GetProductVariant struct {
	IncludeTags             bool     `json:"includeTag"`
	Page                    int      `json:"page"`
	PageSize                int      `json:"pageSize"`
	Deleting                bool     `json:"deleting"`
	Ids                     []int    `json:"ids"`
	IncludeAttributes       bool     `json:"includeAttributes"`
	IncludeBundleComponents bool     `json:"IncludeBundleComponents"`
	IncludeCategories       bool     `json:"includeCategories"`
	IncludeCustomAggregates bool     `json:"includeCustomAggregates"`
	IncludeCustomFields     bool     `json:"includeCustomFields"`
	IncludeImages           bool     `json:"includeImages"`
	IncludeInventoryLinks   bool     `json:"includeInventoryLinks"`
	IncludeListingLinks     bool     `json:"includeListingLinks"`
	IncludeOptions          bool     `json:"includeOptions"`
	IncludeOverwrites       bool     `json:"includeOverwrites"`
	IncludeParent           bool     `json:"includeParent"`
	Skus                    []string `json:"skus"`
	UpdatedAfter            string   `json:"updatedAfter"`
}

type GetListingVariant struct {
	IncludeOverwrites       bool     `json:"includeOverwrites"`
	IncludeTags             bool     `json:"includeTags"`
	Page                    int      `json:"page"`
	PageSize                int      `json:"pageSize"`
	BasicsUpdatedAfter      string   `json:"basicsUpdatedAfter"`
	ChannelId               int      `json:"channelId"`
	Ids                     []int    `json:"ids"`
	IncludeAttributes       bool     `json:"includeAttributes"`
	IncludeCategories       bool     `json:"includeCategories"`
	IncludeCustomAggregates bool     `json:"includeCustomAggregates"`
	IncludeCustomFields     bool     `json:"includeCustomFields"`
	IncludeImages           bool     `json:"includeImages"`
	IncludeOptions          bool     `json:"includeOptions"`
	IncludeParent           bool     `json:"includeParent"`
	Skus                    []string `json:"skus"`
	Sync                    bool     `json:"sync"`
	UpdatedAfter            string   `json:"updatedAfter"`
}

type GetChannels struct {
	IncludeArchived     bool `json:"includeArchived"`
	IncludeIntegrations bool `json:"includeIntegrations"`
}

type GetSources struct {
	FilterArchived        bool     `json:"filterArchived"`
	FilterByIds           []string `json:"filterByIds"`
	FilterFirearmsEnabled bool     `json:"filterFirearmsEnabled"`
	IncludeAddress        bool     `json:"includeAddress"`
	IncludeIntegrations   bool     `json:"includeIntegrations"`
}

type GetSearchInventoryVariants struct {
	FilterByInventoryVariantIds string   `json:"filterByInventoryVariantIds"`
	FilterBySkus                []string `json:"filterBySkus"`
	FilterNeedsDeleting         bool     `json:"filterNeedsDeleting"`
	FilterPageNumber            int      `json:"filterPageNumber"`
	FilterPageSize              int      `json:"filterPageSize"`
	FilterSourceId              int      `json:"filterSourceId"`
	FilterUpdatedAfter          []string `json:"filterUpdatedAfter"`
	Ids                         []int    `json:"ids"`
}

func QueryUrl(data interface{}) string {

	res := "?"

	v := reflect.ValueOf(data)
	typeOfData := v.Type()

	for i := 0; i < v.NumField(); i++ {
		valName := typeOfData.Field(i).Tag.Get("json")
		valType := v.Field(i).Type().Kind()
		value := v.Field(i).Interface()

		switch valType {
		case reflect.Int:
			val := value.(int)
			if val == 0 {
				continue
			}

			res += valName + "=" + strconv.Itoa(val)

		case reflect.String:
			val := value.(string)
			if val == "" {
				continue
			}
			res += valName + "=" + val

		case reflect.Bool:
			val := v.Field(i).Interface().(bool)
			res += valName + "=" + strconv.FormatBool(val)

		case reflect.Array:
			if reflect.SliceOf(v.Field(i).Type()).Kind() == reflect.Int {
				res += valName + "="
				val := v.Field(i).Interface().([]int)
				if len(val) == 0 {
					continue
				}
				for i := 0; i < len(val); i++ {
					res += strconv.Itoa(val[i]) + ","
				}
			} else {
				res += valName + "="
				val := v.Field(i).Interface().([]string)

				if len(val) == 0 {
					continue
				}

				for i := 0; i < len(val); i++ {
					res += val[i] + ","
				}
			}
		}

		if i != v.NumField()-1 {
			res += "&"
		}
	}
	return res
}

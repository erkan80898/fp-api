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
	COUNT_URL_EXT          = "count/"
	SEARCH_URL_EXT         = "search/"
)

const (
	GET_INVENTORY_VARIANTS_PATH = FLX_URL + INVENTORY_URL_EXT + PLURAL_VARIANT_URL_EXT
	GET_PRODUCT_VARIANTS_PATH   = FLX_URL + PRODUCT_URL_EXT + PLURAL_VARIANT_URL_EXT
	GET_LISTING_VARIANTS_PATH   = FLX_URL + LISTING_URL_EXT + PLURAL_VARIANT_URL_EXT
	GET_SOURCES_PATH            = FLX_URL + SOURCES_URL_EXT
	GET_SEARCH_VARIANTS_PATH    = FLX_URL + INVENTORY_URL_EXT + SEARCH_URL_EXT + VARIANT_URL_EXT
)

type GetFamily interface {
	GetInventoryVariant | GetProductVariant | GetListingVariant
	StepPage(int) interface{}
}

type GetInventoryVariant struct {
	Page                         int      `json:"page"`
	PageSize                     int      `json:"pageSize"`
	FilterArchived               bool     `json:"filterArchived"`
	FilterNeedsDeleting          bool     `json:"filterNeedsDeleting"`
	Ids                          []int    `json:"ids"`
	IncludeAttributes            bool     `json:"includeAttributes"`
	IncludeCategories            bool     `json:"includeCategories"`
	IncludeCustomAggregates      bool     `json:"includeCustomAggregates"`
	IncludeCustomFields          bool     `json:"includeCustomFields"`
	IncludeImages                bool     `json:"includeImages"`
	IncludeLinkedProductVariants bool     `json:"includeLinkedProductVariants"`
	IncludeOptions               bool     `json:"includeOptions"`
	IncludeParent                bool     `json:"includeParent"`
	Skus                         []string `json:"skus"`
	SourceId                     int      `json:"sourceId"`
	UpdatedAfter                 string   `json:"updatedAfter"`
}

func (x GetInventoryVariant) StepPage(step int) interface{} {
	x.Page += step
	return x
}

type GetProductVariant struct {
	IncludeTags             bool     `json:"includeTag"`
	Page                    int      `json:"page"`
	PageSize                int      `json:"pageSize"`
	Deleting                bool     `json:"deleting"`
	Ids                     []int    `json:"ids"`
	IncludeAttributes       bool     `json:"includeAttributes"`
	IncludeBundleComponents bool     `json:"includeBundleComponents"`
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

func (x GetProductVariant) StepPage(step int) interface{} {
	x.Page += step
	return x
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

func (x GetListingVariant) StepPage(step int) interface{} {
	x.Page += step
	return x
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

type GetCountListingVariant struct {
	BasicsUpdatedAfter  string   `json:"basicsUpdatedAfter"`
	ChannelId           int      `json:"channelId"`
	IncludeDeleting     bool     `json:"includeDeleting"`
	IncludeImagesCached bool     `json:"includeImagesCached"`
	IncludePaused       bool     `json:"includePaused"`
	MinimumQuantity     int      `json:"minimumQuantity"`
	Publish             bool     `json:"publish"`
	Skus                []string `json:"skus"`
	Sync                bool     `json:"sync"`
	UpdatedAfter        string   `json:"updatedAfter"`
}

type CreateOrUpdateListingVariant struct {
	ModifyChannelData         string `json:"modifyChannelData,omitempty"`
	ModifyContent             string `json:"modifyContent,omitempty"`
	ModifyDimensions          string `json:"modifyDimensions,omitempty"`
	ModifyIdentifiers         string `json:"modifyIdentifiers,omitempty"`
	ModifyImages              string `json:"modifyImages,omitempty"`
	ModifyOptions             string `json:"modifyOptions,omitempty"`
	ModifyPriceOverwrite      string `json:"modifyPriceOverwrite,omitempty"`
	ModifyPublishData         string `json:"modifyPublishData,omitempty"`
	ModifyQuantityOverwrite   string `json:"modifyQuantityOverwrite,omitempty"`
	ModifyStatusData          string `json:"modifyStatusData,omitempty"`
	ModifySyncData            string `json:"modifySyncData,omitempty"`
	ModifyVariantCustomFields string `json:"modifyVariantCustomFields,omitempty"`
	RestrictCreateOrUpdate    string `json:"restrictCreateOrUpdate,omitempty"`
	MasterSkuUpdateContent    bool   `json:"masterSkuUpdateContent,omitempty"`
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

		case reflect.Slice:
			if v.Field(i).Type().Elem().Kind() == reflect.Int {
				val := v.Field(i).Interface().([]int)
				if len(val) == 0 {
					continue
				}
				res += valName + "="
				for i := 0; i < len(val); i++ {
					res += "&" + valName + "=" + strconv.Itoa(val[i])
				}
			} else {
				val := v.Field(i).Interface().([]string)

				if len(val) == 0 {
					continue
				}
				res += valName + "="
				for i := 0; i < len(val); i++ {
					res += "&" + valName + "=" + val[i]
				}
			}
		}

		if i != v.NumField()-1 {
			res += "&"
		}
	}
	return res
}

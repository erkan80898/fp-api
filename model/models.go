package model

const FLX_URL = "https://api.flxpoint.com/"
const LISTING_URL_EXT = "listing/"
const INVENTORY_URL_EXT = "inventory/"
const PARENT_URL_EXT = "parent/"
const VARIANT_URL_EXT = "variant/"
const PLURAL_PARENT_URL_EXT = "parents/"
const PLURAL_VARIANT_URL_EXT = "variants/"

const GET_INVENTORY_VARIANTS_PATH = FLX_URL + LISTING_URL_EXT + PLURAL_VARIANT_URL_EXT
const GET_LISTING_VARIANTS_PATH = FLX_URL + INVENTORY_URL_EXT + PLURAL_VARIANT_URL_EXT

const (
	ALPHA  int = 48226
	SNS    int = 47537
	FRUIT  int = 570935
	SANMAR int = 612515
)

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

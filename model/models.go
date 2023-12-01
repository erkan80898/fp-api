package model

type Number interface {
	int | float64
}

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
)

const (
	GET_INVENTORY_VARIANTS_PATH = FLX_URL + INVENTORY_URL_EXT + PLURAL_VARIANT_URL_EXT
	GET_PRODUCT_VARIANTS_PATH   = FLX_URL + PRODUCT_URL_EXT + PLURAL_VARIANT_URL_EXT
	GET_LISTING_VARIANTS_PATH   = FLX_URL + LISTING_URL_EXT + PLURAL_VARIANT_URL_EXT
	GET_SOURCES_PATH            = FLX_URL + SOURCES_URL_EXT
)

const (
	AMAZON  int = 48226
	WALMART int = 47537
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
	FilterArchived        bool `json:"filterArchived"`
	FilterByIds           bool `json:"filterByIds"`
	FilterFirearmsEnabled bool `json:"filterFirearmsEnabled"`
	IncludeAddress        bool `json:"includeAddress"`
	IncludeIntegrations   bool `json:"includeIntegrations"`
}

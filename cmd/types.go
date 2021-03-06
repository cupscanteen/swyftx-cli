package cmd

import "math"

type MarketsInfoBasicDTO []MarketsInfoBasic
type MarketsInfoDetailDTO []MarketsInfoDetail
type AssetHistoryAllDTO AssetHistoryAll

type AssetHistoryAll struct {
	Items []struct {
		Date             int64  `json:"date"`
		Amount           string `json:"amount"`
		Movement         string `json:"movement"`
		ApproxMovement   string `json:"approxMovement"`
		UserCountryValue string `json:"userCountryValue"`
		UUID             string `json:"uuid"`
		Type             string `json:"type"`
		Status           string `json:"status"`
		StatusRaw        int    `json:"statusRaw"`
		OrderType        int    `json:"orderType"`
		SecondaryAsset   int    `json:"secondaryAsset"`
		PrimaryAsset     int    `json:"primaryAsset"`
		SecondaryAmount  string `json:"secondaryAmount"`
		PrimaryAmount    string `json:"primaryAmount"`
	} `json:"items"`
	RecordCount int      `json:"recordCount,omitempty"`
	Metadata    Metadata `json:"metadata"`
}

type MarketsInfoBasic struct {
	Name      string  `json:"name"`
	AltName   string  `json:"altName"`
	Code      string  `json:"code"`
	ID        int     `json:"id"`
	Rank      int     `json:"rank"`
	Buy       string  `json:"buy,omitempty"`
	Sell      string  `json:"sell,omitempty"`
	Spread    string  `json:"spread,omitempty"`
	Volume24H float64 `json:"volume24H"`
	MarketCap float64 `json:"marketCap"`
}

type MarketsInfoDetail struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Mineable    bool   `json:"mineable"`
	Spread      string `json:"spread"`
	Rank        int    `json:"rank"`
	RankSuffix  string `json:"rankSuffix"`
	Volume      struct {
		Two4H float64 `json:"24H"`
	} `json:"volume"`
	PriceChange struct {
		Week  float64 `json:"week"`
		Month float64 `json:"month"`
	} `json:"priceChange"`
	Urls struct {
		Explorer string `json:"explorer"`
		Reddit   string `json:"reddit"`
		Twitter  string `json:"twitter"`
		Website  string `json:"website"`
	} `json:"urls"`
	Supply struct {
		Circulating float64 `json:"circulating"`
		Total       float64 `json:"total"`
		Max         float64 `json:"max"`
	} `json:"supply"`
}

type LiveRates map[string]interface{}

type Metadata struct {
	CurrentPage  int `json:"current_page,omitempty"`
	PageSize     int `json:"page_size,omitempty"`
	FirstPage    int `json:"first_page,omitempty"`
	LastPage     int `json:"last_page,omitempty"`
	TotalRecords int `json:"total_records,omitempty"`
}

func CalculateMetadata(totalRecords, page, pageSize int) Metadata {
	if totalRecords == 0 {
		return Metadata{}
	}
	return Metadata{
		CurrentPage:  page,
		PageSize:     pageSize,
		FirstPage:    1,
		LastPage:     int(math.Ceil(float64(totalRecords) / float64(pageSize))),
		TotalRecords: totalRecords,
	}
}

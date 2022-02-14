package cmd

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
	RecordCount int `json:"recordCount"`
}

type MarketsBasicInfo []struct {
	Name      string `json:"name"`
	AltName   string `json:"altName"`
	Code      string `json:"code"`
	ID        int    `json:"id"`
	Rank      int    `json:"rank"`
	Buy       string `json:"buy"`
	Sell      string `json:"sell"`
	Spread    string `json:"spread"`
	Volume24H int64  `json:"volume24H"`
	MarketCap int64  `json:"marketCap"`
}
type MarketsDetailInfo []struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Mineable    bool   `json:"mineable"`
	Spread      string `json:"spread"`
	Rank        int    `json:"rank"`
	RankSuffix  string `json:"rankSuffix"`
	Volume      struct {
		Two4H int64 `json:"24H"`
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
		Circulating int `json:"circulating"`
		Total       int `json:"total"`
		Max         int `json:"max"`
	} `json:"supply"`
}

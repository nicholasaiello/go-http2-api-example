package entity

type ProductDetails struct {
	Symbol      string `json:"displaySymbol"`
	Description string `json:"symbolDesc"`
	TypeCode    string `json:"typeCode"`
	SubTypeCode string `json:"subTypeCode"`
}

type BasePosition struct {
	Quantity     int64   `json:"quantity"`
	PricePaid    float64 `json:"pricePaid"`
	TotalCost    float64 `json:"totalCost"`
	Fees         int64   `json:"fees"`
	AcquiredDate string  `json:"dateAcquired"`
}

type PositionLot struct {
	BasePosition
	ID string `json:"lotID"`
}

type Position struct {
	ID            string         `json:"positionID"`
	Symbol        string         `json:"symbol"`
	Details       ProductDetails `json:"productDetails"`
	Lots          []PositionLot  `json:"posLotDetails"`
	LastTradeDate string         `json:"lastTradeTime"`
	// BasePosition
	Quantity     int64   `json:"quantity"`
	PricePaid    float64 `json:"pricePaid"`
	TotalCost    float64 `json:"totalCost"`
	Fees         int64   `json:"fees"`
	AcquiredDate string  `json:"dateAcquired"`
}

type Portfolio struct {
	ID                  string     `json:"portfolioID"`
	Positions           []Position `json:"positionDetails"`
	TotalPositionsCount int        `json:"totalPositionCount"`
	ViewName            string     `json:"viewName"`
	LastUpdated         string     `json:"asOfDate"`
}

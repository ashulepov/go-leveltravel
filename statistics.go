package leveltravel

type cityResponse struct {
	ID     int    `json:"id"`
	NameRu string `json:"name_ru"`
	NameEn string `json:"name_en"`
}

type countryResponse struct {
	ID     int    `json:"id"`
	NameRu string `json:"name_ru"`
	NameEn string `json:"name_en"`
	Iso2   string `json:"iso2"`
}

type fromResponse struct {
	City    cityResponse    `json:"city"`
	Country countryResponse `json:"country"`
}

type toResponse struct {
	City     cityResponse    `json:"city"`
	Country  countryResponse `json:"country"`
	HotelIDs []int           `json:"hotel_ids"`
}

type nightsResponse struct {
	From int `json:"from"`
	To   int `json:"to"`
}

type kidsResponse struct {
	Count int   `json:"count"`
	Ages  []int `json:"ages"`
}

type starsResponse struct {
	From int `json:"from"`
	To   int `json:"to"`
}

type optionsResponse struct {
	FromPrice int  `json:"from_price"`
	FlexDates bool `json:"flex_dates"`
}

type searchParamsResponse struct {
	Query     string          `json:"query"`
	From      fromResponse    `json:"from"`
	To        toResponse      `json:"to"`
	StartDate string          `json:"start_date"`
	Nights    nightsResponse  `json:"nights"`
	Adults    int             `json:"adults"`
	Kids      kidsResponse    `json:"kids"`
	Stars     starsResponse   `json:"stars"`
	Options   optionsResponse `json:"options"`
}

type statisticResponse struct {
	Date      string  `json:"date"`
	Link      string  `json:"link"`
	Price     int     `json:"price"`
	Percent   float32 `json:"percent"`
	BestPrice bool    `json:"best_price"`
}

// StatisticsPrices is the struct for Statistics Prices Response
type StatisticsPrices struct {
	Count        int                  `json:"count"`
	AveragePrice int                  `json:"average_price"`
	MinPrice     int                  `json:"min_price"`
	MaxPrice     int                  `json:"max_price"`
	MinDate      string               `json:"min_date"`
	MaxDate      string               `json:"max_date"`
	SearchParams searchParamsResponse `json:"search_params"`
	Statistic    []statisticResponse  `json:"statistic"`
}

// StatisticsPricesFilter is query params for Statistics Prices Request
type StatisticsPricesFilter struct {
	Count       int    `json:"count"`
	FromCity    string `json:"from_city"`
	FromCountry string `json:"from_country"`
	ToCountry   string `json:"to_country"`
	Nights      int    `json:"nights"`
	Adults      int    `json:"adults"`
	StartDate   string `json:"start_date"`
	Kids        int    `json:"kids"`
	KidsAges    string `json:"kids_ages"`
	StarsFrom   int    `json:"stars_from"`
	StarsTo     int    `json:"stars_to"`
	FlexDates   bool   `json:"flex_dates"`
}

// GetStatisticsPrices requests Statistics Prices
func (a *API) GetStatisticsPrices(args StatisticsPricesFilter) (prices StatisticsPrices, err error) {
	err = a.get("/statistics/prices", convertToMap(args), &prices)
	return prices, err
}

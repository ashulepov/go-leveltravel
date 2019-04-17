package leveltravel

type airportResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Code     string `json:"code"`
	Lat      string `json:"lat"`
	Long     string `json:"long"`
	Timezone string `json:"timezone"`
}

type departureResponse struct {
	ID          int               `json:"id"`
	NameRu      string            `json:"name_ru"`
	NameRuForm1 string            `json:"name_ru_form1"`
	NameRuForm2 string            `json:"name_ru_form2"`
	NameEn      string            `json:"name_en"`
	IATA        string            `json:"iata"`
	Priority    int               `json:"priority"`
	ISO2        string            `json:"iso2"`
	Prioritized bool              `json:"prioritized"`
	Airports    []airportResponse `json:"airports"`
}

// ReferencesDepartures is the struct for References Departures Response
type ReferencesDepartures struct {
	Success          bool                `json:"success"`
	PrioritizedCount int                 `json:"prioritized_count"`
	Departures       []departureResponse `json:"departures"`
}

// GetReferencesDepartures requests References Departures
func (a *API) GetReferencesDepartures() (departures ReferencesDepartures, err error) {
	err = a.get("/references/departures", nil, &departures)
	return departures, err
}

type destinationsCountryResponse struct {
	ID       int            `json:"id"`
	NameRu   string         `json:"name_ru"`
	NameEn   string         `json:"name_en"`
	ISO2     string         `json:"iso2"`
	Priority int            `json:"priority"`
	Cities   []cityResponse `json:"cities"`
}

// ReferencesDestinations is the struct for References Destinations Response
type ReferencesDestinations struct {
	Success   bool                          `json:"success"`
	Countries []destinationsCountryResponse `json:"countries"`
}

// GetReferencesDestinations requests References Destinations
func (a *API) GetReferencesDestinations() (destinations ReferencesDestinations, err error) {
	err = a.get("/references/destinations", nil, &destinations)
	return destinations, err
}

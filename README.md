# go-leveltravel
Level.Travel API wrapper in golang

## Installing

```
go get github.com/ashulepov/go-leveltravel
```

## Using

```
// Init API
lt := leveltravel.NewAPI("APIKey, "AffiliateMark")

// Get Departures
departures, err := lt.GetReferencesDepartures()

// Get Destinations
departures, err := lt.GetReferencesDestinations()

// Get Prices
filter := leveltravel.StatisticsPricesFilter{
	Count:       55,
	FromCity:    "Moscow",
	FromCountry: "RU",
	ToCountry:   "TH",
	Nights:      14,
	Adults:      2,
	StartDate:   "2020-01-01",
	StarsFrom:   1,
	StarsTo:     5,
	FlexDates:   false,
}
prices, err := lt.GetStatisticsPrices(filter)

```

## Testing
Unit-tests:


```
go test -v
```
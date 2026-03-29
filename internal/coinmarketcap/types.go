package coinmarketcap

import "time"

// CryptoCurrencyCoin represents a single coin from CoinMarketCap.
type CryptoCurrencyCoin struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Symbol    string    `json:"symbol"`
	DateAdded time.Time `json:"date_added"`
	Quote     Quote     `json:"quote"`
}

// Quote holds pricing info in different currencies.
type Quote struct {
	USD USDQuote `json:"USD"`
}

// USDQuote contains the USD price and last updated timestamp.
type USDQuote struct {
	Price       float64   `json:"price"`
	LastUpdated time.Time `json:"last_updated"`
}

// QuotesResponse is the JSON returned by the /quotes/latest endpoint.
type QuotesResponse struct {
	Data map[string]CryptoCurrencyCoin `json:"data"`
}

// ListingsResponse is the JSON returned by the /listings/latest endpoint.
type ListingsResponse struct {
	Data []CryptoCurrencyCoin `json:"data"`
}

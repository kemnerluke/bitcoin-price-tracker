package coinmarketcap

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

const baseURL = "https://pro-api.coinmarketcap.com/v1/cryptocurrency"

type Client struct {
	httpClient *http.Client
	apiKey     string
}

func NewClient() (*Client, error) {
	apiKey := os.Getenv("CMC_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("environment variable CMC_API_KEY is not set")
	}

	return &Client{
		httpClient: &http.Client{},
		apiKey:     apiKey,
	}, nil
}

func (c *Client) GetQuote(symbol string) (CryptoCurrencyCoin, error) {
	endpoint := fmt.Sprintf("%s/quotes/latest", baseURL)

	params := url.Values{}
	params.Add("symbol", symbol)

	req, err := http.NewRequest("GET", endpoint+"?"+params.Encode(), nil)
	if err != nil {
		return CryptoCurrencyCoin{}, err
	}

	req.Header.Set("Accepts", "application/json")
	req.Header.Set("X-CMC_PRO_API_KEY", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return CryptoCurrencyCoin{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return CryptoCurrencyCoin{}, err
	}

	var result QuotesResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return CryptoCurrencyCoin{}, err
	}

	coin, ok := result.Data[symbol]
	if !ok {
		return CryptoCurrencyCoin{}, fmt.Errorf("coin %s not found", symbol)
	}

	return coin, nil
}

func (c *Client) GetTopListings(limit int) ([]CryptoCurrencyCoin, error) {
	endpoint := fmt.Sprintf("%s/listings/latest", baseURL)

	params := url.Values{}
	params.Add("limit", fmt.Sprintf("%d", limit))

	req, err := http.NewRequest("GET", endpoint+"?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accepts", "application/json")
	req.Header.Set("X-CMC_PRO_API_KEY", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result ListingsResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

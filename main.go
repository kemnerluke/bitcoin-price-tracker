package main

import (
	"fmt"
	"log"

	"bitcoin_price_checker/internal/coinmarketcap"
)

func main() {
	// Initialize the client (reads API key from env)
	client, err := coinmarketcap.NewClient()
	if err != nil {
		log.Fatalf("Cannot create CoinMarketCap client: %v", err)
	}

	// 1️⃣ Get quotes for specific coins
	symbols := []string{"BTC", "ETH", "XRP"}

	fmt.Println("==== Individual Coin Prices ====")
	for _, symbol := range symbols {
		coin, err := client.GetQuote(symbol)
		if err != nil {
			log.Printf("Failed to fetch %s: %v\n", symbol, err)
			continue
		}
		fmt.Printf("%s (%s): $%.2f | ID: %d\n", coin.Name, coin.Symbol, coin.Quote.USD.Price, coin.ID)
	}

	// 2️⃣ Get top 10 coins by market cap
	topCoins, err := client.GetTopListings(10)
	if err != nil {
		log.Fatalf("Failed to fetch top listings: %v\n", err)
	}

	fmt.Println("\n=== Top 10 Coins ===")
	for i, coin := range topCoins {
		fmt.Printf("%d. %s (%s): $%.2f\n", i+1, coin.Name, coin.Symbol, coin.Quote.USD.Price)
	}
}

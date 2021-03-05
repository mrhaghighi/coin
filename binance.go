package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Colors
const (
	BlackColor   = "\033[1;30m%s\033[0m"
	RedColor     = "\033[1;31m%s\033[0m"
	GreenColor   = "\033[1;32m%s\033[0m"
	YellowColor  = "\033[1;33m%s\033[0m"
	PurpleColor  = "\033[1;34m%s\033[0m"
	MagentaColor = "\033[1;35m%s\033[0m"
	TealColor    = "\033[1;36m%s\033[0m"
	WhiteColor   = "\033[1;37m%s\033[0m"
)

func getPrice() {
	// Init symbols
	symbols := []string{"BTCUSDT", "ETHUSDT", "LTCUSDT", "XRPUSDT", "TRXUSDT"}

	// Show messages
	fmt.Println("-------------------------------")
	fmt.Println("Fetching prices...")
	fmt.Println("-------------------------------")

	// Getting prices from API & printing them
	fmt.Println("Symbol      Price           Change(24h)")
	fmt.Println("----------  --------------  -----------")
	for i := range symbols {
		price, priceChangePercent := getPriceFromApi(symbols[i])
		printPrice(symbols[i], price, priceChangePercent)
	}

}

func getPriceFromApi(symbol string) (string, string) {
	// API URL
	url := "https://testnet.binance.vision/api/v3/ticker/24hr?symbol=" + symbol

	// Get price from API
	res, err := http.Get(url)
	if (err != nil) {
		log.Fatalln(err)
	}

	// Close connection
	defer res.Body.Close()

	// Mark up data
	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)

	// return result
	return result["weightedAvgPrice"].(string), result["priceChangePercent"].(string)
}

func printPrice(symbol, price, priceChangePercent string) {
	fmt.Printf("%-10v  %-14v  %v\n", symbol, price, priceChangePercent)
}
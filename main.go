package main

import (
	// standard
	"fmt"

	// external
	coinApi "github.com/miguelmota/go-coinmarketcap"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	coinname = kingpin.Arg("coin", "Coin to lookup").String()
)

func init() {
	kingpin.Parse()
}

func main() {
	coin, err := coinApi.GetCoinData(*coinname)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s:\n", coin.Name)
	fmt.Printf("USD  %.2f\n", coin.PriceUsd)
	fmt.Printf("Market Cap %.2f\n", coin.MarketCapUsd)
	fmt.Printf("Available Supply %.0f\n", coin.AvailableSupply)
	fmt.Printf("Total Supply %.0f\n", coin.TotalSupply)
	fmt.Printf("Symbol: %s\n", coin.Symbol)
	fmt.Printf("Rank: %d\n", coin.Rank)
}

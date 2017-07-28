package main

import (
	"github.com/toorop/go-bittrex"
	"fmt"
	"log"
	"os"
	"strings"
)

var exchange string
var coinFrom string
var coinTo string
var market string

func init() {
	exchange := os.Getenv("EXCHANGE")
	coinFrom := strings.ToUpper(os.Getenv("COIN_FROM"))
	coinTo := strings.ToUpper(os.Getenv("COIN_TO"))
	market := strings.Join([]string{coinFrom, coinTo}, "-")
	if exchange == "" || coinFrom == "" || coinTo == "" || market == "" {
		log.Fatal("[FATAL] Please set Rex properly with the run script, thanks!")
	}
}

func main() {
	// gather account and market data
	balance := getBalance()
	marketData := getMarketTicker()

	fmt.Printf("[INFO] Market Data fetched for  %s. Beginning trade...", market)
	trade(balance, marketData);
}

func trade(b float64, d bittrex.Ticker) {
	var initialAsk = d.Ask 
	purchasePrice := b - (b * 0.125)
	fmt.Printf("[trade] %d (init ask) %d (purchasePrice)", initialAsk, purchasePrice)
	return
}

func Client() *bittrex.Bittrex {
	if exchange == "bittrex" {
		k := os.Getenv("BITTREX_KEY")
		s := os.Getenv("BITTREX_SECRET")
		if k == "" || s == "" {
			log.Fatal("[FATAL] Environment not properly configured.")
		}

		return bittrex.New(k, s)
	}
	
	log.Fatal("[FATAL] exchange %s unknown...", exchange)
	return nil
}

func getBalance() float64 {
	c := Client()
	if exchange == "bittrex" {
		coinInfo, err := c.GetBalance(coinFrom)
		if err != nil {
			log.Fatalf("[FATAL] failed to get bittrex balance for %s", coinFrom)
		}
		
		if coinInfo.Balance <= 0 {
			log.Fatal("[FATAL] insufficient funds, pal")
		} else {
			return coinInfo.Balance
		}
	}

	return 0
}

func getMarketTicker() bittrex.Ticker {
	c := Client() 
	if exchange == "bittrex" {
		ticker, err := c.GetTicker(market)
		if err != nil {
			log.Fatalf("[FATAL] could not get ticker from ")
		}

		return ticker
	}

	return bittrex.Ticker{}
}  

func sell() {
	// stop sell everything @ 97.5% of INITIAL_ASK
	// set "TRADE_ACTIVE"
	// while TRADE_ACTIVE
	// check ASK repeatedly for changes.
	// check open orders for changes
	// if open orders = 0, exit trade get account && record earnings.
	// if LATEST_ASK / INITIAL_ASK < 0.975, cancel all open orders for that coin and attempt to set 0.25 lower again sell stop.
	// if LATEST_ASL / INITIAL_ASK > 1.00, cancel all orders and set sale to 0.25% less than LATEST_ASK
}

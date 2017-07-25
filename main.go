package main

import (
	"strconv"
	"strings"
	"os"
	"log"
	"github.com/toorop/go-bittrex"
)

func main() {
	key := os.Getenv("KEY")
	secret := os.Getenv("SECRET")
	coinFrom := strings.ToUpper(os.Getenv("COIN_FROM"))
	coinTo := strings.ToUpper(os.Getenv("COIN_TO"))
	if key == ""  || secret == "" {
		log.Fatal("[FATAL] Environment not properly configured.")
	}
	if coinFrom == "" || coinTo == "" {
		log.Fatal("[FATAL] Please use the run script to use rex.")
	}
	
	bittrex := bittrex.New(key, secret)

	// get balance and establish recommended purchasing price
	balance, err := bittrex.GetBalances()
	if err != nil {
		log.Fatalf("[FATAL] could not get balance for '%s', error: %v", coinFrom, err)
	}
	fetchedBalance := strconv.ParseFloat(balance.result.Available)
	purchasePrice := fetchedBalance - (fetchedBalance * 0.125)

	// get product ticker information for market
	market := strings.Join([]string{coinFrom, coinTo}, "-")	
	ticker, err := bittrex.GetTicker(market)
	if err != nil {
		log.Fatalf("[FATAL] could not get ticker for %s. Aborting trade.", market)
	}

	askHistory := []float
	askHistory = append(askHistory, ticker.result.Ask)
	fmt.Printf("[INFO] intiial ask price of %d recieved for %s. Beginning trade...", initialAsk, market)

	// attempt purchase
	_, err := bittrex.BuyLimit(market, purchasePrice, initialAsk)
	if err != nil {
		log.Fatalf("[FATAL] could not purchase at %d, ", initialAsk)
		// TODO: attempt purchases and then go into sales loop
	} else {
		fmt.Print("[SUCCESS] initial purchase successful, watching for best sell opporunities")
		var tradeActive = true
		var stop = askHistory[0] * 0.025
		for tradeActive {
			ticker, err := bittrex.GetTicker(market)
			if err != nil {
				log.Fatalf("[FATAL] could not get ticker for %s. Aborting trade.", market)
			}

			ask := ticker.result.Ask

			if ask <= stop {
				// attempt sale immediately.
				var checkingSale = true
				for checkingSale {
					// get sale
					// if sale, check ticker
					// if ticker ask <= stop, set new stop and attempt sale.
					// repeat
				}
			}
			
			if ask > askHistory[askHistory.length] {
				askHistory = append(askHistory, ask)
				stop = askHistory[askHistory.length] * 0.025
			}			
		}
	}
}


func sell(b) {
	var tradeActive = true
	for tradeActive {
		
	}
	// stop sell everything @ 97.5% of INITIAL_ASK
	// check ASK repeatedly for changes.
	// check open orders for changes
	// if open orders = 0, exit trade get account && record earnings.
	// if LATEST_ASK / INITIAL_ASK < 0.975, cancel all open orders for that coin and attempt to set 0.25 lower again sell stop.
	// if LATEST_ASL / INITIAL_ASK > 1.00, cancel all orders and set sale to 0.25% less than LATEST_ASK
}

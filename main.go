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
	
	balance, err := bittrex.GetBalances()
	if err != nil {
		log.Fatalf("[FATAL] could not get balance for '%s', error: %v", coinFrom, err)
	}

	market := strings.Join([]string{coinFrom, coinTo}, "-")
	
	ticker, err := bittrex.GetTicker(market)
	if err != nil {
		log.Fatalf("[FATAL] could not get ticker for")
	}

	initialAsk := ticker.result.Ask
	fmt.Printf("[INFO] intiial ask price of %d recieved for %s. Beginning trade...", initialAsk, market)

	fetchedBalance := strconv.ParseFloat(balance.result.Available)
	purchaseBalance := fetchedBalance - (fetchedBalance * 0.125)

	initialPurchaseUUID, err := bittrex.BuyLimit(market, purchaseBalance, initialAsk)
	if err != nil {
		fmt.Printf("[WARNING] could not purchase at %d, ", initialAsk)
		// TODO: enter attempt purchase loop for next initialAsk
		// set limit to at where you'll buy it based on a 0.5% increase (2 attempts ideally
		// success, err := attemptPurchase()
		// check error, abort if there is one
		// if success, enter sell function
	} else {
		fmt.Print("[INFO] initial purchase successful, watching for best sell opporunities")
		// set trade active
		// TODO: enter sell function
	}
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

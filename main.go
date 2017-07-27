package main

import (
	"github.com/toorop/go-bittrex"
	"log"
	"os"
	"strconv"
	"strings"
)

func Client(e string) {
	if e == "gdax" {

	}

	if e == "bittrex" {
		k := os.Getenv("BITTREX_KEY")
		s := os.Getenv("BITTREX_SECRET")
		if k == "" || s == "" {
			log.Fatal("[FATAL] Environment not properly configured.")
		}

		return bittrex.New(k, s)
	}

	log.Fatal("[FATAL] exchange %s unknown...", e)
}

func getBalance(e, coinFrom) number {
	c := Client(e)
	if e == "bittrex" {
		balance, err := bittrex.GetBalance(coinFrom)
		if err != nil {
			log.Fatalf("[FATAL] failed to get bittrex balance for %s", coinFrom)
		}
	}

	if e == "gdax" {
		accounts, err := client.GetAccounts()
		if err != nil {
			log.Fatal("[FATAL] failed to get gdax accounts")
		}

		for _, a range accounts {
			if a.currency == coinFrom {
				if a.balance <= 0 {
					log.Fatal("[FATAL] insufficient funds, pal.")
				} else {
					return 
				}
			}
		}

		log.Fatalf("[FATAL] could not find account corresponding to %s.", coinFrom)
	}
}

func main() {
	exchange := os.Getenv("EXCHANGE")
	
	coinFrom := strings.ToUpper(os.Getenv("COIN_FROM"))
	coinTo := strings.ToUpper(os.Getenv("COIN_TO"))
	if coinFrom == "" || coinTo == "" {
		log.Fatal("[FATAL] Please use the run script to use rex.")
	}

	balance := getBalance(e, coinFrom)

	market := strings.Join([]string{coinFrom, coinTo}, "-")

	// todo refactor everything below to support gdax/bittrex
	return
	
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

package main

import (
	"os"
	"fmt"
	"log"
	"github.com/toorop/go-bittrex"
)

func main() {
	key := os.Getenv("KEY")
	secret := os.Getenv("SECRET")
	coinFrom := os.Getenv("COIN_FROM")
	coinTo := os.Getenv("COIN_TO")
	if key == ""  || secret == "" {
		log.Fatal("[FATAL] Environment not properly configured.")
	}

	if coinFrom == "" || coinTo == "" {
		log.Fatal("[FATAL] Please use the run script to use rex.")
	}
	
	bittrex := bittrex.New(key, secret)
	
	balance, err := bittrex.GetBalances()
	if err != nil {
		log.Fatalf("[FATAL] could not get balance for '%s', error: %d", coinFrom, err)
	}

	
	
	// get INITIAL_ASK from COIN_TO
	// purchase 1/8th of COIN_FORM_BALANCE in COIN_TO = ENTRY
	// stop sell everything @ 97.5% of INITIAL_ASK
	// set "TRADE_ACTIVE"
	// while TRADE_ACTIVE
	// check ASK repeatedly for changes.
	// check open orders for changes
	// if open orders = 0, exit trade get account && record earnings.
	// if LATEST_ASK / INITIAL_ASK < 0.975, cancel all open orders for that coin and attempt to set 0.25 lower again sell stop.
	// if LATEST_ASL / INITIAL_ASK > 1.00, cancel all orders and set sale to 0.25% less than LATEST_ASK
}

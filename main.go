package main

import (
	"github.com/toorop/go-bittrex"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	exchange
	coinFrom
	coinTo
	market
)

func init() {
	exchange := os.Getenv("EXCHANGE")
	coinFrom := strings.ToUpper(os.Getenv("COIN_FROM"))
	coinTo := strings.toUpper(os.Getenv("COIN_TO"))
	market := strings.Join([]string{coinFrom, coinTo}, "-")
	if exchange == "" || coinFrom == "" || coinTo == "" || market == "" {
		log.Fatal("[FATAL] Please set Rex properly with the run script, thanks!")
	}
}

func main() {
	// gather account and market data
	balance := getBalance(e, coinFrom)
	market := strings.Join([]string{coinFrom, coinTo}, "-")
	marketData := getMarketTicker(market)

	fmt.Printf("[INFO] intiial ask price of %d recieved for %s. Beginning trade...", marketData.Ask, market)
	trade(balance, marketData);
}

func trade(b number, d bittrex.Ticker) {
	var initialAsk = d.Ask 
	fetchedBalance := strconv.ParseFloat(b.Available)
	purchaseBalance := fetchedBalance - (fetchedBalance * 0.125)
	c := Client(exchange)

	return
	// todo enable trade and conditional selling
	
	initialPurchaseUUID, err := c.BuyLimit(market, purchaseBalance, initialAsk)
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


func Client() {
	if exchange == "bittrex" {
		k := os.Getenv("BITTREX_KEY")
		s := os.Getenv("BITTREX_SECRET")
		if k == "" || s == "" {
			log.Fatal("[FATAL] Environment not properly configured.")
		}

		return bittrex.New(k, s)
	}
	
	log.Fatal("[FATAL] exchange %s unknown...", e)
}

func getBalance() number {
	c := Client(exchange)
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
}

func getMarketTicker() bittrex.Ticker {
	c := Client(exchange) 
	if exchange === "bittrex" {
		ticker, err := c.GetTicker(market)
		if err != nil {
			log.Fatalf("[FATAL] could not get ticker from ")
		}

		return ticker
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

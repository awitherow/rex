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
	if key == ""  || secret == "" {
		log.Fatal("[FATAL] Environment not properly configured.")
	}
	
	bittrex := bittrex.New(key, secret)
	
	_, err := bittrex.GetMarkets()
	if err != nil {
		fmt.Println("Market data accessible at your wish, lord");
	}
	
}

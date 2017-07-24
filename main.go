package main

import (
	"os"
	"fmt"
	"log"
	"github.com/toorop/go-bittrex"
	"github.com/urfave/cli"
)

func main() {
	key := os.Getenv("KEY")
	secret := os.Getenv("SECRET")
	if key == ""  || secret == "" {
		log.Fatal("[FATAL] Environment not properly configured.")
	}
	
	bittrex := bittrex.New(key, secret)
	
	markets, err := bittrex.GetMarkets()
	
	app := cli.NewApp();
	app.Name = "rex"
	app.Usage = "do not fuck up while trading bitcoin"
	app.Action = func(c *cli.Context) error {
		fmt.Println("boom!")
		return nil
	}

	app.Run(os.Args)
}

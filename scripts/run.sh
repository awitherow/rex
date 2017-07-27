#!/bin/bash

parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )
cd "$parent_path"

source ../.env

if [ "$REX_USER_NAME" == "" ]; then
    echo "Welcome!"
    echo "What is your name? "
    read REX_USER_NAME
    echo "REX_USER_NAME=$REX_USER_NAME" >> ../.env
fi

echo "Welcome, $REX_USER_NAME!"

if [ "$DEFAULT_EXCHANGE" == "" ]; then  
    echo "Which exchange would you like to use? options: [gdax] [bittrex]"

    read EXCHANGE

    if [ "$EXCHANGE" == "gdax" ]; then
	echo "KEY: "
	read KEY
	echo "GDAX_KEY=$KEY" >> ../.env

	echo "SECRET: "
	read SECRET
	echo "GDAX_SECRET=$SECRET" >> ../.env

	echo "PASSPHRASE: "
	read PASSPHRASE
	echo "GDAX_PASSPHRASE=$PASSPHRASE" >> ../.env
    elif [ "$EXCHANGE" == "bittrex" ]; then
	echo "KEY: "
	read KEY
	echo "BITTREX_KEY=$KEY" >> ../.env

	echo "SECRET: "
	read SECRET
	echo "BITTREX_SECRET=$SECRET" >> ../.env

    else
	echo "Exchange $EXCHANGE is not yet supported, maybe you can make that happen with a new PR? exiting."
	exit 0
    fi

    echo "Would you like this to be your default exchange? [yes] [no]"

    read DEFAULT_EXCHANGE

    if [ "$DEFAULT_EXCHANGE" == "yes" ]; then    
	echo "DEFAULT_EXCHANGE=$EXCHANGE" >> ../.env
    elif [ "$DEFAULT_ EXCHANGE" == "no" ]; then
	echo "Alright! If you'd like to set it in the future, manually add 'DEFAULT_EXCHANGE' to your .env."
    else
	echo "I was expecting 'yes' or 'no', not $DEFAULT_EXCHANGE... exiting."
	exit 0
    fi  
fi

echo "Which coin are you trading from? "

read COIN_FROM

export "$COIN_FROM"

echo "Which coin do you wish to trade to? "

read COIN_TO

export "$COIN_TO" 

# rex

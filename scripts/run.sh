#!/bin/bash

parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )
cd "$parent_path"

source ../.env

if [ "$REX_USER_NAME" == "" ]
then
    echo "Welcome!"
    echo "What is your name? "
    read REX_USER_NAME
    echo "REX_USER_NAME=$REX_USER_NAME" >> ../.env
fi

echo "Welcome, $REX_USER_NAME!"

echo "Which coin are you trading from? "

read COIN_FORM

export "$COIN_FROM"

echo "Which coin do you wish to trade to? "

read COIN_TO

export "$COIN_TO" 

rex

package main

import (
	//"net/url"
	"net/http"
	"log"
	"os"
	"encoding/json"
	"strconv"
)

const (
	BASE_URL = "https://us.api.battle.net/wow"
	AUCTION_URL = "/auction/"
	ITEM_URL = "/item/"
	EN_US_LOCALE = "en_US"
	DATA = "/data/"
)

// Calls the Blizzard community API to get periodic data dumps.

func BuildItemQueryString(locale string, apiKey string, id int) string {
	url, err := http.NewRequest("GET", BASE_URL + ITEM_URL + strconv.Itoa(id), nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	query := url.URL.Query()
	query.Add("apikey",apiKey)

	query.Add("locale",locale)

	url.URL.RawQuery = query.Encode()

	//println(url.URL.String())

	return url.URL.String()
}

func BuildAuctionLocationQueryString(locale string, apiKey string, realm string) string {
	url, err := http.NewRequest("GET", BASE_URL + AUCTION_URL + DATA + realm, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	query := url.URL.Query()
	query.Add("apikey",apiKey)

	query.Add("locale",locale)

	url.URL.RawQuery = query.Encode()

	println(url.URL.String())

	return url.URL.String()
}

func GetItemRequest(request string, item *Item) {
	response,err := http.Get(request)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(item)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
}

func GetAuctionLocationRequest(request string, auction *AuctionLocation) {
	response,err := http.Get(request)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(auction)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
}

func GetAuctionRequest(request string, auction *Auction) {
	response,err := http.Get(request)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(auction)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
}

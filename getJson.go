package main

import (
	"net/http"

	"encoding/json"
	"strconv"
	"github.com/kr/pretty"
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
	checkErr(err)

	query := url.URL.Query()
	query.Add("apikey",apiKey)

	query.Add("locale",locale)

	url.URL.RawQuery = query.Encode()

	//println(url.URL.String())

	return url.URL.String()
}

func BuildAuctionLocationQueryString(locale string, apiKey string, realm string) string {
	url, err := http.NewRequest("GET", BASE_URL + AUCTION_URL + DATA + realm, nil)
	checkErr(err)

	query := url.URL.Query()
	query.Add("apikey",apiKey)

	query.Add("locale",locale)

	url.URL.RawQuery = query.Encode()

	pretty.Println(url.URL.String())

	return url.URL.String()
}

func GetItemRequest(request string, item *Item) {
	//pretty.Printf("GET %v \n", request)
	response,err := http.Get(request)
	checkErr(err)

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(item)
	checkErr(err)
}

func GetAuctionLocationRequest(request string, auction *AuctionLocation) {
	response,err := http.Get(request)
	checkErr(err)

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(auction)
	checkErr(err)
}

func GetAuctionRequest(request string, auction *Auction) {
	response,err := http.Get(request)
	checkErr(err)

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(auction)
	checkErr(err)
}

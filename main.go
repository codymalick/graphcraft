package main

import (
	"flag"
	//"fmt"
	"github.com/kr/pretty"

	"strconv"
)
const (
	BASE_URL = "https://us.api.battle.net/wow"
	AUCTION_URL = "/auction/"
	ITEM_URL = "/item/"
	EN_US_LOCALE = "en_US"
	DATA = "/data/"
)

func main() {

	// Read api key from cli
	apiKey := flag.String("a", "", "-a <key>")
	itemId := flag.String("i", "", "-i <item id>")
	realm := flag.String("r","","-r <realm_name>")

	flag.Parse()

	if *apiKey == "" {
		pretty.Print("Please provide an apikey using -a")
	}

	pretty.Printf("api key: %v\n", *apiKey)
	pretty.Printf("item id: %v\n", *itemId)
	pretty.Printf("realm: %v\n", *realm)
	//if itemId != nil {
	//	_ = GetItemById(*apiKey, *itemId)
	//}

	if *realm != "" && *realm != "" {
		auc := FetchLatestAuctionData(*apiKey, *realm)

		// NOTE: API limit is 100/second
		items := make([]Item, 10)
		for i := 0; i < 5; i++ {
			items = append(items, *(GetItemById(*apiKey, strconv.Itoa(auc.Auctions[i].Item))))
		}

		pretty.Print(items)
	}



	




}

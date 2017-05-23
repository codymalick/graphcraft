package main

import (
	"flag"
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
	apiKey := flag.String("a", "00000000", "-a <key>")

	flag.Parse()

	// Call the item API

	requestUrl := BuildItemQueryString(EN_US_LOCALE, *apiKey, "18802")

	item := new(Item)
	GetItemRequest(requestUrl, item)

	println(len(item.ItemSpells))

	auctionUrl := BuildAuctionLocationQueryString(EN_US_LOCALE, *apiKey, "Emerald Dream")

	locationUrl := new(AuctionLocation)

	GetAuctionLocationRequest(auctionUrl, locationUrl)

	println(locationUrl.Files[0].URL)

	auction := new(Auction)

	GetAuctionRequest(locationUrl.Files[0].URL, auction)

	for index,auc := range auction.Auctions {
		println(index,auc.Item)
	}


}

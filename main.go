package main

import (
	"flag"
	"github.com/kr/pretty"
	"strconv"
)

const (
	BASE_URL = "https://us.api.battle.net/wow"
	AUCTION_URL = "/auction/"
	ITEM_URL = "/item/"
	EN_US_LOCALE = "en_US"
	DATA = "/data/"

	DB_ADDRESS = "localhost"
	DB_PORT = "3306"
	DB_PROTOCOL = "tcp"
)

func main() {

	// Read api key from cli
	apiKey := flag.String("a", "", "-a <key>")
	itemId := flag.String("i", "", "-i <item id>")
	realm := flag.String("r","","-r <realm_name>")
	user := flag.String("u","","-u <db_username>")
	password := flag.String("-p","","-p <db_password")

	flag.Parse()

	if *apiKey == "" {
		pretty.Print("Please provide an apikey using -a")
	}



	pretty.Printf("api key: %v\n", *apiKey)
	pretty.Printf("item id: %v\n", *itemId)
	pretty.Printf("realm: %v\n", *realm)

	if itemId != nil {
		_ = GetItemById(*apiKey, *itemId)
	}

	if *realm != "" {
		auc := FetchLatestAuctionData(*apiKey, *realm)
		popular := MostPopularAuctions(auc)

		// NOTE: API limit is 100/second
		i := 0

		for _,v := range popular {
			pretty.Printf("id:%v count:%v\n", v.id,v.count)
			i++
			item := GetItemById(*apiKey, strconv.Itoa(v.id))
			pretty.Printf("%v\n",item.Name)

			if i >= 20 {
				break
			}
		}

	}



	




}

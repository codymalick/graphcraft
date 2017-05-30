package main

import (
	"flag"
	"github.com/kr/pretty"
	"log"
)

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {

	// Read api key from cli
	apiKey := flag.String("a", "", "-a <key>")
	itemId := flag.Int("i", 0, "-i <item id>")
	realm := flag.String("r","","-r <realm_name>")
	user := flag.String("u","","-u <db_username>")
	password := flag.String("p","","-p <db_password>")
	debug := flag.Bool("d", false, "-d")

	flag.Parse()

	if *apiKey == "" {
		pretty.Print("Please provide an apikey using -a")
	}

	if *debug {

		pretty.Printf("api key: %v\n", *apiKey)
		pretty.Printf("item id: %v\n", *itemId)
		pretty.Printf("realm: %v\n", *realm)
		pretty.Printf("db_user: %v\n", *user)
	}

	// Connect to db so we don't have to hand db off everywhere
	InitDb(*user, *password)


	// Get single item data
	if *itemId != 0 {
		item := GetItemById(*apiKey,*itemId)
		pretty.Println(item)
	}

	// Get realm data
	if *realm != "" {
		auc := FetchLatestAuctionData(*apiKey, *realm)

		StoreAuctionData(auc, *realm, *apiKey)
		// cache the data
		//
		//popular := MostPopularAuctions(auc)
		//
		//// NOTE: API limit is 100/second
		//i := 0
		//
		//for _,v := range popular {
		//	pretty.Printf("id:%v count:%v\n", v.id , v.count)
		//	i++
		//	item := GetItemByIdApi(*apiKey, v.id)
		//	pretty.Printf("%v\n",item.Name)
		//
		//	if i >= 20 {
		//		break
		//	}
		//}

	}



	




}

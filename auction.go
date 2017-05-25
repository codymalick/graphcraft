package main

import "github.com/kr/pretty"

type AuctionLocation struct {
	Files []struct {
		URL string `json:"url"`
		LastModified int64 `json:"lastModified"`
	} `json:"files"`
}

// Main Auction House data structure. This is the full form of items return from the Blizzard community API
type Auction struct {
	Realms []struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"realms"`
	Auctions []struct {
		Auc int `json:"auc"`
		Item int `json:"item"`
		Owner string `json:"owner"`
		OwnerRealm string `json:"ownerRealm"`
		Bid int `json:"bid"`
		Buyout int `json:"buyout"`
		Quantity int `json:"quantity"`
		TimeLeft string `json:"timeLeft"`
		Rand int `json:"rand"`
		Seed int `json:"seed"`
		Context int `json:"context"`
		Modifiers []struct {
			Type int `json:"type"`
			Value int `json:"value"`
		} `json:"modifiers,omitempty"`
		PetSpeciesID int `json:"petSpeciesId,omitempty"`
		PetBreedID int `json:"petBreedId,omitempty"`
		PetLevel int `json:"petLevel,omitempty"`
		PetQualityID int `json:"petQualityId,omitempty"`
		BonusLists []struct {
			BonusListID int `json:"bonusListId"`
		} `json:"bonusLists,omitempty"`
	} `json:"auctions"`
}

func FetchLatestAuctionData(apiKey string, realm string) *Auction {

	auctionUrl := BuildAuctionLocationQueryString(EN_US_LOCALE, apiKey, realm)

	locationUrl := new(AuctionLocation)

	GetAuctionLocationRequest(auctionUrl, locationUrl)

	println(locationUrl.Files[0].URL)

	auction := new(Auction)

	GetAuctionRequest(locationUrl.Files[0].URL, auction)

	//pretty.Print(auction)
	pretty.Printf("\nGot latest auction data from %v\n", locationUrl.Files[0].LastModified)
	return auction
}



package main

import (
	"github.com/kr/pretty"
	"sort"
)

type AuctionLocation struct {
	Files []struct {
		URL string `json:"url"`
		LastModified int64 `json:"lastModified"`
	} `json:"files"`
}

type Listing struct {
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
}

// Main Auction House data structure. This is the full form of items return from the Blizzard community API
type Auction struct {
	Realms []struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"realms"`
	Listings []Listing `json:"auctions"`
}

type PopularPair struct {
	id int
	count int
}

type PopularPairList []PopularPair

func (p PopularPairList) Len() int { return len(p) }
func (p PopularPairList) Less(i, j int) bool { return p[i].count > p[j].count }
func (p PopularPairList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }

func FetchLatestAuctionData(apiKey string, realm string) *Auction {
	auctionUrl := BuildAuctionLocationQueryString(EN_US_LOCALE, apiKey, realm)
	locationUrl := new(AuctionLocation)
	GetAuctionLocationRequest(auctionUrl, locationUrl)
	println(locationUrl.Files[0].URL)
	auction := new(Auction)
	GetAuctionRequest(locationUrl.Files[0].URL, auction)

	pretty.Printf("\nGot latest auction data from %v\n", locationUrl.Files[0].LastModified)

	return auction
}

func CreatePairs(mapped map[int]int) PopularPairList {
	popular := make(PopularPairList, len(mapped))

	for k,v := range mapped {
		popular = append(popular, PopularPair{k,v})
	}

	sort.Sort(popular)
	return popular
}

// Prints out top 20 most popular auctions to the shell
func MostPopularAuctions(auction *Auction) PopularPairList {
	// Assume half the auctions are unique for performance
	listings := make(map[int]int, len(auction.Listings)/2)

	// group number of instances by id
	for _,item := range auction.Listings {
		listings[item.Item]++
	}

	popular := CreatePairs(listings)
	return popular
}





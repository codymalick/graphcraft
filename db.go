package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"github.com/kr/pretty"
	"fmt"
)

const (
	DB_ADDRESS = ""
	DB_PORT = "3306"
	DB_PROTOCOL = "tcp"
	ITEM_INSERT = "INSERT INTO items"
	ITEM_SELECT = "SELECT * FROM items"

	LISTING_INSERT = "INSERT INTO listings"
	LISTING_SELECT = "SELECT * FROM listings"

	USER_INSERT = "INSERT INTO users"
	USER_SELECT = "SELECT * FROM users"

	TIMESTAMP_INSERT = "INSERT INTO timestamps"
	TIMESTAMP_SELECT = "SELECT * FROM timestamps"
	TIMESTAMP_LAST_ROW = "SELECT * FROM timestamps ORDER BY id DESC LIMIT 1"

)

var db *sql.DB

func InitDb(user string, password string) {
	var err error
	connString := user + ":" + password + "@" + DB_ADDRESS + "/graphcraft"
	pretty.Println(connString)
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	pretty.Printf("Connected to %v/graphcraft\n", DB_ADDRESS)
}

func InsertItem(item *Item) error {
	_, err := db.Exec(ITEM_INSERT +
		"(id, description, name, icon, stackable) VALUES (?,?,?,?,?)",
		item.ID,
		item.Description,
		item.Name,
		item.Icon,
		item.Stackable,
	)

	checkErr(err)

	// pretty.Printf("Inserted item %v\n", item.Name)
	return nil
}

func QueryItem(itemId int) *Item {
	var id int
	var description string
	var name string
	var icon string
	var stackable int

	err := db.QueryRow(ITEM_SELECT + " WHERE id=?",
		itemId,
	).Scan(&id,&description,&name,&icon,&stackable)

	if err != nil {
		if err == sql.ErrNoRows {
			//fmt.Printf("No item %v, found\n", itemId)
			return nil
		}
		checkErr(err)
	}

	var result Item
	result.ID = id
	result.Description = description
	result.Name = name
	result.Icon = icon
	result.Stackable = stackable

	//pretty.Printf("Selected item %v\n", result.Name)
	return &result
}

func InsertUser(username string, realm string) error {
	_, err := db.Exec(USER_INSERT +
		"(username, server) VALUES (?,?)",
		username,
		realm,
	)

	checkErr(err)

	//pretty.Printf("Inserted user %v,%v\n", username, realm)
	return nil
}

func QueryUserById(userId int, realm string) *User {
	rows, err := db.Query(USER_SELECT + " WHERE id=? and server=?",
		userId,
		realm,
	)

	checkErr(err)

	var result User

	for rows.Next() {
		var id int
		var username string
		var server string

		err = rows.Scan(&id, &username, &server)
		checkErr(err)

		if id != 0 {
			result.ID = id
			result.Username = username
			result.Server = server
		} else {
			return nil
		}
	}

	//pretty.Printf("Selected user %v,%v\n", result.Username, result.Server)
	return &result
}


func QueryUserByName(username string, realm string) *User {

	var id int
	var uname string
	var server string

	err := db.QueryRow(USER_SELECT + " WHERE username=? and server=?",
		username,
		realm,
	).Scan(&id,&uname,&server)

	if err != nil {
		if err == sql.ErrNoRows {
			//fmt.Printf("No user %v,%v found\n",username, realm)
			return nil
		}
		checkErr(err)
	}

	result := User{id,uname,server}

	//pretty.Printf("Selected user %v,%v\n", result.Username, result.Server)
	return &result
}

func InsertListing(listing *Listing) error {
	_, err := db.Exec(LISTING_INSERT +
		"(listing_id, item_id, user_id, b_id, buyout, quantity, timestamp_id) VALUES (?,?,?,?,?,?,?)",
		listing.Auc,
		listing.Item,
		listing.Owner_ID,
		listing.Bid,
		listing.Buyout,
		listing.Quantity,
		listing.Timestamp_ID,
	)

	checkErr(err)

	//pretty.Printf("Inserted listing %v,%v\n", listing.Auc, listing.Item)
	return nil
}

func InsertTimestamp(unix int64) error {
	_, err := db.Exec(TIMESTAMP_INSERT +
		"(timestamp) VALUES (?)",
		unix,
	)

	checkErr(err)

	//pretty.Printf("Inserted timestamp %v\n", unix)
	return nil
}

func QueryTimestampLatest() *Timestamp {
	var id int
	var unix int64

	err := db.QueryRow(TIMESTAMP_LAST_ROW).Scan(&id,&unix)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("No timestamp found\n")
			return nil
		}
		checkErr(err)
	}

	result := Timestamp{id,unix}

	//pretty.Printf("Selected latest timestamp %v\n", result.Unix)
	return &result
}



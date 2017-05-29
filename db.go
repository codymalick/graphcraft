package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"github.com/kr/pretty"
)

const (
	DB_ADDRESS = ""
	DB_PORT = "3306"
	DB_PROTOCOL = "tcp"
	ITEM_INSERT = "INSERT INTO items"
	ITEM_SELECT = "SELECT * FROM items"
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

	pretty.Printf("Inserted %v into db\n", item.Name)
	return nil
}

func QueryItem(itemId int) *Item {
	rows, err := db.Query(ITEM_SELECT + " WHERE id=?",
		itemId,
	)

	checkErr(err)

	var result Item

	for rows.Next() {
		var id int
		var description string
		var name string
		var icon string
		var stackable int

		err = rows.Scan(&id, &description, &name, &icon, &stackable)
		checkErr(err)

		if id != 0 {
			result.ID = id
			result.Description = description
			result.Name = name
			result.Icon = icon
			result.Stackable = stackable
		} else {
			return nil
		}
	}

	pretty.Printf("Selected %v from db\n", result.Name)
	return &result
}



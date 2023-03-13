package pkg

import (
	"database/sql"
	"fmt"
)



func GiveRating(name string, newRating string, db *sql.DB) {
	rows, err := db.Query(`SELECT rating FROM lineitem where name = $1`, name)
	CheckError(err)
	defer rows.Close()

	for rows.Next() {
		var rating string
		err = rows.Scan(&rating)
		CheckError(err)

		if (rating == "0") {
			_, err = db.Exec(`update lineitem set rating = $1 where name = $2`, newRating, name)
			CheckError(err)
		} else {
			_, err = db.Exec(`update lineitem set rating = (rating+$1)/2 where name = $2`, newRating, name)
			CheckError(err)
		}
	}
}

func  SearchItemByName(db *sql.DB, find string) *Item {
	item := new(Item)

	rows, err := db.Query(`SELECT name, price, rating FROM lineitem where name = $1`, find)
	CheckError(err)
	defer rows.Close()
	
	for rows.Next() {
		err = rows.Scan(&item.Name, &item.Price, &item.Rating)
		CheckError(err)
	} 
	return item
}


func FilterItemsByRating(db *sql.DB, f string, s string) []Item{

	rows, err := db.Query(`SELECT name, price, rating FROM lineitem where rating between $1 and $2`, f, s)
	CheckError(err)
	defer rows.Close()

	items := []Item{}

	for rows.Next() {
		item := new(Item)
		err = rows.Scan(&item.Name, &item.Price, &item.Rating)
		CheckError(err)

		items = append(items, *item) 
	} 
	return items
}

func FilterItemsByPrice(db *sql.DB, f string, s string) []Item {
	
	rows, err := db.Query(`SELECT name, price, rating FROM lineitem where price between $1 and $2`, f, s)
	CheckError(err)
	defer rows.Close()

	items := []Item{}

	for rows.Next() {
		item := new(Item)
		err = rows.Scan(&item.Name, &item.Price, &item.Rating)
		CheckError(err)

		items = append(items, *item)
	}
	return items
}

func (s *Store) ShowAllItems(db *sql.DB) {
	rows, err := db.Query(`SELECT name, price, rating FROM lineitem `)
	CheckError(err)
	defer rows.Close()
	fmt.Printf("\n %-15s  %-15s %-15s \n", "name ", "price", "raiting")
	for rows.Next() {
		var name string
		var price, raiting float32

		err = rows.Scan(&name, &price, &raiting)
		CheckError(err)

		fmt.Printf(" %-15s  %-15f %-15f \n", name, price, raiting)
	}
}



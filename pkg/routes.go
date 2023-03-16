package pkg

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345"
	dbname   = "db_go"
)

func Routes() {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	rtr := mux.NewRouter()
	rtr.HandleFunc("/log/", Log_in).Methods("Post")
	rtr.HandleFunc("/reg/", Registration).Methods("Post")
	rtr.HandleFunc("/search/{item}/", searchItemByName).Methods("GET")
	rtr.HandleFunc("/giveRating/{item}/", giveRating).Methods("PUT")
	rtr.HandleFunc("/", home).Methods("GET")
	// rtr.HandleFunc("/filterItemsByPrice/", filterItemsByPrice).Methods("GET")
	// rtr.HandleFunc("/filterItemsByPrice/", filterItemsByPricePost).Methods("POST")
	// rtr.HandleFunc("/filterItemsByRating/", filterItemsByRatingPost).Methods("POST")
	// rtr.HandleFunc("/filterItemsByRating/", filterItemsByRating).Methods("GET")
	// rtr.HandleFunc("/searchItem/", searchItemByName).Methods("GET")
	// rtr.HandleFunc("/searchItem/", searchItemByNamePost).Methods("POST")
	// rtr.HandleFunc("/giveRating/", giveRating).Methods("GET")
	// rtr.HandleFunc("/giveRating/", giveRatingPost).Methods("POST")
	// rtr.HandleFunc("/log/", LoginP).Methods("Get")
	// rtr.HandleFunc("/shop/", Shop).Methods("Get")
	// rtr.HandleFunc("/", homePage).Methods("GET")

	http.Handle("/", rtr)
	err = http.ListenAndServe(":8080", rtr)
	CheckError(err)
}
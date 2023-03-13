package pkg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	// "strconv"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
	logged_in = false 
)




func giveRating(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	item := Item{}
	item.Name = params["item"]

	_ = json.NewDecoder(r.Body).Decode(&item)
	// json.NewEncoder(w).Encode(item)
	if (item.Name == "" || item.Rating == 0) {
		fmt.Fprintln(w, "Item name/Item rating cannot be empty")
	} else {
		item2 := SearchItemByName(db, item.Name)
		if item2.Name == "" {
			fmt.Fprintln(w, "This item does not exist")
		} else {
			GiveRating(item.Name, fmt.Sprintf("%f", item.Rating), db)
			fmt.Fprintln(w, "Rating added....")
			item2 = SearchItemByName(db, item.Name)
			fmt.Fprintf(w, "%s with new rating  %.2f", item.Name, item2.Rating)
		}
	}
}

func Registration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	reg := new(Register)
	reg.CreateUser()

	_ = json.NewDecoder(r.Body).Decode(reg.User)
	json.NewEncoder(w).Encode(reg.User)

	if (reg.User.Nickname =="" || reg.User.Password=="") {
		fmt.Fprintf(w, "Nickname/Password cannot be empty")
	} else {
		reg.InsertUserToDB(db)
		fmt.Fprintf(w, "Registration was successful")
	}
}


func Log_in(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	loginU := new(Login)
	loginU.User = new(User)

	_ = json.NewDecoder(r.Body).Decode(loginU.User)
	json.NewEncoder(w).Encode(loginU.User)
	
	if loginU.Validation(db) {
		logged_in = true
		fmt.Fprintf(w, "Logged in successfuly")
	} else {
		logged_in = false
		fmt.Fprintf(w, "Invalid nickname or password, try again or create new account")
	}
}



func searchItemByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	item := SearchItemByName(db, params["item"])
	if item.Name == "" {
		fmt.Fprintln(w, "Nothing found")
	} else {
		json.NewEncoder(w).Encode(item)
	}
}




func filterItemsByPricePost(w http.ResponseWriter, r *http.Request) { // nnot workiing
	f := r.FormValue("first")
	s := r.FormValue("second")
	items := FilterItemsByPrice(db, f, s)

	fmt.Fprintf(w, "%s %s %s\n", "Name", "Price", "Rating")

	for _, item := range items{
		fmt.Fprintf(w, "%s %.f %.f\n", item.Name, item.Price, item.Rating)
	}
	if len(items) == 0{
		fmt.Fprintln(w, "nothig found")
	}
}

func filterItemsByRatingPost(w http.ResponseWriter, r *http.Request) {  // nnot workiing
	f := r.FormValue("first")
	s := r.FormValue("second")

	items := FilterItemsByRating(db, f, s)
	fmt.Fprintf(w, "%s %s %s\n", "Name", "Price", "Rating")

	for _, item := range items{
		fmt.Fprintf(w, "%s %f %f\n", item.Name, item.Price, item.Rating)
	}

	if len(items) == 0{
		fmt.Fprintln(w, "nothig found")
		http.Redirect(w, r, "/", 302)
	}
}



func CheckError(err error) {
	if err != nil {
			panic(err)
	}
}


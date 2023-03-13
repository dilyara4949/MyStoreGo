package pkg

import (
	// "encoding/json"
)

type Item struct {
	Name string 
	Price float32 
	Rating float32 
}
type Login struct {
	User *User
}

type Register struct {
	User *User
}

type Store struct{}

type User struct {
	Nickname string 
	Password string	
}
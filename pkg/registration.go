package pkg

import (
	"database/sql"
	_ "github.com/lib/pq"
	// "net/http"
)


func (r *Register) CreateUser() {
	r.User = new(User)
}

func (r *Register) InsertUserToDB(db *sql.DB) {
	_, err := db.Exec(insertUserStat(), r.User.Nickname, r.User.Password)
  CheckError(err)
}

func insertUserStat() string{
	return `
INSERT INTO users (nickname, pswrd)
VALUES ($1, $2)`
}
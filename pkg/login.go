package pkg

import (
	"database/sql"
)




func (l *Login) Validation(db *sql.DB) bool {
	rows, err := db.Query(`SELECT nickname, pswrd FROM users where nickname = $1 and pswrd = $2`, l.User.Nickname, l.User.Password)
	CheckError(err)
	defer rows.Close()

	for rows.Next() {
		var nick string
		var psw string

		err = rows.Scan(&nick, &psw)
		CheckError(err)

		if nick != "" {
			return true
		}
	}
	CheckError(err)	
	return false
}
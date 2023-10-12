package handler

import (
	"database/sql"
	"pair-project/entity"

	"golang.org/x/crypto/bcrypt"
)

func Login(db *sql.DB, email, password string) (user entity.User, err error) {
	row := db.QueryRow("SELECT UserID, Email, Password, FirstName, LastName FROM Users WHERE Email = ?", email)
	err = row.Scan(&user.Id, &user.Email, &user.Password, &user.FirstName, &user.LastName)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return
	}

	return
}

package handler

import (
	"database/sql"
	"pair-project/entity"

	"golang.org/x/crypto/bcrypt"
)

func Register(db *sql.DB, user entity.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO Users (Email, Password, FirstName, LastName) VALUES (?,?,?,?)", user.Email, hashedPassword, user.FirstName, user.LastName)
	if err != nil {
		return err
	}

	return nil
}

package handler

import (
	"database/sql"
	"pair-project/entity"

	"golang.org/x/crypto/bcrypt"
)

// Login checks given username and password and verify whether those combinations are correct. The application will either give or reject access
func Login(db *sql.DB, email, password string) (user entity.User, err error) {
	row := db.QueryRow(login, email)
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

// Register will create new user in the database
func Register(db *sql.DB, user entity.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = db.Exec(register, user.Email, hashedPassword, user.FirstName, user.LastName)
	if err != nil {
		return err
	}

	return nil
}

// ReadUsers will return list of available users
func ReadUsers(db *sql.DB) ([]entity.User, error) {
	rows, err := db.Query(readUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []entity.User{}
	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

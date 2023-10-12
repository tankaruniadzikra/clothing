package handler

import (
	"database/sql"
	"pair-project/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func TestLogin(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/coba")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	email := "tan"
	password := "tan"

	expectedUser := entity.User{
		Id:        1,
		Email:     email,
		Password:  "$2a$10$WSOXuTabyEd4FYhV4LnevugipbDIMleqnsLS1KLbtmIkn42x97JZq",
		FirstName: "tan",
		LastName:  "karunia",
	}

	user, err := Login(db, email, password)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if user.Email != email {
		t.Errorf("Expected email %s, got %s", email, user.Email)
	}

	// cek apakah password telah divalidasi dengan benar.
	err = bcrypt.CompareHashAndPassword([]byte(expectedUser.Password), []byte(password))
	if err != nil {
		t.Errorf("Expected password validation to succeed, got error: %v", err)
	}
}

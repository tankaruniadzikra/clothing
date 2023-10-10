package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"pair-project/config"
	"pair-project/entity"
	"pair-project/handler"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Selamat datang di Cakra Store")
		fmt.Println("Menu:")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")
		fmt.Print("Masukkan pilihan (1/2/3): ")

		var option int
		scanner.Scan()
		_, err := fmt.Sscanf(scanner.Text(), "%d", &option)
		if err != nil {
			fmt.Println("Input anda bukan merupakan integer")
			continue
		}

		switch option {
		case 1:
			var email, password, firstName, lastName string
			fmt.Printf("\nREGISTER\n")
			fmt.Print("Masukkan email: ")
			fmt.Scanln(&email)
			fmt.Print("Masukkan password: ")
			fmt.Scanln(&password)
			fmt.Print("Masukkan nama depan: ")
			fmt.Scanln(&firstName)
			fmt.Print("Masukkan nama belakang: ")
			fmt.Scanln(&lastName)

			user := entity.User{
				Email:     email,
				Password:  password,
				FirstName: firstName,
				LastName:  lastName,
			}

			if err := handler.Register(user, db); err != nil {
				fmt.Println("Kesalahan saat registrasi", err)
			} else {
				fmt.Println("Registrasi berhasil!")
			}

		case 2:
			var email, password string
			fmt.Printf("\nLOGIN\n")
			fmt.Print("Email: ")
			fmt.Scanln(&email)
			fmt.Print("Password: ")
			fmt.Scanln(&password)

			user, authenticated, err := handler.Login(email, password, db)
			if err != nil {
				fmt.Println("Kesalahan saat login:", err)
			} else if authenticated {
				fmt.Printf("Login berhasil! Selamat datang, %s %s!\n", user.FirstName, user.LastName)
			} else {
				fmt.Println("Login gagal. Periksa kembali email dan password Anda.")
			}

		case 3:
			fmt.Println("Sampai jumpa!")
			os.Exit(0)

		default:
			fmt.Println("Input harus merupakan angka 1-3!")
		}
	}
}

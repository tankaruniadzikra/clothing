package cli

import (
	"database/sql"
	"fmt"
	"os"
	"pair-project/entity"
	"pair-project/handler"
	"text/tabwriter"

	_ "github.com/go-sql-driver/mysql"
)

// MainMenu will return integer value to dictate what the app will do
// 0 - back to MainMenu
// 1 - go to next menu
func MainMenu(db *sql.DB) int {
	fmt.Print(entity.MainMenu)

	var option int
	if err := inputScanner(&option, "Pilih menu yang ingin anda akses (1/2/3): "); err != nil {
		return 0
	}

	switch option {
	case 1:
		user, err := registerUser()
		if err != nil {
			fmt.Println("Terjadi kesalahan teknis. Silahkan hubungi administrator")
			fmt.Println("Kembali ke main menu...")
			fmt.Println()
			return 0
		}

		if err := handler.Register(db, user); err != nil {
			fmt.Println(err)
			fmt.Println("Terjadi kesalahan teknis. Silahkan hubungi administrator")
			fmt.Println("Kembali ke main menu...")
			fmt.Println()
			return 0
		}

		fmt.Printf("\nRegistrasi berhasil. Selamat datang di Cakra Clothing Store, %s %s!\n", user.FirstName, user.LastName)
		return 1
	case 2:
		var email, password string
		if err := inputScanner(&email, "Masukkan email: "); err != nil {
			fmt.Println("Terjadi kesalahan teknis. Silahkan hubungi administrator")
			fmt.Println("Kembali ke main menu...")
			fmt.Println()
			return 0
		}
		if err := inputScanner(&password, "Masukkan password: "); err != nil {
			fmt.Println("Terjadi kesalahan teknis. Silahkan hubungi administrator")
			fmt.Println("Kembali ke main menu...")
			fmt.Println()
			return 0
		}

		user, err := handler.Login(db, email, password)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Login gagal. Periksa kembali email dan password Anda.")
			fmt.Println("Kembali ke main menu...")
			fmt.Println()
			return 0
		}

		fmt.Printf("\nLogin berhasil. Selamat datang kembali, %s %s!\n", user.FirstName, user.LastName)
		return 1
	case 3:
		fmt.Println("Terima kasih telah mengunjungi Cakra Clothing Store. Sampai jumpa!")
		os.Exit(0)
		return 2
	default:
		fmt.Println("Maaf, menu yang anda pilih belum tersedia.")
		fmt.Println()
		return 0
	}
}

// registerUser will receive input from user to register their info, and construct them
// into User entity
func registerUser() (user entity.User, err error) {
	err = inputScanner(&user.Email, "Masukkan email: ")
	if err != nil {
		return
	}

	err = inputScanner(&user.Password, "Masukkan password: ")
	if err != nil {
		return
	}

	err = inputScanner(&user.FirstName, "Masukkan nama depan: ")
	if err != nil {
		return
	}

	err = inputScanner(&user.LastName, "Masukkan nama belakang: ")
	if err != nil {
		return
	}

	return
}

func listUserMenu(db *sql.DB) error {
	users, err := handler.ReadUsers(db)
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(w, "---\t | ------\t | -----------\t | --------------\t")
	fmt.Fprintln(w, "Id \t | Email \t | Nama Depan \t | Nama Belakang \t")
	fmt.Fprintln(w, "---\t | ------\t | -----------\t | --------------\t")

	for i := range users {
		row := fmt.Sprintf("%v\t | %v\t | %v\t | $%v\t", users[i].Id, users[i].Email, users[i].FirstName, users[i].LastName)
		fmt.Fprintln(w, row)
	}
	w.Flush()

	return nil
}

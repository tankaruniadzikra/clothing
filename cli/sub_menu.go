package cli

import (
	"database/sql"
	"fmt"
	"os"
	"pair-project/entity"
)

// SubMenu will wait user input after they are logged in. The prompt will then take the user to their selected menu
func SubMenu(db *sql.DB) int {
	fmt.Println()
	fmt.Print(entity.SubMenu)

	var option int
	if err := inputScanner(&option, "Pilih menu yang ingin anda akses (1/2/3/4): "); err != nil {
		return 0
	}

	// 1. Kelola product
	// 2. Input order
	// 3. Lihat report
	// 4. Exit
	switch option {
	case 1:
		ProductMenu(db)
		return 0
	case 2:
		OrderMenu(db)
		return 0
	case 3:
		ReportMenu(db)
		return 0
	case 4:
		fmt.Println("Terima kasih telah mengunjungi Cakra Clothing Store. Sampai jumpa!")
		os.Exit(0)
		return 2
	default:
		fmt.Println("Maaf, menu yang anda pilih belum tersedia.")
		fmt.Println()
		return 0
	}
}

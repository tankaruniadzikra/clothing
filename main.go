package main

import (
	"flag"
	"fmt"
	"log"
	"pair-project/cli"
	"pair-project/config"
)

func main() {
	defer fmt.Println("Terima kasih telah mengunjungi Cakra Clothing Store. Sampai jumpa!")

	// get connection string from flag
	connString := flag.String("db_url", "root:@tcp(localhost:3306)/cakra_clothing", "db connection string")

	// connect and maintain DB connection
	db, err := config.ConnectDB(*connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// this loop will continuously prompt the menu to user until
	// the user is ready to move to next menu
	var mainOpt int
	for {
		if mainOpt > 0 {
			break
		}

		mainOpt = cli.MainMenu(db)
	}

	// same behavior as mainOpt
	var subOpt int
	for {
		if subOpt > 0 {
			break
		}

		subOpt = cli.SubMenu(db)
	}
}

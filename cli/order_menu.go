package cli

import (
	"database/sql"
	"fmt"
	"pair-project/entity"
)

func OrderMenu(db *sql.DB) int {
	fmt.Println()
	fmt.Print(entity.OrderMenu)

	return 0
}

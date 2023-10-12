package cli

import (
	"database/sql"
	"fmt"
	"pair-project/entity"
)

func ReportMenu(db *sql.DB) int {
	fmt.Println()
	fmt.Print(entity.ReportMenu)

	return 0
}

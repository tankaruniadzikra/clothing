package cli

import (
	"database/sql"
	"fmt"
	"os"
	"pair-project/entity"
	"pair-project/handler"
	"text/tabwriter"
)

func ReportMenu(db *sql.DB) int {
	fmt.Println()
	fmt.Print(entity.ReportMenu)
	var option int
	if err := inputScanner(&option, "Pilih menu yang ingin anda akses (1-8): "); err != nil {
		return 0
	}

	switch option {
	case 1:
		res, err := handler.OrderOverview(db)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Terjadi kesalahan teknis. Silahkan hubungi administrator")
			fmt.Println("Kembali ke menu...")
			return 0
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight)
		table := res.ConvertToTable()
		fmt.Fprintln(w, table)
		w.Flush()
	case 2:
		res, err := handler.OrderSummary(db)
		if err != nil {
			fmt.Println("Terjadi kesalahan teknis. Silahkan hubungi administrator")
			fmt.Println("Kembali ke menu...")
			return 0
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight)
		table := res.ConvertToTable()
		fmt.Fprintln(w, table)
		w.Flush()
	case 3:
		res, err := handler.SalesSummary(db)
		if err != nil {
			fmt.Println("Terjadi kesalahan teknis. Silahkan hubungi administrator")
			fmt.Println("Kembali ke menu...")
			return 0
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight)
		table := res.ConvertToTable()
		fmt.Fprintln(w, table)
		w.Flush()
	case 4:
		res, err := handler.InventorySummary(db)
		if err != nil {
			fmt.Println("Terjadi kesalahan teknis. Silahkan hubungi administrator")
			fmt.Println("Kembali ke menu...")
			return 0
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight)
		table := res.ConvertToTable()
		fmt.Fprintln(w, table)
		w.Flush()
	case 5:
		res, err := handler.TopPurchasedProducts(db)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Terjadi kesalahan teknis. Silahkan hubungi administrator")
			fmt.Println("Kembali ke menu...")
			return 0
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight)
		table := res.ConvertToTable()
		fmt.Fprintln(w, table)
		w.Flush()
	case 6:
		res, err := handler.TopPurchasedBrands(db)
		if err != nil {
			fmt.Println("Terjadi kesalahan teknis. Silahkan hubungi administrator")
			fmt.Println("Kembali ke menu...")
			return 0
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight)
		table := res.ConvertToTable()
		fmt.Fprintln(w, table)
		w.Flush()
	case 7:
		res, err := handler.TopSpender(db)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Terjadi kesalahan teknis. Silahkan hubungi administrator")
			fmt.Println("Kembali ke menu...")
			return 0
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight)
		table := res.ConvertToTable()
		fmt.Fprintln(w, table)
		w.Flush()
	case 8:
		fmt.Println("Terima kasih telah mengunjungi Cakra Clothing Store. Sampai jumpa!")
		os.Exit(0)
		return 2
	}

	return 0
}

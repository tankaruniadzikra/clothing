package cli

import (
	"database/sql"
	"fmt"
	"log"
	"pair-project/entity"
	"pair-project/handler"
)

func ReportMenu(db *sql.DB) int {
	fmt.Println()
	fmt.Print(entity.ReportMenu)
	var option int
	if err := inputScanner(&option, "Pilih menu yang ingin anda akses (1/2/3/4): "); err != nil {
		return 0
	}

	switch option {
	case 1:
		topUsers, err := handler.TopPurchasingUsers(db)
		if err != nil {
			log.Fatal(err)
		}

		for _, user := range topUsers {
			fmt.Println("\nUser yang melakukan total pembelian terbanyak:")
			fmt.Printf("UserID: %d - Email: %s - Name: %s %s - Total Purchase: $%.2f\n", user.Id, user.Email, user.FirstName, user.LastName, user.TotalPurchase)
		}

	case 2:
		orderReports, err := handler.TopSellingDayReport(db)
		if err != nil {
			log.Fatal(err)
		}

		for _, report := range orderReports {
			fmt.Println("\nTotal penjualan terbanyak perhari:")
			fmt.Printf("Order Day: %s\nTotal Quantity: %d\nTotal Sales: %.2f\n\n", report.OrderDay, report.TotalQuantity, report.TotalSales)
		}

	case 3:
		maxProduct, err := handler.GetMaxStockProduct(db)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Produk dengan stok paling banyak:")
		fmt.Printf("Product ID: %d - Name: %s - Description: %s - Price: %.2f - Brand: %s - Size: %s - Color: %s - Stock: %d\n",
			maxProduct.ProductID, maxProduct.ProductName, maxProduct.Description,
			maxProduct.Price, maxProduct.Brand.BrandName, maxProduct.Size.SizeName,
			maxProduct.Color.ColorName, maxProduct.Stock.Quantity,
		)

		minProduct, err := handler.GetMinStockProduct(db)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("\nProduk dengan stok paling sedikit:")
		fmt.Printf("Product ID: %d - Name: %s - Description: %s - Price: %.2f - Brand: %s - Size: %s - Color: %s - Stock: %d\n",
			minProduct.ProductID, minProduct.ProductName, minProduct.Description,
			minProduct.Price, minProduct.Brand.BrandName, minProduct.Size.SizeName,
			minProduct.Color.ColorName, minProduct.Stock.Quantity,
		)

	case 4:
		fmt.Println("Terima kasih telah mengunjungi Cakra Clothing Store. Sampai jumpa!")
		return 2
	}

	return 0
}

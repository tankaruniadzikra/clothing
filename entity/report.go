package entity

import (
	"fmt"
	"time"
)

type OrderOverviews []OrderOverview

func (r OrderOverviews) ConvertToTable() string {
	var table string
	table = table + ("------- \t ------ \t --------- \t ----------- \t \n")
	table = table + ("OrderID \t UserID \t OrderDate \t TotalAmount \t \n")
	table = table + ("------- \t ------ \t --------- \t ----------- \t \n")

	for i := range r {
		table = table + fmt.Sprintf("%v \t %v \t %v \t %v \t \n", r[i].OrderID, r[i].UserID, r[i].OrderDate.Format("2006-01-02"), r[i].TotalAmount)
	}

	return table
}

type OrderOverview struct {
	OrderID     int
	UserID      int
	OrderDate   time.Time
	TotalAmount float64
}

type OrderSummaries []OrderSummary

func (r OrderSummaries) ConvertToTable() string {
	var table string
	table = table + ("--------- \t ----------- \t ---------------- \t \n")
	table = table + ("OrderDate \t TotalOrders \t TotalAmountSpent \t \n")
	table = table + ("--------- \t ----------- \t ---------------- \t \n")

	for _, summary := range r {
		table = table + fmt.Sprintf("%s \t %d \t %.2f \t \n", summary.OrderDate.Format("2006-01-02"), summary.TotalOrders, summary.TotalAmountSpent)
	}

	return table
}

type OrderSummary struct {
	OrderDate        time.Time
	TotalOrders      int
	TotalAmountSpent float64
}

type SalesSummaries []SalesSummary

func (r SalesSummaries) ConvertToTable() string {
	var table string
	table = table + ("--------- \t ---------- \t \n")
	table = table + ("OrderDate \t TotalSales \t \n")
	table = table + ("--------- \t ---------- \t \n")

	for _, summary := range r {
		table = table + fmt.Sprintf("%s \t %.2f \t \n", summary.OrderDate.Format("2006-01-02"), summary.TotalSales)
	}

	return table
}

type SalesSummary struct {
	OrderDate  time.Time
	TotalSales float64
}

type InventorySummaries []InventorySummary

func (r InventorySummaries) ConvertToTable() string {
	var table string
	table = table + ("----------- \t ------- \t \n")
	table = table + ("ProductName \t InStock \t \n")
	table = table + ("----------- \t ------- \t \n")

	for _, summary := range r {
		table = table + fmt.Sprintf("%s \t %d \t \n", summary.ProductName, summary.InStock)
	}

	return table
}

type InventorySummary struct {
	ProductName string
	InStock     int
}

type TopPurchasedProducts []TopPurchasedProduct

func (r TopPurchasedProducts) ConvertToTable() string {
	var table string
	table = table + ("----------- \t -------------- \t \n")
	table = table + ("ProductName \t TotalPurchased \t \n")
	table = table + ("----------- \t -------------- \t \n")

	for _, product := range r {
		table = table + fmt.Sprintf("%s \t %d \t \n", product.ProductName, product.TotalPurchased)
	}

	return table
}

type TopPurchasedProduct struct {
	ProductName    string
	TotalPurchased int
}

type TopPurchasedBrands []TopPurchasedBrand

func (r TopPurchasedBrands) ConvertToTable() string {
	var table string
	table = table + ("--------- \t -------------- \t \n")
	table = table + ("BrandName \t TotalPurchased \t \n")
	table = table + ("--------- \t -------------- \t \n")

	for _, brand := range r {
		table = table + fmt.Sprintf("%s \t %d \t \n", brand.BrandName, brand.TotalPurchased)
	}

	return table
}

type TopPurchasedBrand struct {
	BrandName      string
	TotalPurchased int
}

type TopSpenders []TopSpender

func (r TopSpenders) ConvertToTable() string {
	var table string
	table = table + ("--------- \t -------- \t ---------- \t \n")
	table = table + ("FirstName \t LastName \t TotalSpent \t \n")
	table = table + ("--------- \t -------- \t ---------- \t \n")

	for _, spender := range r {
		table = table + fmt.Sprintf("%s \t %s \t %.2f \t \n", spender.FirstName, spender.LastName, spender.TotalSpent)
	}

	return table
}

type TopSpender struct {
	FirstName  string
	LastName   string
	TotalSpent float64
}

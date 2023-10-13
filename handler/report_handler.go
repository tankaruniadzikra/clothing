package handler

import (
	"database/sql"
	"pair-project/entity"
	"time"
)

// OrderOverview will return OrderOverview entity that represents its report
func OrderOverview(db *sql.DB) (entity.OrderOverviews, error) {
	rows, err := db.Query(orderOverview)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var overview []entity.OrderOverview
	for rows.Next() {
		var order entity.OrderOverview
		var date string
		err := rows.Scan(&order.OrderID, &order.UserID, &date, &order.TotalAmount)
		if err != nil {
			return nil, err
		}

		order.OrderDate, err = time.Parse("2006-01-02", date)
		if err != nil {
			return nil, err
		}

		overview = append(overview, order)
	}

	return overview, nil
}

// OrderSummary will return OrderSummary entity that represents its report
func OrderSummary(db *sql.DB) (entity.OrderSummaries, error) {
	rows, err := db.Query(orderSummary)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var summary []entity.OrderSummary
	for rows.Next() {
		var order entity.OrderSummary
		var date string
		err := rows.Scan(&date, &order.TotalOrders, &order.TotalAmountSpent)
		if err != nil {
			return nil, err
		}

		order.OrderDate, err = time.Parse("2006-01-02", date)
		if err != nil {
			return nil, err
		}

		summary = append(summary, order)
	}

	return summary, nil
}

// SalesSummary will return SalesSummary entity that represents its report
func SalesSummary(db *sql.DB) (entity.SalesSummaries, error) {
	rows, err := db.Query(salesSummary)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var summary []entity.SalesSummary
	for rows.Next() {
		var sales entity.SalesSummary
		var date string
		err := rows.Scan(&date, &sales.TotalSales)
		if err != nil {
			return nil, err
		}

		sales.OrderDate, err = time.Parse("2006-01-02", date)
		if err != nil {
			return nil, err
		}

		summary = append(summary, sales)
	}

	return summary, nil
}

// InventorySummary will return InventorySummary entity that represents its report
func InventorySummary(db *sql.DB) (entity.InventorySummaries, error) {
	rows, err := db.Query(inventorySummary)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var summary []entity.InventorySummary
	for rows.Next() {
		var inventory entity.InventorySummary
		err := rows.Scan(&inventory.ProductName, &inventory.InStock)
		if err != nil {
			return nil, err
		}
		summary = append(summary, inventory)
	}

	return summary, nil
}

// TopPurchasedProducts will return TopPurchasedProducts entity that represents its report
func TopPurchasedProducts(db *sql.DB) (entity.TopPurchasedProducts, error) {
	rows, err := db.Query(topPurchasedProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entity.TopPurchasedProduct
	for rows.Next() {
		var product entity.TopPurchasedProduct
		err := rows.Scan(&product.ProductName, &product.TotalPurchased)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

// TopPurchasedBrands will return TopPurchasedBrands entity that represents its report
func TopPurchasedBrands(db *sql.DB) (entity.TopPurchasedBrands, error) {
	rows, err := db.Query(topPurchasedBrands)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var brands []entity.TopPurchasedBrand
	for rows.Next() {
		var brand entity.TopPurchasedBrand
		err := rows.Scan(&brand.BrandName, &brand.TotalPurchased)
		if err != nil {
			return nil, err
		}
		brands = append(brands, brand)
	}

	return brands, nil
}

// TopSpender will return TopSpender entity that represents its report
func TopSpender(db *sql.DB) (entity.TopSpenders, error) {
	rows, err := db.Query(topSpender)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var spenders []entity.TopSpender
	for rows.Next() {
		var spender entity.TopSpender
		err := rows.Scan(&spender.FirstName, &spender.LastName, &spender.TotalSpent)
		if err != nil {
			return nil, err
		}
		spenders = append(spenders, spender)
	}

	return spenders, nil
}

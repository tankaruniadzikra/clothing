package handler

import (
	"database/sql"
	"pair-project/entity"
)

func TopPurchasingUsers(db *sql.DB) ([]entity.User, error) {
	rows, err := db.Query(`
		SELECT 
			u.UserID, 
			u.Email, 
			u.FirstName, 
			u.LastName, 
			COALESCE(SUM(o.TotalAmount), 0) AS TotalPurchase
		FROM Users u
		LEFT JOIN Orders o ON u.UserID = o.UserID
		GROUP BY u.UserID
		ORDER BY TotalPurchase DESC
		LIMIT 1
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var topUsers []entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName, &user.TotalPurchase)
		if err != nil {
			return nil, err
		}
		topUsers = append(topUsers, user)
	}

	return topUsers, nil
}

func TopSellingDayReport(db *sql.DB) ([]entity.OrderReport, error) {
	rows, err := db.Query(`
		SELECT DATE(OrderDate) AS OrderDay, 
		       SUM(Quantity) AS TotalQuantity, 
		       COALESCE(SUM(TotalAmount), 0) AS TotalSales
		FROM Orders o
		JOIN OrderItems oi ON o.OrderID = oi.OrderID
		GROUP BY DATE(OrderDate)
		ORDER BY TotalSales DESC
		LIMIT 1
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orderReports []entity.OrderReport

	for rows.Next() {
		var orderReport entity.OrderReport
		err := rows.Scan(&orderReport.OrderDay, &orderReport.TotalQuantity, &orderReport.TotalSales)
		if err != nil {
			return nil, err
		}
		orderReports = append(orderReports, orderReport)
	}

	return orderReports, nil
}

func GetMaxStockProduct(db *sql.DB) (entity.Product, error) {
	var maxProduct entity.Product

	// Query untuk mendapatkan produk dengan stok terbanyak
	err := db.QueryRow(`
		SELECT p.ProductID, p.ProductName, p.Description, p.Price, p.Material, p.Weight,
			b.BrandID, b.BrandName, sz.SizeID, sz.SizeName, co.ColorID, co.ColorName,
			s.Quantity
		FROM Products p
		JOIN Brands b ON p.BrandID = b.BrandID
		JOIN Sizes sz ON p.SizeID = sz.SizeID
		JOIN Colors co ON p.ColorID = co.ColorID
		JOIN Stock s ON p.ProductID = s.ProductID
		ORDER BY s.Quantity DESC
		LIMIT 1;
	`).Scan(
		&maxProduct.ProductID, &maxProduct.ProductName, &maxProduct.Description,
		&maxProduct.Price, &maxProduct.Material, &maxProduct.Weight,
		&maxProduct.Brand.BrandID, &maxProduct.Brand.BrandName,
		&maxProduct.Size.SizeID, &maxProduct.Size.SizeName,
		&maxProduct.Color.ColorID, &maxProduct.Color.ColorName,
		&maxProduct.Stock.Quantity,
	)
	if err != nil {
		return entity.Product{}, err
	}

	return maxProduct, nil
}

func GetMinStockProduct(db *sql.DB) (entity.Product, error) {
	var minProduct entity.Product

	// Query untuk mendapatkan produk dengan stok paling sedikit
	err := db.QueryRow(`
		SELECT p.ProductID, p.ProductName, p.Description, p.Price, p.Material, p.Weight,
			b.BrandID, b.BrandName, sz.SizeID, sz.SizeName, co.ColorID, co.ColorName,
			s.Quantity
		FROM Products p
		JOIN Brands b ON p.BrandID = b.BrandID
		JOIN Sizes sz ON p.SizeID = sz.SizeID
		JOIN Colors co ON p.ColorID = co.ColorID
		JOIN Stock s ON p.ProductID = s.ProductID
		ORDER BY s.Quantity ASC
		LIMIT 1;
	`).Scan(
		&minProduct.ProductID, &minProduct.ProductName, &minProduct.Description,
		&minProduct.Price, &minProduct.Material, &minProduct.Weight,
		&minProduct.Brand.BrandID, &minProduct.Brand.BrandName,
		&minProduct.Size.SizeID, &minProduct.Size.SizeName,
		&minProduct.Color.ColorID, &minProduct.Color.ColorName,
		&minProduct.Stock.Quantity,
	)
	if err != nil {
		return entity.Product{}, err
	}

	return minProduct, nil
}

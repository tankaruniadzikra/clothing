package handler

const (
	orderOverview = `
	SELECT
		OrderID,
		UserID,
		OrderDate,
		TotalAmount
	FROM Orders;
`

	orderSummary = `
	SELECT
		DATE(OrderDate) AS OrderDate,
		COUNT(OrderID) AS TotalOrders,
		SUM(TotalAmount) AS TotalAmountSpent
	FROM Orders
	GROUP BY OrderDate;
`

	salesSummary = `
	SELECT
		DATE(OrderDate) AS OrderDate,
		SUM(TotalAmount) AS TotalSales
	FROM Orders
	GROUP BY OrderDate;
`

	inventorySummary = `
	SELECT
		P.ProductName,
		S.Quantity AS InStock
	FROM Stock S
	JOIN Products P ON S.ProductID = P.ProductID;
`

	topPurchasedProducts = `
	SELECT
		P.ProductName,
		COALESCE(0, SUM(OI.Quantity)) AS TotalPurchased
	FROM Products P
	LEFT JOIN OrderItems OI ON P.ProductID = OI.ProductID
	GROUP BY P.ProductName
	ORDER BY TotalPurchased DESC
	LIMIT 10;
`

	topPurchasedBrands = `
	SELECT
		B.BrandName,
		COALESCE(0, SUM(OI.Quantity)) AS TotalPurchased
	FROM Brands B
	LEFT JOIN Products P ON B.BrandID = P.BrandID
	LEFT JOIN OrderItems OI ON P.ProductID = OI.ProductID
	GROUP BY B.BrandName
	ORDER BY TotalPurchased DESC
	LIMIT 10;
`

	topSpender = `
	SELECT
		U.FirstName,
		U.LastName,
		COALESCE(0, SUM(O.TotalAmount)) AS TotalSpent
	FROM Users U
	LEFT JOIN Orders O ON U.UserID = O.UserID
	GROUP BY U.UserID
	ORDER BY TotalSpent DESC
	LIMIT 10;
`
)

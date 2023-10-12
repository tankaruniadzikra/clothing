package handler

const (
	createOrder = `INSERT INTO Orders (UserID, OrderDate, TotalAmount) VALUES (?, ?, ?)`

	createOrderDetail = `INSERT INTO OrderItems (OrderID, ProductID, Quantity, PricePerUnit, TotalPrice) VALUES (?, ?, ?, ?, ?)`

	decreaseStock = `UPDATE Stock SET Quantity = Quantity - ? WHERE ProductID = ?`

	readStocks = `SELECT StockID, ProductID, Quantity FROM Stock`
)

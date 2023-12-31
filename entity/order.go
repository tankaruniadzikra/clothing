package entity

type Order struct {
	OrderID     int
	UserID      int
	OrderDate   string
	TotalAmount float64
	Details     []OrderDetail
}

type OrderDetail struct {
	OrderItemID  int
	OrderID      int
	ProductID    int
	Quantity     int
	PricePerUnit float64
	TotalPrice   float64
}

type Stock struct {
	StockID   int
	ProductID int
	Quantity  int
}

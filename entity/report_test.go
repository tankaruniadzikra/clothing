package entity

import (
	"testing"
	"time"
)

func TestOrderOverviews_ConvertToTable(t *testing.T) {
	tests := []struct {
		name string
		o    OrderOverviews
		want string
	}{
		{
			name: "ok",
			o: OrderOverviews{
				{OrderID: 1, UserID: 1, OrderDate: time.Date(2023, 10, 13, 0, 0, 0, 0, time.UTC), TotalAmount: 100.0},
				{OrderID: 2, UserID: 2, OrderDate: time.Date(2023, 10, 14, 0, 0, 0, 0, time.UTC), TotalAmount: 200.0},
			},
			want: "------- \t ------ \t --------- \t ----------- \t \nOrderID \t UserID \t OrderDate \t TotalAmount \t \n------- \t ------ \t --------- \t ----------- \t \n1 \t 1 \t 2023-10-13 \t 100 \t \n2 \t 2 \t 2023-10-14 \t 200 \t \n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.ConvertToTable(); got != tt.want {
				t.Errorf("OrderOverviews.ConvertToTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderSummaries_ConvertToTable(t *testing.T) {
	tests := []struct {
		name string
		s    OrderSummaries
		want string
	}{
		{
			name: "ok",
			s: OrderSummaries{
				{OrderDate: time.Date(2023, 10, 13, 0, 0, 0, 0, time.UTC), TotalOrders: 5, TotalAmountSpent: 500.0},
				{OrderDate: time.Date(2023, 10, 14, 0, 0, 0, 0, time.UTC), TotalOrders: 10, TotalAmountSpent: 1000.0},
			},
			want: "--------- \t ----------- \t ---------------- \t \nOrderDate \t TotalOrders \t TotalAmountSpent \t \n--------- \t ----------- \t ---------------- \t \n2023-10-13 \t 5 \t 500.00 \t \n2023-10-14 \t 10 \t 1000.00 \t \n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ConvertToTable(); got != tt.want {
				t.Errorf("OrderSummaries.ConvertToTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSalesSummaries_ConvertToTable(t *testing.T) {
	tests := []struct {
		name string
		s    SalesSummaries
		want string
	}{
		{
			name: "ok",
			s: SalesSummaries{
				{OrderDate: time.Date(2023, 10, 13, 0, 0, 0, 0, time.UTC), TotalSales: 1000.0},
				{OrderDate: time.Date(2023, 10, 14, 0, 0, 0, 0, time.UTC), TotalSales: 2000.0},
			},
			want: "--------- \t ---------- \t \nOrderDate \t TotalSales \t \n--------- \t ---------- \t \n2023-10-13 \t 1000.00 \t \n2023-10-14 \t 2000.00 \t \n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ConvertToTable(); got != tt.want {
				t.Errorf("SalesSummaries.ConvertToTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInventorySummaries_ConvertToTable(t *testing.T) {
	tests := []struct {
		name string
		i    InventorySummaries
		want string
	}{
		{
			name: "ok",
			i: InventorySummaries{
				{ProductName: "Product1", InStock: 10},
				{ProductName: "Product2", InStock: 20},
			},
			want: "----------- \t ------- \t \nProductName \t InStock \t \n----------- \t ------- \t \nProduct1 \t 10 \t \nProduct2 \t 20 \t \n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.ConvertToTable(); got != tt.want {
				t.Errorf("InventorySummaries.ConvertToTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTopPurchasedProducts_ConvertToTable(t *testing.T) {
	tests := []struct {
		name string
		p    TopPurchasedProducts
		want string
	}{
		{
			name: "ok",
			p: TopPurchasedProducts{
				{ProductName: "Product1", TotalPurchased: 100},
				{ProductName: "Product2", TotalPurchased: 200},
			},
			want: "----------- \t -------------- \t \nProductName \t TotalPurchased \t \n----------- \t -------------- \t \nProduct1 \t 100 \t \nProduct2 \t 200 \t \n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.ConvertToTable(); got != tt.want {
				t.Errorf("TopPurchasedProducts.ConvertToTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTopPurchasedBrands_ConvertToTable(t *testing.T) {
	tests := []struct {
		name string
		b    TopPurchasedBrands
		want string
	}{
		{
			name: "ok",
			b: TopPurchasedBrands{
				{BrandName: "Brand1", TotalPurchased: 1000},
				{BrandName: "Brand2", TotalPurchased: 2000},
			},
			want: "--------- \t -------------- \t \nBrandName \t TotalPurchased \t \n--------- \t -------------- \t \nBrand1 \t 1000 \t \nBrand2 \t 2000 \t \n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.ConvertToTable(); got != tt.want {
				t.Errorf("TopPurchasedBrands.ConvertToTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTopSpenders_ConvertToTable(t *testing.T) {
	tests := []struct {
		name string
		s    TopSpenders
		want string
	}{
		{
			name: "ok",
			s: TopSpenders{
				{FirstName: "John", LastName: "Doe", TotalSpent: 1000.0},
				{FirstName: "Jane", LastName: "Smith", TotalSpent: 2000.0},
			},
			want: "--------- \t -------- \t ---------- \t \nFirstName \t LastName \t TotalSpent \t \n--------- \t -------- \t ---------- \t \nJohn \t Doe \t 1000.00 \t \nJane \t Smith \t 2000.00 \t \n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ConvertToTable(); got != tt.want {
				t.Errorf("TopSpenders.ConvertToTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

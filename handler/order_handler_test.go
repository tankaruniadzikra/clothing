package handler

import (
	"database/sql"
	"fmt"
	"testing"

	"pair-project/entity"

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
	"github.com/stretchr/testify/assert"
)

func TestCreateOrderIntegration(t *testing.T) {
	// Create a connection to the test database (replace with your database configuration)
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/cakra_clothing_test")
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")

	// Clean up the database (if needed)
	_, err = db.Exec("TRUNCATE TABLE Orders")
	if err != nil {
		t.Fatalf("Failed to truncate Orders table: %v", err)
	}

	// Define the test order and its details
	order := entity.Order{
		UserID:      1,
		OrderDate:   "2023-10-13",
		TotalAmount: 500.0,
		Details: []entity.OrderDetail{
			{ProductID: 1, Quantity: 2, PricePerUnit: 100.0, TotalPrice: 200.0},
			{ProductID: 2, Quantity: 3, PricePerUnit: 50.0, TotalPrice: 150.0},
		},
	}

	// Call the CreateOrder function with the real database connection
	err = CreateOrder(db, order)

	// Use assertions to verify the results
	assert.NoError(t, err, "Expected no error")

	// Verify the data in the database (query and assert the results)
	// Note: This depends on your database structure and schema.
	var orderCount int
	err = db.QueryRow("SELECT COUNT(*) FROM Orders").Scan(&orderCount)
	if err != nil {
		t.Fatalf("Failed to query Orders table: %v", err)
	}
	assert.Equal(t, 1, orderCount, "Expected one order in the database")
}

func init() {
	// Initialize and set up the database connection (only once)
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/cakra_clothing_test")
	if err != nil {
		fmt.Printf("Failed to initialize database connection: %v\n", err)
	}
	defer db.Close()
}

func TestReadStocksIntegration(t *testing.T) {
	// Create a connection to the test database (replace with your database configuration)
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/cakra_clothing_test")
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")

	// Clean up the database (if needed)
	_, err = db.Exec("TRUNCATE TABLE Stock")
	if err != nil {
		t.Fatalf("Failed to truncate Stock table: %v", err)
	}

	// Insert test data into the Stock table
	_, err = db.Exec("INSERT INTO Stock (ProductID, Quantity) VALUES (1, 100), (2, 200)")
	if err != nil {
		t.Fatalf("Failed to insert test data into Stock table: %v", err)
	}

	// Call the ReadStocks function with the real database connection
	stockMap, err := ReadStocks(db)

	// Use assertions to verify the results
	assert.NoError(t, err, "Expected no error")
	assert.NotNil(t, stockMap, "Expected a non-nil stock map")

	// Verify the data in the stock map (assuming the test data you inserted)
	assert.Equal(t, 100, stockMap[1].Quantity, "Expected quantity of 100 for ProductID 1")
	assert.Equal(t, 200, stockMap[2].Quantity, "Expected quantity of 200 for ProductID 2")

	// Clean up the database (if needed)
	_, err = db.Exec("TRUNCATE TABLE Stock")
	if err != nil {
		t.Fatalf("Failed to truncate Stock table: %v", err)
	}
}

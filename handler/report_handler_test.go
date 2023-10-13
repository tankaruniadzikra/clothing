package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderOverviewIntegration(t *testing.T) {
	// Set up the test database
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up the test database: %v", err)
	}
	defer db.Close()

	// Call the OrderOverview function with the real database connection
	_, err = OrderOverview(db)

	// Use assertions to verify the results
	assert.NoError(t, err, "Expected no error")

	// Clean up the test database
	teardownTestDB(db)
}

func TestOrderSummaryIntegration(t *testing.T) {
	// Set up the test database
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up the test database: %v", err)
	}
	defer db.Close()

	// Call the OrderSummary function with the real database connection
	_, err = OrderSummary(db)

	// Use assertions to verify the results
	assert.NoError(t, err, "Expected no error")

	// Clean up the test database
	teardownTestDB(db)
}

func TestSalesSummaryIntegration(t *testing.T) {
	// Set up the test database
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up the test database: %v", err)
	}
	defer db.Close()

	// Call the SalesSummary function with the real database connection
	_, err = SalesSummary(db)

	// Use assertions to verify the results
	assert.NoError(t, err, "Expected no error")

	// Clean up the test database
	teardownTestDB(db)
}

func TestInventorySummaryIntegration(t *testing.T) {
	// Set up the test database
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up the test database: %v", err)
	}
	defer db.Close()

	// Call the InventorySummary function with the real database connection
	_, err = InventorySummary(db)

	// Use assertions to verify the results
	assert.NoError(t, err, "Expected no error")

	// Clean up the test database
	teardownTestDB(db)
}

func TestTopPurchasedProductsIntegration(t *testing.T) {
	// Set up the test database
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up the test database: %v", err)
	}
	defer db.Close()

	// Call the TopPurchasedProducts function with the real database connection
	_, err = TopPurchasedProducts(db)

	// Use assertions to verify the results
	assert.NoError(t, err, "Expected no error")

	// Clean up the test database
	teardownTestDB(db)
}

func TestTopPurchasedBrandsIntegration(t *testing.T) {
	// Set up the test database
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up the test database: %v", err)
	}
	defer db.Close()

	// Call the TopPurchasedBrands function with the real database connection
	_, err = TopPurchasedBrands(db)

	// Use assertions to verify the results
	assert.NoError(t, err, "Expected no error")

	// Clean up the test database
	teardownTestDB(db)
}

func TestTopSpenderIntegration(t *testing.T) {
	// Set up the test database
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up the test database: %v", err)
	}
	defer db.Close()

	// Call the TopSpender function with the real database connection
	_, err = TopSpender(db)

	// Use assertions to verify the results
	assert.NoError(t, err, "Expected no error")

	// Clean up the test database
	teardownTestDB(db)
}

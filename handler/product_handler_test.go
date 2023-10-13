package handler

import (
	"database/sql"
	"fmt"
	"os"
	"pair-project/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
	"github.com/stretchr/testify/assert"
)

func TestCreateProductIntegration(t *testing.T) {
	// Create a connection to the test database (replace with your database configuration)
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/cakra_clothing_test")
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")

	// Clean up the database (if needed)
	_, err = db.Exec("TRUNCATE TABLE Products")
	if err != nil {
		t.Fatalf("Failed to truncate Products table: %v", err)
	}

	// Define the test product
	product := entity.Product{
		ProductName: "Test Product",
		Description: "Description",
		Price:       100.0,
		Material:    "Material",
		Weight:      1.5,
		BrandID:     1,
		SizeID:      1,
		ColorID:     1,
		Categories: []entity.Category{
			{CategoryID: 1},
			{CategoryID: 2},
		},
	}

	// Call the CreateProduct function with the real database connection
	err = CreateProduct(db, product)

	// Use assertions to verify the results
	assert.NoError(t, err, "Expected no error")

	// Verify the data in the database (query and assert the results)
	// Note: This depends on your database structure and schema.

	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")

	// Clean up the database (if needed)
	_, err = db.Exec("TRUNCATE TABLE Products")
	if err != nil {
		t.Fatalf("Failed to truncate Products table: %v", err)
	}
}

func TestReadProductsIntegration(t *testing.T) {
	// Create a connection to the test database (replace with your database configuration)
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/cakra_clothing_test")
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")

	// Clean up the database (if needed)
	_, err = db.Exec("TRUNCATE TABLE Products")
	if err != nil {
		t.Fatalf("Failed to truncate Products table: %v", err)
	}

	// Insert test data into the Products table
	_, err = db.Exec("INSERT INTO Products (ProductName, Description, Price, Material, Weight, BrandID, SizeID, ColorID) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		"Product A", "Description A", 100.0, "Material A", 1.0, 1, 1, 1)
	if err != nil {
		t.Fatalf("Failed to insert test data into Products table: %v", err)
	}

	// Call the ReadProducts function with the real database connection
	products, err := ReadProducts(db)

	// Use assertions to verify the results
	assert.NoError(t, err, "Expected no error")
	assert.NotNil(t, products, "Expected a non-nil product slice")
	assert.Len(t, products, 1, "Expected one product in the slice")

	// Verify the data in the products slice
	assert.Equal(t, "Product A", products[0].ProductName, "Expected Product A")

	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")

	// Clean up the database (if needed)
	_, err = db.Exec("TRUNCATE TABLE Products")
	if err != nil {
		t.Fatalf("Failed to truncate Products table: %v", err)
	}
}

func TestUpdateProductIntegration(t *testing.T) {
	// Create a connection to the test database (replace with your database configuration)
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/cakra_clothing_test")
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")

	// Clean up the database (if needed)
	_, err = db.Exec("TRUNCATE TABLE Products")
	if err != nil {
		t.Fatalf("Failed to truncate Products table: %v", err)
	}

	// Insert a test product into the Products table
	_, err = db.Exec("INSERT INTO Products (ProductName, Description, Price, Material, Weight, BrandID, SizeID, ColorID) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		"Product A", "Description A", 100.0, "Material A", 1.0, 1, 1, 1)
	if err != nil {
		t.Fatalf("Failed to insert a test product into Products table: %v", err)
	}

	// Retrieve the product ID (you might use the real product ID from the database)
	var productID int
	err = db.QueryRow("SELECT ProductID FROM Products WHERE ProductName = ?", "Product A").Scan(&productID)
	if err != nil {
		t.Fatalf("Failed to retrieve ProductID: %v", err)
	}

	// Define the updated product
	updatedProduct := entity.Product{
		ProductID:   productID,
		ProductName: "Updated Product A",
		Description: "Updated Description A",
		Price:       150.0,
		Material:    "Updated Material A",
		Weight:      1.5,
		BrandID:     2,
		SizeID:      2,
		ColorID:     2,
	}

	// Call the UpdateProduct function with the real database connection
	err = UpdateProduct(db, updatedProduct)

	// Use assertions to verify the results
	assert.NoError(t, err, "Expected no error")

	// Verify the data in the database (query and assert the results)
	// Note: This depends on your database structure and schema.

	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")

	// Clean up the database (if needed)
	_, err = db.Exec("TRUNCATE TABLE Products")
	if err != nil {
		t.Fatalf("Failed to truncate Products table: %v", err)
	}
}

func TestDeleteProductIntegration(t *testing.T) {
	// Create a connection to the test database (replace with your database configuration)
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/cakra_clothing_test")
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")

	// Clean up the database (if needed)
	_, err = db.Exec("TRUNCATE TABLE Products")
	if err != nil {
		t.Fatalf("Failed to truncate Products table: %v", err)
	}

	// Insert a test product into the Products table
	_, err = db.Exec("INSERT INTO Products (ProductName, Description, Price, Material, Weight, BrandID, SizeID, ColorID) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		"Product A", "Description A", 100.0, "Material A", 1.0, 1, 1, 1)
	if err != nil {
		t.Fatalf("Failed to insert a test product into Products table: %v", err)
	}

	// Retrieve the product ID (you might use the real product ID from the database)
	var productID int
	err = db.QueryRow("SELECT ProductID FROM Products WHERE ProductName = ?", "Product A").Scan(&productID)
	if err != nil {
		t.Fatalf("Failed to retrieve ProductID: %v", err)
	}

	// Call the DeleteProduct function with the real database connection
	err = DeleteProduct(db, productID)

	// Use assertions to verify the results
	assert.NoError(t, err, "Expected no error")

	// Verify that the product was deleted (query and assert the results)
	// Note: This depends on your database structure and schema.

	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")

	// Clean up the database (if needed)
	_, err = db.Exec("TRUNCATE TABLE Products")
	if err != nil {
		t.Fatalf("Failed to truncate Products table: %v", err)
	}
}

func TestReadSizesIntegration(t *testing.T) {
	// Create a connection to the test database (replace with your database configuration)
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/cakra_clothing_test")
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")

	// Clean up the database (if needed)
	_, err = db.Exec("TRUNCATE TABLE Sizes")
	if err != nil {
		t.Fatalf("Failed to truncate Sizes table: %v", err)
	}

	// Insert test data into the Sizes table
	_, err = db.Exec("INSERT INTO Sizes (SizeName) VALUES ('Small'), ('Medium')")
	if err != nil {
		t.Fatalf("Failed to insert test data into Sizes table: %v", err)
	}

	// Call the ReadSizes function with the real database connection
	sizes, err := ReadSizes(db)

	// Use assertions to verify the results
	assert.NoError(t, err, "Expected no error")
	assert.NotNil(t, sizes, "Expected a non-nil sizes slice")
	assert.Len(t, sizes, 2, "Expected two sizes in the slice")

	// Verify the data in the sizes slice
	assert.Equal(t, "Small", sizes[0].SizeName, "Expected Small size")
	assert.Equal(t, "Medium", sizes[1].SizeName, "Expected Medium size")

	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")

	// Clean up the database (if needed)
	_, err = db.Exec("TRUNCATE TABLE Sizes")
	if err != nil {
		t.Fatalf("Failed to truncate Sizes table: %v", err)
	}
}

func TestReadBrandsIntegration(t *testing.T) {
	// Create a connection to the test database (replace with your database configuration)
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/cakra_clothing_test")
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")

	// Clean up the database (if needed)
	_, err = db.Exec("TRUNCATE TABLE Brands")
	if err != nil {
		t.Fatalf("Failed to truncate Brands table: %v", err)
	}

	// Insert test data into the Brands table
	_, err = db.Exec("INSERT INTO Brands (BrandName) VALUES ('Brand A'), ('Brand B')")
	if err != nil {
		t.Fatalf("Failed to insert test data into Brands table: %v", err)
	}

	// Call the ReadBrands function with the real database connection
	brands, err := ReadBrands(db)

	// Use assertions to verify the results
	assert.NoError(t, err, "Expected no error")
	assert.NotNil(t, brands, "Expected a non-nil brands slice")
	assert.Len(t, brands, 2, "Expected two brands in the slice")

	// Verify the data in the brands slice
	assert.Equal(t, "Brand A", brands[0].BrandName, "Expected Brand A")
	assert.Equal(t, "Brand B", brands[1].BrandName, "Expected Brand B")

	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")

	// Clean up the database (if needed)
	_, err = db.Exec("TRUNCATE TABLE Brands")
	if err != nil {
		t.Fatalf("Failed to truncate Brands table: %v", err)
	}
}

func TestReadColorsIntegration(t *testing.T) {
	// Create a connection to the test database (replace with your database configuration)
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/cakra_clothing_test")
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")

	// Clean up the database (if needed)
	_, err = db.Exec("TRUNCATE TABLE Colors")
	if err != nil {
		t.Fatalf("Failed to truncate Colors table: %v", err)
	}

	// Insert test data into the Colors table
	_, err = db.Exec("INSERT INTO Colors (ColorName) VALUES ('Red'), ('Blue')")
	if err != nil {
		t.Fatalf("Failed to insert test data into Colors table: %v", err)
	}

	// Call the ReadColors function with the real database connection
	colors, err := ReadColors(db)

	// Use assertions to verify the results
	assert.NoError(t, err, "Expected no error")
	assert.NotNil(t, colors, "Expected a non-nil colors slice")
	assert.Len(t, colors, 2, "Expected two colors in the slice")

	// Verify the data in the colors slice
	assert.Equal(t, "Red", colors[0].ColorName, "Expected Red color")
	assert.Equal(t, "Blue", colors[1].ColorName, "Expected Blue color")

	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")

	// Clean up the database (if needed)
	_, err = db.Exec("TRUNCATE TABLE Colors")
	if err != nil {
		t.Fatalf("Failed to truncate Colors table: %v", err)
	}
}

func TestReadCategoriesIntegration(t *testing.T) {
	// Create a connection to the test database (replace with your database configuration)
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/cakra_clothing_test")
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")

	// Clean up the database (if needed)
	_, err = db.Exec("TRUNCATE TABLE Categories")
	if err != nil {
		t.Fatalf("Failed to truncate Categories table: %v", err)
	}

	// Insert test data into the Categories table
	_, err = db.Exec("INSERT INTO Categories (CategoryName) VALUES ('Category A'), ('Category B')")
	if err != nil {
		t.Fatalf("Failed to insert test data into Categories table: %v", err)
	}

	// Call the ReadCategories function with the real database connection
	categories, err := ReadCategories(db)

	// Use assertions to verify the results
	assert.NoError(t, err, "Expected no error")
	assert.NotNil(t, categories, "Expected a non-nil categories slice")
	assert.Len(t, categories, 2, "Expected two categories in the slice")

	// Verify the data in the categories slice
	assert.Equal(t, "Category A", categories[0].CategoryName, "Expected Category A")
	assert.Equal(t, "Category B", categories[1].CategoryName, "Expected Category B")

	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")

	// Clean up the database (if needed)
	_, err = db.Exec("TRUNCATE TABLE Categories")
	if err != nil {
		t.Fatalf("Failed to truncate Categories table: %v", err)
	}
}

// Make sure to replace the database connection string and queries with your actual database details.

func TestMain(m *testing.M) {
	// You can set up any database connection and perform schema setup/teardown here if needed.
	fmt.Println("Test setup")

	// Run the tests
	exitCode := m.Run()

	// You can do any necessary cleanup here.

	fmt.Println("Test teardown")
	os.Exit(exitCode)
}

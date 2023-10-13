package handler

import (
	"pair-project/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestLoginIntegration(t *testing.T) {
	// Set up the test database
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up the test database: %v", err)
	}
	defer db.Close()

	// Insert test user data into the database
	// You can use SQL queries to insert test data.

	// Define a test user
	email := "testuser@example.com"
	password := "testpassword"

	// Create a hashed password for the test user
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("Failed to hash the test password: %v", err)
	}

	// Insert the test user into the database
	_, err = db.Exec("INSERT INTO Users (Email, Password, FirstName, LastName) VALUES (?,?,?,?)", email, hashedPassword, "John", "Doe")
	if err != nil {
		t.Fatalf("Failed to insert the test user: %v", err)
	}

	// Call the Login function with the real database connection
	user, err := Login(db, email, password)

	// Use assertions to verify the results
	assert.NoError(t, err, "Expected no error")
	assert.NotNil(t, user, "Expected a non-nil user")

	// Verify the user's information to ensure a successful login
	assert.Equal(t, email, user.Email, "Expected email to match")

	// Clean up the test database
	teardownTestDB(db)
}

func TestRegisterIntegration(t *testing.T) {
	// Set up the test database
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up the test database: %v", err)
	}
	defer db.Close()

	// Create a test user
	user := entity.User{
		Email:     "newuser@example.com",
		Password:  "newpassword",
		FirstName: "New",
		LastName:  "User",
	}

	// Call the Register function with the real database connection
	err = Register(db, user)

	// Use assertions to verify the results
	assert.NoError(t, err, "Expected no error")

	// Clean up the test database
	teardownTestDB(db)
}

func TestReadUsersIntegration(t *testing.T) {
	// Set up the test database
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up the test database: %v", err)
	}
	defer db.Close()

	// Insert test user data into the database
	// You can use SQL queries to insert test data.

	// Call the ReadUsers function with the real database connection
	users, err := ReadUsers(db)

	// Use assertions to verify the results
	assert.NoError(t, err, "Expected no error")
	assert.NotNil(t, users, "Expected a non-nil users slice")

	// Verify the data in the users slice
	// You should compare the actual results with the expected results.

	// Clean up the test database
	teardownTestDB(db)
}

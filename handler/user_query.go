package handler

const (
	login = `SELECT UserID, Email, Password, FirstName, LastName FROM Users WHERE Email = ?`

	register = `INSERT INTO Users (Email, Password, FirstName, LastName) VALUES (?,?,?,?)`

	readUser = `SELECT UserID, Email, FirstName, LastName FROM Users`
)

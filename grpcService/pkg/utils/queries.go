package utils

var (
	CreateUser  string = "INSERT INTO USER (first_name, id, age, pass) VALUES (?,?,?,?)"
	GetUser     string = "SELECT first_name, age, email FROM USER WHERE id = ?"
	GetPassword string = "SELECT pass FROM USER WHERE id = ?"
)

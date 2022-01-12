package utils

var (
	CreateUser string = "INSERT INTO USER (first_name, id, age, pass) VALUES (?,?,?,?)"
	GetUser    string = "SELECT first_name, age FROM USER WHERE id = ?"
)

package utils

var (
	CreateUser string = "INSERT INTO USER (first_name,pass,age) VALUES (?,?,?)"
	GetUser    string = "SELECT first_name, age FROM USER WHERE ID = ?"
)

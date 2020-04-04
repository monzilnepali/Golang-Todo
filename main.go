package main

import (
	"fmt"
	"net/http"

	"github.com/monzilnepali/Golang-Todo/db"
	todo "github.com/monzilnepali/Golang-Todo/routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db.ConnectDB()
	defer db.DB.Close()
	fmt.Println("db connected")

	http.HandleFunc("/", todo.Home)
	http.HandleFunc("/fetchtodo", todo.GetTodo)
	http.HandleFunc("/addtodo", todo.AddTodo)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

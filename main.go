package main

import (
	"fmt"
	"net/http"
	"todo/db"
	route "todo/routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db.ConnectDB()
	defer db.DB.Close()
	fmt.Println("db connected")

	http.HandleFunc("/", route.Home)
	http.HandleFunc("/fetchtodo", route.GetTodo)
	http.HandleFunc("/addtodo", route.AddTodo)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

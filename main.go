package main

import (
	"fmt"
	"net/http"

	"github.com/monzilnepali/Golang-Todo/db"
	"github.com/monzilnepali/Golang-Todo/middleware"
	todo "github.com/monzilnepali/Golang-Todo/routes"
	user "github.com/monzilnepali/Golang-Todo/routes"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	db.ConnectDB()
	fmt.Println("db connected")
	defer db.DB.Close()
}

func main() {
	finalHandler := http.HandlerFunc(todo.Home)
	http.Handle("/", middleware.AuthMiddleware(finalHandler))
	http.HandleFunc("/fetchtodo", todo.GetTodo)
	http.HandleFunc("/addtodo", todo.AddTodo)
	http.HandleFunc("/login", user.Login)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

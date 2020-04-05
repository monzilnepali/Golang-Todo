package main

import (
	"fmt"
	"net/http"

	"github.com/monzilnepali/Golang-Todo/db"
	"github.com/monzilnepali/Golang-Todo/middleware"
	"github.com/monzilnepali/Golang-Todo/routes"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	db.ConnectDB()
	fmt.Println("db connected")

}

func main() {
	defer db.DB.Close()

	http.HandleFunc("/", middleware.Auth(routes.Home))
	// http.HandleFunc("/fetchtodo", middleware.AuthMiddleware( http.HandlerFunc(todo.GetTodo)))
	http.HandleFunc("/fetchtodo", middleware.Auth(routes.GetAllTodo))

	// http.HandleFunc("/addtodo", middleware.Auth(todo.AddTodo))
	http.HandleFunc("/signup", routes.Signup)
	http.HandleFunc("/login", routes.Login)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

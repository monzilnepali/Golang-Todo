package main

import (
	"fmt"
	"net/http"
	"path"

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
	fmt.Println("path", path.Base)
	http.HandleFunc("/signup", routes.Signup)
	http.HandleFunc("/login", routes.Login)
	http.HandleFunc("/", middleware.Auth(routes.Home))
	// http.HandleFunc("/fetchtodo", middleware.AuthMiddleware( http.HandlerFunc(todo.GetTodo)))
	http.HandleFunc("/fetchtodo", middleware.Auth(routes.GetAllTodo))
	http.HandleFunc("/addtodo", middleware.Auth(routes.AddTodo))
	http.HandleFunc("/updatetodo/:id", middleware.Auth(routes.UpdateTodo))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

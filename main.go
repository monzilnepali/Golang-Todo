package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
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
	router := httprouter.New()
	router.POST("/signup", routes.Signup)
	router.POST("/login", routes.Login)
	// router.GET("/:id", middleware.Auth(routes.Home))
	router.GET("/fetchtodo", middleware.Auth(routes.GetAllTodo))

	//	http.HandleFunc("/", middleware.Auth(routes.Home))
	// http.HandleFunc("/fetchtodo", middleware.Auth(routes.GetAllTodo))
	// http.HandleFunc("/addtodo", middleware.Auth(routes.AddTodo))
	// http.HandleFunc("/updatetodo/:id", middleware.Auth(routes.UpdateTodo))

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}

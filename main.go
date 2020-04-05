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
	router.GET("/", routes.Home)
	router.POST("/auth/signup", routes.Signup)
	router.POST("/auth/login", routes.Login)
	router.GET("/api/fetchtodo", middleware.Auth(routes.GetAllTodo))
	router.PUT("/api/updatetodo/:id", middleware.Auth(routes.UpdateTodo))
	router.DELETE("/api/deletetodo/:id", middleware.Auth(routes.DeleteTodo))
	router.POST("/api/addtodo", middleware.Auth(routes.AddTodo))

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}

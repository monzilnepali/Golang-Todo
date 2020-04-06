package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/monzilnepali/Golang-Todo/middleware"
)

//Routes list
func Routes() http.Handler {

	router := httprouter.New()
	router.GET("/", handlers.Home)
	router.POST("/auth/signup", handlers.Signup)
	router.POST("/auth/login", handlers.Login)

	router.GET("/api/fetchtodo", middleware.Auth(handers))
	router.PUT("/api/updatetodo/:id", middleware.Auth(handlers.UpdateTodo))
	//	router.DELETE("/api/deletetodo/:id", middleware.Auth(handlers.DeleteTodo))
	//	router.POST("/api/addtodo", middleware.Auth(handlers.AddTodo))
	return router

}

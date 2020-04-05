package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/monzilnepali/Golang-Todo/handler"
	"github.com/monzilnepali/Golang-Todo/model"
)

//Home route
func Home(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "root page")
	fmt.Println("parama", ps.ByName("id"))
	// fmt.Println("id of current active use", activeUser.UserID)
	s, ok := r.Context().Value("settings").(string)

	if !ok {
		fmt.Println("s is not type string")
	}

	fmt.Println(s) // "/var/bin"
}

//GetTodo route
func GetAllTodo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//for get request only

	//getting user id via context
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		fmt.Println("s is not type int")
	}

	w.Header().Set("Content-Type", "application/json")
	todolist := handler.GetTodoList(userID)
	json.NewEncoder(w).Encode(todolist)

}

//AddTodo route
func AddTodo(w http.ResponseWriter, r *http.Request, activeUser *model.User) {
	if r.Method != "POST" {
		fmt.Fprint(w, r.Method+r.URL.Path+" cannot be resolve")
	}
	//getting title from reqest body
	decoder := json.NewDecoder(r.Body)
	var newTodo model.Todo
	err := decoder.Decode(&newTodo)
	if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	//addTodoHandler
	err = handler.AddTodoHandler(activeUser.UserID, newTodo.Title)
	if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	fmt.Fprint(w, "Todo add")
	return

}

//UpdateTodo handler
func UpdateTodo(w http.ResponseWriter, r *http.Request, activeUser *model.User) {
	if r.Method != "PUT" {
		fmt.Fprint(w, r.Method+r.URL.Path+" cannot be resolve")
	}

	fmt.Fprint(w, "update todo end point")

}

//DeleteTodo handler
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Fprint(w, r.Method+r.URL.Path+" cannot be resolve")
	}
	fmt.Fprint(w, "delete todo end point")

}

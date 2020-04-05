package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/monzilnepali/Golang-Todo/handler"
	"github.com/monzilnepali/Golang-Todo/model"
)

//Home route
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "root page")
}

//GetTodo route
func GetTodo(w http.ResponseWriter, r *http.Request) {
	//for get request only

	if r.Method != "GET" {
		fmt.Fprint(w, r.Method+r.URL.Path+" cannot be resolve")
	}
	w.Header().Set("Content-Type", "application/json")
	todolist := handler.GetTodoList()
	json.NewEncoder(w).Encode(todolist)

}

//AddTodo route
func AddTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprint(w, r.Method+r.URL.Path+" cannot be resolve")
	}
	fmt.Fprint(w, "addtodo end point hit")
	decoder := json.NewDecoder(r.Body)
	var mytodo model.Todo
	err := decoder.Decode(&mytodo)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("todo title ", mytodo)

}

//UpdateTodo handler
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "UPDATE" {
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

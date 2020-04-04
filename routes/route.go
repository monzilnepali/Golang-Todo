package route

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
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		todolist := handler.GetTodoList()
		json.NewEncoder(w).Encode(todolist)
	default:
		fmt.Fprint(w, r.Method+r.URL.Path+" cannot be resolve")
	}
}

//AddTodo route
func AddTodo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		fmt.Fprint(w, "addtodo end point hit")
		decoder := json.NewDecoder(r.Body)
		var mytodo model.Todo
		err := decoder.Decode(&mytodo)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("todo title ", mytodo)
	default:
		fmt.Fprint(w, r.Method+r.URL.Path+" cannot be resolve")
	}
}

//UpdateTodo handler
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "UPDATE":
		fmt.Fprint(w, "update todo end point")
	default:
		fmt.Fprint(w, r.Method+r.URL.Path+" cannot be resolve")
	}
}

//DeleteTodo handler
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "DELETE":
		fmt.Fprint(w, "delete todo end point")
	default:
		fmt.Fprint(w, r.Method+r.URL.Path+" cannot be resolve")
	}
}

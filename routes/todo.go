package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/monzilnepali/Golang-Todo/handler"
	"github.com/monzilnepali/Golang-Todo/model"
)

//Home route
func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "home page")
	return
}

//GetAllTodo handler
func GetAllTodo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//for get request only

	//getting user id via context
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		log.Fatal(ok)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	todolist := handler.GetTodoList(userID)
	json.NewEncoder(w).Encode(todolist)
	return

}

//AddTodo route
func AddTodo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	//getting user id via context
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		log.Fatal(ok)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	//getting title from reqest body
	decoder := json.NewDecoder(r.Body)
	var newTodo model.Todo
	err := decoder.Decode(&newTodo)
	if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	//addTodoHandler
	err = handler.AddTodoHandler(userID, newTodo.Title)
	if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Todo add")
	return

}

//UpdateTodo handler
func UpdateTodo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	//getting user id via context
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		log.Fatal(ok)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	//parse int
	todoID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	//update todo db operation
	err = handler.UpdateTodoStatus(userID, todoID)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, "Todo updated")
	return

}

//DeleteTodo handler
func DeleteTodo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	//getting user id via context
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		log.Fatal(ok)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	//parse int
	todoID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = handler.DeleteTodoStatus(userID, todoID)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "Todo deleted")
	return

}

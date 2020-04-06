package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	handler "github.com/monzilnepali/Golang-Todo/domain"
)

//Home route
func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "home page")
	return
}

//GetAllTodoHandler return todo list
func GetAllTodoHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//for get request only

	//getting user id via context
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		log.Fatal(ok)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	todolist, err := handler.GetTodoList(userID)
	switch Err := err.(type) {
	case *handler.HTTPError:
		http.Error(w, Err.Message, Err.StatusCode)
	case nil:
		json.NewEncoder(w).Encode(todolist)
		return
	default:
		panic(err.Error())
	}

}

//AddTodoHandler add todo
func AddTodoHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

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
	err = handlers.AddTodoHandler(userID, r.Body)
	switch Err := err.(type) {
	case *handler.HTTPError:
		http.Error(w, Err.Message, Err.StatusCode)
	case nil:
		fmt.Fprint(w, "Todo add")
		return
	default:
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		panic(err.Error())
	}

}

//UpdateTodoHandler Update todo status(toggle iscompleted field)
func UpdateTodoHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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
	switch Err := err.(type) {
	case *handler.HTTPError:
		http.Error(w, Err.Message, Err.StatusCode)
	case nil:
		fmt.Fprint(w, "Todo UPDATED")
		return
	default:
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		panic(err.Error())
	}

}

//DeleteTodoHandler delete todo
func DeleteTodoHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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

	switch Err := err.(type) {
	case *handler.HTTPError:
		http.Error(w, Err.Message, Err.StatusCode)
	case nil:
		fmt.Fprint(w, "Todo DELETED")
		return
	default:
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		panic(err.Error())
	}

}

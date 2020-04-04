package handler

import (
	"fmt"

	"github.com/monzilnepali/Golang-Todo/db"
	"github.com/monzilnepali/Golang-Todo/model"
)

func GetTodoList() []model.Todo {
	//getting todo list
	fmt.Println("getrodolist called")
	result, err := db.DB.Query("SELECT * FROM todo")
	if err != nil {
		panic(err.Error())
	}

	var todoList []model.Todo
	for result.Next() {
		var todo model.Todo
		err := result.Scan(&todo.ID, &todo.Title, &todo.Iscompleted)
		if err != nil {
			panic(err.Error())
		}
		todoList = append(todoList, todo)

	}
	return todoList

}

package handler

import (
	"errors"
	"fmt"
	"log"

	"github.com/monzilnepali/Golang-Todo/db"
	"github.com/monzilnepali/Golang-Todo/model"
)

//GetTodoList handler
func GetTodoList(userID int) []model.Todo {
	//getting todo list
	fmt.Println("getrodolist called")
	result, err := db.DB.Query("SELECT * FROM todo 	WHERE UserID=?", userID)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var todoList []model.Todo
	for result.Next() {
		var todo model.Todo
		err := result.Scan(&todo.TodoID, &todo.Title, &todo.Iscompleted)
		if err != nil {
			panic(err.Error())
		}
		todoList = append(todoList, todo)

	}
	return todoList

}

func AddTodo(userId int, todoTitle string) error {
	fmt.Println("add todo called")
	stmt,err:db.DB.Prepare("INSERT INTO todo(UserID,Title) VALUES(?,?")
  if err!=nil{
		log.Fatal(err)
	}
	res,err:=stmt.Exec(userId,todoTitle)
	if err !=nil{
		log.Fatal(err)
		return err
	}
	return nil


}

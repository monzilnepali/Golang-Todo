package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/monzilnepali/Golang-Todo/db"
	"github.com/monzilnepali/Golang-Todo/model"
)

//GetTodoList handler
func GetTodoList(userID int) ([]model.Todo, error) {
	//getting todo list
	fmt.Println("getrodolist called")
	result, err := db.DB.Query("SELECT TodoID,Title,Iscompleted FROM todo 	WHERE UserID=?", userID)
	if err != nil {
		// panic(err.Error())
		return nil, errors.New(err.Error())
	}
	defer result.Close()
	var todoList []model.Todo
	for result.Next() {
		var todo model.Todo
		err := result.Scan(&todo.TodoID, &todo.Title, &todo.Iscompleted)
		if err != nil {
			return nil, NewHTTPError("cannot GET todo", http.StatusInternalServerError)
		}
		todoList = append(todoList, todo)

	}
	return todoList, nil

}

//AddTodoHandler handler
func AddTodoHandler(userID int, todoTitle string) error {
	fmt.Println("add todo called")
	stmt, err := db.DB.Prepare("INSERT INTO todo(UserID,Title) VALUES(?,?)")
	defer stmt.Close()
	if err != nil {
		return errors.New(err.Error())

	}
	_, err = stmt.Exec(userID, todoTitle)
	if err != nil {
		return NewHTTPError("cannot ADD todo", http.StatusInternalServerError)
	}
	return nil

}

//UpdateTodoStatus handler
func UpdateTodoStatus(userID, todoID int) error {
	fmt.Println("update todo db ops")

	stmt, err := db.DB.Prepare("UPDATE todo SET Iscompleted = NOT Iscompleted WHERE UserID=? AND TodoID=?")
	defer stmt.Close()
	if err != nil {
		return errors.New(err.Error())
	}
	_, err = stmt.Exec(userID, todoID)
	if err != nil {
		return NewHTTPError("cannot UPDATE todo", http.StatusInternalServerError)
	}
	return nil

}

//DeleteTodoStatus handler
func DeleteTodoStatus(userID, todoID int) error {
	fmt.Println("delete todo db osp")

	stmt, err := db.DB.Prepare("DELETE FROM todo WHERE UserID=? AND TodoID=?")
	defer stmt.Close()
	if err != nil {
		return errors.New(err.Error())

	}
	res, err := stmt.Exec(userID, todoID)

	if err != nil {
		return errors.New(err.Error())

	}
	//0 row affected error
	count, err := res.RowsAffected()
	if err != nil {
		return errors.New(err.Error())

	}
	if count == 0 {
		return NewHTTPError("cannot DELETE todo", http.StatusConflict)
	} else {
		return nil
	}

}

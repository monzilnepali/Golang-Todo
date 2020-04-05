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
	result, err := db.DB.Query("SELECT TodoID,Title,Iscompleted FROM todo 	WHERE UserID=?", userID)
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

//AddTodoHandler
func AddTodoHandler(userID int, todoTitle string) error {
	fmt.Println("add todo called")
	stmt, err := db.DB.Prepare("INSERT INTO todo(UserID,Title) VALUES(?,?)")
	defer stmt.Close()
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(userID, todoTitle)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil

}

func UpdateTodoStatus(userID, todoID int) error {
	fmt.Println("update todo db ops")

	stmt, err := db.DB.Prepare("UPDATE todo SET Iscompleted = NOT Iscompleted WHERE UserID=? AND TodoID=?")
	defer stmt.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}
	_, err = stmt.Exec(userID, todoID)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil

}

//DeleteTodoStatus handler
func DeleteTodoStatus(userID, todoID int) error {
	fmt.Println("delete todo db osp")

	stmt, err := db.DB.Prepare("DELETE FROM todo WHERE UserID=? AND TodoID=?")
	defer stmt.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}
	res, err := stmt.Exec(userID, todoID)

	if err != nil {
		log.Fatal(err)
		return err
	}
	//0 row affected error
	count, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
		return err
	}
	if count == 0 {
		return errors.New("cannot delete")
	} else {

		return nil
	}

}

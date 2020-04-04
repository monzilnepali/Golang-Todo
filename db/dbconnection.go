package db

import "database/sql"

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "lapzap:root@tcp(127.0.0.1:3306)/todo")

	if err != nil {
		panic(err.Error())
	}
	DB = db

}

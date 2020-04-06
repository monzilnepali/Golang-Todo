package main

import (
	"fmt"
	"net/http"

	"github.com/monzilnepali/Golang-Todo/db"
	"github.com/monzilnepali/Golang-Todo/router"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	db.ConnectDB()
	fmt.Println("db connected")

}

func main() {
	defer db.DB.Close()
	r := router.Routes()

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}

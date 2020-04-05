package routes

import (
	"fmt"
	"net/http"
)

//Login handler
func Login(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "hello from login ssytem")

}

//Signup Handler
func Signup(w http.ResponseWriter, r *http.Request) {
	fmt.FPrintf(w, "hello from signup")
}

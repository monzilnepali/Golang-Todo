package routes

import (
	"fmt"
	"net/http"
)

//Login handler
func Login(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "login system")

}

//Signup Handler
func Signup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello from signup")
}

package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Login handler
func Login(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "login system")

}

//Signup Handler
func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprint(w, r.Method+r.URL.Path+" cannot be resolve")
	}
	//getting email and password from req.body
	decoder := json.NewDecoder(r.Body)
	user := User{}
	fmt.Fprint(w, "hello from signup")
}

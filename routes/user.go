package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	user "github.com/monzilnepali/Golang-Todo/model"
	formValidation "github.com/monzilnepali/Golang-Todo/utils"
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
	var newUser user.User
	err := decoder.Decode(&newUser)
	if err != nil {
		panic(err.Error())
	}

	//validate email pattern
	emailError := formValidation.ValidateEmail(newUser.Email)
	if emailError != nil {
		http.Error(w, emailError.Error(), http.StatusBadRequest)
		return
	}
	//validate password strength
	passwordError := formValidation.ValidatePassword(newUser.Password)
	if passwordError != nil {
		http.Error(w, passwordError.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "hello from signup")
}

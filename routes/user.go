package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	userHandler "github.com/monzilnepali/Golang-Todo/handler"
	user "github.com/monzilnepali/Golang-Todo/model"
)

//http custom error

//Login handler
func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		fmt.Fprint(w, r.Method+r.URL.Path+" cannot be resolve")
	}
	//getting email and password from req.body
	decoder := json.NewDecoder(r.Body)
	var newUser user.User
	err := decoder.Decode(&newUser)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "something Went wrong", http.StatusInternalServerError)

	}

	if newUser.Email != "" && newUser.Password != "" {
		tokenString, loginErr := userHandler.LoginHandler(newUser)
		switch loginErr := loginErr.(type) {
		case *userHandler.HttpError:
			http.Error(w, loginErr.Message, loginErr.StatusCode)

		case nil:
			//signnup completed

			token := make(map[string]string)
			token["Token"] = tokenString
			json.NewEncoder(w).Encode(token)

		}
	} else {
		//empty email and password field
		http.Error(w, "Invalid email and password field", http.StatusBadRequest)

	}

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

	signupErr := userHandler.SignupHandler(newUser)
	switch signupErr := signupErr.(type) {
	case *userHandler.HttpError:
		http.Error(w, signupErr.Message, signupErr.StatusCode)

	case nil:
		//signnup completed
		//send success feedback
		fmt.Fprint(w, "sign up completed", http.StatusOK)
	}

}

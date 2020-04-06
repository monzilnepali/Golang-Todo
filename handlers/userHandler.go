package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/user"

	"github.com/julienschmidt/httprouter"
	userHandler "github.com/monzilnepali/Golang-Todo/domain"
)

//http custom error

//Login handler
func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	//getting email and password from req.body
	decoder := json.NewDecoder(r.Body)
	var newUser user.User
	err := decoder.Decode(&newUser)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "something Went wrong", http.StatusInternalServerError)

	}

	fmt.Println("newuserer", newUser)
	if newUser.Email != "" && newUser.Password != "" {
		tokenString, loginErr := userHandler.LoginHandler(newUser)
		switch loginErr := loginErr.(type) {
		case *userHandler.HTTPError:
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
		return

	}

}

//Signup Handler
func Signup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	//getting email and password from req.body
	decoder := json.NewDecoder(r.Body)
	var newUser user.User
	err := decoder.Decode(&newUser)
	if err != nil {
		panic(err.Error())
	}

	signupErr := userHandler.SignupHandler(newUser)
	switch signupErr := signupErr.(type) {
	case *userHandler.HTTPError:
		http.Error(w, signupErr.Message, signupErr.StatusCode)
		return

	case nil:
		//signnup completed
		//send success feedback
		fmt.Fprint(w, "sign up completed", http.StatusOK)
	}

}

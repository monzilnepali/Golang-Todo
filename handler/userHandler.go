package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/monzilnepali/Golang-Todo/db"
	user "github.com/monzilnepali/Golang-Todo/model"
	jwt "github.com/monzilnepali/Golang-Todo/services"
	formValidation "github.com/monzilnepali/Golang-Todo/utils"
)

//check exiting user

//add new user

//check email and pasword validation

//FindUserViaEmail handler
func findUserViaEmail(email string) (user.User, error) {

	var existingUser user.User
	err := db.DB.QueryRow("SELECT UserID, Email FROM user WHERE Email=?", email).Scan(&existingUser.UserID, &existingUser.Email)
	if err == sql.ErrNoRows {
		fmt.Println("user doesnot exist")
		return user.User{}, nil

	} else if err != nil {
		fmt.Println("something wrong happedn")
		panic(err.Error())
	}
	fmt.Println("user exist already")
	return existingUser, errors.New("user already exist")

}

//FindUserViaEmailAndPassword handler
func findUserViaEmailAndPassword(email, password string) (user.User, error) {

	var existingUser user.User
	err := db.DB.QueryRow("SELECT UserID, Email FROM user WHERE Email=? AND Password=?", email, password).Scan(&existingUser.UserID, &existingUser.Email)
	if err == sql.ErrNoRows {
		return user.User{}, errors.New("Email address and password doesnot match")

	} else if err != nil {
		fmt.Println("something wrong happend")
		panic(err.Error())
	}
	fmt.Println("email & password matched")
	return existingUser, nil

}

//SignupHandler handler
func SignupHandler(newUser user.User) error {

	//
	//validate email pattern
	emailError := formValidation.ValidateEmail(newUser.Email)
	if emailError != nil {
		return NewHTTPError(emailError.Error(), http.StatusBadRequest)
	}
	//validate password strength
	passwordError := formValidation.ValidatePassword(newUser.Password)
	if passwordError != nil {
		return NewHTTPError(passwordError.Error(), http.StatusBadRequest)
	}

	//check whether email address already exist or not
	_, err1 := findUserViaEmail(newUser.Email)
	if err1 != nil {
		return NewHTTPError(err1.Error(), http.StatusBadRequest)
	}

	//if not add new user
	err2 := addNewUser(newUser.Email, newUser.Password)
	if err2 != nil {
		return NewHTTPError(err2.Error(), http.StatusInternalServerError)
	}
	//successful signup
	return nil
}

func addNewUser(email, password string) error {

	// db.Exec() for db transaction that doesnot return any sql.row
	// db.Query() that return sql.row. sql.row reserve db connection whil row is closed.

	stmt, err := db.DB.Prepare("INSERT INTO user(Email,Password) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
		return errors.New("something went wrong")
	}
	_, err1 := stmt.Exec(email, password)
	return err1

}

//LoginHandler
func LoginHandler(user user.User) (string, error) {

	//!check the validate the password and email

	//validate email pattern
	emailError := formValidation.ValidateEmail(user.Email)
	if emailError != nil {
		return "", NewHTTPError(emailError.Error(), http.StatusBadRequest)
	}
	//validate password strength
	passwordError := formValidation.ValidatePassword(user.Password)
	if passwordError != nil {
		return "", NewHTTPError(passwordError.Error(), http.StatusBadRequest)
	}

	//!check entry in database
	existingUser, loginError := findUserViaEmailAndPassword(user.Email, user.Password)
	if loginError != nil {
		fmt.Println(loginError.Error())
		return "", NewHTTPError(loginError.Error(), http.StatusUnauthorized)

	}

	fmt.Println("user id", existingUser)
	//!generate JWT token
	tokenString, err := jwt.GenerateToken(existingUser.UserID)
	if err != nil {
		return "", NewHTTPError(err.Error(), http.StatusInternalServerError)

	}

	//!send response back
	return tokenString, nil

}

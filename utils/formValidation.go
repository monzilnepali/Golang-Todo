package utils

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	errInvalidEmailFormat = errors.New("Invalid Email format")
)

//ValidateEmail handler
func ValidateEmail(email string) error {

	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	fmt.Println("match", re.MatchString(email))
	if !re.MatchString(email) {
		return errInvalidEmailFormat
	}
	return nil

}

package utils

import (
	"errors"
	"regexp"
)

var (
	errInvalidEmailFormat = errors.New("Invalid Email format")
	errPassword           = errors.New("At least 6 character long")
)

//ValidateEmail handler
func ValidateEmail(email string) error {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !re.MatchString(email) {
		return errInvalidEmailFormat
	}
	return nil

}

//ValidatePassword handler
func ValidatePassword(password string) error {
	re := regexp.MustCompile(".{6,}")
	if !re.MatchString(password) {
		return errPassword
	}
	return nil

}

package model

//User model
type User struct {
	id       int
	Email    string `json.string`
	Password string `json.string`
}

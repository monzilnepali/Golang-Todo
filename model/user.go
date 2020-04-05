package model

//User model
type User struct {
	Id       int    `json.id`
	Email    string `json.email`
	Password string `json.password`
}

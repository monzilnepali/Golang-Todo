package model

//User model
type User struct {
	UserID   int    `json:"userid"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

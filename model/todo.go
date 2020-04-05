package model

//Todo model
type Todo struct {
	TodoID      int    `json.todoid`
	Title       string `json.title`
	Iscompleted bool   `json.isCompleted`
}

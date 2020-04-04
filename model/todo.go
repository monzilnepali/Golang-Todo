package model

//Todo model
type Todo struct{
	ID int `json.id`
	Title string `json.title`
	Iscompleted bool `json.isCompleted`
}

package models

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Sex      string `json:"sex"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}

package dto

type Body struct {
	Name string `form:"nama" json:"nama"`
	Age  int8   `form:"umur" json:"umur"`
}

type User struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

var Users []User

type LoginInput struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

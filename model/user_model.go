package model

type UserModel struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

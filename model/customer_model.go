package model

type CustomerModel struct {
	Id            string `json:"id"`
	Username      string `json:"username"`
	Fullname      string `json:"fullname"`
	AccountNumber string `json:"account_number"`
	Email         string `json:"email"`
	Address       string `json:"address"`
}

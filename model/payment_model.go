package model

type PaymentModel struct {
	Id                    string `json:"id"`
	CustomerId            string `json:"customer_id"`
	CustomerName          string `json:"customer_name"`
	AccountNumberCustomer string `json:"account_number_customer"`
	Amount                int    `json:"amount"`
	MerchantId            string `json:"merchant_id"`
	MerchantName          string `json:"merchant_name"`
	AccountNumberMerchant string `json:"account_number_merchant"`
	Date                  string `json:"date"`
}

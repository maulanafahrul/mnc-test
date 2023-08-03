package model

import "time"

type PaymentModel struct {
	Id           string
	CustomerId   string
	CustomerName string
	Amount       int
	MerchantId   string
	MerchantName string
	Date         time.Time
}

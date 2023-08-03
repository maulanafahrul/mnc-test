package repository

import (
	"encoding/json"
	"os"

	"github.com/maulanafahrul/mnc-test/model"
)

type PaymentRepository interface {
	Create(*model.PaymentModel) error
	GetPaymentById(string) *model.PaymentModel
}

type paymentRepositoryImpl struct {
	payments []model.PaymentModel
}

func (r *paymentRepositoryImpl) GetPaymentById(id string) *model.PaymentModel {

	for _, pay := range r.payments {
		if pay.Id == id {
			return &pay
		}
	}
	return nil
}

func (r *paymentRepositoryImpl) Create(payment *model.PaymentModel) error {
	// Check if the user exists in the slice
	found := false
	for i, u := range r.payments {
		if u.Id == payment.Id {
			r.payments[i] = *payment
			found = true
			break
		}
	}
	if !found {
		r.payments = append(r.payments, *payment)
	}

	// Open the JSON file
	file, err := os.OpenFile("data/payments.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(r.payments)
	if err != nil {
		return err
	}

	return nil
}

func NewPaymentRepository() PaymentRepository {
	repo := &paymentRepositoryImpl{}

	// Open the JSON file
	file, err := os.Open("data/payments.json")
	if err != nil {

		return nil
	}
	defer file.Close()

	// Decode the file into the users slice
	err = json.NewDecoder(file).Decode(&repo.payments)
	if err != nil {
		return nil
	}

	return repo
}

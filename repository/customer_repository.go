package repository

import (
	"encoding/json"
	"os"

	"github.com/maulanafahrul/mnc-test/model"
)

type CustomerRepository interface {
	GetCustomerByFullname(string) *model.CustomerModel
	GetCustomerByUsername(string) *model.CustomerModel
	Create(*model.CustomerModel) error
	Update(*model.CustomerModel) error
	Delete(*model.CustomerModel) error
}

type customerRepositoryImpl struct {
	customers []model.CustomerModel
}

func (r *customerRepositoryImpl) GetCustomerByFullname(fullname string) *model.CustomerModel {

	for _, cus := range r.customers {
		if cus.Fullname == fullname {
			return &cus
		}
	}
	return nil
}
func (r *customerRepositoryImpl) GetCustomerByUsername(username string) *model.CustomerModel {

	for _, cus := range r.customers {
		if cus.Username == username {
			return &cus
		}
	}
	return nil
}

func (r *customerRepositoryImpl) Create(customer *model.CustomerModel) error {
	// Check if the user exists in the slice
	found := false
	for i, u := range r.customers {
		if u.Id == customer.Id {
			r.customers[i] = *customer
			found = true
			break
		}
	}
	if !found {
		r.customers = append(r.customers, *customer)
	}

	// Open the JSON file
	file, err := os.OpenFile("data/customers.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(r.customers)
	if err != nil {
		return err
	}

	return nil
}

func (r *customerRepositoryImpl) Update(customer *model.CustomerModel) error {
	// Update the user in the slice
	for i, usr := range r.customers {
		if usr.Id == customer.Id {
			r.customers[i] = *customer
			break
		}
	}

	// Open the JSON file
	file, err := os.OpenFile("data/customers.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode the users slice back into the file
	err = json.NewEncoder(file).Encode(r.customers)
	if err != nil {
		return err
	}

	return nil
}

func (r *customerRepositoryImpl) Delete(customer *model.CustomerModel) error {
	// Remove the user from the slice
	for i, u := range r.customers {
		if u.Id == customer.Id {
			r.customers = append(r.customers[:i], r.customers[i+1:]...)
			break
		}
	}

	// Open the JSON file
	file, err := os.OpenFile("data/customers.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode the users slice back into the file
	err = json.NewEncoder(file).Encode(r.customers)
	if err != nil {
		return err
	}

	return nil
}

func NewCustomerRepository() CustomerRepository {
	repo := &customerRepositoryImpl{}

	// Open the JSON file
	file, err := os.Open("data/customers.json")
	if err != nil {
		return nil
	}
	defer file.Close()

	// Decode the file into the users slice
	err = json.NewDecoder(file).Decode(&repo.customers)
	if err != nil {
		return nil
	}

	return repo
}

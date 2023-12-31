package service

import (
	"fmt"

	"github.com/maulanafahrul/mnc-test/apperror"
	"github.com/maulanafahrul/mnc-test/model"
	"github.com/maulanafahrul/mnc-test/repository"
)

type CustomerService interface {
	GetCustomerByFullname(string) (*model.CustomerModel, error)
	InsertCustomer(*model.CustomerModel) error
	UpdateCustomer(*model.CustomerModel) error
	DeleteCustomer(string) error
}

type customerServiceImpl struct {
	customerRepo repository.CustomerRepository
}

func NewCustomerService(customerRepo repository.CustomerRepository) CustomerService {
	return &customerServiceImpl{
		customerRepo: customerRepo,
	}
}

func (cs *customerServiceImpl) GetCustomerByFullname(fullname string) (*model.CustomerModel, error) {
	cust := cs.customerRepo.GetCustomerByFullname(fullname)
	if cust == nil {
		return nil, &apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: "Data not found",
		}
	}
	return cust, nil
}

func (cs *customerServiceImpl) InsertCustomer(customer *model.CustomerModel) error {
	if len(customer.Fullname) < 5 {
		return &apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: "fullname at lease 5 character",
		}
	}
	if len(customer.AccountNumber) < 15 && len(customer.AccountNumber) > 20 {
		return &apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: "AccountNumber at lease 15 character or below 20 character",
		}
	}
	cust := cs.customerRepo.GetCustomerByFullname(customer.Fullname)
	if cust != nil {
		return &apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: fmt.Sprintf("customer data with the fullname %v already exists", customer.Fullname),
		}
	}
	existUser := cs.customerRepo.GetCustomerByUsername(customer.Username)
	if cust == nil && existUser == nil {
		return cs.customerRepo.Create(customer)
	}

	return &apperror.AppError{
		ErrorCode:    400,
		ErrorMassage: fmt.Sprintf("user with username : %v already has a customer account with name : %v", customer.Username, customer.Fullname),
	}

}

func (cs *customerServiceImpl) UpdateCustomer(newCustomer *model.CustomerModel) error {
	if len(newCustomer.Fullname) < 5 {
		return &apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: "fullname at lease 5 character",
		}
	}
	if len(newCustomer.AccountNumber) < 15 && len(newCustomer.AccountNumber) > 20 {
		return &apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: "AccountNumber at lease 15 character or below 20 character",
		}
	}

	existDataCus := cs.customerRepo.GetCustomerByFullname(newCustomer.Fullname)
	existUser := cs.customerRepo.GetCustomerByUsername(newCustomer.Username)
	if existUser == nil {
		return &apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: fmt.Sprintf("Customer data with the username %v not found", newCustomer.Username),
		}
	}
	if existUser != nil && existDataCus.Fullname == newCustomer.Fullname {
		return &apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: fmt.Sprintf("Customer data with the fullname %v already exists", newCustomer.Fullname),
		}
	}

	return cs.customerRepo.Update(newCustomer)
}

func (cs *customerServiceImpl) DeleteCustomer(fullname string) error {
	cust := cs.customerRepo.GetCustomerByFullname(fullname)
	if cust == nil {
		return &apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: "fullname doesn't exist",
		}
	}
	return cs.customerRepo.Delete(cust)
}

package service

import (
	"fmt"
	"time"

	"github.com/maulanafahrul/mnc-test/apperror"
	"github.com/maulanafahrul/mnc-test/model"
	"github.com/maulanafahrul/mnc-test/repository"
	"github.com/maulanafahrul/mnc-test/utils"
)

type PaymentService interface {
	GetPaymentById(string) (*model.PaymentModel, error)
	InsertPayment(*model.PaymentModel) error
}

type paymentServiceImpl struct {
	paymentRepo  repository.PaymentRepository
	customerRepo repository.CustomerRepository
}

func NewPaymentService(paymentRepo repository.PaymentRepository, customerRepo repository.CustomerRepository) PaymentService {
	return &paymentServiceImpl{
		paymentRepo:  paymentRepo,
		customerRepo: customerRepo,
	}
}

func (ps *paymentServiceImpl) InsertPayment(payment *model.PaymentModel) error {
	if payment.Amount <= 0 {
		return &apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: "amount cant less then 0 or minus",
		}
	}
	existCustomer := ps.customerRepo.GetCustomerByFullname(payment.CustomerName)
	if existCustomer == nil {
		return &apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: "Transfer only do for registered customer",
		}
	}

	existMerchant := ps.customerRepo.GetCustomerByFullname(payment.MerchantName)
	if existMerchant == nil {
		return &apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: fmt.Sprintf("customer data with the fullname %v not found", payment.MerchantName),
		}
	}

	payment.Id = utils.UuidGenerate()
	payment.CustomerId = existCustomer.Id
	payment.AccountNumberCustomer = existCustomer.AccountNumber
	payment.MerchantId = existMerchant.Id
	payment.AccountNumberMerchant = existMerchant.AccountNumber
	payment.Date = time.Now().Format("2006-01-02")
	fmt.Println(payment)
	return ps.paymentRepo.Create(payment)
}

func (ps *paymentServiceImpl) GetPaymentById(id string) (*model.PaymentModel, error) {
	pay := ps.paymentRepo.GetPaymentById(id)
	if pay == nil {
		return nil, &apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: "Data not found",
		}
	}
	return pay, nil
}

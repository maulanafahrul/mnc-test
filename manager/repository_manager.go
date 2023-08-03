package manager

import (
	"sync"

	"github.com/maulanafahrul/mnc-test/repository"
)

type RepositoryManager interface {
	GetUserRepo() repository.UserRepository
	GetCustomerRepo() repository.CustomerRepository
	GetPaymentRepo() repository.PaymentRepository
}

type repositoryManager struct {
	usrRepo     repository.UserRepository
	custRepo    repository.CustomerRepository
	paymentRepo repository.PaymentRepository
}

var onceLoadUserRepo sync.Once
var onceLoadCustomerRepo sync.Once
var onceLoadPaymentRepo sync.Once

func (rm *repositoryManager) GetUserRepo() repository.UserRepository {
	onceLoadUserRepo.Do(func() {
		rm.usrRepo = repository.NewUserRepository()
	})
	return rm.usrRepo
}
func (rm *repositoryManager) GetCustomerRepo() repository.CustomerRepository {
	onceLoadCustomerRepo.Do(func() {
		rm.custRepo = repository.NewCustomerRepository()
	})
	return rm.custRepo
}
func (rm *repositoryManager) GetPaymentRepo() repository.PaymentRepository {
	onceLoadPaymentRepo.Do(func() {
		rm.paymentRepo = repository.NewPaymentRepository()
	})
	return rm.paymentRepo
}

func NewRepoManager() RepositoryManager {
	return &repositoryManager{}
}

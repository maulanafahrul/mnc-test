package manager

import (
	"sync"

	"github.com/maulanafahrul/mnc-test/service"
)

type ServiceManager interface {
	GetUserService() service.UserService
	GetLoginService() service.LoginService
	GetCustomerService() service.CustomerService
	GetPaymentService() service.PaymentService
}

type serviceManager struct {
	repoManager RepositoryManager

	usrService   service.UserService
	custService  service.CustomerService
	payService   service.PaymentService
	loginService service.LoginService
}

func NewServiceManager(repoManager RepositoryManager) ServiceManager {
	return &serviceManager{
		repoManager: repoManager,
	}
}

var onceLoadUserService sync.Once
var onceLoadCustomerService sync.Once
var onceLoadPaymentService sync.Once
var onceLoadLoginService sync.Once

func (sm *serviceManager) GetUserService() service.UserService {
	onceLoadUserService.Do(func() {
		sm.usrService = service.NewUserService(sm.repoManager.GetUserRepo())
	})
	return sm.usrService
}
func (sm *serviceManager) GetCustomerService() service.CustomerService {
	onceLoadCustomerService.Do(func() {
		sm.custService = service.NewCustomerService(sm.repoManager.GetCustomerRepo())
	})
	return sm.custService
}
func (sm *serviceManager) GetPaymentService() service.PaymentService {
	onceLoadPaymentService.Do(func() {
		sm.payService = service.NewPaymentService(sm.repoManager.GetPaymentRepo(), sm.repoManager.GetCustomerRepo())
	})
	return sm.payService
}
func (sm *serviceManager) GetLoginService() service.LoginService {
	onceLoadLoginService.Do(func() {
		sm.loginService = service.NewLoginService(sm.repoManager.GetUserRepo())
	})
	return sm.loginService
}

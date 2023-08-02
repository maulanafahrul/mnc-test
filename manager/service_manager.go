package manager

import (
	"sync"

	"github.com/maulanafahrul/mnc-test/service"
)

type ServiceManager interface {
	GetUserService() service.UserService
	GetLoginService() service.LoginService
}

type serviceManager struct {
	repoManager RepositoryManager

	usrService   service.UserService
	loginService service.LoginService
}

func NewServiceManager(repoManager RepositoryManager) ServiceManager {
	return &serviceManager{
		repoManager: repoManager,
	}
}

var onceLoadUserService sync.Once
var onceLoadLoginService sync.Once

func (sm *serviceManager) GetUserService() service.UserService {
	onceLoadUserService.Do(func() {
		sm.usrService = service.NewUserService(sm.repoManager.GetUserRepo())
	})
	return sm.usrService
}
func (sm *serviceManager) GetLoginService() service.LoginService {
	onceLoadLoginService.Do(func() {
		sm.loginService = service.NewLoginService(sm.repoManager.GetUserRepo())
	})
	return sm.loginService
}

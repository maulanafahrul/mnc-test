package manager

import (
	"sync"

	"github.com/maulanafahrul/mnc-test/repository"
)

type RepositoryManager interface {
	GetUserRepo() repository.UserRepository
}

type repositoryManager struct {
	usrRepo repository.UserRepository
}

var onceLoadUserRepo sync.Once

func (rm *repositoryManager) GetUserRepo() repository.UserRepository {
	onceLoadUserRepo.Do(func() {
		rm.usrRepo = repository.NewUserRepository()
	})
	return rm.usrRepo
}

func NewRepoManager() RepositoryManager {
	return &repositoryManager{}
}

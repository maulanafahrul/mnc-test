package service

import (
	"fmt"

	"github.com/maulanafahrul/mnc-test/apperror"
	"github.com/maulanafahrul/mnc-test/model"
	"github.com/maulanafahrul/mnc-test/repository"
	"github.com/maulanafahrul/mnc-test/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	InsertUser(*model.UserModel) error
	DeleteUser(string) error
}

type userServiceImpl struct {
	usrRepo repository.UserRepository
}

func NewUserService(usrRepo repository.UserRepository) UserService {
	return &userServiceImpl{
		usrRepo: usrRepo,
	}
}

func (us *userServiceImpl) InsertUser(usr *model.UserModel) error {
	if len(usr.Username) < 5 {
		return &apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: "username at lease 5 character",
		}
	}
	if len(usr.Password) < 5 {
		return &apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: "username at lease 5 character",
		}
	}
	user := us.usrRepo.GetUserByUsername(usr.Username)
	if user != nil {
		return &apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: fmt.Sprintf("User data with the name %v already exists", usr.Username),
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	usr.Id = utils.UuidGenerate()
	usr.Password = string(hashedPassword)
	return us.usrRepo.Create(usr)
}

func (us *userServiceImpl) DeleteUser(username string) error {
	user := us.usrRepo.GetUserByUsername(username)
	if user == nil {
		return &apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: "username doesn't exist",
		}
	}
	return us.usrRepo.Delete(user)
}

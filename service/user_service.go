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
	GetUserByUsername(string) (*model.UserModel, error)
	InsertUser(*model.UserModel) error
	DeleteUser(string) error
	UpdateUser(*model.UserModel) error
}

type userServiceImpl struct {
	usrRepo repository.UserRepository
}

func NewUserService(usrRepo repository.UserRepository) UserService {
	return &userServiceImpl{
		usrRepo: usrRepo,
	}
}

func (us *userServiceImpl) GetUserByUsername(username string) (*model.UserModel, error) {
	usr := us.usrRepo.GetUserByUsernameForView(username)
	if usr == nil {
		return nil, &apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: "Data not found",
		}
	}
	return usr, nil
}

func (us *userServiceImpl) InsertUser(usr *model.UserModel) error {
	if len(usr.Username) <= 5 {
		return &apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: "username at lease 5 character",
		}
	}
	if len(usr.Password) <= 5 {
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

func (us *userServiceImpl) UpdateUser(newUser *model.UserModel) error {
	if len(newUser.Username) <= 5 {
		return &apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: "username at lease 5 character",
		}
	}
	if len(newUser.Password) <= 5 {
		return &apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: "username at lease 5 character",
		}
	}

	existDataUsr := us.usrRepo.GetUserByUsername(newUser.Username)

	if existDataUsr != nil && existDataUsr.Id != newUser.Id {
		return &apperror.AppError{
			ErrorCode:    1,
			ErrorMassage: fmt.Sprintf("User data with the username %v already exists", newUser.Username),
		}
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("userUsecaseImpl.GenerateFromPassword(): %w", err)
	}
	newUser.Password = string(passHash)

	return us.usrRepo.Update(newUser)
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

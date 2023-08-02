package service

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/maulanafahrul/mnc-test/apperror"
	"github.com/maulanafahrul/mnc-test/model"
	"github.com/maulanafahrul/mnc-test/repository"
	"golang.org/x/crypto/bcrypt"
)

type LoginService interface {
	Login(*model.LoginModel, *gin.Context) (*model.UserModel, error)
	Logout(*gin.Context)
}

type loginServiceImpl struct {
	userRepo repository.UserRepository
}

func NewLoginService(userRepo repository.UserRepository) LoginService {
	return &loginServiceImpl{
		userRepo: userRepo,
	}
}

func (ls *loginServiceImpl) Login(usr *model.LoginModel, c *gin.Context) (*model.UserModel, error) {
	session := sessions.Default(c)
	// validate session exist
	existSession := session.Get("Username")
	if existSession != nil {
		return nil, &apperror.AppError{
			ErrorCode:    1,
			ErrorMassage: fmt.Sprintf("You are already logged in as %v", existSession),
		}
	}

	existData := ls.userRepo.GetUserByUsername(usr.Username)
	if existData == nil {
		return nil, &apperror.AppError{
			ErrorCode:    1,
			ErrorMassage: "Username is not registered",
		}
	}

	err := bcrypt.CompareHashAndPassword([]byte(existData.Password), []byte(usr.Password))
	if err != nil {
		return nil, &apperror.AppError{
			ErrorCode:    1,
			ErrorMassage: "Password does not match",
		}
	}

	// Login session
	session.Set("Username", existData.Username)
	session.Save()

	existData.Password = ""
	return existData, nil
}

func (ls *loginServiceImpl) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

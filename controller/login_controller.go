package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maulanafahrul/mnc-test/apperror"
	"github.com/maulanafahrul/mnc-test/model"
	"github.com/maulanafahrul/mnc-test/service"
	"github.com/maulanafahrul/mnc-test/utils"
)

type LoginController struct {
	loginService service.LoginService
}

func NewLoginController(srv *gin.Engine, loginService service.LoginService) *LoginController {
	controller := &LoginController{
		loginService: loginService,
	}
	srv.POST("/login", controller.Login)
	srv.POST("/logout")

	return controller
}

func (lc LoginController) Login(c *gin.Context) {
	loginReq := &model.LoginModel{}
	err := c.ShouldBindJSON(&loginReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}
	if loginReq.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "username cant be empty",
		})
		return
	}
	if loginReq.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "password cant be empty",
		})
		return
	}

	err = lc.loginService.Login(loginReq, c)

	if err != nil {
		appError := &apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("loginService.Login() 1: %v", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("loginService.Login() 2: %v", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred during login",
				"errorCode":    err.Error(),
			})
		}
		return
	}
	tokenJwt, err := utils.GenerateToken(loginReq.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid Token",
		})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    tokenJwt,
		HttpOnly: true,
	})
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": tokenJwt,
	})
}

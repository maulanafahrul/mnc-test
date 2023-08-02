package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maulanafahrul/mnc-test/apperror"
	"github.com/maulanafahrul/mnc-test/middleware"
	"github.com/maulanafahrul/mnc-test/model"
	"github.com/maulanafahrul/mnc-test/service"
)

type UserController struct {
	usrService service.UserService
}

func NewUserConroller(srv *gin.Engine, usrService service.UserService) *UserController {
	controller := &UserController{
		usrService: usrService,
	}
	srv.POST("/user", controller.AddUser)
	srv.DELETE("/user/:username", middleware.RequireToken(), controller.DeleteUser)
	return controller
}

func (uc UserController) AddUser(c *gin.Context) {
	usr := &model.UserModel{}
	err := c.ShouldBindJSON(&usr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}
	// validate
	if usr.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "username cant be empty",
		})
		return
	}
	if usr.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "password cant be empty",
		})
		return
	}

	err = uc.usrService.InsertUser(usr)
	if err != nil {
		appError := &apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("usrService.InsertUser() 1: %v", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("usrService.InsertUser() 2: %v", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while saving user data",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "success insert data user",
	})
}

func (uc UserController) DeleteUser(c *gin.Context) {
	username := c.Param("username")

	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "username cant be empty",
		})
		return
	}

	if err := uc.usrService.DeleteUser(username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User deleted successfully",
	})
}

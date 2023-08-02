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
	srv.GET("/user/:username", middleware.RequireToken(), controller.GetUserByUsername)
	srv.POST("/user", controller.AddUser)
	srv.PUT("/user", middleware.RequireToken(), controller.UpdateUser)
	srv.DELETE("/user/:username", middleware.RequireToken(), controller.DeleteUser)
	return controller
}

func (uc UserController) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	usr, err := uc.usrService.GetUserByUsername(username)
	if err != nil {
		appError := &apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("usrService.GetUserByUsername() 1: %v", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("usrService.GetUserByUsername() 2: %v", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while fetching user data",
			})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    usr,
	})
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

func (uc UserController) UpdateUser(c *gin.Context) {
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

	id, _ := c.Get("id")
	userID := id.(string)

	usr.Id = userID

	err = uc.usrService.UpdateUser(usr)
	if err != nil {
		appError := &apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("usrService.UpdateUser() 1: %v", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("usrService.UpdateUser() 2: %v", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while updating user data",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "success update user data",
	})
}

func (uc UserController) DeleteUser(c *gin.Context) {
	username := c.Param("username")

	if err := uc.usrService.DeleteUser(username); err != nil {
		appError := &apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("usrService.DeleteUser() 1: %v", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("usrService.DeleteUser() 2: %v", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while delete user data",
			})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User deleted successfully",
	})
}

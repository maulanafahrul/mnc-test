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

type CustomerController struct {
	customerService service.CustomerService
}

func NewCustomerController(srv *gin.Engine, customerService service.CustomerService) *CustomerController {
	controller := &CustomerController{
		customerService: customerService,
	}
	srv.GET("/customer/:fullname", middleware.RequireToken(), controller.GetCustomerByFullname)
	srv.POST("/customer", middleware.RequireToken(), controller.AddCustomer)
	srv.PUT("/customer", middleware.RequireToken(), controller.UpdateCustomer)
	srv.DELETE("/customer/:fullname", middleware.RequireToken(), controller.DeleteCustomer)
	return controller
}

func (cc CustomerController) GetCustomerByFullname(c *gin.Context) {
	fullname := c.Param("fullname")
	usr, err := cc.customerService.GetCustomerByFullname(fullname)
	if err != nil {
		appError := &apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("customerService.GetCustomerByFullname() 1: %v", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("customerService.GetCustomerByFullname() 2: %v", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while fetching customer data",
			})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    usr,
	})
}

func (cc CustomerController) AddCustomer(c *gin.Context) {
	cust := &model.CustomerModel{}
	err := c.ShouldBindJSON(&cust)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}
	// validate
	if cust.Fullname == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "fullname cant be empty",
		})
		return
	}
	if cust.AccountNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "accountNumber cant be empty",
		})
		return
	}
	if cust.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "email cant be empty",
		})
		return
	}
	if cust.Address == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "address cant be empty",
		})
		return
	}
	username, _ := c.Get("username")
	usernameStr := username.(string)
	cust.Username = usernameStr
	id, _ := c.Get("id")
	customerId := id.(string)
	cust.Id = customerId

	err = cc.customerService.InsertCustomer(cust)
	if err != nil {
		appError := &apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("customerService.InsertCustomer() 1: %v", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("customerService.InsertCustomer() 2: %v", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while saving customer data",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "success insert customer data",
	})
}

func (cc CustomerController) UpdateCustomer(c *gin.Context) {
	Newcust := &model.CustomerModel{}
	err := c.ShouldBindJSON(&Newcust)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}
	// validate
	if Newcust.Fullname == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "fullname cant be empty",
		})
		return
	}
	if Newcust.AccountNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "accountNumber cant be empty",
		})
		return
	}
	if Newcust.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "email cant be empty",
		})
		return
	}
	if Newcust.Address == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "address cant be empty",
		})
		return
	}

	id, _ := c.Get("id")
	customerId := id.(string)
	Newcust.Id = customerId

	err = cc.customerService.UpdateCustomer(Newcust)
	if err != nil {
		appError := &apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("customerService.UpdateCustomer() 1: %v", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("customerService.UpdateCustomer() 2: %v", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while updating customer data",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "success update customer data",
	})
}

func (cc CustomerController) DeleteCustomer(c *gin.Context) {
	fullname := c.Param("fullname")

	if err := cc.customerService.DeleteCustomer(fullname); err != nil {
		appError := &apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("customerService.DeleteCustomer() 1: %v", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("customerService.DeleteCustomer() 2: %v", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while delete customer data",
			})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "customer deleted successfully",
	})
}

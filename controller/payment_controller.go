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

type PaymentController struct {
	paymentService service.PaymentService
}

func NewPaymentController(srv *gin.Engine, paymentService service.PaymentService) *PaymentController {
	controller := &PaymentController{
		paymentService: paymentService,
	}
	srv.GET("/payment/:id", middleware.RequireToken(), controller.GetPaymentById)
	srv.POST("/payment", middleware.RequireToken(), controller.AddPayment)
	return controller
}

func (pc PaymentController) GetPaymentById(c *gin.Context) {
	id := c.Param("id")
	pay, err := pc.paymentService.GetPaymentById(id)
	if err != nil {
		appError := &apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("paymentService.GetPaymentById() 1: %v", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("paymentService.GetPaymentById() 2: %v", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while fetching customer data",
			})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    pay,
	})
}

func (pc PaymentController) AddPayment(c *gin.Context) {
	pay := &model.PaymentModel{}
	err := c.ShouldBindJSON(&pay)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}
	// validate
	if pay.CustomerName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "customer name cant be empty",
		})
		return
	}
	if pay.MerchantName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "merchant name cant be empty",
		})
		return
	}
	err = pc.paymentService.InsertPayment(pay)
	if err != nil {
		appError := &apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("paymentService.InsertPayment() 1: %v", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("paymentService.InsertPayment() 2: %v", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while saving payment data",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "success insert payment data",
	})
}

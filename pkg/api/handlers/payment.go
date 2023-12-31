package handlers

import (
	services "Zhooze/pkg/usecase/interface"
	"Zhooze/pkg/utils/response"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	PaymentUseCase services.PaymentUseCase
	OrderUseCase   services.OrderUseCase
}

func NeWPaymentHandler(useCase services.PaymentUseCase, orderUsecase services.OrderUseCase) *PaymentHandler {
	return &PaymentHandler{
		PaymentUseCase: useCase,
		OrderUseCase:   orderUsecase,
	}
}

func (py *PaymentHandler) MakePaymentRazorPay(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Query("order_id"))
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "error from orderID", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	paymentMethodID, err := py.OrderUseCase.PaymentMethodID(orderID)
	if err != nil {
		err := response.ClientResponse(http.StatusInternalServerError, "error from paymentId ", nil, err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if paymentMethodID == 2 {
		payment, _ := py.PaymentUseCase.PaymentAlreadyPaid(orderID)
		if payment {
			c.HTML(http.StatusOK, "pay.html", nil)
			return

		}
		orderDetail, razorID, err := py.PaymentUseCase.MakePaymentRazorPay(orderID)
		if err != nil {
			errs := response.ClientResponse(http.StatusInternalServerError, "could not generate order details", nil, err.Error())
			c.JSON(http.StatusInternalServerError, errs)
			return
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"final_price": orderDetail.FinalPrice * 100,
			"razor_id":    razorID,
			"order_id":    orderDetail.OrderId,
			"user_name":   orderDetail.Firstname,
			"total":       int(orderDetail.FinalPrice),
		})
	}
	c.HTML(http.StatusNotFound, "notfound.html", nil)
}
func (py *PaymentHandler) VerifyPayment(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Query("order_id"))
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "error from orderID", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	paymentID := c.Query("payment_id")
	err = py.PaymentUseCase.SavePaymentDetails(orderID, paymentID)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "could not update payment details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully updated payment details", nil, nil)
	c.JSON(http.StatusOK, success)
}

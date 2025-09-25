package handlers

import (
	"net/http"

	"github.com/mahdifr17/CheckoutService/internals/dtos"
	"github.com/mahdifr17/CheckoutService/internals/services"
	"github.com/mahdifr17/CheckoutService/internals/services/appmodel"

	"github.com/gin-gonic/gin"
)

type CheckoutHandler struct {
	checkoutService services.CheckoutService
}

func NewCheckoutHandler(checkoutService services.CheckoutService) *CheckoutHandler {
	return &CheckoutHandler{
		checkoutService: checkoutService,
	}
}

func (h *CheckoutHandler) GetProducts(c *gin.Context) {
	products, err := h.checkoutService.GetAllProducts(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *CheckoutHandler) Checkout(c *gin.Context) {
	var req dtos.CheckoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// [ ] validate request

	response, err := h.checkoutService.ProcessCheckout(
		c, appmodel.CheckoutCommand{},
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

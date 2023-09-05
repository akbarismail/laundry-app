package api

import (
	"clean-code/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	customerUseCase usecase.CustomerUseCase
	routerGroup     *gin.RouterGroup
}

func (c *CustomerController) findByPhoneCustomer(ctx *gin.Context) {
	phone := ctx.Param("phone")

	c2, err := c.customerUseCase.GetByPhone(phone)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "successfully",
		"data":    c2,
	})
}

func (c *CustomerController) findCustomers(ctx *gin.Context) {
	c2, err := c.customerUseCase.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "successfully",
		"data":    c2,
	})
}

func (c *CustomerController) Route() {
	c.routerGroup.GET("/customers/:phone", c.findByPhoneCustomer)
	c.routerGroup.GET("/customers", c.findCustomers)
}

func NewCustomerController(customerUseCase usecase.CustomerUseCase, routerGroup *gin.RouterGroup) *CustomerController {
	return &CustomerController{
		customerUseCase: customerUseCase,
		routerGroup:     routerGroup,
	}
}

package api

import (
	"clean-code/model/dto"
	"clean-code/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authUseCase usecase.AuthUseCase
	userUseCase usecase.UserUseCase
	rg          *gin.RouterGroup
}

func (a *AuthController) login(c *gin.Context) {
	var authReq dto.AuthRequest
	if err := c.ShouldBindJSON(&authReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	authRes, err := a.authUseCase.Login(authReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully login",
		"data":    authRes,
	})
}

func (a *AuthController) register(c *gin.Context) {
	var authReq dto.AuthRequest
	if err := c.ShouldBindJSON(&authReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := a.userUseCase.Register(authReq); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully register",
	})
}

func (a *AuthController) Route() {
	a.rg.POST("/auth/login", a.login)
	a.rg.POST("/auth/register", a.register)
}

func NewAuthController(authUseCase usecase.AuthUseCase, userUseCase usecase.UserUseCase, rg *gin.RouterGroup) *AuthController {
	return &AuthController{
		authUseCase: authUseCase,
		userUseCase: userUseCase,
		rg:          rg,
	}
}

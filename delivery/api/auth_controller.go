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

// AuthController godoc
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        Body body dto.AuthRequest  true  "Auth login"
// @Success      201  {object}  dto.AuthRequest
// @Router       /auth/login [post]
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

	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully login",
		"data":    authRes,
	})
}

// AuthController godoc
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        Body body dto.AuthRequest  true  "Auth register"
// @Success      201  {object}  string
// @Router       /auth/register [post]
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

	c.JSON(http.StatusCreated, gin.H{
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

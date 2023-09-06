package middleware

import (
	"clean-code/util/security"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var authHeader authHeader
		if err := ctx.ShouldBindHeader(&authHeader); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": fmt.Sprintf("unauthorized: %v", err.Error()),
			})
			return
		}

		token := strings.Replace(authHeader.AuthorizationHeader, "Bearer ", "", 1)
		fmt.Println("1st", token)
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			return
		}

		claims, err := security.VerifyJwtToken(token)
		fmt.Println("2st", claims)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": fmt.Sprintf("unauthorized: %v", err.Error()),
			})
			return
		}

		ctx.Set("claims", claims)
		ctx.Next()
	}
}

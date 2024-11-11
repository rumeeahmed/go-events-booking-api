package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-events-booking-api/utils"
	"net/http"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	// abort the request and do not allow any subsequent middleware or request handlers to be executed.
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	context.Set("userId", userId)

	// Ensure the next request handler/middleware executes.
	context.Next()
}

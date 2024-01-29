package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("APP_TOKEN"))

func AuthJWTProtected() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization Header"})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		bearerToken := authHeader[7:]
		token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Expired Token"})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse claims"})
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		data, _ := claims["data"].(map[string]interface{})
		user, _ := data["user"].(map[string]interface{})

		username := user["Username"]

		fmt.Println(username)
		ctx.Next()
	}
}

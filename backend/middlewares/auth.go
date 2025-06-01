package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ducktordanny/cubeit/backend/configs"
	"github.com/ducktordanny/cubeit/backend/types"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func UserAuthSessionMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		secret := configs.Envs.JWTSecret
		auth := context.GetHeader("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Missing token",
			})
			return
		}
		tokenString := strings.TrimPrefix(auth, "Bearer ")

		claims := &types.AuthClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return secret, nil
		})
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}
		if !token.Valid {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		context.Set("auth", claims)
		context.Next()
	}
}

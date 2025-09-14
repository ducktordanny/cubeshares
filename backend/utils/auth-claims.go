package utils

import (
	"net/http"

	"github.com/ducktordanny/cubeshares/backend/types"
	"github.com/gin-gonic/gin"
)

func GetAuthClaims(context *gin.Context) *types.AuthClaims {
	claims, ok := context.Get("auth")
	if !ok {
		handleNotFoundClaims(context)
		return nil
	}
	typed, ok := claims.(*types.AuthClaims)
	if !ok {
		handleNotFoundClaims(context)
		return nil
	}
	return typed
}

func handleNotFoundClaims(context *gin.Context) {
	context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Couldn't fetch or process auth claims"})
}

package user

import (
	"net/http"

	"github.com/ducktordanny/cubeshares/backend/middlewares"
	"github.com/ducktordanny/cubeshares/backend/types"
	"github.com/ducktordanny/cubeshares/backend/utils"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (handler *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("user/me", middlewares.UserAuthSessionMiddleware(), handler.handleUserMe)
}

func (handler *Handler) handleUserMe(context *gin.Context) {
	claims := utils.GetAuthClaims(context)
	user, err := handler.store.GetUserById(claims.Sub)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Couldn't find authenticated user"})
		return
	}
	context.IndentedJSON(http.StatusOK, user)
}

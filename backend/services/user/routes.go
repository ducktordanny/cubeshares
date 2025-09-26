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
	router.PUT("user/me/bio", middlewares.UserAuthSessionMiddleware(), handler.handleUserMeBio)
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

func (handler *Handler) handleUserMeBio(context *gin.Context) {
	claims := utils.GetAuthClaims(context)
	var request types.UpdateUserBioRequestBody
	if err := context.ShouldBindJSON(&request); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	err := handler.store.UpdateUserBio(claims.Sub, request.Bio)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update user bio: " + err.Error()})
		return
	}
	context.Status(http.StatusNoContent)
}

package user

import (
	"net/http"
	"strconv"

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
	router.GET("user/:id", handler.handleUserById)
}

func (handler *Handler) handleUserMe(context *gin.Context) {
	claims := utils.ReadAuthClaims(context)
	user, err := handler.store.ReadUserById(claims.UserId)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Couldn't find authenticated user"})
		return
	}
	context.IndentedJSON(http.StatusOK, user)
}

func (handler *Handler) handleUserMeBio(context *gin.Context) {
	claims := utils.ReadAuthClaims(context)
	var request types.UpdateUserBioRequestBody
	if err := context.ShouldBindJSON(&request); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	err := handler.store.UpdateUserBio(claims.UserId, request.Bio)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update user bio: " + err.Error()})
		return
	}
	context.Status(http.StatusNoContent)
}

func (handler *Handler) handleUserById(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Couldn't parse provided user ID"})
	}
	user, err := handler.store.ReadUserById(id)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Couldn't find user by provided ID"})
		return
	}
	context.IndentedJSON(http.StatusOK, user)
}

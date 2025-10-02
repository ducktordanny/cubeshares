package posts

import (
	"net/http"
	"strconv"

	"github.com/ducktordanny/cubeshares/backend/middlewares"
	"github.com/ducktordanny/cubeshares/backend/types"
	"github.com/ducktordanny/cubeshares/backend/utils"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	store types.PostsStore
}

func NewHandler(store types.PostsStore) *Handler {
	return &Handler{store: store}
}

func (handler *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("post", handler.handleReadPost)
	router.GET("post/:id", handler.handleReadPostById)
	router.GET("posts", handler.handleReadPostList)
	router.POST("post", middlewares.UserAuthSessionMiddleware(), handler.handleCreatePost)
}

func (handler *Handler) handleReadPost(context *gin.Context) {
	context.IndentedJSON(http.StatusNotImplemented, gin.H{"message": "[GET] /api/v1/post WIP"})
}

func (handler *Handler) handleReadPostById(context *gin.Context) {
	context.IndentedJSON(http.StatusNotImplemented, gin.H{"message": "[GET] /api/v1/post/:id WIP"})
}

func (handler *Handler) handleReadPostList(context *gin.Context) {
	userIdString := context.Query("userId")
	if userIdString == "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing userId"})
		return
	}
	userId, err := strconv.ParseInt(userIdString, 10, 64)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid userId"})
		return
	}
	response, err := handler.store.ReadPostsOfUser(context.Request.Context(), userId)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Could not read posts: " + err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, response)
}

func (handler *Handler) handleCreatePost(context *gin.Context) {
	claims := utils.ReadAuthClaims(context)
	var request types.CreatePostRequestBody
	if err := context.ShouldBindJSON(&request); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}
	response, err := handler.store.CreateNewPost(context.Request.Context(), claims.UserId, request)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Could not create post: " + err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, response)
}

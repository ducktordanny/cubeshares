package posts

import (
	"net/http"

	"github.com/ducktordanny/cubeshares/backend/middlewares"
	"github.com/ducktordanny/cubeshares/backend/types"
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
	context.IndentedJSON(http.StatusNotImplemented, gin.H{"message": "[GET] /api/v1/posts WIP"})
}

func (handler *Handler) handleCreatePost(context *gin.Context) {
	context.IndentedJSON(http.StatusNotImplemented, gin.H{"message": "[POST] /api/v1/post WIP"})
	var request types.CreatePostRequestBody
	if err := context.ShouldBindJSON(&request); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body" + err.Error()})
		return
	}
	context.IndentedJSON(http.StatusNotImplemented, gin.H{"message": "[POST] /api/v1/post WIP", "request": request})
}

package user

import (
	"github.com/ducktordanny/cubeit/backend/types"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (handler *Handler) RegisterRoutes(router *gin.RouterGroup) {}

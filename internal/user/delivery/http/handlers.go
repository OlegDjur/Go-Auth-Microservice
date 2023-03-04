package http

import "github.com/gin-gonic/gin"

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) createUser(c *gin.Context) {
}

func (h *UserHandler) loginUser(c *gin.Context) {
}

package http

import (
	"net/http"

	"github.com/OlegDjur/Go-Auth-Microservice/internal/user"
	"github.com/OlegDjur/Go-Auth-Microservice/internal/user/dto"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service user.Service
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) createUser(c *gin.Context) {
	var user dto.CreateUserRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func (h *UserHandler) loginUser(c *gin.Context) {
}

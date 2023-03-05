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

func NewUserHandler() *handler {
	return &handler{}
}

func (h *handler) createUser(c *gin.Context) {
	var req dto.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := dto.NewUserResponse(user)

	c.JSON(http.StatusOK, res)
}

func (h *handler) loginUser(c *gin.Context) {
	var req *dto.LoginUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, accessToken, err := h.service.Login(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := dto.NewLoginUserResponse(user, accessToken)

	c.JSON(http.StatusOK, res)
}

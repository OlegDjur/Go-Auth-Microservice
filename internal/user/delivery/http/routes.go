package http

import "github.com/gin-gonic/gin"

func (h *handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/user", h.createUser)
	router.POST("/user/login", h.loginUser)

	return router
}

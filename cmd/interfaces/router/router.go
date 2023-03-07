package router

import (
	"github.com/gin-gonic/gin"
	"micro-user-management/cmd/interfaces/handler"
)

func SetUpRoutes(router *gin.Engine, userHandler *handler.UserHandler) {
	router.POST("/users/register", userHandler.Register)
	router.POST("/users/login", userHandler.Login)
	router.GET("/users/:id", userHandler.GetUser)
}

package router

import (
	"let-you-cook/handler"

	"github.com/gin-gonic/gin"
)

func AuthRoute(r *gin.RouterGroup, authHandler *handler.AuthHandler) {

	authRoute := r.Group("/auth")
	{
		authRoute.POST("/register", authHandler.Register)
		authRoute.POST("/login", authHandler.Login)
	}
}

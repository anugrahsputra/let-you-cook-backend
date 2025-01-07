package router

import (
	"let-you-cook/handler"
	"let-you-cook/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.RouterGroup, userHandler *handler.UserHandler) {
	userRoutes := r.Group("/users", middleware.AuthMiddleware())
	{
		userRoutes.GET("", userHandler.GetUsers)
		userRoutes.GET("/:id", userHandler.GetUserByID)
	}
}

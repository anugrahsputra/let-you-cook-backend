package router

import (
	"let-you-cook/handler"
	"let-you-cook/middleware"

	"github.com/gin-gonic/gin"
)

func SessionRoute(r *gin.RouterGroup, sessionHandler *handler.SessionHandler) {
	sessionRoutes := r.Group("/session", middleware.AuthMiddleware())
	{
		sessionRoutes.POST("", sessionHandler.CreateSession)
		sessionRoutes.POST("/start/:id", sessionHandler.StartSession)
		sessionRoutes.POST("/end/:id", sessionHandler.EndSession)
	}
}

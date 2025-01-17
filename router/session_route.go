package router

import (
	"let-you-cook/handler"
	"let-you-cook/middleware"

	"github.com/gin-gonic/gin"
)

func SessionRoute(r *gin.RouterGroup, sessionHandler *handler.SessionHandler) {
	sessionRoutes := r.Group("/session", middleware.AuthMiddleware())
	{
		sessionRoutes.GET("", sessionHandler.GetAllSessions)
		sessionRoutes.POST("", sessionHandler.CreateSession)
		sessionRoutes.PATCH("/:id", sessionHandler.UpdateSession)
	}
}

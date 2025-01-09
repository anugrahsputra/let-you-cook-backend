package router

import (
	"let-you-cook/handler"
	"let-you-cook/middleware"

	"github.com/gin-gonic/gin"
)

func ProfileRoute(r *gin.RouterGroup, profileHandler *handler.ProfileHandler) {
	profileRoutes := r.Group("/profile", middleware.AuthMiddleware())
	{
		profileRoutes.POST("", profileHandler.CreateProfile)
		profileRoutes.GET("", profileHandler.GetProfileByAccountID)
		profileRoutes.PATCH("", profileHandler.UpdateProfile)
	}
}

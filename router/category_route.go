package router

import (
	"let-you-cook/handler"
	"let-you-cook/middleware"

	"github.com/gin-gonic/gin"
)

func CategoryRoute(r *gin.RouterGroup, categoryHandler *handler.CategoryHandler) {
	categroyRoutes := r.Group("/category", middleware.AuthMiddleware())
	{
		categroyRoutes.POST("", categoryHandler.CreateCategory)
		categroyRoutes.GET("", categoryHandler.GetCategories)
		categroyRoutes.GET("/:id", categoryHandler.GetCategoryById)
		categroyRoutes.PATCH("/:id", categoryHandler.UpdateCategory)
		categroyRoutes.DELETE("/:id", categoryHandler.DeleteCategory)
	}
}

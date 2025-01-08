package router

import (
	"let-you-cook/handler"
	"let-you-cook/middleware"

	"github.com/gin-gonic/gin"
)

func TaskRoute(r *gin.RouterGroup, taskHandler *handler.TaskHandler) {
	taskRoutes := r.Group("/tasks", middleware.AuthMiddleware())
	{
		taskRoutes.POST("", taskHandler.CreateTask)
		taskRoutes.GET("", taskHandler.GetTasks)
	}
}

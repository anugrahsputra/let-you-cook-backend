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
		taskRoutes.GET("/category", taskHandler.GetTaskGroupedByCategory)
		taskRoutes.PATCH("/:id", taskHandler.UpdateTask)
		taskRoutes.DELETE("/:id", taskHandler.DeleteTask)
	}
}

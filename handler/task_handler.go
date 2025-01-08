package handler

import (
	"let-you-cook/domain/dto"
	"let-you-cook/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskService service.ITaskService
}

func NewTaskHandler(taskHandler service.ITaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskHandler,
	}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	userId := c.MustGet("user_id").(string)
	var reqTask dto.ReqTask

	if err := c.ShouldBindJSON(&reqTask); err != nil {
		c.JSON(http.StatusBadRequest, dto.Resp{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	if err := h.taskService.CreateTask(userId, reqTask); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Resp{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, dto.Resp{
		Status:  http.StatusOK,
		Message: "task created success",
		Data:    nil,
	})
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	userId := c.MustGet("user_id").(string)

	tasks, err := h.taskService.GetTasks(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Resp{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, dto.Resp{
		Status:  http.StatusOK,
		Message: "success",
		Data:    tasks,
	})

}

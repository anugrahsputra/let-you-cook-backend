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

func NewTaskHandler(taskService service.ITaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
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

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	userId := c.MustGet("user_id").(string)

	if id == "" {
		c.JSON(http.StatusBadRequest, dto.Resp{
			Status:  http.StatusBadRequest,
			Message: "id is required",
			Data:    nil,
		})
		return
	}

	var reqTask map[string]interface{}

	if err := c.ShouldBindJSON(&reqTask); err != nil {
		c.JSON(http.StatusBadRequest, dto.Resp{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	task, err := h.taskService.UpdateTask(id, userId, reqTask)
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
		Message: "task updated success",
		Data:    task,
	})
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	userId := c.MustGet("user_id").(string)

	if id == "" {
		c.JSON(http.StatusBadRequest, dto.Resp{
			Status:  http.StatusBadRequest,
			Message: "id is required",
		})
		return
	}

	_, err := h.taskService.DeleteTask(id, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Resp{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Resp{
		Status:  http.StatusOK,
		Message: "task deleted success",
	})
}

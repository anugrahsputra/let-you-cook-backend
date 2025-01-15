package handler

import (
	"let-you-cook/domain/dto"
	"let-you-cook/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	USER_ID = "user_id"
)

type SessionHandler struct {
	sessionService service.ISessionService
}

func NewSessionHandler(sessionService service.ISessionService) *SessionHandler {
	return &SessionHandler{
		sessionService: sessionService,
	}
}

func (h *SessionHandler) CreateSession(c *gin.Context) {
	userId := c.MustGet(USER_ID).(string)
	var req dto.ReqCreateSession

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Resp{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	if req.Name == "" {
		c.JSON(http.StatusBadRequest, dto.Resp{
			Status:  http.StatusBadRequest,
			Message: "session name cannot be empty",
			Data:    nil,
		})
		return
	}
	if req.FocusDuration <= 0 || req.BreakDuration <= 0 {
		c.JSON(http.StatusBadRequest, dto.Resp{
			Status:  http.StatusBadRequest,
			Message: "focus duration and break duration should be greater than 0",
			Data:    nil,
		})
		return
	}

	if err := h.sessionService.CreateSession(userId, req); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Resp{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, dto.Resp{
		Status:  http.StatusOK,
		Message: "session created succes",
		Data:    nil,
	})
}

func (h *SessionHandler) StartSession(c *gin.Context) {
	userId := c.MustGet(USER_ID).(string)
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, dto.Resp{
			Status:  http.StatusBadRequest,
			Message: "id is required",
			Data:    nil,
		})
		return
	}

	err := h.sessionService.StartSession(id, userId)
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
		Message: "session started",
		Data:    nil,
	})
}

func (h *SessionHandler) EndSession(c *gin.Context) {
	userId := c.MustGet(USER_ID).(string)
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, dto.Resp{
			Status:  http.StatusBadRequest,
			Message: "id is required",
			Data:    nil,
		})
		return
	}

	err := h.sessionService.EndSession(id, userId)
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
		Message: "session ended",
		Data:    nil,
	})
}

func (h *SessionHandler) GetAllSessions(c *gin.Context) {
	userId := c.MustGet(USER_ID).(string)

	sessions, err := h.sessionService.GetAllSessions(userId)

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
		Data:    sessions,
	})
}

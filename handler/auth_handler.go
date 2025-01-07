package handler

import (
	"let-you-cook/domain/dto"
	"let-you-cook/domain/model"
	"let-you-cook/service"
	"let-you-cook/utils/validator"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthHandler struct {
	authService service.IAuthService
}

func NewAuthHandler(authService service.IAuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var reqUser dto.ReqRegister
	if err := c.ShouldBindJSON(&reqUser); err != nil {
		c.JSON(400, dto.Resp{
			Status:  400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	validationError := validator.ValidateStruct(reqUser)
	if validationError != nil {
		c.JSON(400, dto.Resp{
			Status:  400,
			Message: "invalid request body",
			Data:    nil,
		})
		return
	}

	user := model.User{
		Id:        uuid.New().String(),
		Username:  reqUser.Username,
		Password:  reqUser.Password,
		Email:     reqUser.Email,
		CreatedAt: int(time.Now().Unix()),
		UpdatedAt: int(time.Now().Unix()),
	}

	if err := h.authService.Register(user); err != nil {
		c.JSON(400, dto.Resp{
			Status:  400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(200, dto.Resp{
		Status:  200,
		Message: "Register Success",
		Data:    user,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.ReqLogin
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, dto.Resp{
			Status:  400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	validationError := validator.ValidateStruct(req)
	if validationError != nil {
		c.JSON(400, dto.Resp{
			Status:  400,
			Message: "invalid request body",
			Data:    nil,
		})
		return
	}

	token, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(400, dto.Resp{
			Status:  400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.Header("Authorization", token)

	c.JSON(200, dto.Resp{
		Status:  200,
		Message: "Login Success",
		Data:    nil,
	})

}

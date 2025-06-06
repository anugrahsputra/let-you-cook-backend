package handler

import (
	"let-you-cook/domain/dto"
	"let-you-cook/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("main")

type ProfileHandler struct {
	profileService service.IProfileService
}

func NewProfileHanlder(profileService service.IProfileService) *ProfileHandler {
	return &ProfileHandler{
		profileService: profileService,
	}
}

func (h *ProfileHandler) CreateProfile(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	email := c.MustGet("email").(string)

	var reqProfile dto.ReqProfile
	err := c.ShouldBindJSON(&reqProfile)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Resp{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	err = h.profileService.CreateProfile(userID, email, reqProfile)
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
		Data:    nil,
	})
}

func (h *ProfileHandler) GetProfileByAccountID(c *gin.Context) {
	userID := c.MustGet("user_id").(string)

	profile, err := h.profileService.GetProfileByAccountId(userID)
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
		Data:    profile,
	})
}

func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	userID := c.MustGet("user_id").(string)

	var reqProfile dto.ReqPatchProfile
	if err := c.ShouldBindJSON(&reqProfile); err != nil {
		c.JSON(http.StatusBadRequest, dto.Resp{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	profile, err := h.profileService.UpdateProfile(userID, reqProfile)
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
		Message: "profile updated success",
		Data:    profile,
	})
}

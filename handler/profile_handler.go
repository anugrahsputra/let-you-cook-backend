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
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	err = h.profileService.CreateProfile(userID, email, reqProfile)
	if err != nil {
		c.JSON(http.StatusConflict, dto.ErrorResponse{
			Status:  http.StatusConflict,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    nil,
	})
}

func (h *ProfileHandler) GetProfileByAccountID(c *gin.Context) {
	userID := c.MustGet("user_id").(string)

	profile, err := h.profileService.GetProfileByAccountId(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    profile,
	})
}

func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	userID := c.MustGet("user_id").(string)

	var reqProfile dto.ReqPatchProfile
	if err := c.ShouldBindJSON(&reqProfile); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	profile, err := h.profileService.UpdateProfile(userID, reqProfile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "profile updated success",
		Data:    profile,
	})
}

func (h *ProfileHandler) UploadProfilePicture(c *gin.Context) {
	userID := c.MustGet("user_id").(string)

	file, err := c.FormFile("photo_profile")
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "photo_profile is required",
		})
		return
	}

	// Check MIME type
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "failed to open file",
		})
		return
	}
	defer src.Close()

	buffer := make([]byte, 512)
	_, err = src.Read(buffer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "failed to read file",
		})
		return
	}

	contentType := http.DetectContentType(buffer)
	if contentType != "image/jpeg" && contentType != "image/png" {
		c.JSON(http.StatusUnsupportedMediaType, dto.ErrorResponse{
			Status:  http.StatusUnsupportedMediaType,
			Message: "unsupported file type. only jpeg and png are allowed",
		})
		return
	}

	profile, err := h.profileService.UploadProfilePicture(userID, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    profile,
	})
}

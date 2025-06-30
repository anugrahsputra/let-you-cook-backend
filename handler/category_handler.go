package handler

import (
	"let-you-cook/domain/dto"
	"let-you-cook/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryService service.ICategoryService
}

func NewCategoryHandler(categoryService service.ICategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	userId := c.MustGet("user_id").(string)
	var reqCategory dto.ReqCategory

	if err := c.ShouldBindJSON(&reqCategory); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	if reqCategory.Name == "" {
		c.JSON(http.StatusBadRequest, dto.Response{
			Status:  http.StatusBadRequest,
			Message: "category name cannot be empty",
			Data:    nil,
		})
		return
	}

	if err := h.categoryService.CreateCategory(userId, reqCategory); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "category created succes",
		Data:    nil,
	})
}

func (h CategoryHandler) GetCategories(c *gin.Context) {
	userId := c.MustGet("user_id").(string)
	id := c.Query("id")
	name := c.Query("name")

	req := dto.ReqCategory{
		Id:   id,
		Name: name,
	}

	categories, err := h.categoryService.GetCategories(userId, req)
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
		Message: "category created succes",
		Data:    categories,
	})
}

func (h *CategoryHandler) GetCategoryById(c *gin.Context) {
	userId := c.MustGet("user_id").(string)
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, dto.Response{
			Status:  http.StatusBadRequest,
			Message: "id is required",
			Data:    nil,
		})
		return
	}

	category, err := h.categoryService.GetCategoryById(id, userId)
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
		Message: "success",
		Data:    category,
	})
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	userId := c.MustGet("user_id").(string)
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, dto.Response{
			Status:  http.StatusBadRequest,
			Message: "id is required",
			Data:    nil,
		})
		return
	}

	var reqUpdateCategory dto.ReqPatchCategory
	if err := c.ShouldBindJSON(&reqUpdateCategory); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	category, err := h.categoryService.UpdateCategory(id, userId, reqUpdateCategory)
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
		Message: "category updated success",
		Data:    category,
	})

}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	userId := c.MustGet("user_id").(string)
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, dto.Response{
			Status:  http.StatusBadRequest,
			Message: "id is required",
			Data:    nil,
		})
		return
	}

	if _, err := h.categoryService.DeleteCategory(id, userId); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "category deleted success",
		Data:    nil,
	})

}

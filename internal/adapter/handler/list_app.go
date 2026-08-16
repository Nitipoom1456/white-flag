package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/white-flag/internal/adapter/dto"
	"github.com/white-flag/internal/adapter/response"
)

func (h *Handler) ListApps(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)
	apps, total, err := h.AppUsecase.List(c, pageInt, pageSizeInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:        response.CODE_INTERNAL_SERVER_ERROR,
			Description: response.MSG_INTERNAL_SERVER_ERROR,
		})
		return
	}

	result := make([]dto.AppResponse, 0, len(apps))
	for _, app := range apps {
		result = append(result, dto.AppResponse{
			ID:        app.ID,
			Name:      app.Name,
			CreatedAt: app.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, response.PaginationResponse[dto.AppResponse]{
		Page:     pageInt,
		PageSize: pageSizeInt,
		Total:    total,
		Data:     result,
	})
}

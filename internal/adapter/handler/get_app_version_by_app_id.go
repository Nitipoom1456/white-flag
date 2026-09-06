package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/white-flag/internal/adapter/dto"
	"github.com/white-flag/internal/adapter/response"
)

func (h *Handler) GetAppVersionByAppID(c *gin.Context) {
	appID := c.Param("id")
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)

	appVersions, total, usecaseErr := h.AppVersionUsecase.GetByAppID(c, appID, pageInt, pageSizeInt)
	if usecaseErr != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:        usecaseErr.Code,
			Description: usecaseErr.Message,
		})
		return
	}

	appVersionResponses := make([]dto.AppVersionResponse, len(appVersions))
	for i, appVersion := range appVersions {
		appVersionResponses[i] = dto.AppVersionResponse{
			ID:      appVersion.ID,
			Version: appVersion.Version,
			AppID:   appVersion.AppID,
		}
	}

	c.JSON(http.StatusOK, response.PaginationResponse[dto.AppVersionResponse]{
		Page:     pageInt,
		PageSize: pageSizeInt,
		Total:    total,
		Data:     appVersionResponses,
	})
}

package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/white-flag/internal/adapter/dto"
	"github.com/white-flag/internal/adapter/response"
	"github.com/white-flag/internal/infrastructure/logger"
)

func (h *Handler) GetAppFeatureFlagByEnvironmentID(c *gin.Context) {
	appID := c.Param("id")
	environmentID := c.Param("environmentID")
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)

	var query dto.GetAppFeatureFlagByEnvironmentIDQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		logger.Error(c, "invalid query", err)
		c.JSON(http.StatusBadRequest, response.ValidationError(err))
		return
	}

	featureFlagSettings, total, usecaseErr := h.FeatureFlagSettingUsecase.GetFeatureFlagSettingsByAppIDAndEnvironmentID(c, appID, environmentID, pageInt, pageSizeInt)
	if usecaseErr != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:        "internal_server_error",
			Description: usecaseErr.Message,
		})
		return
	}

	featureFlagResponses := make([]dto.FeatureFlagResponse, len(featureFlagSettings))
	for i, setting := range featureFlagSettings {
		featureFlagResponses[i] = dto.FeatureFlagResponse{
			ID:            setting.ID,
			Name:          setting.Name,
			Description:   setting.Description,
			Active:        setting.Active,
			Tags:          setting.Tags,
			MinAppVersion: setting.MinAppVersionID,
			MaxAppVersion: setting.MaxAppVersionID,
			CreatedAt:     setting.CreatedAt,
			UpdatedAt:     setting.UpdatedAt,
		}
	}

	c.JSON(http.StatusOK, response.PaginationResponse[dto.FeatureFlagResponse]{
		Page:     pageInt,
		PageSize: pageSizeInt,
		Total:    total,
		Data:     featureFlagResponses,
	})
}

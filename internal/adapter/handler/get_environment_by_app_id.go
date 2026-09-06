package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/white-flag/internal/adapter/dto"
	"github.com/white-flag/internal/adapter/response"
)

func (h *Handler) GetEnvironmentByAppID(c *gin.Context) {
	appID := c.Param("id")

	environments, usecaseErr := h.EnvironmentUsecase.GetByAppID(c, appID)
	if usecaseErr != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:        usecaseErr.Code,
			Description: usecaseErr.Message,
		})
		return
	}

	result := make([]dto.EnvironmentResponse, len(environments))
	for i, env := range environments {
		result[i] = dto.EnvironmentResponse{
			ID:    env.ID,
			Name:  env.Name,
			AppID: env.AppID,
		}
	}

	c.JSON(http.StatusOK, response.SuccessDataResponse[[]dto.EnvironmentResponse]{
		Message: "success",
		Data:    result,
	})
}

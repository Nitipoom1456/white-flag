package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/white-flag/internal/adapter/dto"
	"github.com/white-flag/internal/adapter/response"
	"github.com/white-flag/internal/infrastructure/logger"
)

func (h *Handler) CreateEnvironment(c *gin.Context) {
	appID := c.Param("id")

	var req dto.CreateEnvironmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error(c, "invalid request", err)
		c.JSON(http.StatusBadRequest, response.ValidationError(err))
		return
	}

	usecaseErr := h.EnvironmentUsecase.Create(c, appID, req.Name)
	if usecaseErr != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:        usecaseErr.Code,
			Description: usecaseErr.Message,
		})
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Message: "environment created successfully",
	})
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/white-flag/internal/adapter/dto"
	"github.com/white-flag/internal/adapter/response"
	"github.com/white-flag/internal/infrastructure/logger"
)

func (h *Handler) CreateAppVersion(c *gin.Context) {
	var req dto.CreateAppVersionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error(c, "invalid request", err)
		c.JSON(http.StatusBadRequest, response.ValidationError(err))
		return
	}

	usecaseErr := h.AppVersionUsecase.Create(c, req.AppID, req.Version)
	if usecaseErr != nil {
		c.JSON(http.StatusInternalServerError, response.ValidateErrorResponse{
			Code:        usecaseErr.Code,
			Description: usecaseErr.Message,
		})
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Message: "app created successfully",
	})
}

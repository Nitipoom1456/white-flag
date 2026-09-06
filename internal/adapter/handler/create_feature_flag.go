package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/white-flag/internal/adapter/dto"
	"github.com/white-flag/internal/adapter/response"
	"github.com/white-flag/internal/infrastructure/logger"
	featureflag "github.com/white-flag/internal/usecase/feature_flag"
)

func (h *Handler) CreateFeatureFlag(c *gin.Context) {
	var req dto.CreateFeatureFlagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error(c, "invalid request", err)
		c.JSON(http.StatusBadRequest, response.ValidationError(err))
		return
	}

	usecaseErr := h.FeatureFlagUsecase.Create(c, featureflag.CreateFeatureFlagParams{
		AppID:         req.AppID,
		Name:          req.Name,
		Description:   req.Description,
		Active:        req.Active,
		Tags:          req.Tags,
		MinAppVersion: req.MinAppVersion,
		MaxAppVersion: req.MaxAppVersion,
	})
	if usecaseErr != nil {
		if usecaseErr.Code == "environment_not_found" {
			c.JSON(http.StatusNotFound, response.ErrorResponse{
				Code:        usecaseErr.Code,
				Description: usecaseErr.Message,
			})
			return
		}
		if usecaseErr.Code == "duplicate_key" {
			c.JSON(http.StatusConflict, response.ErrorResponse{
				Code:        usecaseErr.Code,
				Description: usecaseErr.Message,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:        usecaseErr.Code,
			Description: usecaseErr.Message,
		})
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Message: "feature flag created successfully",
	})
}

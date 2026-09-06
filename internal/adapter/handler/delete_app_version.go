package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/white-flag/internal/adapter/response"
)

func (h *Handler) DeleteAppVersion(c *gin.Context) {
	id := c.Param("id")
	usecaseErr := h.AppVersionUsecase.Delete(c, id)
	if usecaseErr != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:        usecaseErr.Code,
			Description: usecaseErr.Message,
		})
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Message: "app version deleted successfully",
	})
}

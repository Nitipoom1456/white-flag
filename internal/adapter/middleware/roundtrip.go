package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/white-flag/internal/infrastructure/logger"
)

const requestIDHeader = "X-Request-ID"

func Roundtrip() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		requestID := c.GetHeader(requestIDHeader)
		if requestID == "" {
			uuid, _ := uuid.NewV7()
			requestID = uuid.String()
		}
		c.Writer.Header().Set(requestIDHeader, requestID)

		ctx := logger.With(c.Request.Context(), logger.F("request_id", requestID))
		c.Request = c.Request.WithContext(ctx)

		path := c.Request.URL.Path
		if raw := c.Request.URL.RawQuery; raw != "" {
			path += "?" + raw
		}

		c.Next()

		fields := []logger.Field{
			logger.F("method", c.Request.Method),
			logger.F("path", path),
			logger.F("status", c.Writer.Status()),
			logger.F("latency_ms", time.Since(start).Milliseconds()),
			logger.F("client_ip", c.ClientIP()),
			logger.F("bytes", c.Writer.Size()),
		}

		if err := c.Errors.Last(); err != nil {
			logger.Error(ctx, "request", err, fields...)
			return
		}
		logger.Info(ctx, "request", fields...)
	}
}

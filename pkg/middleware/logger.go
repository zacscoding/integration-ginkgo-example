package middleware

import (
	"github.com/gin-gonic/gin"
	"integration-ginkgo-example/pkg/logging"
)

func NewLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := logging.WithLogger(c.Request.Context(), logging.DefaultLogger())
		c.Request = c.Request.WithContext(ctx)
	}
}

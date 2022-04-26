package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func (m Middleward) Cost() gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()
		c.Next()
		m.statsd.ReportHTTPCost(c.Request.Method, c.Request.RequestURI, time.Since(now))
	}
}

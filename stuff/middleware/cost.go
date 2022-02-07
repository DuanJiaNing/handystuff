package middleware

import (
	"handystuff/common/statsd"
	"time"

	"github.com/gin-gonic/gin"
)

func Cost() gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()
		c.Next()
		statsd.ReportHTTPCost(c.Request.Method, c.Request.RequestURI, time.Since(now))
	}
}

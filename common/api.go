package common

import (
	"handystuff/common/log"

	"github.com/gin-gonic/gin"
)

func AbortWithError(c *gin.Context, code int, err error) {
	nerr := c.AbortWithError(code, err)
	if nerr != nil {
		log.Printlnf("Error when AbortWithError, err: %v", nerr)
	}
}

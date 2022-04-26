package api

import (
	"github.com/gin-gonic/gin"
	xhttp "handystuff/common/http"
	"math/rand"
	"net/http"
	"time"
)

func (h Handler) GenerateToken(c *gin.Context) {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const length = 8

	rand.Seed(time.Now().UnixNano())
	chs := make([]byte, 0, length)
	for i := 0; i < length; i++ {
		chs = append(chs, chars[rand.Intn(len(chars))])
	}

	h.logger.Debugf(string(chs))
	c.Data(http.StatusOK, xhttp.PlainTextContentType, chs)
}

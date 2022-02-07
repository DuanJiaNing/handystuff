package tools

import (
	xhttp "handystuff/common/http"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	chars  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	length = 32
)

func GenerateToken(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	chs := make([]byte, 0, length)
	for i := 0; i < length; i++ {
		chs = append(chs, chars[rand.Intn(len(chars))])
	}

	log.Println(string(chs))
	c.Data(http.StatusOK, xhttp.PlainTextContentType, chs)
}

package api

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Handler struct {
	logger LogClient
	statsd StatsdClient

	decryptService DecryptService
}

// DecryptService 定义了加解密服务。
type DecryptService interface {
	// IsAESKeySupport 检查指定 aesKeyName 是否被支持。
	IsAESKeySupport(aesKeyName string) bool

	// DecryptAESString 解密 aes 字符串。
	DecryptAESString(str, aesKeyName string) ([]byte, error)

	// EncryptAESString 用 aes 加密字符串。
	EncryptAESString(str, aesKeyName string) ([]byte, error)
}

// StatsdClient 定义 statsd 客户端。
type StatsdClient interface {
	// ReportHTTPCost 报告指定 api 耗时。
	ReportHTTPCost(method, uri string, dur time.Duration)
}

type LogClient interface {
	Debugf(format string, a ...interface{})
	Infof(format string, a ...interface{})
	Warnf(format string, a ...interface{})
}

func NewHandler(
	logger LogClient,
	statsd StatsdClient,
	decryptService DecryptService,
) Handler {
	return Handler{
		logger:         logger,
		statsd:         statsd,
		decryptService: decryptService,
	}
}

func (h Handler) AbortWithError(c *gin.Context, code int, err error) {
	nerr := c.AbortWithError(code, err)
	if nerr != nil {
		h.logger.Warnf("Error when AbortWithError, err: %v", nerr)
	}
}

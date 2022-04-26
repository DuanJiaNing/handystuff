package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	xhttp "handystuff/common/http"
	"net/http"
)

func (h Handler) Encrypt(c *gin.Context) {
	app := c.Param("app")
	if h.decryptService.IsAESKeySupport(app) {
		h.AbortWithError(c, http.StatusBadRequest, errors.New("aes key not found"))
		return
	}

	bs, err := c.GetRawData()
	if err != nil {
		h.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	raw := string(bs)
	res, err := h.decryptService.EncryptAESString(raw, app)
	if err != nil {
		h.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	c.Data(http.StatusOK, xhttp.PlainTextContentType, res)
}

func (h Handler) Decrypt(c *gin.Context) {
	app := c.Param("app")
	if h.decryptService.IsAESKeySupport(app) {
		h.AbortWithError(c, http.StatusBadRequest, errors.New("aes key not found"))
		return
	}

	bs, err := c.GetRawData()
	if err != nil {
		h.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	raw := string(bs)
	res, err := h.decryptService.DecryptAESString(raw, app)
	if err != nil {
		h.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	c.Data(http.StatusOK, xhttp.JSONContentType, res)
}

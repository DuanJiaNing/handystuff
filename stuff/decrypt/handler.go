package decrypt

import (
	"encoding/base64"
	"errors"
	"handystuff/common"
	"handystuff/common/http"
	"handystuff/config"
	"handystuff/stuff/decrypt/internal"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Encrypt(c *gin.Context) {
	app := c.Param("app")
	if _, ok := config.Conf.AESCrypt.Keys[app]; !ok {
		common.AbortWithError(c, http.StatusBadRequest, errors.New("aes key not found"))
		return
	}

	bs, err := c.GetRawData()
	if err != nil {
		common.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	raw := string(bs)
	res, err := encrypt(raw, app)
	if err != nil {
		common.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	c.Data(http.StatusOK, xhttp.PlainTextContentType, res)
}

func Decrypt(c *gin.Context) {
	app := c.Param("app")
	if _, ok := config.Conf.AESCrypt.Keys[app]; !ok {
		common.AbortWithError(c, http.StatusBadRequest, errors.New("aes key not found"))
		return
	}

	bs, err := c.GetRawData()
	if err != nil {
		common.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	raw := string(bs)
	res, err := decrypt(raw, app)
	if err != nil {
		common.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	c.Data(http.StatusOK, xhttp.JSONContentType, res)
}

func decrypt(str, appName string) ([]byte, error) {
	cbcCipher, err := internal.NewAesCrypter(config.Conf.AESCrypt.Keys[appName], internal.Ecb, nil, "")
	if err != nil {
		return nil, err
	}
	plain, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}
	res, err := cbcCipher.Decrypt(string(plain))
	if err != nil {
		return nil, err
	}

	return res.RByte, nil
}

func encrypt(str, appName string) ([]byte, error) {
	encodedStr := base64.StdEncoding.EncodeToString([]byte(str))
	cbcCipher, err := internal.NewAesCrypter(config.Conf.AESCrypt.Keys[appName], internal.Ecb, nil, "")
	if err != nil {
		return nil, err
	}
	res, err := cbcCipher.Encrypt(encodedStr)
	if err != nil {
		return nil, err
	}

	return res.RByte, nil
}

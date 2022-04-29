package decrypt

import (
	"encoding/base64"
	"handystuff/config"
	"handystuff/stuff/decrypt/internal"
)

type Stuff struct {
	conf config.AESCryptConfig
}

func NewStuff(conf config.AESCryptConfig) Stuff {
	return Stuff{conf: conf}
}

func (s Stuff) IsAESKeySupport(aesKeyName string) bool {
	_, ok := s.conf.Keys[aesKeyName]
	return ok
}

func (s Stuff) DecryptAESString(str, aesKeyName string) ([]byte, error) {
	cbcCipher, err := internal.NewAesCrypter(s.conf.Keys[aesKeyName], internal.Ecb, nil, "")
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

func (s Stuff) EncryptAESString(str, aesKeyName string) ([]byte, error) {
	encodedStr := base64.StdEncoding.EncodeToString([]byte(str))
	cbcCipher, err := internal.NewAesCrypter(s.conf.Keys[aesKeyName], internal.Ecb, nil, "")
	if err != nil {
		return nil, err
	}
	res, err := cbcCipher.Encrypt(encodedStr)
	if err != nil {
		return nil, err
	}

	return res.RByte, nil
}

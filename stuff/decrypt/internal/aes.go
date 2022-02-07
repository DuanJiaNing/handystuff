package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"runtime"
)

type aesOption string

var Cbc aesOption = "cbc"
var Ecb aesOption = "ecb"

type aesCodingType string

var AesBase64 aesCodingType = "base64"
var AesHex aesCodingType = "hex"

type aesCrypter struct {
	key       string
	aesOption aesOption
	iv        []byte
	cType     aesCodingType
}

type aesCryptRes struct {
	RByte   []byte
	RBase64 string
	RHex    string
}

func NewAesCrypter(key string, option aesOption, iv []byte, cType aesCodingType) (*aesCrypter, error) {

	return &aesCrypter{
		key:       key,
		aesOption: option,
		iv:        iv,
		cType:     cType,
	}, nil
}

func (a *aesCrypter) cbcEncrypt(plain string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(a.key))
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				fmt.Println("runtime err:", err, "Check that the key is correct")
			default:
				fmt.Println("error:", err)
			}
		}
	}()

	blockSize := block.BlockSize()

	rawData := PKCS5Padding([]byte(plain), blockSize)
	cipherText := make([]byte, blockSize+len(rawData))

	var iv []byte

	if len(a.iv) != 0 {
		if len(a.iv) != 16 {
			return nil, errors.New("The length of iv should be 16 ")
		} else {
			iv = a.iv
		}
	} else {
		iv = cipherText[:blockSize]
	}
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[blockSize:], rawData)

	return cipherText, nil
}

func (a *aesCrypter) cbcDecrypt(plain string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(a.key))
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()

	encryptData := []byte(plain)
	if len(encryptData) < blockSize {
		return nil, errors.New("Cipher text too short ")
	}

	var iv []byte

	if len(a.iv) != 0 {
		if len(a.iv) != 16 {
			return nil, errors.New("The length of iv should be 16 ")
		} else {
			iv = a.iv
		}
	} else {
		iv = encryptData[:blockSize]
	}

	encryptData = encryptData[blockSize:]

	if len(encryptData)%blockSize != 0 {
		return nil, errors.New("Cipher text is not a multiple of the block size ")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(encryptData, encryptData)
	plainText, err := PKCS5UnPadding(encryptData)

	if err != nil {
		return nil, err
	}
	return plainText, nil
}

func (a *aesCrypter) ecbEncrypt(plain string) ([]byte, error) {
	block, _ := aes.NewCipher([]byte(a.key))
	data := []byte(plain)

	data = PKCS5Padding(data, block.BlockSize())
	encrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Encrypt(encrypted[bs:be], data[bs:be])
	}

	return encrypted, nil
}

func (a *aesCrypter) ecbDecrypt(plain string) ([]byte, error) {
	block, _ := aes.NewCipher([]byte(a.key))
	data := []byte(plain)
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Decrypt(decrypted[bs:be], data[bs:be])
	}

	cipherText, err := PKCS5UnPadding(decrypted)
	if err != nil {
		return nil, errors.New("Unpadding failed. ")
	}

	return cipherText, nil
}

func (a *aesCrypter) Encrypt(plain string) (*aesCryptRes, error) {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				fmt.Println("runtime err:", err, "Check that the input is correct")
			default:
				fmt.Println("error:", err)
			}
		}
	}()

	res := aesCryptRes{}

	var enByte []byte
	var err error

	switch a.aesOption {
	case Cbc:
		enByte, err = a.cbcEncrypt(plain)
	case Ecb:
		enByte, err = a.ecbEncrypt(plain)
	default:
		return nil, errors.New("The type of aes option is not valid. ")
	}

	if err != nil {
		return nil, err
	}

	switch a.cType {

	case AesBase64:
		str := base64.StdEncoding.EncodeToString(enByte)
		if err != nil {
			return nil, err
		}
		res.RBase64 = str
	case AesHex:
		str := hex.EncodeToString(enByte)
		if err != nil {
			return nil, err
		}
		res.RHex = str
	default:
		res.RByte = enByte
	}

	return &res, nil

}

func (a *aesCrypter) Decrypt(plain string) (*aesCryptRes, error) {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				fmt.Println("runtime err:", err, "Check that the input is correct")
			default:
				fmt.Println("error:", err)
			}
		}
	}()

	res := aesCryptRes{}

	var deByte []byte
	var err error

	switch a.aesOption {
	case Cbc:
		deByte, err = a.cbcDecrypt(plain)
	case Ecb:
		deByte, err = a.ecbDecrypt(plain)
	default:
		return nil, errors.New("The type of aes option is not valid. ")

	}
	if err != nil {
		return nil, fmt.Errorf("AES %s decrypt failure: %w", a.aesOption, err)
	}

	switch a.cType {
	case AesBase64:
		str := base64.StdEncoding.EncodeToString(deByte)
		if err != nil {
			return nil, err
		}
		res.RBase64 = str
	case AesHex:
		str := hex.EncodeToString(deByte)
		if err != nil {
			return nil, err
		}
		res.RHex = str
	default:
		res.RByte = deByte
	}

	return &res, nil
}

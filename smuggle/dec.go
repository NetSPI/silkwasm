package smuggle

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func pkcs5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

func aesDecrypt(key string, buf string) ([]byte, error) {

	encKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}

	encBuf, err := base64.StdEncoding.DecodeString(buf)
	if err != nil {
		return nil, err
	}

	var block cipher.Block

	block, err = aes.NewCipher(encKey)
	if err != nil {
		return nil, err
	}

	if len(encBuf) < aes.BlockSize {

		return nil, nil
	}
	iv := encBuf[:aes.BlockSize]
	encBuf = encBuf[aes.BlockSize:]

	cbc := cipher.NewCBCDecrypter(block, iv)
	cbc.CryptBlocks(encBuf, encBuf)
	decBuf := pkcs5Trimming(encBuf)

	return decBuf, nil

}

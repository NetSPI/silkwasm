package smuggle

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"io"
	"math/big"
)

func Sum256(msg []byte) ([]byte, error) {

	msgHash := sha256.New()
	_, err := msgHash.Write(msg)
	if err != nil {
		return nil, err
	}

	msgHashSum := msgHash.Sum(nil)
	return msgHashSum, nil

}

func GenKey(Size int) ([]byte, error) {
	key := make([]byte, Size)
	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func RandInt(min int, max int) (int, error) {
	max64 := big.NewInt(int64(max))
	outbig, err := rand.Int(rand.Reader, max64)
	if err != nil {
		return 0, err

	}

	out64 := outbig.Int64()
	return int(out64), nil
}

// Crypto padding
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// Encrypt and Decrypt
func AESCBCEncrypt(key, plainbytes []byte) ([]byte, error) {

	if len(plainbytes)%aes.BlockSize != 0 {
		plainbytes = PKCS5Padding(plainbytes, aes.BlockSize)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	cipherbytes := make([]byte, aes.BlockSize+len(plainbytes))
	iv := make([]byte, aes.BlockSize)
	_, err = rand.Read(iv)
	if err != nil {
		return nil, err

	}

	copy(cipherbytes[:aes.BlockSize], iv)

	cbc := cipher.NewCBCEncrypter(block, iv)
	cbc.CryptBlocks(cipherbytes[aes.BlockSize:], plainbytes)

	return cipherbytes, nil
}

// Generate an AES Key for use with encryption/decryption.
func (s *Smuggler) NewAESKey() ([]byte, error) {
	var err error
	s.Key, err = GenKey(32)
	return s.Key, err

}

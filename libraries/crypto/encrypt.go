package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"

	"github.com/donghquinn/blog_back_go/configs"
)

// 암호화 #요청
func EncryptString(rawString string) (string, error) {
	globalConfig := configs.GlobalConfig

	key := []byte(globalConfig.AesKey)
	plaintext := []byte(rawString)

	block, err := aes.NewCipher(key)

	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, len(plaintext))
	iv := []byte(globalConfig.AesIv)
	
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCBCEncrypter(block, iv)
	stream.CryptBlocks(ciphertext, plaintext)
	// stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// Encode the ciphertext in base64 to make it easier to handle as a string
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}


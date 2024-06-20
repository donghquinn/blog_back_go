package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

	"github.com/donghquinn/blog_back_go/configs"
)

// 복호화 요청 #회원가입 #로그인
func DecryptString(encodedString string) (string, error) {
	globalConfig := configs.GlobalConfig

    key := []byte(globalConfig.AesKey)

	decodedCiphertext, err := base64.StdEncoding.DecodeString(encodedString)

	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	iv := []byte(globalConfig.AesIv)

	decrypter := cipher.NewCBCDecrypter(block, []byte(iv))
	plaintext := make([]byte, len(decodedCiphertext))
	
	decrypter.CryptBlocks(plaintext, decodedCiphertext)

    
	return string(plaintext), nil
}


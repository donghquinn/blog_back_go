package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/donghquinn/blog_back_go/configs"
)

// Encrypt AES-256-CBC
func EncryptString(rawString string) (string, error) {
	globalConfig := configs.GlobalConfig

	if (len(globalConfig.AesIv) != aes.BlockSize) {
		log.Printf("Block Size and IV Length is Not Match\n AES IV: %v\n BlockSize: %v\n",len(globalConfig.AesIv), aes.BlockSize)

		return "", fmt.Errorf("Block Size and IV Length is Not Match")
	}

	byteString := PKCS5Padding([]byte(rawString), aes.BlockSize, len(rawString))

	block, err := aes.NewCipher([]byte(globalConfig.AesKey))

	if err != nil {
		log.Printf("Creating Decrypt Error: %v", err)

		return "", err
	}
	
	cipherText := make([]byte, len(byteString))

	mode := cipher.NewCBCEncrypter(block, []byte(globalConfig.AesIv))
	mode.CryptBlocks(cipherText, byteString)

	encodedText := hex.EncodeToString(cipherText)

	return encodedText, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

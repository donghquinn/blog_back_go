package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
	"log"

	"github.com/donghquinn/blog_back_go/configs"
)

// Decrypt AES-256-CBC
func DecryptString(encodedString string) (string, error) {
	globalConfig := configs.GlobalConfig

    if len(globalConfig.AesIv) != aes.BlockSize {
        log.Printf("[DECRYPT] Block Size and IV Length Do Not Match\n AES IV: %v\n BlockSize: %v\n", len(globalConfig.AesIv), aes.BlockSize)
        return "", fmt.Errorf("%s", "block size and iv length do not match")
    }

	// Validate input string
    if len(encodedString)%2 != 0 {
        log.Printf("[DECRYPT] invalid input length")
        return "", errors.New("invalid input length")
    }

    for _, char := range encodedString {
        if (char < '0' || char > '9') && (char < 'a' || char > 'f') && (char < 'A' || char > 'F') {
            log.Printf("[DECRYPT] invalid hexadecimal character: %c", char)
            return "", fmt.Errorf("invalid hexadecimal character: %c", char)
        }
    }

    // Decode hexadecimal encoded ciphertext
    cipherText, hexErr := hex.DecodeString(encodedString)

    if hexErr != nil {
        log.Printf("[DECRYPT] Hex Decoding Error: %v", hexErr)
        return "", hexErr
    }

    block, cipherErr := aes.NewCipher([]byte(globalConfig.AesKey))

    if cipherErr != nil {
        log.Printf("[DECRYPT] Creating Cipher Error: %v", cipherErr)
        return "", cipherErr
    }

    // Create a CBC mode decrypter
    mode := cipher.NewCBCDecrypter(block, []byte(globalConfig.AesIv))

    // Create a buffer for the plaintext
    plainText := make([]byte, len(cipherText))

    // Decrypt the ciphertext to the buffer
    mode.CryptBlocks(plainText, cipherText)

    // Unpad the plaintext to remove padding
    plainText = PKCS5UnPadding(plainText)

    return string(plainText), nil
}


func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])

	return src[:(length - unpadding)]
}


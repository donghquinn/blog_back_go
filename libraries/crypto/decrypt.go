package crypto

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
        log.Printf("Block Size and IV Length Do Not Match\n AES IV: %v\n BlockSize: %v\n", len(globalConfig.AesIv), aes.BlockSize)
        return "", fmt.Errorf("Block Size and IV Length Do Not Match")
    }

	// Validate input string
    if len(encodedString)%2 != 0 {
        log.Printf("invalid input length")
        return "", errors.New("invalid input length")
    }

    for _, char := range encodedString {
        if (char < '0' || char > '9') && (char < 'a' || char > 'f') && (char < 'A' || char > 'F') {
            log.Printf("invalid hexadecimal character: %c", char)
            return "", fmt.Errorf("invalid hexadecimal character: %c", char)
        }
    }

    // Decode hexadecimal encoded ciphertext
    cipherText, err := hex.DecodeString(encodedString)

    if err != nil {
        log.Printf("Decoding Error: %v", err)
        return "", err
    }

    block, err := aes.NewCipher([]byte(globalConfig.AesKey))

    if err != nil {
        log.Printf("Creating Cipher Error: %v", err)
        return "", err
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


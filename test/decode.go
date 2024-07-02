package testlogic

import (
	"log"

	crypt "github.com/donghquinn/blog_back_go/libraries/crypto"
)

func Test() {
	email := "ehdgus1524@gmail.com"
	name := "김동현"
	password := "samquinnWkd1"
	encoded , _ := crypt.EncryptString(email)
	encodedName, _ := crypt.EncryptString(name)
	encodedPass, _ := crypt.EncryptString(password)

	log.Println(encoded)
	log.Println(encodedName)
	log.Println(encodedPass)
	// receivedPw:=  "Z8CcFfKXIcOkTq1aBn/tew=="
	// dbPwd := "$2a$10$hKBV01zNJjuojmhUzdJ6z.2V/Ua4QLr/NxP86jUvj70M2jUSzM7h6"
	
	// decoded, _ := crypt.DecryptString(receivedPw)
	// isMatch, err := crypt.PasswordCompare(dbPwd, decoded)

	// if err != nil {
	// 	log.Printf("ERR: %v", err)
	// }
	// log.Printf("[복호화된] email: %v", isMatch)
}

// func generateRandomString(length int) (string, error) {
// 	randomBytes := make([]byte, length)
// 	_, err := rand.Read(randomBytes)
// 	if err != nil {
// 		return "", err
// 	}
// 	return hex.EncodeToString(randomBytes), nil
// }
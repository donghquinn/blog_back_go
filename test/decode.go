package testlogic

import (
	"log"

	crypt "github.com/donghquinn/blog_back_go/libraries/crypto"
)

func Test() {
	// textData, _ := generateRandomString(16)
	// log.Println(textData)
	email:="WFHhVq6smMIeDskaIHg4Ow=="
	// name := "김동현"
	// password := "패스워드"
	
	// enc, encErr := crypt.EncryptString(email)


	em, err := crypt.DecryptString(email)

	if err != nil {
	log.Println(err)
	}

	// na, naErr := crypt.DecryptString(encN)

	// if naErr != nil {
	// log.Println(naErr)
	// }

	// pa, paErr := crypt.DecryptString(encP)

	// if paErr != nil {
	// log.Println(paErr)
	// }

	log.Printf("[복호화된] email: %s", em)
}

// func generateRandomString(length int) (string, error) {
// 	randomBytes := make([]byte, length)
// 	_, err := rand.Read(randomBytes)
// 	if err != nil {
// 		return "", err
// 	}
// 	return hex.EncodeToString(randomBytes), nil
// }
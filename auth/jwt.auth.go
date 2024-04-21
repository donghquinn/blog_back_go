package auth

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/donghquinn/blog_back_go/configs"
	"github.com/donghquinn/blog_back_go/types"
	"github.com/golang-jwt/jwt/v5"
)

// JWT 토큰 생성
func CreateJwtToken(userId string, userEmail string, userStatus string) (string, error) {
	globalConfig := configs.GlobalConfig

	// block, _ := pem.Decode([]byte(globalConfig.JwtKey))

	// privKey, parseErr := x509.ParsePKCS8PrivateKey(block.Bytes);

	// if parseErr != nil {
	// 	log.Printf("Parse PK Error: %v", parseErr)

	// 	return "", parseErr
	// }

	jwtToken := jwt.New(jwt.SigningMethodHS256)

	claims := jwtToken.Claims.(jwt.MapClaims)

	claims["userId"] = userId
	claims["userEmail"] = userEmail
	claims["userStatus"] = userStatus
	// 만료 시간 - 3시간
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix()

	token, err := jwtToken.SignedString([]byte(globalConfig.JwtKey))

	if err != nil {
		log.Printf("[JWT] Create Token Error: %v", err)

		return "", err
	}
	log.Printf("[JWT] Created Token: %s", token)
	
	return token, nil
}


// LoadPublicKey loads RSA public key from PEM format
// func LoadPublicKey(publicKeyPEM string) (any, error) {
//     block, _ := pem.Decode([]byte(publicKeyPEM))
//     if block == nil {
//         return nil, errors.New("failed to decode PEM block containing public key")
//     }

//     publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
//     if err != nil {
//         return nil, err
//     }

//     return publicKey, nil
// }

// JWT 키  검증
func ValidateJwtToken(req *http.Request) (string, string, string, error) {
	token := strings.Split(req.Header["Authorization"][0], "Bearer ")[1]
	// var publicKey *rsa.PublicKey

	globalConfig := configs.GlobalConfig

	// 퍼블릭 키 검증
	// publicKey, pubErr :=  LoadPublicKey(globalConfig.SecretKey)

	// if pubErr != nil {
	// 	log.Printf("[JWT] Pub Key Parse Error: %v", pubErr)

	// 	return "", "", "", pubErr
	// }

	// JWT 토큰 파싱
	parseToken, err := jwt.ParseWithClaims(token, &types.JwtInterface{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			parseErr := fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])

			log.Printf("[JWT] Parse With Claims Error: %v", parseErr)
			return nil, parseErr
		}

		return []byte(globalConfig.JwtKey), nil
	})

	if err != nil {
		log.Printf("[JWT] Parsing JWT Validation Error: %v", err)

		return "","","",err
	}

	claim, ok := parseToken.Claims.(*types.JwtInterface)

	if !ok {
		claimErr := fmt.Errorf("Can't Parse Values from TOKEN")
		log.Printf("[JWT] Parse Token with Claims: %v", claimErr)
		return "", "", "", claimErr
	}

	return claim.UserId, claim.UserEmail, claim.UserType, nil
}

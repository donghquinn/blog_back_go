package users

import (
	"log"
	"net/http"

	"github.com/donghquinn/blog_back_go/dto"
	"github.com/donghquinn/blog_back_go/libraries/crypto"
	"github.com/donghquinn/blog_back_go/types"
	"github.com/donghquinn/blog_back_go/utils"
)

func LoginController(res http.ResponseWriter, req *http.Request) {
	var loginRequst types.UserLoginRequest

	decodeErr := utils.DecodeBody(req, &loginRequst)

	if decodeErr != nil {
		log.Printf("[LOGIN] Parse Body Error: %v", decodeErr)

		dto.SetErrorResponse(res, 401, "01", "SignUp Parsing Error", decodeErr)
	}
}

func DecodeLoginRequest(loginRequest types.UserLoginRequest) (string, string, error) {
	decodeEmail, decodeEmailErr := crypto.DecryptString(loginRequest.UserEmail)

	if decodeEmailErr != nil {
		log.Printf("[LOGIN] Decode Email Err: %v", decodeEmailErr)

		return "", "", decodeEmailErr
	}

	decodePassword, decodePassErr := crypto.DecryptString(loginRequest.UserPassword)

	if decodePassErr!= nil {
		log.Printf("[LOGIN] Decode Password Err: %v", decodePassErr)
		return "","",decodePassErr
	}

	return decodeEmail, decodePassword, nil
}
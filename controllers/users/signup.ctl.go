package users

import (
	"log"
	"net/http"

	"github.com/donghquinn/blog_back_go/dto"
	"github.com/donghquinn/blog_back_go/libraries/crypto"
	"github.com/donghquinn/blog_back_go/libraries/database"
	queries "github.com/donghquinn/blog_back_go/queries/users"
	"github.com/donghquinn/blog_back_go/types"
	"github.com/donghquinn/blog_back_go/utils"
	"github.com/google/uuid"
)

func SignupController(res http.ResponseWriter, req *http.Request) {
	var signupRequestBody types.UserSignupRequest

	parseErr := utils.DecodeBody(req, &signupRequestBody)

	if parseErr != nil {
		log.Printf("[SIGN_UP] Parse Body Error: %v", parseErr)

		dto.SetErrorResponse(res, 401, "01", "SignUp Parsing Error", parseErr)
		return
	}
	
	decodedEmail, decodedName, decodedPassword, decodeErr := decodeSignupUserRequest(signupRequestBody)

	if decodeErr != nil {
		dto.SetErrorResponse(res, 402, "02", "Decode Received User Info", decodeErr)
		return
	}

	connect, dbErr := database.InitDatabaseConnection()

	if dbErr != nil {
		dto.SetErrorResponse(res, 403, "03", "Database Connect Error", dbErr)
		return
	}

	userId, encodedEmail, encodedName, encodedPassword, enocodErr := encodeSignupUserInfo(decodedEmail, decodedPassword, decodedName)

	if enocodErr != nil {
		dto.SetErrorResponse(res, 404, "04", "Encoding Process Error", enocodErr)
		return
	}
	
	_, insertErr := database.InsertQuery(connect, queries.InsertSignupUser, userId, encodedEmail, encodedName, encodedPassword)

	if insertErr != nil {
		dto.SetErrorResponse(res, 405, "05", "Insert New User Info Error", insertErr)
		return
	}

	dto.SetResponse(res, 200, "01")
}

func decodeSignupUserRequest(signupRequest types.UserSignupRequest) (string, string, string, error) {
	decodedEmail, encodeEmailErr := crypto.DecryptString(signupRequest.UserEmail)

	if encodeEmailErr != nil {
		return "","","",encodeEmailErr
	}

	decodedName, encodeNameErr := crypto.EncryptString(signupRequest.UserName)

	if encodeNameErr != nil {
		return "","","",encodeNameErr
	}

	decodedPassword, encodePasswordErr := crypto.DecryptString(signupRequest.UserPassword)

	if encodePasswordErr != nil {
		return "","","",encodePasswordErr
	}

	return decodedEmail, decodedName, decodedPassword, nil
}

// 인코딩
func encodeSignupUserInfo(decodeEmail string, decodePassword string, decodeName string) (string, string, string, string, error) {
	userId, uuidErr := uuid.NewV7()

	if uuidErr != nil {
		log.Printf("[SIGN_UP] Creating User UUID Error: %v", uuidErr)

		return "", "", "", "", uuidErr
	}

	encodedEmail, encodeEmailErr := crypto.EncryptString(decodeEmail)

	if encodeEmailErr != nil {
		return "","","","",encodeEmailErr
	}

	encodedName, encodeNameErr := crypto.EncryptString(decodeName)

	if encodeNameErr != nil {
		return "","","","",encodeNameErr
	}

	encodedPassword, encodePasswordErr := crypto.EncryptHashPassword(decodePassword)

	if encodePasswordErr != nil {
		return "","","","",encodePasswordErr
	}

	return userId.String(), encodedEmail, encodedName, encodedPassword, nil
}

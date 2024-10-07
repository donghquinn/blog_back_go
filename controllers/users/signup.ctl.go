package controllers

import (
	"log"
	"net/http"

	"github.com/donghquinn/blog_back_go/dto"
	crypt "github.com/donghquinn/blog_back_go/libraries/crypto"
	"github.com/donghquinn/blog_back_go/libraries/database"
	queries "github.com/donghquinn/blog_back_go/queries/users"
	"github.com/donghquinn/blog_back_go/types"
	"github.com/donghquinn/blog_back_go/utils"
	"github.com/google/uuid"
)

func SignupController(res http.ResponseWriter, req *http.Request) {
	var signupRequestBody types.UserSignupRequest

	// BODY 파싱
	parseErr := utils.DecodeBody(req, &signupRequestBody)

	if parseErr != nil {
		log.Printf("[SIGN_UP] Parse Body Error: %v", parseErr)

		dto.Response(res, types.ResponseSignupType{
			Status:  http.StatusBadRequest,
			Code:    "USG001",
			Result:  false,
			Message: "Parsing Error",
		})

		return
	}

	// 요청 회원가입 데이터 복호화 (패스워드는 암호화의 암호화된 상태로 전달됨)
	decodedEmail, decodedName, decodedPassword, decodeErr := decodeSignupUserRequest(signupRequestBody)

	// log.Printf("[SIGNUP] decodedEmail: %s, decodedName: %s, decodedPassword: %s", decodedEmail, decodedName, decodedPassword)

	if decodeErr != nil {
		dto.Response(res, types.ResponseSignupType{
			Status:  http.StatusInternalServerError,
			Code:    "USG002",
			Result:  false,
			Message: "Decode Error",
		})
		return
	}

	connect, dbErr := database.InitDatabaseConnection()

	if dbErr != nil {
		dto.Response(res, types.ResponseSignupType{
			Status:  http.StatusInternalServerError,
			Code:    "USG003",
			Result:  false,
			Message: "Db Connection Error",
		})
		return
	}

	// 암호화해서 업로드
	userId, encodedEmail, encodedName, encodedPassword, enocodErr := encodeSignupUserInfo(decodedEmail, decodedPassword, decodedName)

	if enocodErr != nil {
		dto.Response(res, types.ResponseSignupType{
			Status:  http.StatusInternalServerError,
			Code:    "USG004",
			Result:  false,
			Message: "Encrypt User Data Error",
		})
		return
	}

	// log.Printf("[SIGNUP] userId: %s, encodedEmail: %s, encodedName: %s, encodedPassword: %s",userId, encodedEmail, encodedName, encodedPassword)
	// 새로운 유저 데이터 입력
	_, insertErr := connect.InsertQuery(queries.InsertSignupUser, userId, encodedEmail, encodedPassword, encodedName, signupRequestBody.BlogId)

	if insertErr != nil {
		dto.Response(res, types.ResponseSignupType{
			Status:  http.StatusBadRequest,
			Code:    "USG005",
			Result:  false,
			Message: "Insert User Info Error",
		})
		return
	}

	dto.Response(res, types.ResponseSignupType{
		Status:  http.StatusOK,
		Code:    "0000",
		Result:  true,
		Message: "Success",
	})
}

func decodeSignupUserRequest(signupRequest types.UserSignupRequest) (string, string, string, error) {
	decodedEmail, decodeEmailErr := crypt.DecryptString(signupRequest.Email)

	if decodeEmailErr != nil {
		log.Printf("[SIGNUP] Decode Email Error: %v", decodeEmailErr)
		return "", "", "", decodeEmailErr
	}

	decodedName, decodeNameErr := crypt.DecryptString(signupRequest.Name)

	if decodeNameErr != nil {
		log.Printf("[SIGNUP] Decode Name Error: %v", decodeNameErr)
		return "", "", "", decodeNameErr
	}

	decodedPassword, decodePasswordErr := crypt.DecryptString(signupRequest.Password)

	if decodePasswordErr != nil {
		log.Printf("[SIGNUP] Decode Password Error: %v", decodePasswordErr)
		return "", "", "", decodePasswordErr
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

	encodedEmail, encodeEmailErr := crypt.EncryptString(decodeEmail)

	if encodeEmailErr != nil {
		return "", "", "", "", encodeEmailErr
	}

	encodedName, encodeNameErr := crypt.EncryptString(decodeName)

	if encodeNameErr != nil {
		return "", "", "", "", encodeNameErr
	}

	encodedPassword, encodePasswordErr := crypt.EncryptHashPassword(decodePassword)

	if encodePasswordErr != nil {
		return "", "", "", "", encodePasswordErr
	}

	return userId.String(), encodedEmail, encodedName, encodedPassword, nil
}

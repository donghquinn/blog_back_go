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

	// BODY 파싱
	parseErr := utils.DecodeBody(req, &signupRequestBody)

	if parseErr != nil {
		log.Printf("[SIGN_UP] Parse Body Error: %v", parseErr)

		dto.SetErrorResponse(res, 401, "01", "SignUp Parsing Error", parseErr)
		return
	}

	// 요청 회원가입 데이터 복호화 (패스워드는 암호화의 암호화된 상태로 전달됨)
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

	// 암호화해서 업로드
	userId, encodedEmail, encodedName, encodedPassword, enocodErr := encodeSignupUserInfo(decodedEmail, decodedPassword, decodedName)

	if enocodErr != nil {
		dto.SetErrorResponse(res, 404, "04", "Encoding Process Error", enocodErr)
		return
	}
	
	// 새로운 유저 데이터 입력
	_, insertErr := database.InsertQuery(connect, queries.InsertSignupUser, userId, encodedEmail, encodedName, encodedPassword)

	if insertErr != nil {
		dto.SetErrorResponse(res, 405, "05", "Insert New User Info Error", insertErr)
		return
	}

	dto.SetResponse(res, 200, "01")
	return
}

func decodeSignupUserRequest(signupRequest types.UserSignupRequest) (string, string, string, error) {
	decodedEmail, encodeEmailErr := crypto.DecryptString(signupRequest.Email)

	if encodeEmailErr != nil {
		log.Printf("[SIGNUP] Decode Email Error: %v", encodeEmailErr)
		return "","","",encodeEmailErr
	}

	decodedName, encodeNameErr := crypto.EncryptString(signupRequest.Name)

	if encodeNameErr != nil {
		log.Printf("[SIGNUP] Decode Name Error: %v", encodeNameErr)
		return "","","",encodeNameErr
	}

	decodedPassword, encodePasswordErr := crypto.DecryptString(signupRequest.Password)

	if encodePasswordErr != nil {
		log.Printf("[SIGNUP] Decode Password Error: %v", encodePasswordErr)
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

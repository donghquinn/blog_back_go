package controllers

import (
	"log"
	"net/http"

	"github.com/donghquinn/blog_back_go/auth"
	"github.com/donghquinn/blog_back_go/dto"
	crypt "github.com/donghquinn/blog_back_go/libraries/crypto"
	"github.com/donghquinn/blog_back_go/libraries/database"
	queries "github.com/donghquinn/blog_back_go/queries/users"
	"github.com/donghquinn/blog_back_go/types"
	"github.com/donghquinn/blog_back_go/utils"
	"github.com/google/uuid"
)

func LoginController(res http.ResponseWriter, req *http.Request) {
	var loginRequst types.UserLoginRequest

	parseErr := utils.DecodeBody(req, &loginRequst)

	if parseErr != nil {
		log.Printf("[LOGIN] Parse Body Error: %v", parseErr)

		dto.Response(res, types.ResponseTokenType{
			Status:  http.StatusBadRequest,
			Code:    "ULG001",
			Token:   "",
			Message: "Parsing Request Error",
		})

		return
	}

	// 복호화
	decodeEmail, decodePassword, decodeErr := decodeLoginRequest(loginRequst)

	if decodeErr != nil {
		log.Printf("[LOGIN] Decode Requested User Info Error: %v", decodeErr)

		dto.Response(res, types.ResponseTokenType{
			Status:  http.StatusBadRequest,
			Code:    "ULG002",
			Token:   "",
			Message: "Decoding Request Error",
		})
		return
	}

	// DB에서 유저 데이터 체크
	queryResult, queryErr := getUserInfo(loginRequst.Email)

	if queryErr != nil {

		dto.Response(res, types.ResponseTokenType{
			Status:  http.StatusInternalServerError,
			Code:    "ULG003",
			Token:   "",
			Message: "Query Error",
		})
		return
	}

	// 패스워드 비교 (암호화 해싱된 패스워드)
	isMatch, matchErr := crypt.PasswordCompare(queryResult.UserPassword, decodePassword)

	if matchErr != nil {
		log.Printf("[LOGIN] Match Hashed Password Error: %v", matchErr)

		dto.Response(res, types.ResponseTokenType{
			Status:  http.StatusInternalServerError,
			Code:    "ULG004",
			Token:   "",
			Message: "Camparing Password Error",
		})
		return
	}

	// 패스워드 일치하지 않을 때
	if !isMatch {
		log.Printf("[LOGIN] Password Does Not Match: %v", isMatch)

		dto.Response(res, types.ResponseTokenType{
			Status:  http.StatusBadRequest,
			Code:    "ULG005",
			Token:   "",
			Message: "Password Does Not Match",
		})
		return
	}

	uuid, uuidErr := uuid.NewUUID()

	if uuidErr != nil {
		log.Printf("[REDIS] Create UUID Error: %v", uuidErr)

		dto.Response(res, types.ResponseTokenType{
			Status:  http.StatusInternalServerError,
			Code:    "ULG006",
			Token:   "",
			Message: "Create UUID Error",
		})
	}

	// JWT 토큰 생성
	token, tokenErr := auth.CreateJwtToken(queryResult.UserId, uuid.String(), decodeEmail, queryResult.UserStatus, queryResult.BlogId)

	if tokenErr != nil {
		dto.Response(res, types.ResponseTokenType{
			Status:  http.StatusBadRequest,
			Code:    "ULG007",
			Token:   "",
			Message: "Create JWT Token Error",
		})
		return
	}

	dto.Response(res, types.ResponseTokenType{
		Status: http.StatusOK,
		Code:   "0000",
		Token:  token,
	})
}

func decodeLoginRequest(loginRequest types.UserLoginRequest) (string, string, error) {
	decodeEmail, decodeEmailErr := crypt.DecryptString(loginRequest.Email)

	if decodeEmailErr != nil {
		log.Printf("[LOGIN] Decode Email Err: %v", decodeEmailErr)
		return "", "", decodeEmailErr
	}

	decodePassword, decodePassErr := crypt.DecryptString(loginRequest.Password)

	if decodePassErr != nil {
		log.Printf("[LOGIN] Decode Password Err: %v", decodePassErr)
		return "", "", decodePassErr
	}

	return decodeEmail, decodePassword, nil
}

func getUserInfo(encodedEmail string) (types.UserLoginQueryResult, error) {
	connect, connectErr := database.InitDatabaseConnection()

	if connectErr != nil {
		return types.UserLoginQueryResult{}, connectErr
	}

	result, queryErr := connect.QueryOne(queries.SelectUserInfo, encodedEmail)

	if queryErr != nil {
		return types.UserLoginQueryResult{}, queryErr
	}

	defer connect.Close()

	var queryUserInfoResult types.UserLoginQueryResult

	result.Scan(
		&queryUserInfoResult.UserId,
		&queryUserInfoResult.UserPassword,
		&queryUserInfoResult.UserStatus,
		&queryUserInfoResult.BlogId)

	return queryUserInfoResult, nil
}

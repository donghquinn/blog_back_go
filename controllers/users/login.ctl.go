package users

import (
	"fmt"
	"log"
	"net/http"

	"github.com/donghquinn/blog_back_go/auth"
	"github.com/donghquinn/blog_back_go/dto"
	"github.com/donghquinn/blog_back_go/libraries/crypto"
	"github.com/donghquinn/blog_back_go/libraries/database"
	queries "github.com/donghquinn/blog_back_go/queries/users"
	"github.com/donghquinn/blog_back_go/types"
	"github.com/donghquinn/blog_back_go/utils"
)

func LoginController(res http.ResponseWriter, req *http.Request) {
	var loginRequst types.UserLoginRequest

	parseErr := utils.DecodeBody(req, &loginRequst)

	if parseErr != nil {
		log.Printf("[LOGIN] Parse Body Error: %v", parseErr)

		dto.SetErrorResponse(res, 401, "01", "SignUp Parsing Error", parseErr)
		return
	}

	decodeEmail, decodePassword, decodeErr := decodeLoginRequest(loginRequst)

	if decodeErr != nil {
		log.Printf("[LOGIN] Decode Requested User Info Error: %v", decodeErr)
		dto.SetErrorResponse(res, 402, "02", "Decode Login Request", decodeErr)
		return
	}

	queryResult, queryErr := getUserInfo(decodeEmail)

	if queryErr != nil {
		dto.SetErrorResponse(res, 403, "03", "Query User Info Error", queryErr)
		return
	}
	
	isMatch, matchErr := crypto.PasswordCompare(queryResult.User_password, decodePassword)

	if matchErr != nil {
		log.Printf("[LOGIN] Match Hashed Password Error: %v", matchErr)
		dto.SetErrorResponse(res, 404, "04", "Matching User Password Error", matchErr)
		return
	}

	if !isMatch {
		log.Printf("[LOGIN] Password Does Not Match: %v", isMatch)
		dto.SetErrorResponse(res, 405, "05", "Password Does not Match", fmt.Errorf("Password Does not match"))
		return
	}

	token, tokenErr := auth.CreateJwtToken(queryResult.User_id, decodeEmail, queryResult.User_status)

	if tokenErr != nil {
		dto.SetErrorResponse(res, 406, "06", "Create JWT Token Error", tokenErr)
		return
	}

	dto.SetTokenResponse(res, 200, "01", token)
}

func decodeLoginRequest(loginRequest types.UserLoginRequest) (string, string, error) {
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

func getUserInfo(decodeEmail string) (types.UserLoginQueryResult, error){
	connect, connectErr := database.InitDatabaseConnection()

	if connectErr != nil {
		return types.UserLoginQueryResult{}, connectErr
	}

	result, queryErr := database.Query(connect, queries.SelectUserInfo, decodeEmail)

	if queryErr != nil {
		return types.UserLoginQueryResult{}, queryErr
	}

	var queryUserInfoResult types.UserLoginQueryResult

	result.Scan(
		&queryUserInfoResult.User_id,
		&queryUserInfoResult.User_password,
		&queryUserInfoResult.User_status)

	return queryUserInfoResult, nil
}
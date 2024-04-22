package users

import (
	"net/http"

	"github.com/donghquinn/blog_back_go/dto"
	"github.com/donghquinn/blog_back_go/libraries/database"
	queries "github.com/donghquinn/blog_back_go/queries/users"
	"github.com/donghquinn/blog_back_go/types"
	"github.com/donghquinn/blog_back_go/utils"
)

// 이메일 찾기
func SearchEmailController(res http.ResponseWriter, req *http.Request){ 
	var findEmailRequest types.UserSearchEmailRequest

	parsErr := utils.DecodeBody(req, &findEmailRequest)

	if parsErr != nil {
		dto.SetErrorResponse(res, 401, "01", "Parse Find Email Request Body", parsErr)
		return
	}

	foundUserEmail, findErr := getUserEmail(findEmailRequest.Name)

	if findErr != nil {
		dto.SetErrorResponse(res, 402, "02",  "Could Not Found User Email Error", findErr)
		return
	}

	dto.SetEmailResponse(res, 200, "01", foundUserEmail.UserEmail)
}

func getUserEmail(userName string) (types.SelectUserSearchEmailResult, error) {
	var emailQueryResult types.SelectUserSearchEmailResult

	connect, connectErr := database.InitDatabaseConnection()

	if connectErr != nil {
		return types.SelectUserSearchEmailResult{}, connectErr
	}

	queryResult, queryErr := database.QueryOne(connect, queries.SelectUserEmail, userName)

	if queryErr != nil {
		return types.SelectUserSearchEmailResult{}, queryErr
	}

	queryResult.Scan(
		&emailQueryResult.UserEmail)
	
	return emailQueryResult, nil
}

// 페스워드 찾기
func SearchPasswordController(res http.ResponseWriter, req *http.Request){ 
	var findEmailRequest types.UserSearchPasswordRequest

	parsErr := utils.DecodeBody(req, &findEmailRequest)

	if parsErr != nil {
		dto.SetErrorResponse(res, 401, "01", "Parse Find Email Request Body", parsErr)
		return
	}

	foundUserPassword, findErr := getUserPassword(findEmailRequest.Email, findEmailRequest.Name)

	if findErr != nil {
		dto.SetErrorResponse(res, 402, "02",  "Could Not Found User Email Error", findErr)
		return
	}

	dto.SetEmailResponse(res, 200, "01", foundUserPassword.UserPassword)
}

func getUserPassword(userEmail string, userName string) (types.SelectUserSearchPasswordResult, error) {
	var emailQueryResult types.SelectUserSearchPasswordResult

	connect, connectErr := database.InitDatabaseConnection()

	if connectErr != nil {
		return types.SelectUserSearchPasswordResult{}, connectErr
	}

	queryResult, queryErr := database.QueryOne(connect, queries.SelectUserPassword, userName, userEmail)

	if queryErr != nil {
		return types.SelectUserSearchPasswordResult{}, queryErr
	}

	queryResult.Scan(
		&emailQueryResult.UserPassword)
	
	return emailQueryResult, nil
}
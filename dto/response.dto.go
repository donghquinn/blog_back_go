package dto

import (
	"encoding/json"
	"net/http"

	"github.com/donghquinn/blog_back_go/types"
)

// 기본 응답
func SetResponse(res http.ResponseWriter, statusCode int, code string) {
	responseObject, _ := json.Marshal(types.ResponseType{Code: code, Status: true})

	res.WriteHeader(200)
	res.Write(responseObject)
}

// 메세지 담은 응답
func SetResponseWithMessage(res http.ResponseWriter, statusCode int, code string, message string) {
	responseObject, _ := json.Marshal(types.ResponseMessageType{Code: code, Status: true, Message: message})

	res.WriteHeader(200)
	res.Write(responseObject)
}

// 토큰 담은 메세지
func SetTokenResponse(res http.ResponseWriter, statusCode int, code string, token string) {
	responseObject, _ := json.Marshal(types.ResponseTokenType{Code: code, Status: true, Token: token})

	res.WriteHeader(200)
	res.Write(responseObject)
}

// 이메일 담음 응답
func SetEmailResponse(res http.ResponseWriter, statusCode int, code string, email string) {
	responseObject, _ := json.Marshal(types.ResponseFoundEmailType{Code: code, Status: true, Email: email})

	res.WriteHeader(200)
	res.Write(responseObject)
}

func SetPasswordResponse(res http.ResponseWriter, statusCode int, code string, password string) {
	responseObject, _ := json.Marshal(types.ResponseFoundPasswordType{Code: code, Status: true, Password: password})

	res.WriteHeader(200)
	res.Write(responseObject)
}

// 게시글 리스트 담음 응답
func SetPostListResponse(res http.ResponseWriter, statusCode int, code string, data []types.SelectAllPostDataResult) {
	responseObject, _ := json.Marshal(types.ResponsePostListType{Code: code, Status: true, Result: data})

	res.WriteHeader(200)
	res.Write(responseObject)
}

// 게시글 리스트 담음 응답
func SetFileInsertIdResponse(res http.ResponseWriter, statusCode int, code string, insertId string) {
	responseObject, _ := json.Marshal(types.ResponseInsertIdType{Code: code, Status: true, InsertId: insertId})

	res.WriteHeader(200)
	res.Write(responseObject)
}


// 에러 응답
func SetErrorResponse(res http.ResponseWriter, statusCode int, code string, message string, err error ) {
	responseObject, _ := json.Marshal(types.ErrorResponseType{Code: code, Status: false, Message: message})

	res.WriteHeader(500)
	res.Write(responseObject)
}
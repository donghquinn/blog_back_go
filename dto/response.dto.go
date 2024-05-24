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

func SetImageUrlResponse(res http.ResponseWriter, statusCode int, code string, urls []string) {
	responseObject, _ := json.Marshal(types.ResponseImageUrl{Code: code, Status: true, Result: urls})

	res.WriteHeader(200)
	res.Write(responseObject)
}

// 기본 응답
func SetProfileResponse(res http.ResponseWriter, statusCode int, code string, profile types.UserProfileDataResponseType) {
	responseObject, _ := json.Marshal(types.ResponseProfileType{Code: code, Status: true, Result: profile})

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

// 에러 응답
func SetErrorResponse(res http.ResponseWriter, statusCode int, code string, message string, err error ) {
	responseObject, _ := json.Marshal(types.ErrorResponseType{Code: code, Status: false, Message: message})

	res.WriteHeader(500)
	res.Write(responseObject)
}
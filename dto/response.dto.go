package dto

import (
	"encoding/json"
	"net/http"

	"github.com/donghquinn/blog_back_go/types"
)

func SetResponse(res http.ResponseWriter, statusCode int, code string) {
	responseObject, _ := json.Marshal(types.ResponseType{Code: code, Status: true})

	res.WriteHeader(200)
	res.Write(responseObject)
}

func SetPostListResponse(res http.ResponseWriter, statusCode int, code string, data []types.SelectAllPostData) {
	responseObject, _ := json.Marshal(types.ResponsePostListType{Code: code, Status: true, Result: data})

	res.WriteHeader(200)
	res.Write(responseObject)
}


func SetErrorResponse(res http.ResponseWriter, statusCode int, code string, message string, err error ) {
	responseObject, _ := json.Marshal(types.ErrorResponseType{Code: code, Status: false, Message: message})

	res.WriteHeader(500)
	res.Write(responseObject)
}
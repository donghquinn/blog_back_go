package dto

import (
	"encoding/json"
	"net/http"

	types "github.com/donghquinn/blog_back_go/types/post"
)

// 게시글 리스트 담음 응답
func SetPostListResponse(res http.ResponseWriter, statusCode int, code string, data []types.SelectAllPostDataResponse) {
	responseObject, _ := json.Marshal(types.ResponsePostListType{Code: code, Status: true, Result: data})

	res.WriteHeader(200)
	res.Write(responseObject)
}

// 게시글 리스트 담음 응답
func SetPostContentsResponse(res http.ResponseWriter, statusCode int, code string, posts types.ViewSpecificPostContentsResponse) {
	responseObject, _ := json.Marshal(types.ResponsePostContentsType{Code: code, Status: true, Posts: posts})

	res.WriteHeader(200)
	res.Write(responseObject)
}

// 게시글 리스트 담음 응답
func SetFileInsertIdResponse(res http.ResponseWriter, statusCode int, code string, insertId string) {
	responseObject, _ := json.Marshal(types.ResponseInsertIdType{Code: code, Status: true, InsertId: insertId})

	res.WriteHeader(200)
	res.Write(responseObject)
}

// 태그로 게시글 리스트 담음 응답
func SetPostByTagResponse(res http.ResponseWriter, statusCode int, code string, posts []types.PostsByTagsResponseType) {
	responseObject, _ := json.Marshal(types.ResponsePostByTagListType{Code: code, Status: true, Result: posts})

	res.WriteHeader(200)
	res.Write(responseObject)
}


// 카테고리 게시글 리스트 담음 응답
func SetPostByCategoryResponse(res http.ResponseWriter, statusCode int, code string, posts []types.PostByCategoryResponseType) {
	responseObject, _ := json.Marshal(types.ResponsePostByCategoryListType{Code: code, Status: true, Result: posts})

	res.WriteHeader(200)
	res.Write(responseObject)
}
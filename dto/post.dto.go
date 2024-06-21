package dto

import (
	"encoding/json"
	"net/http"
	"strconv"

	types "github.com/donghquinn/blog_back_go/types/post"
)

// 게시글 리스트 담음 응답
func SetPostListResponse(res http.ResponseWriter, statusCode int, code string, data []types.SelectAllPostDataResponse, postCount int) {
	responseObject, _ := json.Marshal(types.ResponsePostListType{Code: code, Result: true, PostList: data, PostCount: strconv.Itoa(postCount)})

	res.WriteHeader(200)
	res.Write(responseObject)
}

// 게시글 리스트 담음 응답
func SetPostContentsResponse(res http.ResponseWriter, statusCode int, code string, posts types.ViewSpecificPostContentsResponse) {
	responseObject, _ := json.Marshal(types.ResponsePostContentsType{Code: code, Result: true, PostList: posts})

	res.WriteHeader(200)
	res.Write(responseObject)
}

// 게시글 리스트 담음 응답
func SetFileInsertIdResponse(res http.ResponseWriter, statusCode int, code string, insertId string) {
	responseObject, _ := json.Marshal(types.ResponseInsertIdType{Code: code, Result: true, InsertId: insertId})

	res.WriteHeader(200)
	res.Write(responseObject)
}

// 태그로 게시글 리스트 담음 응답
func SetPostByTagResponse(res http.ResponseWriter, statusCode int, code string, posts []types.PostsByTagsResponseType) {
	responseObject, _ := json.Marshal(types.ResponsePostByTagListType{Code: code, Result: true, PostList: posts})

	res.WriteHeader(200)
	res.Write(responseObject)
}


// 카테고리 게시글 리스트 담음 응답
func SetPostByCategoryResponse(res http.ResponseWriter, statusCode int, code string, posts []types.PostByCategoryResponseType) {
	responseObject, _ := json.Marshal(types.ResponsePostByCategoryListType{Code: code, Result: true, PostList: posts})

	res.WriteHeader(200)
	res.Write(responseObject)
}
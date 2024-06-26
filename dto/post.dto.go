package dto

import (
	"encoding/json"
	"net/http"

	types "github.com/donghquinn/blog_back_go/types/post"
)

// 게시글 리스트 담음 응답
func SetPostListResponse(res http.ResponseWriter, statusCode int, code string, unpinnedData []types.SelectAllPostDataResponse, pinnedData []types.SelectAllPostDataResponse, postCount string, page int, size int) {
	responseObject, _ := json.Marshal(types.ResponsePostListType{Code: code, Result: true, PinnedPostList: pinnedData, UnpinnedPostList: unpinnedData, PostCount: postCount, Page: page, Size: size})

	res.WriteHeader(200)
	res.Write(responseObject)
}

func SetPinnedPostListResponse(res http.ResponseWriter, statusCode int, code string, pinnedData []types.SelectAllPostDataResponse, postCount string, page int, size int) {
	responseObject, _ := json.Marshal(types.ResponsePinnedPostListType{Code: code, Result: true, PinnedPostList: pinnedData,  PostCount: postCount, Page: page, Size: size})

	res.WriteHeader(200)
	res.Write(responseObject)
}
// 게시글 리스트 담음 응답
func SetPostRegisterResponse(res http.ResponseWriter, statusCode int, code string, postSeq int64) {
	responseObject, _ := json.Marshal(types.ResponsePostRegisterType{Code: code, Result: true, PostSeq: postSeq})

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
func SetPostByTagResponse(res http.ResponseWriter, statusCode int, code string, posts []types.PostsByTagsResponseType, postCount string) {
	responseObject, _ := json.Marshal(types.ResponsePostByTagListType{Code: code, Result: true, PostList: posts, PostCount: postCount})

	res.WriteHeader(200)
	res.Write(responseObject)
}


// 카테고리 게시글 리스트 담음 응답
func SetPostByCategoryResponse(res http.ResponseWriter, statusCode int, code string, posts []types.PostByCategoryResponseType, postCount string) {
	responseObject, _ := json.Marshal(types.ResponsePostByCategoryListType{Code: code, Result: true, PostList: posts, PostCount: postCount})

	res.WriteHeader(200)
	res.Write(responseObject)
}

func SetCategoryResponse(res http.ResponseWriter, statusCode int, code string, categoryList []string) {
	responseObject, _ := json.Marshal(types.ResponseCategoryResponseType{Result: true, Code: code, CategoryList: categoryList})

	res.WriteHeader(200)
	res.Write(responseObject)
}
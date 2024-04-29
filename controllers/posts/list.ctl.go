package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/donghquinn/blog_back_go/dto"
	"github.com/donghquinn/blog_back_go/libraries/crypto"
	post "github.com/donghquinn/blog_back_go/libraries/post"
	types "github.com/donghquinn/blog_back_go/types/post"
	"github.com/donghquinn/blog_back_go/utils"
)

// 전체 포스트 가져오기 - 페이징
func GetPostController(res http.ResponseWriter, req *http.Request) {
	page, _ := strconv.Atoi(req.URL.Query().Get("page"))
	size, _ := strconv.Atoi(req.URL.Query().Get("size"))

	queryResult, queryErr := post.QueryAllPostData(page, size)

	if queryErr != nil {
		dto.SetErrorResponse(res, 402, "02", "Query Post Data Error", queryErr)

		return
	}

	var returnDecodedData []types.SelectAllPostDataResponse

	// 이름 디코딩 위해
	for _, data := range(queryResult){
		decodedName, decodeErr := crypto.DecryptString(data.UserName)

		if decodeErr != nil {
			log.Printf("[LIST] Decoding User Name Error: %v", decodeErr)
			dto.SetErrorResponse(res, 403,"03", "Decode Name Error", decodeErr)
			return
		}

		returnDecodedData = append(returnDecodedData, types.SelectAllPostDataResponse{
			PostSeq: data.PostSeq,
			PostTitle: data.PostTitle,
			PostContents: data.PostContents,
			UserName: decodedName,
			IsPinned: data.IsPinned,
			Viewed: data.Viewed,
			RegDate: data.RegDate,
			ModDate: data.ModDate,
		})
	}

	dto.SetPostListResponse(res, 200, "01", returnDecodedData)
}



// 태그로 포스트 찾기
func GetPostsByTagController(res http.ResponseWriter, req *http.Request ) {
	var getPostByTagRequest types.GetPostsByTagRequest

	parseErr := utils.DecodeBody(req, &getPostByTagRequest)

	if parseErr != nil {
		dto.SetErrorResponse(res, 201, "01", "Parse Request Body Error", parseErr)
		return
	}

	page, _ := strconv.Atoi(req.URL.Query().Get("page"))
	size, _ := strconv.Atoi(req.URL.Query().Get("size"))

	postList, postErr := post.GetPostByTag(getPostByTagRequest, page, size)

	if postErr != nil {
		dto.SetErrorResponse(res, 202, "02", "Get Post List By Tag Error", postErr)
		return
	}

	dto.SetPostByTagResponse(res, 200, "01", postList)
}

// 태그로 포스트 찾기
func GetPostsByCategoryController(res http.ResponseWriter, req *http.Request ) {
	var getPostByCategoryRequest types.GetPostsByCategoryRequest

	parseErr := utils.DecodeBody(req, &getPostByCategoryRequest)

	if parseErr != nil {
		dto.SetErrorResponse(res, 201, "01", "Parse Request Body Error", parseErr)
		return
	}

	page, _ := strconv.Atoi(req.URL.Query().Get("page"))
	size, _ := strconv.Atoi(req.URL.Query().Get("size"))

	postList, postErr := post.GetPostByCategory(getPostByCategoryRequest, page, size)

	if postErr != nil {
		dto.SetErrorResponse(res, 202, "02", "Get Post List By Tag Error", postErr)
		return
	}

	dto.SetPostByCategoryResponse(res, 200, "01", postList)
}
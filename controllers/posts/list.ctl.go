package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/donghquinn/blog_back_go/dto"
	crypt "github.com/donghquinn/blog_back_go/libraries/crypto"
	post "github.com/donghquinn/blog_back_go/libraries/post"
	types "github.com/donghquinn/blog_back_go/types/post"
	"github.com/donghquinn/blog_back_go/utils"
)

// 전체 포스트 가져오기 - 페이징
func GetPostController(res http.ResponseWriter, req *http.Request) {
	var getPostListRequest types.GetPostListRequest

	page, _ := strconv.Atoi(req.URL.Query().Get("page"))
	size, _ := strconv.Atoi(req.URL.Query().Get("size"))

	parseErr := utils.DecodeBody(req, &getPostListRequest)

	if parseErr != nil {
		dto.Response(res, types.ResponsePostListType{
			Status:  http.StatusBadRequest,
			Code:    "PCL001",
			Result:  false,
			Message: "Parse Error",
		})
		return
	}

	unpinnedQueryResult, queryErr := post.QueryUnpinnedPostData(getPostListRequest.BlogId, page, size)

	if queryErr != nil {
		dto.Response(res, types.ResponsePostListType{
			Status:  http.StatusInternalServerError,
			Code:    "PCL002",
			Result:  false,
			Message: "Query Unpinned Posts Error",
		})
		return
	}

	pinnedQueryResult, pinnedErr := post.QueryisPinnedPostData(getPostListRequest.BlogId)

	if pinnedErr != nil {
		dto.Response(res, types.ResponsePostListType{
			Status:  http.StatusInternalServerError,
			Code:    "PCL003",
			Result:  false,
			Message: "Query Pinned Posts Error",
		})
		return
	}

	unpinnedTotalCount, unpinnedTotalCountErr := post.GetTotalUnPinnedPostCount(getPostListRequest.BlogId)

	if unpinnedTotalCountErr != nil {
		dto.Response(res, types.ResponsePostListType{
			Status:  http.StatusInternalServerError,
			Code:    "PCL004",
			Result:  false,
			Message: "Total Unpinned Post Count Error",
		})
		return
	}

	var pinnedData []types.SelectAllPostDataResponse
	var unpinnedData []types.SelectAllPostDataResponse

	for _, data := range unpinnedQueryResult {
		decodedName, decodeErr := crypt.DecryptString(data.UserName)

		if decodeErr != nil {
			log.Printf("[LIST] Decoding User Name Error: %v", decodeErr)
			dto.Response(res, types.ResponsePostListType{
				Status:  http.StatusInternalServerError,
				Code:    "PCL005",
				Result:  false,
				Message: "Decode Error",
			})
			return
		}

		unpinnedData = append(unpinnedData, types.SelectAllPostDataResponse{
			PostSeq:      data.PostSeq,
			PostTitle:    data.PostTitle,
			PostContents: data.PostContents,
			CategoryName: data.CategoryName,
			UserName:     decodedName,
			IsPinned:     data.IsPinned,
			Viewed:       data.Viewed,
			RegDate:      data.RegDate,
			ModDate:      data.ModDate,
		})
	}

	// 이름 디코딩 위해
	for _, data := range pinnedQueryResult {
		decodedName, decodeErr := crypt.DecryptString(data.UserName)

		if decodeErr != nil {
			log.Printf("[LIST] Decoding User Name Error: %v", decodeErr)
			dto.Response(res, types.ResponsePostListType{
				Status:  http.StatusInternalServerError,
				Code:    "PCL006",
				Result:  false,
				Message: "Decode Error",
			})
			return
		}

		pinnedData = append(pinnedData, types.SelectAllPostDataResponse{
			PostSeq:      data.PostSeq,
			PostTitle:    data.PostTitle,
			PostContents: data.PostContents,
			CategoryName: data.CategoryName,
			UserName:     decodedName,
			IsPinned:     data.IsPinned,
			Viewed:       data.Viewed,
			RegDate:      data.RegDate,
			ModDate:      data.ModDate,
		})
	}

	dto.Response(res, types.ResponsePostListType{
		Status:           http.StatusOK,
		Code:             "0000",
		PinnedPostList:   pinnedData,
		UnpinnedPostList: unpinnedData,
		PostCount:        unpinnedTotalCount.Count,
		Size:             size,
		Page:             page,
		Result:           false,
		Message:          "Parse Error",
	})
}

// 전체 포스트 가져오기 - 페이징
func GetPinnedPostController(res http.ResponseWriter, req *http.Request) {
	page, _ := strconv.Atoi(req.URL.Query().Get("page"))
	size, _ := strconv.Atoi(req.URL.Query().Get("size"))

	var getPinnedPostRequest types.GetPostListRequest

	parseErr := utils.DecodeBody(req, &getPinnedPostRequest)

	if parseErr != nil {
		dto.Response(res, types.ResponsePinnedPostListType{
			Status:  http.StatusBadRequest,
			Code:    "PPL001",
			Result:  false,
			Message: "Decode Error",
		})
		return
	}

	pinnedQueryResult, pinnedErr := post.QueryisPinnedPostList(getPinnedPostRequest.BlogId, page, size)

	if pinnedErr != nil {
		dto.Response(res, types.ResponsePinnedPostListType{
			Status:  http.StatusInternalServerError,
			Code:    "PPL002",
			Result:  false,
			Message: "Query Error",
		})

		return
	}

	pinnedTotalCount, pinnedTotalCountErr := post.GetTotalPinnedPostCount(getPinnedPostRequest.BlogId)

	if pinnedTotalCountErr != nil {
		dto.Response(res, types.ResponsePinnedPostListType{
			Status:  http.StatusInternalServerError,
			Code:    "PPL003",
			Result:  false,
			Message: "Get Total Count Error",
		})

		return
	}

	var pinnedData []types.SelectAllPostDataResponse

	// 이름 디코딩 위해
	for _, data := range pinnedQueryResult {
		decodedName, decodeErr := crypt.DecryptString(data.UserName)

		if decodeErr != nil {
			log.Printf("[LIST] Decoding User Name Error: %v", decodeErr)
			dto.Response(res, types.ResponsePinnedPostListType{
				Status:  http.StatusInternalServerError,
				Code:    "PPL004",
				Result:  false,
				Message: "Decrypt Error",
			})
			return
		}

		pinnedData = append(pinnedData, types.SelectAllPostDataResponse{
			PostSeq:      data.PostSeq,
			PostTitle:    data.PostTitle,
			PostContents: data.PostContents,
			CategoryName: data.CategoryName,
			UserName:     decodedName,
			IsPinned:     data.IsPinned,
			Viewed:       data.Viewed,
			RegDate:      data.RegDate,
			ModDate:      data.ModDate,
		})
	}
	dto.Response(res, types.ResponsePinnedPostListType{
		Status:         http.StatusOK,
		Code:           "0000",
		PinnedPostList: pinnedData,
		PostCount:      pinnedTotalCount.Count,
		Page:           page,
		Size:           size,
		Result:         false,
		Message:        "Success",
	})
}

// 태그로 포스트 찾기
func GetPostsByTagController(res http.ResponseWriter, req *http.Request) {
	var getPostByTagRequest types.GetPostsByTagRequest

	parseErr := utils.DecodeBody(req, &getPostByTagRequest)

	if parseErr != nil {
		dto.SetErrorResponse(res, 401, "01", "Parse Request Body Error", parseErr)
		return
	}

	page, _ := strconv.Atoi(req.URL.Query().Get("page"))
	size, _ := strconv.Atoi(req.URL.Query().Get("size"))

	postList, totalPostCount, postErr := post.GetPostByTag(getPostByTagRequest, page, size)

	if postErr != nil {
		dto.SetErrorResponse(res, 402, "02", "Get Post List By Tag Error", postErr)
		return
	}

	dto.SetPostByTagResponse(res, 200, "01", postList, totalPostCount.Count)
}

// 태그로 포스트 찾기
func GetPostsByCategoryController(res http.ResponseWriter, req *http.Request) {
	var getPostByCategoryRequest types.GetPostsByCategoryRequest

	parseErr := utils.DecodeBody(req, &getPostByCategoryRequest)

	if parseErr != nil {
		dto.SetErrorResponse(res, 401, "01", "Parse Request Body Error", parseErr)
		return
	}

	page, _ := strconv.Atoi(req.URL.Query().Get("page"))
	size, _ := strconv.Atoi(req.URL.Query().Get("size"))

	postList, totalCount, postErr := post.GetPostByCategory(getPostByCategoryRequest, page, size)

	if postErr != nil {
		dto.SetErrorResponse(res, 402, "02", "Get Post List By Tag Error", postErr)
		return
	}

	dto.SetPostByCategoryResponse(res, 200, "01", postList, totalCount.Count)
}

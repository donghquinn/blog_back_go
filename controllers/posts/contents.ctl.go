package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/donghquinn/blog_back_go/dto"
	crypt "github.com/donghquinn/blog_back_go/libraries/crypto"
	"github.com/donghquinn/blog_back_go/libraries/database"
	"github.com/donghquinn/blog_back_go/libraries/post"
	queries "github.com/donghquinn/blog_back_go/queries/posts"
	types "github.com/donghquinn/blog_back_go/types/post"
	"github.com/donghquinn/blog_back_go/utils"
)

// 게시글 컨텐츠 컨트롤러
func PostContentsController(res http.ResponseWriter, req *http.Request) {
	var postContentsRequest types.ViewPostContents

	parseErr := utils.DecodeBody(req, &postContentsRequest)

	if parseErr != nil {
		dto.Response(res, types.ResponsePostContentsType{
			Status:  http.StatusBadRequest,
			Code:    "PCT001",
			Result:  false,
			Message: "Parse Error",
		})
		return
	}

	// 게시글 쿼리
	queryResult, queryErr := post.GetPostData(postContentsRequest.PostSeq, postContentsRequest.BlogId)

	if queryErr != nil {
		dto.Response(res, types.ResponsePostContentsType{
			Status:  http.StatusInternalServerError,
			Code:    "PCT002",
			Result:  false,
			Message: "Get Post Data Error",
		})
		return
	}

	imageData, imageErr := GetImageData(postContentsRequest.PostSeq)

	if imageErr != nil {
		dto.Response(res, types.ResponsePostContentsType{
			Status:  http.StatusBadRequest,
			Code:    "PCT003",
			Result:  false,
			Message: "Get Image Data Error",
		})
		return
	}

	var urlArray []string

	if len(imageData) > 0 {
		// 게시글 URL 배열 만들기
		for _, data := range imageData {
			url, getErr := database.GetImageUrl(data.ObjectName, data.FileFormat)

			if getErr != nil {
				dto.Response(res, types.ResponsePostContentsType{
					Status:  http.StatusBadRequest,
					Code:    "PCT004",
					Result:  false,
					Message: "Get Presigned Data Error",
				})
				return
			}

			if url == nil {
				urlArray = make([]string, 0)
			}

			urlArray = append(urlArray, url.String())
		}
	}

	userName, _ := crypt.DecryptString(queryResult.UserName)

	// 특정 게시글 태그 배열 가공해서 담아 응답
	var tagsArray []string

	if queryResult.Tags != nil {
		jsonErr := json.Unmarshal([]byte(*queryResult.Tags), &tagsArray)
		if jsonErr != nil {
			log.Printf("[CONTENTS] JSON Unmarsh tag array Error: %v", jsonErr)
			dto.Response(res, types.ResponsePostContentsType{
				Status:  http.StatusBadRequest,
				Code:    "PCT005",
				Result:  false,
				Message: "Unmarshal Error",
			})
			return
		}
	} else {
		tagsArray = make([]string, 0)
	}

	var categoryName string

	if queryResult.CategoryName != nil {
		categoryName = *queryResult.CategoryName
	} else {
		categoryName = ""
	}

	// 게시글 컨텐츠 데이터
	postContentsData := types.ViewSpecificPostContentsResponse{
		PostSeq:      queryResult.PostSeq,
		PostTitle:    queryResult.PostTitle,
		Tags:         tagsArray,
		PostContents: queryResult.PostContents,
		CategoryName: categoryName,
		UserName:     userName,
		Urls:         urlArray,
		Viewed:       queryResult.Viewed,
		IsPinned:     queryResult.IsPinned,
		RegDate:      queryResult.RegDate,
		ModDate:      queryResult.ModDate,
	}

	dto.Response(res, types.ResponsePostContentsType{
		Status:   http.StatusOK,
		Code:     "0000",
		Result:   false,
		PostList: postContentsData,
		Message:  "Success",
	})
}

// 게시글 번호에 맞는 file 데이터 전부 가져오기
func GetImageData(postSeq string) ([]types.SelectPostImageData, error) {
	var returnImageDate []types.SelectPostImageData

	connect, connectErr := database.InitDatabaseConnection()

	if connectErr != nil {
		log.Printf("[CONTENTS] Init Database Connection Error for Image Data: %v", connectErr)
		return []types.SelectPostImageData{}, connectErr
	}

	result, queryErr := connect.GetMultiple(queries.SelectImageData, postSeq, "POST_IMAGE")

	if queryErr != nil {
		log.Printf("[CONTENTS] Query Image Data Error: %v", queryErr)
		return []types.SelectPostImageData{}, queryErr
	}

	defer connect.Close()

	for result.Next() {
		var row types.SelectPostImageData

		scanErr := result.Scan(
			&row.ObjectName,
			&row.FileFormat,
			&row.TargetPurpose,
			&row.TargetSeq)

		if scanErr != nil {
			if scanErr == sql.ErrNoRows {
				returnImageDate = append(returnImageDate, types.SelectPostImageData{})
			} else {
				log.Printf("[CONTENTS] Scan Files Error: %v", scanErr)
				return []types.SelectPostImageData{}, nil
			}
		}

		returnImageDate = append(returnImageDate, row)
	}

	return returnImageDate, nil
}

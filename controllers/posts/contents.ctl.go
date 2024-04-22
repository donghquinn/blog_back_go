package posts

import (
	"log"
	"net/http"

	"github.com/donghquinn/blog_back_go/dto"
	"github.com/donghquinn/blog_back_go/libraries/database"
	queries "github.com/donghquinn/blog_back_go/queries/posts"
	"github.com/donghquinn/blog_back_go/types"
	"github.com/donghquinn/blog_back_go/utils"
)

// 게시글 컨텐츠 컨트롤러
func PostContentsController(res http.ResponseWriter, req *http.Request) {
	var postContentsRequest types.ViewPostContents

	parseErr := utils.DecodeBody(req, &postContentsRequest)

	if parseErr != nil {
		dto.SetErrorResponse(res, 401, "01", "Parse View Specific Post Contents Error", parseErr)
		return
	}

	log.Printf("[CONTENTS] Post Seq: %s", postContentsRequest.PostSeq)
	
	// 게시글 쿼리
	queryResult, queryErr := GetPostData(postContentsRequest.PostSeq)

	if queryErr != nil {
		dto.SetErrorResponse(res, 402, "02", "Query Specific Contents Error", queryErr)
		return
	}

	log.Printf("[CONTENTS] Post Contents Data: %v",queryResult)

	var urlArray []string

	// 게시글 URL 배열 만들기
	for idx, objectName := range(queryResult.ObjectName) {
		url, getErr := database.GetImageUrl(objectName, queryResult.FileFormat[idx])

		if getErr != nil {
			dto.SetErrorResponse(res, 404, "04", "Get Presigned URL Error", getErr)
			return
		}

		urlArray = append(urlArray, url.String())
	}

	// 게시글 컨텐츠 데이터
	postContentsData := types.ViewSpecificPostContentsResponse{
		PostSeq: queryResult.PostSeq,
		PostTitle: queryResult.PostTitle,
		PostContents: queryResult.PostContents,
		UserId: queryResult.UserId,
		UserName: queryResult.UserName,
		Urls: urlArray,
		RegDate: queryResult.RegDate,
		ModDate: queryResult.ModDate,
	}
	
	dto.SetPostContentsResponse(res, 200, "01", postContentsData)
}

func GetPostData(postSeq string) (types.SelectSpecificPostDataResult, error){
	connect, connectErr := database.InitDatabaseConnection()

	if connectErr != nil {
		return types.SelectSpecificPostDataResult{}, connectErr
	}

	result, queryErr := database.QueryOne(connect, queries.SelectSpecificPostContents, postSeq, postSeq)

	if queryErr != nil {
		return types.SelectSpecificPostDataResult{}, queryErr
	}

	var queryResult types.SelectSpecificPostDataResult

	result.Scan(
		&queryResult.PostSeq,
		&queryResult.PostTitle, 
		&queryResult.PostContents, 
		&queryResult.PostStatus,
		&queryResult.UserId, 
		&queryResult.UserName, 
		&queryResult.ObjectName,
		&queryResult.FileFormat,
		&queryResult.TargetSeq,
		&queryResult.RegDate,
		&queryResult.ModDate)

	return queryResult, nil
}
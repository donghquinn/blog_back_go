package posts

import (
	"net/http"

	"github.com/donghquinn/blog_back_go/dto"
	"github.com/donghquinn/blog_back_go/libraries/database"
	queries "github.com/donghquinn/blog_back_go/queries/posts"
	"github.com/donghquinn/blog_back_go/types"
	"github.com/donghquinn/blog_back_go/utils"
)

func PostContentsController(res http.ResponseWriter, req *http.Request) {
	var postContentsRequest types.ViewPostContents

	parseErr := utils.DecodeBody(req, &postContentsRequest)

	if parseErr != nil {
		dto.SetErrorResponse(res, 401, "01", "Parse View Specific Post Contents Error", parseErr)
		return
	}

	queryResult, queryErr := GetPostData(postContentsRequest.PostSeq)

	if queryErr != nil {
		dto.SetErrorResponse(res, 402, "02", "Query Specific Contents Error", queryErr)
		return
	}

	var urlArray []string

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

	result, queryErr := database.QueryOne(connect, queries.SelectSpecificPostContents, postSeq)

	if queryErr != nil {
		return types.SelectSpecificPostDataResult{}, queryErr
	}

	var queryResult types.SelectSpecificPostDataResult

	result.Scan(
		&queryResult.PostTitle, 
		&queryResult.PostContents, 
		&queryResult.UserId, 
		&queryResult.UserName, 
		&queryResult.ObjectName,
		&queryResult.FileFormat, 
		&queryResult.RegDate,
		&queryResult.ModDate)

	return queryResult, nil
}
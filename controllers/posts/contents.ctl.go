package posts

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/donghquinn/blog_back_go/dto"
	"github.com/donghquinn/blog_back_go/libraries/crypto"
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

	// 게시글 쿼리
	queryResult, tagResults, queryErr := GetPostData(postContentsRequest.PostSeq)

	if queryErr != nil {
		dto.SetErrorResponse(res, 402, "02", "Query Specific Contents Error", queryErr)
		return
	}

	imageData, imageErr := GetImageData(postContentsRequest.PostSeq)

	if imageErr != nil {
		dto.SetErrorResponse(res, 403, "03", "Image Data Error", imageErr)
		return
	}

	var urlArray []string

	// 게시글 URL 배열 만들기
	for _, data := range(imageData) {
		url, getErr := database.GetImageUrl(data.ObjectName, data.FileFormat)

		if getErr != nil {
			dto.SetErrorResponse(res, 404, "04", "Get Presigned URL Error", getErr)
			return
		}

		urlArray = append(urlArray, url.String())
	}

	userName, _ := crypto.DecryptString(queryResult.UserName)

	// 특정 게시글 태그 배열 가공해서 담아 응답
	var tagsArray []string

	for _, t := range(tagResults) {
		jsonErr := json.Unmarshal([]byte(t.TagName), &tagsArray)

		if jsonErr != nil {
			log.Printf("[CONTENTS] JSON Unmarsh tag array Error: %v", jsonErr)
			dto.SetErrorResponse(res, 405, "05", "Unmarshing Tags Error", jsonErr)
			return
		}
	} 

	// 게시글 컨텐츠 데이터
	postContentsData := types.ViewSpecificPostContentsResponse{
		PostSeq: queryResult.PostSeq,
		PostTitle: queryResult.PostTitle,
		Tags: tagsArray,
		PostContents: queryResult.PostContents,
		UserId: queryResult.UserId,
		UserName: userName,
		Urls: urlArray,
		Viewed: queryResult.Viewed,
		IsPinned: queryResult.IsPinned,
		RegDate: queryResult.RegDate,
		ModDate: queryResult.ModDate,
	}
	
	dto.SetPostContentsResponse(res, 200, "01", postContentsData)
}

func GetPostData(postSeq string) (types.SelectSpecificPostDataResult, []types.SelectSpeicificPostTagDataResult, error){
	connect, connectErr := database.InitDatabaseConnection()

	if connectErr != nil {
		log.Printf("[CONTENTS] Init Database Connection Error for Post Data: %v", connectErr)
		return types.SelectSpecificPostDataResult{}, []types.SelectSpeicificPostTagDataResult{},connectErr
	}

	// 조회수 업데이트
	_, updateErr := database.InsertQuery(connect, queries.UpdateViewCount, postSeq)

	if updateErr != nil {
		log.Printf("[CONTENTS] Update View Count Error: %v", updateErr)
		return types.SelectSpecificPostDataResult{},[]types.SelectSpeicificPostTagDataResult{}, updateErr
	}

	// 특정 게시글 조회
	result, queryErr := database.QueryOne(connect, queries.SelectSpecificPostContents, postSeq)

	if queryErr != nil {
		log.Printf("[CONTENTS] Query A Post Contents Error: %v", queryErr)
		return types.SelectSpecificPostDataResult{}, []types.SelectSpeicificPostTagDataResult{}, queryErr
	}

	// 태그들 조회
	tagResult, tagErr := database.Query(connect, queries.SelectPostTags, postSeq)

	if tagErr != nil {
		log.Printf("[CONTENTS] Query Tags Error: %v", tagErr)
		return types.SelectSpecificPostDataResult{}, []types.SelectSpeicificPostTagDataResult{}, tagErr
	}

	defer connect.Close()

	var queryResult types.SelectSpecificPostDataResult

	postScanErr := result.Scan(
		&queryResult.PostSeq,
		&queryResult.PostTitle, 
		&queryResult.PostContents, 
		&queryResult.PostStatus,
		&queryResult.UserId, 
		&queryResult.UserName,
		&queryResult.Viewed,
		&queryResult.IsPinned,
		&queryResult.RegDate,
		&queryResult.ModDate)

	if postScanErr != nil {
		log.Printf("[CONTENTS] Can Post Data Error: %v", postScanErr)
		return queryResult, []types.SelectSpeicificPostTagDataResult{}, postScanErr
	}

	// 태그 쿼리
	var tagQueryResult []types.SelectSpeicificPostTagDataResult

	for tagResult.Next() {
		var row types.SelectSpeicificPostTagDataResult

		scanErr := tagResult.Scan(
			&row.TagName)

		if scanErr != nil {
			log.Printf("[CONTENTS] Scan Tag Query Data Error: %v", scanErr)
			return queryResult, tagQueryResult, scanErr
		}

		tagQueryResult = append(tagQueryResult, row)
	}

	return queryResult, tagQueryResult, nil
}

// 게시글 번호에 맞는 file 데이터 전부 가져오기
func GetImageData(postSeq string) ([]types.SelectPostImageData, error){
	var returnImageDate []types.SelectPostImageData

	connect, connectErr := database.InitDatabaseConnection()

	if connectErr != nil {
		log.Printf("[CONTENTS] Init Database Connection Error for Image Data: %v", connectErr)
		return []types.SelectPostImageData{}, connectErr
	}

	result, queryErr := database.Query(connect, queries.SelectImageData, postSeq, "POST_IMAGE")

	if queryErr != nil {
		log.Printf("[CONTENTS] Query Image Data Error: %v", queryErr)
		return []types.SelectPostImageData{}, queryErr
	}

	for result.Next() {
		var row types.SelectPostImageData

		result.Scan(
			&row.ObjectName,
			&row.FileFormat,
			&row.TargetPurpose,
			&row.TargetSeq)

		returnImageDate = append(returnImageDate, row)
	}

	return returnImageDate, nil
}
package posts

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/donghquinn/blog_back_go/dto"
	"github.com/donghquinn/blog_back_go/libraries/crypto"
	"github.com/donghquinn/blog_back_go/libraries/database"
	"github.com/donghquinn/blog_back_go/libraries/postlib"
	quries "github.com/donghquinn/blog_back_go/queries/posts"
	"github.com/donghquinn/blog_back_go/types"
	"github.com/donghquinn/blog_back_go/utils"
)

// 전체 포스트 가져오기 - 페이징
func GetPostController(res http.ResponseWriter, req *http.Request) {
	// parseBodyErr :=utils.DecodeBody(&req.Body)
	connect, dbErr := database.InitDatabaseConnection()

	if dbErr != nil {
		dto.SetErrorResponse(res, 401, "01", "Database Initiating Error",dbErr)

		return
	}

	page, _ := strconv.Atoi(req.URL.Query().Get("page"))
	size, _ := strconv.Atoi(req.URL.Query().Get("size"))

	queryResult, queryErr := QueryAllPostData(connect, page, size)

	if queryErr != nil {
		dto.SetErrorResponse(res, 402, "02", "Query Post Data Error", queryErr)

		return
	}

	var returnDecodedData []types.SelectAllPostDataResult

	// 이름 디코딩 위해
	for _, data := range(queryResult){
		decodedName, decodeErr := crypto.DecryptString(data.UserName)

		if decodeErr != nil {
			log.Printf("[LIST] Decoding User Name Error: %v", decodeErr)
			dto.SetErrorResponse(res, 403,"03", "Decode Name Error", decodeErr)
			return
		}

		returnDecodedData = append(returnDecodedData, types.SelectAllPostDataResult{
			PostSeq: data.PostSeq,
			PostTitle: data.PostTitle,
			PostContents: data.PostContents,
			UserId: data.UserId,
			UserName: decodedName,
			IsPinned: data.IsPinned,
			Viewed: data.Viewed,
			RegDate: data.RegDate,
			ModDate: data.ModDate,
		})
	}

	dto.SetPostListResponse(res, 200, "01", returnDecodedData)
}

// 포스트들 가져오기 - 모듈함수
func QueryAllPostData(connect *sql.DB, page int, size int) ([]types.SelectAllPostDataResult, error) {
	// 페이징 파라미터 파싱
	result, queryErr := database.Query(connect, quries.SelectAllPosts,  fmt.Sprintf("%d", size), fmt.Sprintf("%d", (page - 1) * size))

	if queryErr != nil {
		log.Printf("[LIST] Get Post Data Error: %v", queryErr)

		return nil, queryErr
	}

	var queryResult = []types.SelectAllPostDataResult{}

	for result.Next() {
		var row types.SelectAllPostDataResult

		scanErr := result.Scan(
			&row.PostSeq,
			&row.PostTitle,
			&row.PostContents,
			&row.UserId,
			&row.UserName,
			&row.IsPinned,
			&row.Viewed,
			&row.RegDate,
			&row.ModDate)

		if scanErr != nil {
			log.Printf("[LIST] Scan and Assign Query Result Error: %v", scanErr)

			return nil, scanErr
		}

		queryResult = append(queryResult, row)
	}

	defer connect.Close()

	return queryResult, nil
}

// 태그로 포스트 찾기
func GetPostsByTagController(res http.ResponseWriter, req *http.Request ) {
	var getPostByTagRequest types.GetPostsByTagRequest

	parseErr := utils.DecodeBody(req, &getPostByTagRequest)

	if parseErr != nil {
		dto.SetErrorResponse(res, 201, "01", "Parse Request Body Error", parseErr)
		return
	}

	postList, postErr := postlib.GetPostTag(getPostByTagRequest)

	if postErr != nil {
		dto.SetErrorResponse(res, 202, "02", "Get Post List By Tag Error", postErr)
		return
	}

	dto.SetPostByTagResponse(res, 200, "01", postList)
}
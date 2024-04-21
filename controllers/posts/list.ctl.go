package posts

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/donghquinn/blog_back_go/dto"
	"github.com/donghquinn/blog_back_go/libraries/database"
	quries "github.com/donghquinn/blog_back_go/queries/posts"
	"github.com/donghquinn/blog_back_go/types"
)

// 전체 포스트 가져오기 - 페이징
func GetPostController(res http.ResponseWriter, req *http.Request) {
	// parseBodyErr :=utils.DecodeBody(&req.Body)
	connect, dbErr := database.InitDatabaseConnection()

	if dbErr != nil {
		dto.SetErrorResponse(res, 401, "01", "Database Initiating Error",dbErr)

		return
	}

	log.Printf("[LIST] Request URL PAth - page: %s, size: %s", req.URL.Query().Get("page"), req.URL.Query().Get("size"))

	page, _ := strconv.Atoi(req.URL.Query().Get("page"))
	size, _ := strconv.Atoi(req.URL.Query().Get("size"))

	queryResult, queryErr := QueryAllPostData(connect, page, size)

	if queryErr != nil {
		dto.SetErrorResponse(res, 403, "03", "Query Post Data Error", queryErr)

		return
	}

	dto.SetPostListResponse(res, 200, "01", queryResult)
}

// 포스트들 가져오기 - 모듈함수
func QueryAllPostData(connect *sql.DB, page int, size int) ([]types.SelectAllPostData, error) {
	// 페이징 파라미터 파싱
	result, queryErr := database.Query(connect, quries.GetAllPosts,  fmt.Sprintf("%d", size), fmt.Sprintf("%d", (page - 1) * size))

	if queryErr != nil {
		log.Printf("[LIST] Get Post Data Error: %v", queryErr)

		return nil, queryErr
	}

	var queryResult = []types.SelectAllPostData{}

	for result.Next() {
		var row types.SelectAllPostData

		scanErr := result.Scan(
			&row.PostTitle,
			&row.PostContents,
			&row.UserId,
			&row.RegDate,
			&row.ModDate)

		if scanErr != nil {
			log.Printf("[LIST] Scan and Assign Query Result Error: %v", scanErr)

			return nil, scanErr
		}

		queryResult = append(queryResult, row)
	}

	return queryResult, nil
}
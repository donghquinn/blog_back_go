package posts

import (
	"database/sql"
	"log"
	"net/http"
	"net/url"

	"github.com/donghquinn/blog_back_go/dto"
	"github.com/donghquinn/blog_back_go/libraries/database"
	quries "github.com/donghquinn/blog_back_go/queries/posts"
	"github.com/donghquinn/blog_back_go/types"
)

// 전체 포스트 가져오기 - 페이징
func GetPost(req *http.Request, res http.ResponseWriter) {
	// parseBodyErr :=utils.DecodeBody(&req.Body)
	connect, dbErr := database.InitDatabaseConnection()

	if dbErr != nil {
		dto.SetErrorResponse(res, 401, "01", "Database Initiating Error",dbErr)

		return
	}

	parameters, parseErr := url.ParseQuery(req.URL.Path)

	if parseErr != nil {
		log.Printf("[POST] Get Post Parse Parameters Error: %v", parseErr)

		dto.SetErrorResponse(res, 402, "02", "Parse Parameter Error", parseErr)

		return
	}

	log.Printf("[POST] Got URL Paramter: %s", parameters.Encode())

	queryResult, queryErr := QueryAllPostData(connect, parameters)

	if queryErr != nil {
		dto.SetErrorResponse(res, 403, "03", "Query Post Data Error", queryErr)

		return
	}

	dto.SetPostListResponse(res, 200, "01", queryResult)
}

// 포스트들 가져오기
func QueryAllPostData(connect *sql.DB, parameters url.Values) ([]types.SelectAllPostData, error) {
	// 페이징 파라미터 파싱
	result, queryErr := database.Query(connect, quries.GetAllPosts, parameters.Get("page"), parameters.Get("size"))

	if queryErr != nil {
		log.Printf("[POST] Get Post Data Error: %v", queryErr)

		return nil, queryErr
	}

	var queryResult = []types.SelectAllPostData{}

	for result.Next() {
		var row types.SelectAllPostData

		scanErr := result.Scan(
			&row.Post_title,
			&row.Post_contents,
			&row.User_id,
			&row.Reg_date,
			&row.Mod_date)

		if scanErr != nil {
			log.Printf("[POST] Scan and Assign Query Result Error: %v", scanErr)

			return nil, scanErr
		}

		queryResult = append(queryResult, row)
	}

	return queryResult, nil
}
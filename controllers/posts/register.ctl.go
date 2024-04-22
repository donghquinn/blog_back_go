package posts

import (
	"log"
	"net/http"
	"strconv"

	"github.com/donghquinn/blog_back_go/auth"
	"github.com/donghquinn/blog_back_go/dto"
	"github.com/donghquinn/blog_back_go/libraries/database"
	queries "github.com/donghquinn/blog_back_go/queries/posts"
	"github.com/donghquinn/blog_back_go/types"
	"github.com/donghquinn/blog_back_go/utils"
)

// 게시글 등록
func RegisterPostController(res http.ResponseWriter, req *http.Request) {
	userId, _, _, err := auth.ValidateJwtToken(req)

	if err != nil {
		dto.SetErrorResponse(res, 401, "01", "JWT Verifying Error", err)

		return
	}

	var registerPostRequest types.RegisterPostRequest

	parseErr := utils.DecodeBody(req, &registerPostRequest)

	if parseErr != nil {
		dto.SetErrorResponse(res, 401, "01", "Parsing Request Body", parseErr)
		return
	}

	connect, dbErr := database.InitDatabaseConnection()

	if dbErr != nil {
		dto.SetErrorResponse(res, 402, "02", "Database Connection", dbErr)
		return
	}

	// 데이터 입력
	insertId, queryErr := database.InsertQuery(connect, queries.InsertPost, userId, registerPostRequest.PostTitle, registerPostRequest.PostContents)
	postSeq := strconv.Itoa(int(insertId))

	for _, seq := range(registerPostRequest.ImageSeqs) {
		// 파일 데이터 업데이트
		_, insertUpdateRr := database.Query(connect, queries.InsertUpdatePostImage, postSeq, seq)

		if insertUpdateRr != nil {
			log.Printf("[REGISTER] Insert Update File Data Error: %v", insertUpdateRr)
			dto.SetErrorResponse(res, 403, "03", "Insert Update File Data Error", insertUpdateRr)
			return
		}
	}

	defer connect.Close()

	if queryErr != nil {
		dto.SetErrorResponse(res, 404, "04", "Data Insert Error", queryErr)
		return
	}

	dto.SetResponse(res, 200, "01")
}
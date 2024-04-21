package posts

import (
	"net/http"

	"github.com/donghquinn/blog_back_go/auth"
	"github.com/donghquinn/blog_back_go/dto"
	"github.com/donghquinn/blog_back_go/libraries/database"
	queries "github.com/donghquinn/blog_back_go/queries/posts"
	"github.com/donghquinn/blog_back_go/types"
	"github.com/donghquinn/blog_back_go/utils"
)

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

	_, queryErr := database.InsertQuery(connect, queries.InsertPost, userId, registerPostRequest.PostTitle, registerPostRequest.PostContents)

	if queryErr != nil {
		dto.SetErrorResponse(res, 403, "03", "Data Insert Error", queryErr)
		return
	}

	dto.SetResponse(res, 200, "01")
}
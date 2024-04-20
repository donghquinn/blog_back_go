package posts

import (
	"log"
	"net/http"
	"net/url"

	"github.com/donghquinn/blog_back_go/dto"
	"github.com/donghquinn/blog_back_go/libraries/database"
	quries "github.com/donghquinn/blog_back_go/queries/posts"
)


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

		dto.SetErrorResponse(res, 401, "02", "Parse Parameter Error", parseErr)

		return
	}

	log.Printf("[POST] Got URL Paramter: %s", parameters.Encode())

	database.Query(connect, quries.GetAllPosts, parameters.Get("page"), parameters.Get("size") )
}
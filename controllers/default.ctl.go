package controllers

import (
	"net/http"

	"github.com/donghquinn/blog_back_go/auth"
	"github.com/donghquinn/blog_back_go/dto"
)

func DefaultController(res http.ResponseWriter, req *http.Request) {
	_, _, _, err := auth.ValidateJwtToken(req)

	if err != nil {
		dto.SetErrorResponse(res, 401, "01", "JWT Verifying Error", err)

		return
	}

	dto.SetResponseWithMessage(res, 200, "01", "Hello World") 
}

func CorsTestController(res http.ResponseWriter, req *http.Request) {
	dto.SetResponseWithMessage(res, 200, "01",  "Hi")
}
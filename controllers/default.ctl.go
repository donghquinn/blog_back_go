package controllers

import (
	"net/http"

	"github.com/donghquinn/blog_back_go/dto"
)

func DefaultController(res http.ResponseWriter, req *http.Request) {
	dto.SetResponseWithMessage(res, 200, "01", "Hello World")
	return 
}
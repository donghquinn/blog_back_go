package controllers

import (
	"net/http"

	"github.com/donghquinn/blog_back_go/dto"
	"github.com/donghquinn/blog_back_go/libraries/post"
)

func GetCategoryController(res http.ResponseWriter, req *http.Request) {
	categoryList, categoryErr := post.GetAllCategoryList()

	if categoryErr != nil {
		dto.SetErrorResponse(res, 400, "01", "Get All Category Error", categoryErr)
		return
	}

	dto.SetCategoryResponse(res, 200, "01", categoryList)
}
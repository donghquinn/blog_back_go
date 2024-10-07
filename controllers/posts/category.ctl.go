package controllers

import (
	"net/http"

	"github.com/donghquinn/blog_back_go/dto"
	"github.com/donghquinn/blog_back_go/libraries/post"
	types "github.com/donghquinn/blog_back_go/types/post"
	"github.com/donghquinn/blog_back_go/utils"
)

func GetCategoryController(res http.ResponseWriter, req *http.Request) {
	var getCategoryListRequest types.GetPostListRequest

	parseErr := utils.DecodeBody(req, &getCategoryListRequest)

	if parseErr != nil {
		dto.Response(res, types.ResponseCategoryResponseType{
			Status:  http.StatusBadRequest,
			Code:    "CAT001",
			Result:  false,
			Message: "Parse Error",
		})
		return
	}

	categoryList, categoryErr := post.GetAllCategoryList(getCategoryListRequest.BlogId)

	if categoryErr != nil {
		dto.Response(res, types.ResponseCategoryResponseType{
			Status:  http.StatusInternalServerError,
			Code:    "CAT002",
			Result:  false,
			Message: "Get Categories Error",
		})
		return
	}

	dto.Response(res, types.ResponseCategoryResponseType{
		Status:       http.StatusOK,
		Code:         "0000",
		Result:       false,
		CategoryList: categoryList,
		Message:      "Success",
	})
}

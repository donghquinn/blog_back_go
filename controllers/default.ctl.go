package controllers

import (
	"net/http"

	"github.com/donghquinn/blog_back_go/auth"
	"github.com/donghquinn/blog_back_go/dto"
)

func DefaultController(res http.ResponseWriter, req *http.Request) {
	_, _, _, _, err := auth.ValidateJwtToken(req)

	if err != nil {
		dto.Response(res, dto.CommonResponseWithMessage{
			Status: http.StatusBadRequest,
			Code: "01",
			Message: "JWT Validation Error",
		})

		return
	}

	dto.Response(res, dto.CommonResponseWithMessage{
			Status: http.StatusOK,
			Code: "0000",
			Message: "Validation success",
		})
}

func CorsTestController(res http.ResponseWriter, req *http.Request) {
	dto.Response(res, dto.CommonResponseWithMessage{
			Status: http.StatusOK,
			Code: "0000",
			Message: "Hi",
		})
}
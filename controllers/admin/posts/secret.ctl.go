package admincontrollers

import (
	"net/http"

	"github.com/donghquinn/blog_back_go/auth"
	"github.com/donghquinn/blog_back_go/dto"
	post "github.com/donghquinn/blog_back_go/libraries/post/admin"
	types "github.com/donghquinn/blog_back_go/types/admin/posts"
	"github.com/donghquinn/blog_back_go/utils"
)

// 비공개 게시글로 변경
func ChangeToSecretPostController(response http.ResponseWriter, request *http.Request) {
	_, _, _, _, err := auth.ValidateJwtToken(request)

	if err != nil {
		dto.Response(response, dto.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "CSP001",
			Message: "JWT Validation Error",
		})
		return
	}

	var changeRequest types.ChangeToRequest

	parseErr := utils.DecodeBody(request, &changeRequest)

	if parseErr != nil {
		dto.Response(response, dto.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "CSP002",
			Message: "Parse Error",
		})
		return
	}

	changeErr := post.ChangeToSecretPost(changeRequest.PostSeq)

	if changeErr != nil {
		dto.Response(response, dto.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "CSP003",
			Message: "Chnage to Secret Error",
		})
		return
	}

	dto.Response(response, dto.CommonResponseWithMessage{
		Status:  http.StatusOK,
		Code:    "0000",
		Message: "Success",
	})
}

// 비공개 게시글에서 공개 게시글로 변경
func ChangeToNotSecretPostController(response http.ResponseWriter, request *http.Request) {
	_, _, _, _, err := auth.ValidateJwtToken(request)

	if err != nil {
		dto.Response(response, dto.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "CNS001",
			Message: "JWT Validation Error",
		})
		return
	}

	var changeRequest types.ChangeToRequest

	parseErr := utils.DecodeBody(request, &changeRequest)

	if parseErr != nil {
		dto.Response(response, dto.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "CNS002",
			Message: "Parse Error",
		})
		return
	}

	changeErr := post.ChangeToNotSecretPost(changeRequest.PostSeq)

	if changeErr != nil {
		dto.Response(response, dto.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "CNS003",
			Message: "Chnage to Not Secret Error",
		})
		return
	}

	dto.Response(response, dto.CommonResponseWithMessage{
		Status:  http.StatusOK,
		Code:    "0000",
		Message: "Success",
	})
}

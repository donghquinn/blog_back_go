package admincontrollers

import (
	"log"
	"net/http"

	"github.com/donghquinn/blog_back_go/auth"
	"github.com/donghquinn/blog_back_go/dto"
	post "github.com/donghquinn/blog_back_go/libraries/post/admin"
	types "github.com/donghquinn/blog_back_go/types/admin/posts"
	"github.com/donghquinn/blog_back_go/utils"
)

// 게시글 등록
func RegisterPostController(res http.ResponseWriter, req *http.Request) {
	userId, _, _, blogId, err := auth.ValidateJwtToken(req)

	if err != nil {
		dto.Response(res, types.ResponsePostRegisterType{
			Status:  http.StatusBadRequest,
			Code:    "RGP001",
			Message: "JWT Validate Error",
		})

		return
	}

	var registerPostRequest types.RegisterPostRequest

	parseErr := utils.DecodeBody(req, &registerPostRequest)

	if parseErr != nil {
		dto.Response(res, types.ResponsePostRegisterType{
			Status:  http.StatusBadRequest,
			Code:    "RGP002",
			Message: "Parse Error",
		})
		return
	}

	postSeq, insertErr := post.InsertPostData(registerPostRequest, userId, blogId)

	if insertErr != nil {
		dto.Response(res, types.ResponsePostRegisterType{
			Status:  http.StatusInternalServerError,
			Code:    "RGP003",
			Message: "Query Error",
		})
		return
	}

	dto.Response(res, types.ResponsePostRegisterType{
		Status:  http.StatusOK,
		Code:    "0000",
		PostSeq: postSeq,
		Message: "Success",
	})
}

func DeletePostController(res http.ResponseWriter, req *http.Request) {
	_, _, _, blogId, err := auth.ValidateJwtToken(req)

	if err != nil {
		dto.Response(res, dto.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "DTP001",
			Message: "JWT Validate Error",
		})

		return
	}

	var deleteRequest types.DeletePostRequest

	parseErr := utils.DecodeBody(req, &deleteRequest)

	if parseErr != nil {
		log.Printf("[DELETE] Parse Delete Request Error: %v", parseErr)
		dto.Response(res, dto.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "DTP002",
			Message: "Parse Error",
		})
		return
	}

	deleteErr := post.DeletePost(deleteRequest.PostSeq, blogId)

	if deleteErr != nil {
		dto.Response(res, dto.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "DTP003",
			Message: "Delete Error",
		})
		return
	}

	dto.Response(res, dto.CommonResponseWithMessage{
		Status:  http.StatusOK,
		Code:    "0000",
		Message: "Success",
	})
}

// 고정 게시글 데이터 업데이트
func UpdatePinPostController(res http.ResponseWriter, req *http.Request) {
	_, _, _, blogId, err := auth.ValidateJwtToken(req)

	if err != nil {
		dto.Response(res, dto.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "UPP001",
			Message: "JWT Validation Error",
		})
		return
	}

	var updatePinRequest types.UpdatePinRequest

	parseErr := utils.DecodeBody(req, &updatePinRequest)

	if parseErr != nil {
		log.Printf("[PIN] Parse Pin Request Error: %v", parseErr)
		dto.Response(res, dto.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "UPP002",
			Message: "Parse Error",
		})
		return
	}

	updateErr := post.UpdatePinPost(updatePinRequest, blogId)

	if updateErr != nil {
		dto.Response(res, dto.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "UPP003",
			Message: "Update Error",
		})
		return
	}

	dto.Response(res, dto.CommonResponseWithMessage{
		Status:  http.StatusOK,
		Code:    "0000",
		Message: "Success",
	})
}

// 고정 게시글 해제 데이터 업데이트
func UpdateUnPinPostController(res http.ResponseWriter, req *http.Request) {
	_, _, _, blogId, err := auth.ValidateJwtToken(req)

	if err != nil {
		dto.Response(res, dto.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "UUP001",
			Message: "JWT Validation Error",
		})
		return
	}

	var updateUnPinRequest types.UpdatePinRequest

	parseErr := utils.DecodeBody(req, &updateUnPinRequest)

	if parseErr != nil {
		log.Printf("[PIN] Parse Un-Pin Request Error: %v", parseErr)
		dto.Response(res, dto.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "UUP002",
			Message: "Parse Error",
		})
		return
	}

	updateErr := post.UpdateUnPinPost(updateUnPinRequest, blogId)

	if updateErr != nil {
		dto.Response(res, dto.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "UUP003",
			Message: "Update Unpin Error",
		})
		return
	}

	dto.Response(res, dto.CommonResponseWithMessage{
		Status:  http.StatusOK,
		Code:    "0000",
		Message: "Success",
	})
}

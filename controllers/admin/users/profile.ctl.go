package admincontrollers

import (
	"log"
	"net/http"

	"github.com/donghquinn/blog_back_go/auth"
	"github.com/donghquinn/blog_back_go/dto"
	"github.com/donghquinn/blog_back_go/libraries/profile"
	types "github.com/donghquinn/blog_back_go/types/admin/users"
	"github.com/donghquinn/blog_back_go/utils"
)

// 프로필 변경 컨트롤러
func UpdateProfileController(res http.ResponseWriter, req *http.Request) {
	var updateProfile types.UserChangeProfileRequest

	userId, _, _, blogId, err := auth.ValidateJwtToken(req)

	if err != nil {
		dto.Response(res, dto.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "UPF001",
			Message: "JWT Token Validation Error",
		})
		return
	}

	parseErr := utils.DecodeBody(req, &updateProfile)

	if parseErr != nil {
		log.Printf("[PROFILE] Change Profile Request Error: %v", parseErr)
		dto.Response(res, dto.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "UPF002",
			Message: "Parse Error",
		})
		return
	}

	updateErr := profile.ChangeProfile(updateProfile, userId, blogId)

	if updateErr != nil {
		dto.Response(res, dto.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "UPF003",
			Message: "Change Profile Error",
		})
		return
	}

	dto.Response(res, dto.CommonResponseWithMessage{
		Status:  http.StatusOK,
		Code:    "0000",
		Message: "Success",
	})
}

// 색상 변경 컨트롤러
func UpdateColorController(res http.ResponseWriter, req *http.Request) {
	userId, _, _, blogId, err := auth.ValidateJwtToken(req)

	if err != nil {
		dto.Response(res, dto.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "UPC001",
			Message: "JWT Token Validation Error",
		})
		return
	}

	var changeColorRequest types.UserUpdateProfileColorRequest

	parseErr := utils.DecodeBody(req, &changeColorRequest)

	if parseErr != nil {
		log.Printf("[COLOR] Change Color Request Error: %v", parseErr)
		dto.Response(res, dto.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "UPC002",
			Message: "Parse Error",
		})
		return
	}

	changeColorErr := profile.ChangeColor(changeColorRequest, userId, blogId)

	if changeColorErr != nil {
		dto.Response(res, dto.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "UPC003",
			Message: "Change Color Error",
		})
		return
	}

	dto.Response(res, dto.CommonResponseWithMessage{
		Status:  http.StatusOK,
		Code:    "0000",
		Message: "Success",
	})
}

// 블로그 타이틀 변경 컨트롤러
func UpdateTitleController(res http.ResponseWriter, req *http.Request) {
	userId, _, _, blogId, err := auth.ValidateJwtToken(req)

	if err != nil {
		dto.Response(res, dto.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "UPT001",
			Message: "JWT Token Validation Error",
		})
		return
	}

	var changeTitleRequest types.UserUpdateBlogTitleRequest

	parseErr := utils.DecodeBody(req, &changeTitleRequest)

	if parseErr != nil {
		log.Printf("[TITLE] Change Title Request Error: %v", parseErr)
		dto.Response(res, dto.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "UPT002",
			Message: "Parse Error",
		})
		return
	}

	changeTitleErr := profile.ChangeBlogTitle(changeTitleRequest, userId, blogId)

	if changeTitleErr != nil {
		dto.Response(res, dto.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "UPT003",
			Message: "Change Title Error",
		})
		return
	}

	dto.Response(res, dto.CommonResponseWithMessage{
		Status:  http.StatusOK,
		Code:    "0000",
		Message: "Success",
	})
}

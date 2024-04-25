package users

import (
	"log"
	"net/http"

	"github.com/donghquinn/blog_back_go/auth"
	"github.com/donghquinn/blog_back_go/dto"
	"github.com/donghquinn/blog_back_go/libraries/profile"
	"github.com/donghquinn/blog_back_go/types"
	"github.com/donghquinn/blog_back_go/utils"
)

// 프로필 변경 컨트롤러
func UpdateProfileController(res http.ResponseWriter, req *http.Request) {
	var updateProfile types.UserChangeProfileRequest

	userId, _, _, err := auth.ValidateJwtToken(req)

	if err != nil {
		dto.SetErrorResponse(res, 401, "01", "JWT Verifying Error", err)
		return
	}
	
	parseErr := utils.DecodeBody(req, &updateProfile)

	if parseErr != nil {
		log.Printf("[PROFILE] Change Profile Request Error: %v", parseErr)
		dto.SetErrorResponse(res, 401, "01", "Change Profile Request Error", parseErr)
		return
	}

	updateErr := profile.ChangeProfile(updateProfile, userId)

	if updateErr != nil {
		dto.SetErrorResponse(res, 402, "02", "Insert Update Error", updateErr)
		return
	}

	dto.SetResponse(res, 200, "01")
}

// 색상 변경 컨트롤러
func UpdateColorController(res http.ResponseWriter, req *http.Request) {
	userId, _, _, err := auth.ValidateJwtToken(req)

	if err != nil {
		dto.SetErrorResponse(res, 401, "01", "JWT Verifying Error", err)
		return
	}

	var changeColorRequest types.UserUpdateProfileColorRequest

	parseErr := utils.DecodeBody(req, &changeColorRequest)

	if parseErr != nil {
		log.Printf("[COLOR] Change Color Request Error: %v", parseErr)
		dto.SetErrorResponse(res, 201, "01", "Change Color Request Error", parseErr)
		return
	}

	changeColorErr := profile.ChangeColor(changeColorRequest, userId)

	if changeColorErr != nil {
		dto.SetErrorResponse(res, 202, "02", "Change Color Error", changeColorErr)
		return 
	}

	dto.SetResponse(res, 200, "01")
}

package users

import (
	"net/http"

	"github.com/donghquinn/blog_back_go/dto"
	"github.com/donghquinn/blog_back_go/libraries/profile"
	"github.com/donghquinn/blog_back_go/types"
	"github.com/donghquinn/blog_back_go/utils"
)

func GetUserProfileController(res http.ResponseWriter, req *http.Request) {
	var getProfileUser types.UserProfileRequest

	parseErr := utils.DecodeBody(req, &getProfileUser)

	if parseErr != nil {
		dto.SetErrorResponse(res, 402, "02", "Parse Body Error", parseErr)
		return
	}

	profile, querErr := profile.GetUserProfile(getProfileUser.UserId)

	if querErr != nil {
		dto.SetErrorResponse(res, 402, "02", "Profile query Error", querErr)
		return
	}

	dto.SetProfileResponse(res, 200, "01", profile)
}
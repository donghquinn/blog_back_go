package users

import (
	"log"
	"net/http"

	"github.com/donghquinn/blog_back_go/dto"
	"github.com/donghquinn/blog_back_go/libraries/profile"
	"github.com/donghquinn/blog_back_go/types"
	"github.com/donghquinn/blog_back_go/utils"
)

func UpdateProfileController(res http.ResponseWriter, req *http.Request) {
	var updateProfile types.UserChangeProfileRequest

	parseErr := utils.DecodeBody(req, &updateProfile)

	if parseErr != nil {
		log.Printf("[PROFILE] Change Profile Request Error: %v", parseErr)
		dto.SetErrorResponse(res, 401, "01", "Change Profile Request Error", parseErr)
		return
	}

	updateErr := profile.ChangeProfile(updateProfile)

	if updateErr != nil {
		dto.SetErrorResponse(res, 402, "02", "Insert Update Error", updateErr)
		return
	}

	dto.SetResponse(res, 200, "01")
}
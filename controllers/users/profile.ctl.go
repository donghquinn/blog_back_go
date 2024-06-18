package controllers

import (
	"log"
	"net/http"

	"github.com/donghquinn/blog_back_go/dto"
	"github.com/donghquinn/blog_back_go/libraries/crypto"
	"github.com/donghquinn/blog_back_go/libraries/profile"
	"github.com/donghquinn/blog_back_go/types"
	"github.com/donghquinn/blog_back_go/utils"
)

func GetUserProfileController(res http.ResponseWriter, req *http.Request) {
	var getProfileUser types.UserProfileRequest

	parseErr := utils.DecodeBody(req, &getProfileUser)

	if parseErr != nil {
		dto.SetErrorResponse(res, 401, "01", "Parse Body Error", parseErr)
		return
	}

	profile, querErr := profile.GetUserProfile(getProfileUser.UserId)

	if querErr != nil {
		dto.SetErrorResponse(res, 402, "02", "Profile query Error", querErr)
		return
	}

	decodedName, nameErr := crypto.DecryptString(profile.UserName)

	if nameErr != nil {
		log.Printf("[PROFILE] Decode User Name: %v", nameErr)
		dto.SetErrorResponse(res, 403, "03", "Decode User Name Error", nameErr)
		return
	}

	decodedEmail, emailErr := crypto.DecryptString(profile.UserEmail)

	if emailErr != nil {
		log.Printf("[PROFILE] Decode User Email: %v", emailErr)
		dto.SetErrorResponse(res, 404, "04", "Decode User Email Error", emailErr)
		return
	}
	profile.UserName = decodedName
	profile.UserEmail = decodedEmail

	dto.SetProfileResponse(res, 200, "01", profile)
}
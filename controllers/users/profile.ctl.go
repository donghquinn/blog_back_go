package controllers

import (
	"log"
	"net/http"

	"github.com/donghquinn/blog_back_go/dto"
	crypt "github.com/donghquinn/blog_back_go/libraries/crypto"
	"github.com/donghquinn/blog_back_go/libraries/profile"
)

func GetUserProfileController(res http.ResponseWriter, req *http.Request) {
	profile, querErr := profile.GetUserProfile()

	if querErr != nil {
		dto.SetErrorResponse(res, 402, "02", "Profile query Error", querErr)
		return
	}

	decodedName, nameErr := crypt.DecryptString(profile.UserName)

	if nameErr != nil {
		log.Printf("[PROFILE] Decode User Name: %v", nameErr)
		dto.SetErrorResponse(res, 403, "03", "Decode User Name Error", nameErr)
		return
	}

	decodedEmail, emailErr := crypt.DecryptString(profile.UserEmail)

	if emailErr != nil {
		log.Printf("[PROFILE] Decode User Email: %v", emailErr)
		dto.SetErrorResponse(res, 404, "04", "Decode User Email Error", emailErr)
		return
	}

	profile.UserName = decodedName
	profile.UserEmail = decodedEmail

	dto.SetProfileResponse(res, 200, "01", profile)
}
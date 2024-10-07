package controllers

import (
	"log"
	"net/http"

	"github.com/donghquinn/blog_back_go/dto"
	crypt "github.com/donghquinn/blog_back_go/libraries/crypto"
	"github.com/donghquinn/blog_back_go/libraries/profile"
	types "github.com/donghquinn/blog_back_go/types/admin/users"
	"github.com/donghquinn/blog_back_go/utils"
)

func GetUserProfileController(res http.ResponseWriter, req *http.Request) {
	var getUserProfileRequest types.UserGetProfileRequest

	parseErr := utils.DecodeBody(req, &getUserProfileRequest)

	if parseErr != nil {
		log.Printf("[LOGIN] Parse Body Error: %v", parseErr)

		dto.Response(res, types.ResponseProfileType{
			Status:        http.StatusBadRequest,
			Code:          "UPF001",
			ProfileResult: types.UserProfileDataResponseType{},
			Message:       "Parsing Error",
		})
		return
	}

	profile, querErr := profile.GetUserProfile(getUserProfileRequest.BlogId, getUserProfileRequest.UserId)

	if querErr != nil {
		dto.Response(res, types.ResponseProfileType{
			Status:        http.StatusInternalServerError,
			Code:          "UPF002",
			ProfileResult: types.UserProfileDataResponseType{},
			Message:       "Query Error",
		})
		return
	}

	decodedName, nameErr := crypt.DecryptString(profile.UserName)

	if nameErr != nil {
		log.Printf("[PROFILE] Decode User Name: %v", nameErr)
		dto.Response(res, types.ResponseProfileType{
			Status:        http.StatusInternalServerError,
			Code:          "UPF003",
			ProfileResult: types.UserProfileDataResponseType{},
			Message:       "Decrypt Error",
		})
		return
	}

	decodedEmail, emailErr := crypt.DecryptString(profile.UserEmail)

	if emailErr != nil {
		log.Printf("[PROFILE] Decode User Email: %v", emailErr)
		dto.Response(res, types.ResponseProfileType{
			Status:        http.StatusInternalServerError,
			Code:          "UPF004",
			ProfileResult: types.UserProfileDataResponseType{},
			Message:       "Decrypt Error",
		})
		return
	}

	profile.UserName = decodedName
	profile.UserEmail = decodedEmail

	dto.Response(res, types.ResponseProfileType{
		Status:        http.StatusOK,
		Code:          "0000",
		ProfileResult: profile,
		Message:       "Success",
	})
}

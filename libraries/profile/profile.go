package profile

import (
	"log"

	"github.com/donghquinn/blog_back_go/libraries/database"
	queries "github.com/donghquinn/blog_back_go/queries/users"
	"github.com/donghquinn/blog_back_go/types"
)


func GetUserProfile(userId string) (types.SelectUserProfileQueryResult, error) {
	connect, dbErr := database.InitDatabaseConnection()

	if dbErr != nil {
		return types.SelectUserProfileQueryResult{}, dbErr
	}

	profile, profileErr := database.Query(connect, queries.SelectUserProfile, userId, "USER_PROFILE", "USER_BACKGROUND", userId)

	if profileErr != nil {
		log.Printf("[PROFILE] Get Profile Error: %v", profileErr)
		return types.SelectUserProfileQueryResult{}, profileErr
	}

	var userProfileData types.SelectUserProfileQueryResult

	profile.Scan(
		&userProfileData.FileFormat,
		&userProfileData.FileType,
		&userProfileData.TargetId,
		&userProfileData.ObjectName,
		&userProfileData.UserId,
		&userProfileData.UserEmail,
		&userProfileData.ProfileSeq,
		&userProfileData.BackgroundSeq,
		&userProfileData.Color,
		&userProfileData.Title,
		&userProfileData.GithubUrl,
		&userProfileData.PersonalUrl,
		&userProfileData.Memo)

	return userProfileData, nil

}
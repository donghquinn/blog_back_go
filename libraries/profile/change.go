package profile

import (
	"log"

	"github.com/donghquinn/blog_back_go/libraries/database"
	queries "github.com/donghquinn/blog_back_go/queries/users"
	"github.com/donghquinn/blog_back_go/types"
)

func ChangeProfile(data types.UserChangeProfileRequest, userId string) error {
	connect, connectErr := database.InitDatabaseConnection()

	if connectErr != nil {
		return connectErr
	}

	_ , insertErr := database.InsertQuery(
		connect, 
		queries.InsertUpdateProfileInfo,
		data.ProfileImage,
		data.BackgroundImage,
		data.Color,
		data.Title,
		data.Instagram,
		data.GithubUrls,
		data.PersonalUrls,
		data.Memo)

	if insertErr != nil {
		log.Printf("[PROFILE] Insert Profile Error: %v", insertErr)

		return insertErr
	}

	return nil
}

func ChangeColor(data types.UserUpdateProfileColorRequest, userId string) error {
	connect, connectErr := database.InitDatabaseConnection()

	if connectErr != nil {
		return connectErr
	}

	_, updateErr := database.InsertQuery(connect, queries.UpdateProfileColor, data.Color, userId)

	if updateErr!= nil {
		log.Printf("[COLOR] Update Color Error: %v", updateErr)

		return updateErr
	}

	return nil
}
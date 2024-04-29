package profile

import (
	"log"

	"github.com/donghquinn/blog_back_go/libraries/crypto"
	"github.com/donghquinn/blog_back_go/libraries/database"
	queries "github.com/donghquinn/blog_back_go/queries/users"
	types "github.com/donghquinn/blog_back_go/types/admin/users"
)

// 프로필 전체 변경
func ChangeProfile(data types.UserChangeProfileRequest, userId string) error {
	connect, connectErr := database.InitDatabaseConnection()

	if connectErr != nil {
		return connectErr
	}

	encodedName, encodeErr := crypto.EncryptString(data.Name)

	if encodeErr != nil {
		log.Printf("[PROFILE] Encode Name Error: %v", encodeErr)
		return encodeErr
	}

	_ , insertErr := database.InsertQuery(
		connect, 
		queries.InsertUpdateProfileInfo,
		encodedName,
		data.Color,
		data.Title,
		data.Instagram,
		data.GithubUrls,
		data.PersonalUrls,
		data.Memo,
		userId)

	if insertErr != nil {
		log.Printf("[PROFILE] Insert Profile Error: %v", insertErr)

		return insertErr
	}

	// _, insertImgErr := database.InsertQuery(
	// 	connect)

	defer connect.Close()

	return nil
}

// 색상 변경
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

// 블로그 타이틀 변경
func ChangeBlogTitle(data types.UserUpdateBlogTitleRequest, userId string) error {
	connect, connectErr := database.InitDatabaseConnection()

	if connectErr != nil {
		return connectErr
	}

	_, updateErr := database.InsertQuery(connect, queries.UpdateTitle, data.Title, userId)

	if updateErr != nil {
		log.Printf("[TITLE] Change Title Error: %v", updateErr)
		return updateErr
	}

	return nil
}

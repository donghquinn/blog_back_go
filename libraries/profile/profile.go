package profile

import (
	"log"

	"github.com/donghquinn/blog_back_go/libraries/database"
	queries "github.com/donghquinn/blog_back_go/queries/users"
	"github.com/donghquinn/blog_back_go/types"
)


func GetUserProfile(userId string) (types.UserProfileDataResponseType, error) {
	connect, dbErr := database.InitDatabaseConnection()

	if dbErr != nil {
		return types.UserProfileDataResponseType{}, dbErr
	}

	profile, profileErr := database.QueryOne(connect, queries.SelectUserProfile, userId)

	if profileErr != nil {
		log.Printf("[PROFILE] Get Profile Error: %v", profileErr)
		return types.UserProfileDataResponseType{}, profileErr
	}

	images, imagesErr := database.Query(connect, queries.SelectUserProfileProfileAndBackground, userId, "USER_PROFILE", "USER_BACKGROUND")
	
	if imagesErr != nil {
		log.Printf("[PROFILE] Get Profile And Background Images Error: %v", imagesErr)
		return types.UserProfileDataResponseType{}, imagesErr
	}

	defer connect.Close()

	var userProfileData types.SelectUserProfileQueryResult

	profile.Scan(
		&userProfileData.UserId,
		&userProfileData.UserEmail,
		&userProfileData.UserName,
		&userProfileData.Color,
		&userProfileData.Title,
		&userProfileData.GithubUrl,
		&userProfileData.PersonalUrl,
		&userProfileData.Memo)

	// 이미지 데이터 url 가져오기 시작
	var userImageData []types.SelectFileQueryResult

	for images.Next() {
		var row types.SelectFileQueryResult

		images.Scan(
			&row.FileFormat,
			&row.FileType,
			&row.TargetPurpose,
			&row.TargetId,
			&row.ObjectName)

		userImageData = append(userImageData, row)
	}

	imageUrls, urlErr := getImages(userImageData)

	if urlErr != nil {
		return types.UserProfileDataResponseType{}, urlErr
	}

	 userProfileResult := types.UserProfileDataResponseType {
		UserId: userProfileData.UserId,
		UserName: userProfileData.UserName,
		UserEmail: userProfileData.UserEmail,
		Color: userProfileData.Color,
		Title: userProfileData.Title,
		GithubUrl: userProfileData.GithubUrl,
		PersonalUrl: userProfileData.PersonalUrl,
		Memo: userProfileData.Memo,
		Images: imageUrls,
	 }

	return userProfileResult, nil
}

func getImages(imageData []types.SelectFileQueryResult) (types.UserImageFileData, error){
	var imageUrlData types.UserImageFileData

	for _, row := range(imageData) {
		imageUrl, err := database.GetImageUrl(row.ObjectName, row.FileType)

		if err != nil {
			log.Printf("[PROFILE] Get Profile IMAGE URL: %v", err)
			return types.UserImageFileData{}, err
		}

		if row.TargetPurpose == "USER_PROFILE" {
			imageUrlData.ProfileImage = imageUrl.String()
		}

		if row.TargetPurpose == "USER_BACKGROUND" {
			imageUrlData.BackgroundImage = imageUrl.String()
		}
	}

	return imageUrlData, nil
}
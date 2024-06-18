package profile

import (
	"log"

	"github.com/donghquinn/blog_back_go/libraries/database"
	queries "github.com/donghquinn/blog_back_go/queries/users"
	"github.com/donghquinn/blog_back_go/types"
)


func GetUserProfile(userId string) (types.UserProfileDataResponseType, error) {
	var userProfileResult types.UserProfileDataResponseType

	userProfileData, profileErr := GetUserProfileByUserId(userId)

	if profileErr != nil {
		return types.UserProfileDataResponseType{}, profileErr
	}

	imageUrlList, imageUrlErr := GetUserProfileImageList(userId)
	
	if imageUrlErr != nil {
		return types.UserProfileDataResponseType{}, imageUrlErr
	}

	 userProfileResult = types.UserProfileDataResponseType {
		UserId: userProfileData.UserId,
		UserName: userProfileData.UserName,
		UserEmail: userProfileData.UserEmail,
		Color: userProfileData.Color,
		Title: userProfileData.Title,
		GithubUrl: userProfileData.GithubUrl,
		PersonalUrl: userProfileData.PersonalUrl,
		Memo: userProfileData.Memo,
		Images: imageUrlList,
	 }

	return userProfileResult, nil
}


func GetUserProfileByUserId(userId string) (types.SelectUserProfileQueryResult, error) {
	var userProfileData types.SelectUserProfileQueryResult

	connect, dbErr := database.InitDatabaseConnection()

	if dbErr != nil {
		return types.SelectUserProfileQueryResult{}, dbErr
	}

	profile, profileErr := database.QueryOne(connect, queries.SelectUserProfile, userId)

	if profileErr != nil {
		log.Printf("[PROFILE] Get Profile Error: %v", profileErr)
		return types.SelectUserProfileQueryResult{}, profileErr
	}

	profile.Scan(
		&userProfileData.UserId,
		&userProfileData.UserEmail,
		&userProfileData.UserName,
		&userProfileData.Color,
		&userProfileData.Title,
		&userProfileData.GithubUrl,
		&userProfileData.PersonalUrl,
		&userProfileData.Memo)

	return userProfileData, nil
}

func GetUserProfileImageList(userId string) (types.UserImageFileData, error){
		// 이미지 데이터 url 가져오기 시작
	var userImageData []types.SelectFileQueryResult
	var imageUrlList types.UserImageFileData

	connect, dbErr := database.InitDatabaseConnection()

	if dbErr != nil {
		return types.UserImageFileData{}, dbErr
	}


	images, imagesErr := database.Query(connect, queries.SelectUserProfileProfileAndBackground, userId, "USER_PROFILE", "USER_BACKGROUND")
	
	if imagesErr != nil {
		log.Printf("[PROFILE] Get Profile And Background Images Error: %v", imagesErr)
		return types.UserImageFileData{}, imagesErr
	}

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

	imageUrlList = imageUrls

	if urlErr != nil {
		return types.UserImageFileData{}, urlErr
	}

	return imageUrlList, nil
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
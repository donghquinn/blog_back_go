package post

import (
	"encoding/json"
	"log"

	"github.com/donghquinn/blog_back_go/libraries/database"
	queries "github.com/donghquinn/blog_back_go/queries/admin/posts"
	types "github.com/donghquinn/blog_back_go/types/admin/posts"
	"github.com/donghquinn/blog_back_go/utils"
)

// 게시글 수정
func EditPost(data types.EditPostRequest, userId string) error {
	// 카테고리 유효성 검증
	isValidCategory := utils.ValidateRequestValue(data.Category)

	editCategoryErr := InsertUpdateCategory(data.Category, isValidCategory)

	if editCategoryErr != nil {
		return editCategoryErr
	}

	tags := data.Tags

	editTagErr := InsertUpdateTagList(tags, data.PostSeq)

	if editTagErr != nil {
		return editTagErr
	}

	editImageErr := InsertImageSeqList(data.ImageSeqs, data.PostSeq)

	if editImageErr != nil {
		return editImageErr
	}

	editPost := UpdatePostEdit(data.PostTitle, data.PostContents, data.IsPinned, data.PostSeq)

	if editPost != nil {
		return editPost
	}

	return nil
}

func UpdatePostEdit(postTitle string, postContents string, isPinned string, postSeq string) error {
	connect, connectErr := database.InitDatabaseConnection()

	if connectErr != nil {
		return connectErr
	}

	// 데이터 업데이트
	_, resultErr := database.InsertQuery(
		connect,
		queries.UpdateEditPost, 
		postTitle, 
		postContents,
		isPinned,
		postSeq)
	
	defer connect.Close()

	if resultErr != nil {
		log.Printf("[EDIT] INSERT/UPDATE post Error: %v", resultErr)
		return resultErr
	}

	return nil
}

func InsertUpdateCategory(category string, isValidCategory bool) error {
	connect, connectErr := database.InitDatabaseConnection()

	if connectErr != nil {
		return connectErr
	}

	if isValidCategory {
			_, categoryErr := database.InsertQuery(connect, queries.InsertUpdateCategory, category)

		if categoryErr != nil {
			log.Printf("[EDIT] INSERT/UPDATE category data Error: %v", categoryErr)
			return categoryErr
		}
	} else {
		// 요청에 태그 데이터가 없다면 기존 카테고리 제거
		_, deleteCategoryErr := database.InsertQuery(connect, queries.DeletePostCategory, category)

		if deleteCategoryErr != nil {
			log.Printf("[EDIT] DELETE category data Error: %v", deleteCategoryErr)
			return deleteCategoryErr
		}
	}

	defer connect.Close()

	return nil
}

func InsertUpdateTagList(tagList []string, postSeq string) error {
	connect, connectErr := database.InitDatabaseConnection()

	if connectErr != nil {
		return connectErr
	}

	if len(tagList) > 0 {
		tagArray, _ := json.Marshal(tagList)

		_, tagQueryErr := database.InsertQuery(connect, queries.InsertTag, postSeq, string(tagArray))

		if tagQueryErr != nil {
			log.Printf("[EDIT] Insert Tag Data Error: %v", tagQueryErr)

			return tagQueryErr
		}
	} else {
		// 요청에 태그 데이터가 없다면 기존 태그 제거
		_, deleteTagErr := database.InsertQuery(connect, queries.DeletePostTag, postSeq)

		if deleteTagErr != nil {
			log.Printf("[EDIT] DELETE Tag data Error: %v", deleteTagErr)
			return deleteTagErr
		}
	}

	defer connect.Close()

	return nil
}

func InsertImageSeqList(imageSeqs []string, postSeq string) error {
	connect, connectErr := database.InitDatabaseConnection()

	if connectErr != nil {
		return connectErr
	}

	for _, seq := range(imageSeqs) {
		// 파일 데이터 업데이트
		_, insertUpdateErr := database.InsertQuery(connect, queries.InsertUpdatePostImage, postSeq, seq)

		if insertUpdateErr != nil {
			log.Printf("[EDIT] Insert Update File Data Error: %v", insertUpdateErr)
			return insertUpdateErr
		}
	}

	defer connect.Close()

	return nil
}
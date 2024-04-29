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
	connect, connectErr := database.InitDatabaseConnection()

	if connectErr != nil {
		return connectErr
	}

	// 카테고리 유효성 검증
	isValidCategory := utils.ValidateRequestValue(data.Category)

	if isValidCategory {
		_, categoryErr := database.InsertQuery(connect, queries.InsertUpdateCategory, data.Category)

		if categoryErr != nil {
			log.Printf("[EDIT] INSERT/UPDATE category data Error: %v", categoryErr)
			return categoryErr
		}
	} else {
		// 요청에 태그 데이터가 없다면 기존 카테고리 제거
		_, deleteCategoryErr := database.InsertQuery(connect, queries.DeletePostCategory, data.PostSeq)

		if deleteCategoryErr != nil {
			log.Printf("[EDIT] DELETE category data Error: %v", deleteCategoryErr)
			return deleteCategoryErr
		}
	}

	tags := data.Tags

	if len(tags) > 0 {
		tagArray, _ := json.Marshal(data.Tags)

		_, tagQueryErr := database.InsertQuery(connect, queries.InsertTag, data.PostSeq, string(tagArray))

		if tagQueryErr != nil {
			log.Printf("[EDIT] Insert Tag Data Error: %v", tagQueryErr)

			return tagQueryErr
		}
	} else {
		// 요청에 태그 데이터가 없다면 기존 태그 제거
		_, deleteTagErr := database.InsertQuery(connect, queries.DeletePostTag, data.PostSeq)

		if deleteTagErr != nil {
			log.Printf("[EDIT] DELETE Tag data Error: %v", deleteTagErr)
			return deleteTagErr
		}
	}

	for _, seq := range(data.ImageSeqs) {
		// 파일 데이터 업데이트
		_, insertUpdateErr := database.InsertQuery(connect, queries.InsertUpdatePostImage, data.PostSeq, seq)

		if insertUpdateErr != nil {
			log.Printf("[EDIT] Insert Update File Data Error: %v", insertUpdateErr)
			return insertUpdateErr
		}
	}

	// 데이터 업데이트
	_, resultErr := database.InsertQuery(
		connect,
		queries.UpdateEditPost, 
		data.PostTitle, 
		data.PostContents,
		data.IsPinned,
		data.PostSeq)
	
	if resultErr != nil {
		log.Printf("[EDIT] INSERT/UPDATE post Error: %v", resultErr)
		return resultErr
	}

	defer connect.Close()

	return nil
}
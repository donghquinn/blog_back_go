package postlib

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

	isValidCategory := utils.ValidateRequestValue(data.Category)

	if isValidCategory {
		_, categoryErr := database.InsertQuery(connect, queries.InsertUpdateCategory, data.Category)

		if categoryErr != nil {
			log.Printf("[EDIT] INSERT/UPDATE category data Error: %v", categoryErr)
			return categoryErr
		}
	}

	tags := data.Tags

	if len(tags) >0 {
		tagArray, _ := json.Marshal(data.Tags)

		_, tagQueryErr := database.InsertQuery(connect, queries.InsertTag, data.PostSeq, string(tagArray))

		if tagQueryErr != nil {
			log.Printf("[EDIT] Insert Tag Data Error: %v", tagQueryErr)

			return tagQueryErr
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

	_, resultErr := database.InsertQuery(
		connect,
		queries.UpdateEditPost, 
		data.PostTitle, 
		data.PostContents,
		data.IsPinned)
	
	if resultErr != nil {
		log.Printf("[EDIT] INSERT/UPDATE post Error: %v", resultErr)
		return resultErr
	}

	defer connect.Close()

	return nil
}
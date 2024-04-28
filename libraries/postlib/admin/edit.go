package postlib

import (
	"log"
	"strconv"

	"github.com/donghquinn/blog_back_go/libraries/database"
	queries "github.com/donghquinn/blog_back_go/queries/admin/posts"
	types "github.com/donghquinn/blog_back_go/types/admin/posts"
)

// 게시글 수정
func EditPost(data types.EditPostRequest, userId string) error {
	connect, connectErr := database.InitDatabaseConnection()

	if connectErr != nil {
		return connectErr
	}

	var categorySeq string

	if data.Category != "" {
		categoryId, categoryErr := database.InsertQuery(connect, queries.InsertUpdateCategory, data.Category)

		if categoryErr != nil {
			log.Printf("[EDIT] INSERT/UPDATE category data Error: %v", categoryErr)
			return categoryErr
		}

		categorySeq = strconv.Itoa(int(categoryId))
	}

	_, resultErr := database.InsertQuery(
		connect,
		queries.InsertUpdatePost, 
		data.PostTitle, 
		data.PostContents,
		categorySeq,
		data.IsPinned)
	
	if resultErr != nil {
		log.Printf("[EDIT] INSERT/UPDATE post Error: %v", resultErr)
		return resultErr
	}

	return nil
}
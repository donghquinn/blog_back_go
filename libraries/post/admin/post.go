package post

import (
	"log"

	"github.com/donghquinn/blog_back_go/libraries/database"
	queries "github.com/donghquinn/blog_back_go/queries/admin/posts"
	types "github.com/donghquinn/blog_back_go/types/admin/posts"
)

// 게시글 삭제
func DeletePost(data types.DeletePostRequest) error {
	connect, dbErr := database.InitDatabaseConnection()

	if dbErr != nil {
		return dbErr
	}

	_, deleteErr := database.InsertQuery(connect, queries.DeletePost, "0", data.PostSeq)

	if deleteErr != nil {
		log.Printf("[DELETE] Delete Post Error: %v", deleteErr)
		return deleteErr
	}

	defer connect.Close()

	return nil
}

// 고정 게시글 업데이트
func UpdatePinPost(data types.UpdatePinRequest) error {
	connect, dbErr := database.InitDatabaseConnection()

	if dbErr != nil {
		return dbErr
	}

	_, updateErr := database.InsertQuery(connect, queries.UpdatePinPost, "1", data.PostSeq)

	if updateErr != nil {
		log.Printf("[PIN] Update Pin Post Error: %v", updateErr)
		return updateErr
	}

	defer connect.Close()

	return nil
}

// 고정 게시글 해제
func UpdateUnPinPost(data types.UpdatePinRequest) error {
	connect, dbErr := database.InitDatabaseConnection()

	if dbErr != nil {
		return dbErr
	}

	_, updateErr := database.InsertQuery(connect, queries.UpdatePinPost, "0", data.PostSeq)

	if updateErr != nil {
		log.Printf("[PIN] Update Un-Pin Post Error: %v", updateErr)
		return updateErr
	}

	defer connect.Close()

	return nil
}


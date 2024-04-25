package postlib

import (
	"log"
	"strconv"

	"github.com/donghquinn/blog_back_go/libraries/database"
	queries "github.com/donghquinn/blog_back_go/queries/posts"
	"github.com/donghquinn/blog_back_go/types"
)

// 게시글 데이터 입력
func InsertPostData(registerPostRequest types.RegisterPostRequest, userId string) error {
	connect, dbErr := database.InitDatabaseConnection()

	if dbErr != nil {
		return dbErr
	}

	// 데이터 입력
	insertId, queryErr := database.InsertQuery(
		connect, 
		queries.InsertPost, 
		userId, 
		registerPostRequest.PostTitle, 
		registerPostRequest.PostContents,
		registerPostRequest.IsPinned)

	if queryErr != nil {
		log.Printf("[REGISTER] Insert Post Data Error: %v", queryErr)
		return queryErr
	}
	postSeq := strconv.Itoa(int(insertId))

	for _, t := range(registerPostRequest.Tags) {
		_, tagQueryErr := database.InsertQuery(connect, queries.InsertTag, postSeq, t)

		if tagQueryErr != nil {
			log.Printf("[REGISTER] Insert Tag Data Error: %v", tagQueryErr)

			return tagQueryErr
		}
	}

	for _, seq := range(registerPostRequest.ImageSeqs) {
		// 파일 데이터 업데이트
		_, insertUpdateErr := database.Query(connect, queries.InsertUpdatePostImage, postSeq, seq)

		if insertUpdateErr != nil {
			log.Printf("[REGISTER] Insert Update File Data Error: %v", insertUpdateErr)
			return insertUpdateErr
		}
	}

	defer connect.Close()

	return nil
}

func DeletePost(data types.DeletePostRequest, userId string) error {
	connect, dbErr := database.InitDatabaseConnection()

	if dbErr != nil {
		return dbErr
	}

	_, deleteErr := database.InsertQuery(connect, queries.DeletePost, "0", data.PostSeq, userId)

	if deleteErr != nil {
		log.Printf("[DELETE] Delete Post Error: %v", deleteErr)
		return deleteErr
	}

	return nil
}

func UpdatePinPost(data types.UpdatePinRequest, userId string) error {
	connect, dbErr := database.InitDatabaseConnection()

	if dbErr != nil {
		return dbErr
	}

	_, updateErr := database.InsertQuery(connect, queries.UpdatePinPost, "1", data.PostSeq, userId)

	if updateErr != nil {
		log.Printf("[PIN] Update Pin Post Error: %v", updateErr)
		return updateErr
	}

	return nil
}

func UpdateUnPinPost(data types.UpdatePinRequest, userId string) error {
	connect, dbErr := database.InitDatabaseConnection()

	if dbErr != nil {
		return dbErr
	}

	_, updateErr := database.InsertQuery(connect, queries.UpdatePinPost, "0", data.PostSeq, userId)

	if updateErr != nil {
		log.Printf("[PIN] Update Un-Pin Post Error: %v", updateErr)
		return updateErr
	}

	return nil
}

func GetPostTag(data types.GetPostsByTagRequest) ([]types.SelectPostsByTags, error) {
	connect, dbErr := database.InitDatabaseConnection()

	if dbErr != nil {
		return []types.SelectPostsByTags{}, dbErr
	}

	posts, selectErr := database.Query(connect, queries.SelectPostByTags, data.TagName)

	if selectErr != nil {
		log.Printf("[POST_TAG] GEt Post by TagName Error: %v", selectErr)
		return []types.SelectPostsByTags{}, selectErr
	}

	var postsData []types.SelectPostsByTags

	for posts.Next() {
		var row types.SelectPostsByTags

		scanErr := posts.Scan(
			&row.Tag_name,
			&row.Post_seq,
			&row.Post_title,
			&row.Viewed,
			&row.Reg_date,
			&row.Mod_date)
		
		if scanErr != nil {
			log.Printf("[POST_TAG] Scan Query Result Error: %v", scanErr)
			return []types.SelectPostsByTags{}, scanErr
		}

		postsData = append(postsData, row)
	}

	return postsData, nil
}
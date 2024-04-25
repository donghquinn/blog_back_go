package postlib

import (
	"fmt"
	"log"
	"strconv"

	"github.com/donghquinn/blog_back_go/libraries/database"
	queries "github.com/donghquinn/blog_back_go/queries/posts"
	"github.com/donghquinn/blog_back_go/types"
)

// 포스트들 가져오기 - 모듈함수
func QueryAllPostData(page int, size int) ([]types.SelectAllPostDataResult, error) {
		// parseBodyErr :=utils.DecodeBody(&req.Body)
	connect, dbErr := database.InitDatabaseConnection()

	if dbErr != nil {
		return nil, dbErr
	}

	// 페이징 파라미터 파싱
	result, queryErr := database.Query(connect, queries.SelectAllPosts,  fmt.Sprintf("%d", size), fmt.Sprintf("%d", (page - 1) * size))

	if queryErr != nil {
		log.Printf("[LIST] Get Post Data Error: %v", queryErr)

		return nil, queryErr
	}

	var queryResult = []types.SelectAllPostDataResult{}

	for result.Next() {
		var row types.SelectAllPostDataResult

		scanErr := result.Scan(
			&row.PostSeq,
			&row.PostTitle,
			&row.PostContents,
			&row.UserId,
			&row.UserName,
			&row.IsPinned,
			&row.Viewed,
			&row.RegDate,
			&row.ModDate)

		if scanErr != nil {
			log.Printf("[LIST] Scan and Assign Query Result Error: %v", scanErr)

			return nil, scanErr
		}

		queryResult = append(queryResult, row)
	}

	defer connect.Close()

	return queryResult, nil
}

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

	_, deleteErr := database.InsertQuery(connect, queries.DeletePost, "0", data.PostSeq)

	if deleteErr != nil {
		log.Printf("[DELETE] Delete Post Error: %v", deleteErr)
		return deleteErr
	}

	defer connect.Close()

	return nil
}

func UpdatePinPost(data types.UpdatePinRequest, userId string) error {
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

	defer connect.Close()

	return nil
}

func GetPostTag(data types.GetPostsByTagRequest, page int, size int) ([]types.SelectPostsByTags, error) {
	connect, dbErr := database.InitDatabaseConnection()

	if dbErr != nil {
		return []types.SelectPostsByTags{}, dbErr
	}

	posts, selectErr := database.Query(connect, queries.SelectPostByTags, data.TagName, fmt.Sprintf("%d", size), fmt.Sprintf("%d", (page - 1) * size))

	if selectErr != nil {
		log.Printf("[POST_TAG] GEt Post by TagName Error: %v", selectErr)
		return []types.SelectPostsByTags{}, selectErr
	}

	defer connect.Close()
	
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
package post

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/donghquinn/blog_back_go/libraries/database"
	queries "github.com/donghquinn/blog_back_go/queries/posts"
	types "github.com/donghquinn/blog_back_go/types/post"
)

// 포스트들 가져오기 - 모듈함수
func QueryAllPostData(page int, size int) ([]types.SelectAllPostDataResponse, error) {
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
	
	defer connect.Close()

	var queryResult = []types.SelectAllPostDataResponse{}

	for result.Next() {
		var row types.SelectAllPostDataResponse

		scanErr := result.Scan(
			&row.PostSeq,
			&row.PostTitle,
			&row.PostContents,
			&row.CategoryName,
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

	return queryResult, nil
}


// 게시글 태그로 조회
func GetPostByTag(data types.GetPostsByTagRequest, page int, size int) ([]types.PostsByTagsResponseType, error) {
	connect, dbErr := database.InitDatabaseConnection()

	if dbErr != nil {
		return []types.PostsByTagsResponseType{}, dbErr
	}

	tagArray, _ := json.Marshal(data.TagName)

	posts, selectErr := database.Query(connect, queries.SelectPostByTags, "%"+string(tagArray)+"%", fmt.Sprintf("%d", size), fmt.Sprintf("%d", (page - 1) * size))

	if selectErr != nil {
		log.Printf("[POST_TAG] GET Post by TagName Error: %v", selectErr)
		return []types.PostsByTagsResponseType{}, selectErr
	}

	defer connect.Close()

	var postsData []types.SelectPostsByTags

	// Array https://stackoverflow.com/questions/14477941/read-select-columns-into-string-in-go
	for posts.Next() {
		var row types.SelectPostsByTags

		scanErr := posts.Scan(
			&row.TagName,
			&row.CategoryName,
			&row.UserName,
			&row.PostSeq,
			&row.PostTitle,
			&row.PostContents,
			&row.Viewed,
			&row.RegDate,
			&row.ModDate)
		
		if scanErr != nil {
			log.Printf("[POST_TAG] Scan Query Result Error: %v", scanErr)
			return []types.PostsByTagsResponseType{}, scanErr
		}

		postsData = append(postsData, row)
	}

	// stringify된 array를 array로
	var postByTagsList []types.PostsByTagsResponseType

	for _, d := range(postsData) {
		var tempTag []string

		jsonErr :=  json.Unmarshal([]byte(d.TagName), &tempTag)

		if jsonErr != nil {
			log.Printf("[POST_TAG] Unmarshing Array Error: %v", jsonErr)
			return []types.PostsByTagsResponseType{}, jsonErr
		}

		data := types.PostsByTagsResponseType{
			TagName: tempTag,
			CategoryName: d.CategoryName,
			PostTitle: d.PostTitle,
			PostContents: d.PostContents,
			PostSeq: d.PostSeq,
			Viewed: d.Viewed,
			RegDate: d.RegDate,
			ModDate: d.ModDate}

		postByTagsList = append(postByTagsList, data)
	}

	return postByTagsList, nil
}


// 게시글 카테고리로 조회
func GetPostByCategory(data types.GetPostsByCategoryRequest, page int, size int) ([]types.PostByCategoryResponseType, error) {
	connect, dbErr := database.InitDatabaseConnection()

	if dbErr != nil {
		return []types.PostByCategoryResponseType{}, dbErr
	}

	posts, selectErr := database.Query(connect, queries.SelectPostByCategory, "%"+data.CategoryName+"%", fmt.Sprintf("%d", size), fmt.Sprintf("%d", (page - 1) * size))

	if selectErr != nil {
		log.Printf("[POST_CATEGORY] GET Post by CategoryName Error: %v", selectErr)
		return []types.PostByCategoryResponseType{}, selectErr
	}

	defer connect.Close()

	var postsData []types.SelectPostsByTags

	// Array https://stackoverflow.com/questions/14477941/read-select-columns-into-string-in-go
	for posts.Next() {
		var row types.SelectPostsByTags

		scanErr := posts.Scan(
			&row.TagName,
			&row.CategoryName,
			&row.UserName,
			&row.PostSeq,
			&row.PostTitle,
			&row.PostContents,
			&row.Viewed,
			&row.RegDate,
			&row.ModDate)
		
		if scanErr != nil {
			log.Printf("[POST_CATEGORY] Scan Query Result Error: %v", scanErr)
			return []types.PostByCategoryResponseType{}, scanErr
		}

		postsData = append(postsData, row)
	}

	// stringify된 array를 array로
	var postByCategoryList []types.PostByCategoryResponseType

	for _, d := range(postsData) {
		var tempTag []string

		jsonErr :=  json.Unmarshal([]byte(d.TagName), &tempTag)

		if jsonErr != nil {
			log.Printf("[POST_CATEGORY] Unmarshing Array Error: %v", jsonErr)
			return []types.PostByCategoryResponseType{}, jsonErr
		}

		data := types.PostByCategoryResponseType{
			TagName: tempTag,
			CategoryName: d.CategoryName,
			PostTitle: d.PostTitle,
			PostContents: d.PostContents,
			PostSeq: d.PostSeq,
			Viewed: d.Viewed,
			RegDate: d.RegDate,
			ModDate: d.ModDate}

		postByCategoryList = append(postByCategoryList, data)
	}

	return postByCategoryList, nil
}
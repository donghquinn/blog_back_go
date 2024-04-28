package postlib

import (
	"log"

	"github.com/donghquinn/blog_back_go/libraries/database"
	queries "github.com/donghquinn/blog_back_go/queries/posts"
	types "github.com/donghquinn/blog_back_go/types/post"
)

// 특정 게시글 가져오기
func GetPostData(postSeq string) (types.SelectSpecificPostDataResult, []types.SelectSpeicificPostTagDataResult, error){
	connect, connectErr := database.InitDatabaseConnection()

	if connectErr != nil {
		log.Printf("[CONTENTS] Init Database Connection Error for Post Data: %v", connectErr)
		return types.SelectSpecificPostDataResult{}, []types.SelectSpeicificPostTagDataResult{},connectErr
	}

	// 조회수 업데이트
	_, updateErr := database.InsertQuery(connect, queries.UpdateViewCount, postSeq)

	if updateErr != nil {
		log.Printf("[CONTENTS] Update View Count Error: %v", updateErr)
		return types.SelectSpecificPostDataResult{},[]types.SelectSpeicificPostTagDataResult{}, updateErr
	}

	// 특정 게시글 조회
	result, queryErr := database.QueryOne(connect, queries.SelectSpecificPostContents, postSeq)

	if queryErr != nil {
		log.Printf("[CONTENTS] Query A Post Contents Error: %v", queryErr)
		return types.SelectSpecificPostDataResult{}, []types.SelectSpeicificPostTagDataResult{}, queryErr
	}

	// 태그들 조회
	tagResult, tagErr := database.Query(connect, queries.SelectPostTags, postSeq)

	if tagErr != nil {
		log.Printf("[CONTENTS] Query Tags Error: %v", tagErr)
		return types.SelectSpecificPostDataResult{}, []types.SelectSpeicificPostTagDataResult{}, tagErr
	}

	defer connect.Close()

	var queryResult types.SelectSpecificPostDataResult

	postScanErr := result.Scan(
		&queryResult.PostSeq,
		&queryResult.PostTitle, 
		&queryResult.PostContents, 
		&queryResult.PostStatus,
		&queryResult.UserId, 
		&queryResult.UserName,
		&queryResult.Viewed,
		&queryResult.IsPinned,
		&queryResult.RegDate,
		&queryResult.ModDate)

	if postScanErr != nil {
		log.Printf("[CONTENTS] Can Post Data Error: %v", postScanErr)
		return queryResult, []types.SelectSpeicificPostTagDataResult{}, postScanErr
	}

	// 태그 쿼리
	var tagQueryResult []types.SelectSpeicificPostTagDataResult

	for tagResult.Next() {
		var row types.SelectSpeicificPostTagDataResult

		scanErr := tagResult.Scan(
			&row.TagName)

		if scanErr != nil {
			log.Printf("[CONTENTS] Scan Tag Query Data Error: %v", scanErr)
			return queryResult, tagQueryResult, scanErr
		}

		tagQueryResult = append(tagQueryResult, row)
	}

	return queryResult, tagQueryResult, nil
}
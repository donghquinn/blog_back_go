package types

// 게시글 전체 가져오기 쿼리 결과 타입
type SelectAllPostDataResult struct {
	PostTitle string
	PostContents string
	UserId string
	RegDate string
	ModDate string
}

// 게시글 등록 요청
type RegisterPostRequest struct {
	PostTitle string `json:"postTitle"`
	PostContents string `json:"postContents"`
}
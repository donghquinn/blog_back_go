package types

// 게시글 전체 가져오기 쿼리 결과 타입
type SelectAllPostDataResult struct {
	PostSeq string
	PostTitle string
	PostContents string
	UserId string
	UserName string
	RegDate string
	ModDate string
	// VersionId []string
}

type SelectSpecificPostDataResult struct {
	PostSeq string
	PostTitle string
	PostContents string
	PostStatus string
	UserId string
	UserName string
	RegDate string
	ModDate string
}

type SelectPostImageData struct {
	ObjectName string
	FileFormat string
	TargetSeq string
}

type ViewSpecificPostContentsResponse struct {
	PostSeq string `json:"postSeq"`
	PostTitle string `json:"postTitle"`
	PostContents string `json:"postContents"`
	UserId string `json:"userId"`
	UserName string `json:"userName"`
	Urls []string `json:"urls"`
	RegDate string `json:"regDate"`
	ModDate string `json:"modDate"`
}

// 게시글 등록 요청
type RegisterPostRequest struct {
	PostTitle string `json:"postTitle"`
	PostContents string `json:"postContents"`
	ImageSeqs  []string `json:"imageSeqs"`
}

type ViewPostContents struct {
	PostSeq string `json:"postSeq"`
}
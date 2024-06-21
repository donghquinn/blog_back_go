package types

// 카테고리 이름으로 게시글 조회
type PostByCategoryResponseType struct {
	TagName []string		`json:"tagName"`
	CategoryName string `json:"category"`
	PostSeq string		`json:"postSeq"`
	PostTitle string	`json:"postTitle"`
	PostContents string `json:"postContents"`
	Viewed string	`json:"viewed"`
	RegDate string	`json:"regDate"`
	ModDate string	`json:"modDate"`
}

// 태그로 특정 게시글 조회 응답
type PostsByTagsResponseType struct {
	TagName []string		`json:"tagName"`
	CategoryName string `json:"category"`
	PostSeq string		`json:"postSeq"`
	PostTitle string	`json:"postTitle"`
	PostContents string  `json:"postContents"`
	Viewed string	`json:"viewed"`
	RegDate string	`json:"regDate"`
	ModDate string	`json:"modDate"`
}


// 게시글 전체 가져오기 쿼리 결과 타입
type SelectAllPostDataResponse struct {
	PostSeq string	`json:"postSeq"`
	PostTitle string	`json:"postTitle"`
	PostContents string	`json:"postContents"`
	CategoryName *string `json:"categoryName"`
	UserName string	`json:"userName"`
	IsPinned string	`json:"isPinned"`
	Viewed string	`json:"viewed"`
	RegDate string	`json:"regDate"`
	ModDate string	`json:"modDate"`
	// VersionId []string
}


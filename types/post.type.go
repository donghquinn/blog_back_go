package types


type ViewSpecificPostContentsResponse struct {
	PostSeq string `json:"postSeq"`
	PostTitle string `json:"postTitle"`
	PostContents string `json:"postContents"`
	Tags []string `json:"tags"`
	UserId string `json:"userId"`
	UserName string `json:"userName"`
	Viewed string `json:"viewed"`
	IsPinned string `json:"isPinned"`
	Urls []string `json:"urls"`
	RegDate string `json:"regDate"`
	ModDate string `json:"modDate"`
}


// 상세 포스트 조회 요청
type ViewPostContents struct {
	PostSeq string `json:"postSeq" binding:"required"`
}

// 태그로 특정 게시글 조희 요청
type GetPostsByTagRequest struct {
	TagName string `json:"tag" binding:"required"`
}

// 카테고리 이름으로 게시글 조회
type GetPostsByCategoryRequest struct {
	CategoryName string `json:"categoryName" binding:"required"`
}


// 게시글 전체 가져오기 쿼리 결과 타입
type SelectAllPostDataResult struct {
	PostSeq string	`json:"postSeq"`
	PostTitle string	`json:"postTitle"`
	PostContents string	`json:"postContents"`
	CategoryName string `json:"categoryName"`
	UserName string	`json:"userName"`
	IsPinned string	`json:"isPinned"`
	Viewed string	`json:"viewed"`
	RegDate string	`json:"regDate"`
	ModDate string	`json:"modDate"`
	// VersionId []string
}

// 특정 게시글 데이터 쿼리 결과
type SelectSpecificPostDataResult struct {
	PostSeq string
	PostTitle string
	PostContents string
	PostStatus string
	UserId string
	UserName string
	Viewed string
	IsPinned string
	RegDate string
	ModDate string
}

// 특정 게시글 태그 데이터 쿼리 결과
type SelectSpeicificPostTagDataResult struct {
	TagName string
}

// 태그로 특정 게시글 쿼리
type SelectPostsByTags struct {
	TagName string		
	CategoryName string
	UserName string
	PostSeq string		
	PostTitle string	
	PostContents string
	Viewed string	
	RegDate string	
	ModDate string 
}

type PostByCategoryResponse struct {
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
type PostsByTagsResponse struct {
	TagName []string		`json:"tagName"`
	CategoryName string `json:"category"`
	PostSeq string		`json:"postSeq"`
	PostTitle string	`json:"postTitle"`
	PostContents string  `json:"postContents"`
	Viewed string	`json:"viewed"`
	RegDate string	`json:"regDate"`
	ModDate string	`json:"modDate"`
}

// 이미지 데이터 가져오기
type SelectPostImageData struct {
	ObjectName string
	FileFormat string
	TargetPurpose string
	TargetSeq string
}

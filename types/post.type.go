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

// 게시글 등록 요청
type RegisterPostRequest struct {
	PostTitle string `json:"postTitle" binding:"required"`
	PostContents string `json:"postContents" binding:"required"`
	ImageSeqs  []string `json:"imageSeqs" binding:"required"`
	Tags []string `json:"tags" binding:"required"`
	IsPinned string `json:"isPinned" binding:"required"`
}

// 상세 포스트 조회 요청
type ViewPostContents struct {
	PostSeq string `json:"postSeq" binding:"required"`
}

// 게시글 삭제 요청
type DeletePostRequest struct {
	PostSeq string `json:"postSeq" binding:"required"`
}

// 고정 요청
type UpdatePinRequest struct {
	PostSeq string `json:"postSeq" binding:"required"`
}

type GetPostsByTagRequest struct {
	TagName string `json:"tag" binding:"required"`
}

// 게시글 전체 가져오기 쿼리 결과 타입
type SelectAllPostDataResult struct {
	PostSeq string	`json:"postSeq"`
	PostTitle string	`json:"postTitle"`
	PostContents string	`json:"postContents"`
	UserId string	`json:"userId"`
	UserName string	`json:"userName"`
	IsPinned string	`json:"isPinned"`
	Viewed string	`json:"viewed"`
	RegDate string	`json:"regDate"`
	ModDate string	`json:"modDate"`
	// VersionId []string
}

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

type SelectSpeicificPostTagDataResult struct {
	TagName string
}

type SelectPostsByTags struct {
	Tag_name string		`json:"tagName"`
	Post_seq string		`json:"postSeq"`
	Post_title string	`json:"postTitle"`
	Viewed string	`json:"viewed"`
	Reg_date string	`json:"regDate"`
	Mod_date string	`json:"modDate"`
}

type SelectPostImageData struct {
	ObjectName string
	FileFormat string
	TargetPurpose string
	TargetSeq string
}

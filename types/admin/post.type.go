package types

// 게시글 등록 요청
type RegisterPostRequest struct {
	PostTitle string `json:"postTitle" binding:"required"`
	PostContents string `json:"postContents" binding:"required"`
	ImageSeqs  []string `json:"imageSeqs" binding:"required"`
	Tags []string `json:"tags" binding:"required"`
	IsPinned string `json:"isPinned" binding:"required"`
}

// 게시글 삭제 요청
type DeletePostRequest struct {
	PostSeq string `json:"postSeq" binding:"required"`
}

// 고정 요청
type UpdatePinRequest struct {
	PostSeq string `json:"postSeq" binding:"required"`
}
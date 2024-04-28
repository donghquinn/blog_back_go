package types

// 게시글 등록 요청
type RegisterPostRequest struct {
	PostTitle string `json:"postTitle" binding:"required"`
	PostContents string `json:"postContents" binding:"required"`
	ImageSeqs  []string `json:"imageSeqs" binding:"optional"`
	Category string `json:"category" binding:"optional"`
	Tags []string `json:"tags" binding:"optional"`
	IsPinned string `json:"isPinned" binding:"required"`
}

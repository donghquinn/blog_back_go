package types

type EditPostRequest struct {
	PostSeq string `json:"postSeq" binding:"required"`
	PostTitle string `json:"postTitle" binding:"required"`
	PostContents string `json:"postContents" binding:"required"`
	Category 	string `json:"category" binding:"optional"`
	ImageSeqs  []string `json:"imageSeqs" binding:"required"`
	Tags []string `json:"tags" binding:"required"`
	IsPinned string `json:"isPinned" binding:"required"`
}

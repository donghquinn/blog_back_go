package types

// 태그로 특정 게시글 조희 요청
type GetPostsByTagRequest struct {
	TagName string `json:"tag" binding:"required"`
}

// 카테고리 이름으로 게시글 조회
type GetPostsByCategoryRequest struct {
	CategoryName string `json:"category" binding:"required"`
}


type GetPostByPostSeq struct {
	PostSeq string `json:"postSeq"`
}
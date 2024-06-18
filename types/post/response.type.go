package types

// 게시글 리스트 응답 구조체
type ResponsePostContentsType struct {
	Code string `json:"code"`
	Result bool `json:"result"`
	PostList ViewSpecificPostContentsResponse `json:"postList"`
}

// 게시글 리스트 응답 구조체
type ResponsePostListType struct {
	Code string `json:"code"`
	Result bool `json:"result"`
	PostList []SelectAllPostDataResponse `json:"postList"`
}

// 게시글 리스트 응답 구조체
type ResponsePostByTagListType struct {
	Code string `json:"code"`
	Result bool `json:"result"`
	PostList []PostsByTagsResponseType `json:"postList"`
}

// 카테고리로 게시글 조회
type ResponsePostByCategoryListType struct {
	Code string `json:"code"`
	Result bool `json:"result"`
	PostList []PostByCategoryResponseType `json:"postList"`
}

// 게시글 리스트 응답 구조체
type ResponseInsertIdType struct {
	Code string `json:"code"`
	Result bool `json:"result"`
	InsertId string `json:"insertId"`
}

package types

// 게시글 리스트 응답 구조체
type ResponsePostContentsType struct {
	Code string `json:"code"`
	Status bool `json:"status"`
	Posts ViewSpecificPostContentsResponse `json:"posts"`
}

// 게시글 리스트 응답 구조체
type ResponsePostListType struct {
	Code string `json:"code"`
	Status bool `json:"status"`
	Result []SelectAllPostDataResponse `json:"result"`
}

// 게시글 리스트 응답 구조체
type ResponsePostByTagListType struct {
	Code string `json:"code"`
	Status bool `json:"status"`
	Result []PostsByTagsResponseType `json:"result"`
}

// 카테고리로 게시글 조회
type ResponsePostByCategoryListType struct {
	Code string `json:"code"`
	Status bool `json:"status"`
	Result []PostByCategoryResponseType `json:"result"`
}

// 게시글 리스트 응답 구조체
type ResponseInsertIdType struct {
	Code string `json:"code"`
	Status bool `json:"status"`
	InsertId string `json:"insertId"`
}

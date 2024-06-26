package types

// 게시글 리스트 응답 구조체
type ResponsePostContentsType struct {
	Code string `json:"code"`
	Result bool `json:"result"`
	PostList ViewSpecificPostContentsResponse `json:"postList"`
}

type ResponsePostRegisterType struct {
	Code string `json:"code"`
	Result bool `json:"result"`
	PostSeq int64 `json:"postSeq"`
}

// 게시글 리스트 응답 구조체
type ResponsePostListType struct {
	Code string `json:"code"`
	Result bool `json:"result"`
	PinnedPostList []SelectAllPostDataResponse `json:"pinnedPostList"`
	UnpinnedPostList []SelectAllPostDataResponse `json:"unpinnedPostList"`
	PostCount string `json:"postCount"`
	Page int `json:"page"`
	Size int `json:"size"`
}

type ResponsePinnedPostListType struct {
	Code string `json:"code"`
	Result bool `json:"result"`
	PinnedPostList []SelectAllPostDataResponse `json:"pinnedPostList"`
	PostCount string `json:"postCount"`
	Page int `json:"page"`
	Size int `json:"size"`
}


// 게시글 리스트 응답 구조체
type ResponsePostByTagListType struct {
	Code string `json:"code"`
	Result bool `json:"result"`
	PostList []PostsByTagsResponseType `json:"postList"`
	PostCount string `json:"postCount"`
}

// 카테고리로 게시글 조회
type ResponsePostByCategoryListType struct {
	Code string `json:"code"`
	Result bool `json:"result"`
	PostList []PostByCategoryResponseType `json:"postList"`
	PostCount string `json:"postCount"`
}

// 게시글 리스트 응답 구조체
type ResponseInsertIdType struct {
	Code string `json:"code"`
	Result bool `json:"result"`
	InsertId string `json:"insertId"`
}

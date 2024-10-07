package types

// 게시글 리스트 응답 구조체
type ResponsePostContentsType struct {
	Status   int                              `json:"status"`
	Code     string                           `json:"code"`
	Result   bool                             `json:"result"`
	PostList ViewSpecificPostContentsResponse `json:"postList"`
	Message  string                           `json:"message"`
}

// 게시글 리스트 응답 구조체
type ResponsePostListType struct {
	Status           int                         `json:"status"`
	Code             string                      `json:"code"`
	Result           bool                        `json:"result"`
	PinnedPostList   []SelectAllPostDataResponse `json:"pinnedPostList"`
	UnpinnedPostList []SelectAllPostDataResponse `json:"unpinnedPostList"`
	PostCount        string                      `json:"postCount"`
	Page             int                         `json:"page"`
	Size             int                         `json:"size"`
	Message          string                      `json:"message"`
}

type ResponsePinnedPostListType struct {
	Status         int                         `json:"status"`
	Code           string                      `json:"code"`
	Result         bool                        `json:"result"`
	PinnedPostList []SelectAllPostDataResponse `json:"pinnedPostList"`
	PostCount      string                      `json:"postCount"`
	Page           int                         `json:"page"`
	Size           int                         `json:"size"`
	Message        string                      `json:"message"`
}

// 게시글 리스트 응답 구조체
type ResponsePostByTagListType struct {
	Status    int                       `json:"status"`
	Code      string                    `json:"code"`
	Result    bool                      `json:"result"`
	PostList  []PostsByTagsResponseType `json:"postList"`
	PostCount string                    `json:"postCount"`
	Message   string                    `json:"message"`
}

// 카테고리로 게시글 조회
type ResponsePostByCategoryListType struct {
	Status    int                          `json:"status"`
	Code      string                       `json:"code"`
	Result    bool                         `json:"result"`
	PostList  []PostByCategoryResponseType `json:"postList"`
	PostCount string                       `json:"postCount"`
	Message   string                       `json:"message"`
}

// 게시글 리스트 응답 구조체
type ResponseInsertIdType struct {
	Status   int    `json:"status"`
	Code     string `json:"code"`
	Result   bool   `json:"result"`
	InsertId string `json:"insertId"`
	Message  string `json:"message"`
}

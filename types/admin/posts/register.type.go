package types

// 게시글 등록 요청
type RegisterPostRequest struct {
	PostTitle    string   `json:"postTitle"`
	PostContents string   `json:"postContents"`
	ImageSeqs    []string `json:"imageSeqs,omitempty"`
	Category     string   `json:"category,omitempty"`
	Tags         []string `json:"tags,omitempty"`
	IsPinned     string   `json:"isPinned"`
	// IsSecret int `json:"isSecret"`
}

type ResponsePostRegisterType struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Result  bool   `json:"result"`
	PostSeq int64  `json:"postSeq"`
	Message string `json:"message"`
}

package types

// 기본 응답 구조체
type ResponseType struct {
	Code string `json:"code"`
	Status bool `json:"status"`
}

// 게시글 리스트 응답 구조체
type ResponseInsertIdType struct {
	Code string `json:"code"`
	Status bool `json:"status"`
	InsertId string `json:"insertId"`
}

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
	Result []SelectAllPostDataResult `json:"result"`
}

// 메세지를 담은 응답
type ResponseMessageType struct {
	Code string `json:"code"`
	Status bool `json:"status"`
	Message string	`json:"message"`
}

// 찾은 이메일 담은 응답
type ResponseFoundEmailType struct {
	Code string `json:"code"`
	Status bool `json:"status"`
	Email string	`json:"email"`
}


// 찾은 패스워드 담은 응답
type ResponseFoundPasswordType struct {
	Code string `json:"code"`
	Status bool `json:"status"`
	Password string	`json:"password"`
}


// JWT 토큰을 담은 응답
type ResponseTokenType struct {
	Code string `json:"code"`
	Status bool `json:"status"`
	Token string `json:"token"`
}

// 에러 응답 타입
type ErrorResponseType struct {
	Code string `json:"code"`
	Status bool `json:"status"`
	Message string	`json:"message"`
}

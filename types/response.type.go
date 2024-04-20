package types

type ResponseType struct {
	Code string `json:"code"`
	Status bool `json:"status"`
}

type ResponsePostListType struct {
	Code string `json:"code"`
	Status bool `json:"status"`
	Result []SelectAllPostData `json:"result"`
}


// 메세지를 담은 응답
type ResponseMessageType struct {
	Code string `json:"code"`
	Status bool `json:"status"`
	Message string	`json:"message"`
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

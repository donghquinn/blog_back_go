package types

// 기본 응답 구조체
type ResponseType struct {
	Status int    `json:"status"`
	Code   string `json:"code"`
	Result bool   `json:"result"`
}

type ResponseSignupType struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Result  bool   `json:"result"`
	BlogId  string `json:"blogId"`
	Message string `json:"message"`
}

type ResponseImageUrl struct {
	Status      int      `json:"status"`
	Code        string   `json:"code"`
	Result      bool     `json:"result"`
	ImageResult []string `json:"imageResult"`
}

// 메세지를 담은 응답
type ResponseMessageType struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Result  bool   `json:"result"`
	Message string `json:"message"`
}

// 찾은 이메일 담은 응답
type ResponseFoundEmailType struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Result  bool   `json:"result"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

// 찾은 패스워드 담은 응답
type ResponseFoundPasswordType struct {
	Status   int    `json:"status"`
	Code     string `json:"code"`
	Result   bool   `json:"result"`
	Password string `json:"password"`
}

// JWT 토큰을 담은 응답
type ResponseTokenType struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Result  bool   `json:"result"`
	Token   string `json:"token"`
	Message string `json:"message"`
}

// 에러 응답 타입
type ErrorResponseType struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Result  bool   `json:"result"`
	Message string `json:"message"`
}

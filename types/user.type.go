package types

// 회원가입 요청
type UserSignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	BlogId   string `json:"blogId"`
}

// 로그인 요청
type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// 유저 프로필 조회 요청
type UserProfileRequest struct {
	UserId string `json:"userId"`
}

// 이메일 찾기 요청
type UserSearchEmailRequest struct {
	Name string `json:"name"`
}

// 패스워드 찾기 요청
type UserSearchPasswordRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// 로그인 쿼리 결과
type UserLoginQueryResult struct {
	UserId       string
	UserPassword string
	UserStatus   string
	BlogId       string
}

// 유저 이메일 찾기 쿼리 결과
type SelectUserSearchEmailResult struct {
	UserEmail string
}

// 유저 패스워드 찾기 쿼리 결과
type SelectUserSearchPasswordResult struct {
	UserPassword string
}

type LoginRedisStruct struct {
	Email      string `json:"email"`
	UserStatus string `json:"userStatus"`
	UserId     string `json:"userId"`
}

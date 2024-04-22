package types

// 회원가입 요청
type UserSignupRequest struct {
	Email string 		`json:"email"`
	Password string 	`json:"password"`
	Name string 		`json:"name"`
 }

 // 로그인 요청
 type UserLoginRequest struct {
	Email string 		`json:"email"`
	Password string 	`json:"password"`
 }


 // 이메일 찾기 요청
 type UserSearchEmailRequest struct {
	Name string `json:"name"`
 }

 // 로그인 쿼리 결과
 type UserLoginQueryResult struct {
	UserId string	
	UserPassword string
	UserStatus string
 }

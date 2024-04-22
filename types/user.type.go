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

  // 패스워드 찾기 요청
 type UserSearchPasswordRequest struct {
	Email string `json:"email"`
	Name string `json:"name"`
 }


 // 로그인 쿼리 결과
 type UserLoginQueryResult struct {
	UserId string	
	UserPassword string
	UserStatus string
 }

 // 유저 이메일 찾기 쿼리 결과
 type SelectUserSearchEmailResult struct {
	UserEmail string
 }

 // 유저 패스워드 찾기 쿼리 결과
 type SelectUserSearchPasswordResult struct {
	UserPassword string
 }
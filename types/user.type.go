package types

// 회원가입 요청
type UserSignupRequest struct {
	Email string 		`json:"email" binding:"required"`
	Password string 	`json:"password" binding:"required"`
	Name string 		`json:"name" binding:"required"`
 }

 // 로그인 요청
 type UserLoginRequest struct {
	Email string 		`json:"email" binding:"required"`
	Password string 	`json:"password" binding:"required"`
 }
 
// 유저 프로필 조회 요청
type UserProfileRequest struct {
	UserId string `json:"userId"`
}

 // 이메일 찾기 요청
 type UserSearchEmailRequest struct {
	Name string `json:"name" binding:"required"`
 }

  // 패스워드 찾기 요청
 type UserSearchPasswordRequest struct {
	Email string `json:"email" binding:"required"`
	Name string `json:"name" binding:"required"`
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

type LoginRedisStruct struct {
	Name string `json:"name"`
	Email string `json:"email"`
	UserId string `json:"userId"`
}

 // 유저 프로필 쿼리 결과
 type SelectUserProfileQueryResult struct {
	UserId string
	UserEmail string
	UserName string
	Color string
	Title string
	GithubUrl string
	PersonalUrl string
	Memo string
 }
 
 // 파일 데이터 쿼리 결과
 type SelectFileQueryResult struct {
	FileFormat string
	FileType string
	TargetPurpose string
	TargetId string
	ObjectName string
 }

 // 유저 프로필 이미지 파일 응답
 type UserImageFileData struct {
	ProfileImage string	`json:"profileImage"`
	BackgroundImage string	`json:"backgroundImage"`
 }

 // 우저 프로필 데이터 응답
 type UserProfileDataResponseType struct {
	UserId string		`json:"userId"`
	UserEmail string	`json:"userEmail"`
	UserName string		`json:"userName"`
	Color string		`json:"color"`
	Title string		`json:"title"`
	GithubUrl string	`json:"githubUrl"`
	PersonalUrl string	`json:"personalUrl"`
	Memo string			`json:"memo"`
	Images UserImageFileData	`json:"images"`
 }
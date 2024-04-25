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

 // 프로필 색 변경 요청
 type UserUpdateProfileColorRequest struct {
	Color string `json:"color" binding:"required"`
 }

// 유저 패스워드 변경 요청
 type UserChangePasswordRequest struct {
	Password string `json:"password" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
 }

// 블로그 타이틀 변경 요청
 type UserUpdateBlogTitleRequest struct {
	Title string `json:"title" binding:"required"`
 }

 // 프로필 수정
 type UserChangeProfileRequest struct {
	Name string `json:"name" binding:"optional"`
	Color string `json:"color" binding:"optional"`
	Title string `json:"title" binding:"optional"`
	BackgroundImage string `json:"backgroundImage" binding:"optional"`
	ProfileImage string `json:"profileImage" binding:"optional"`
	Memo string `json:"memo" binding:"optional"`
	Instagram string `json:"instagram" binding:"optional"`
	GithubUrls string `json:"githubUrls" binding:"optional"`
	PersonalUrls string `json:"personalUrls" binding:"optional"`
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
 
 type SelectFileQueryResult struct {
	FileFormat string
	FileType string
	TargetPurpose string
	TargetId string
	ObjectName string
 }

 type UserImageFileData struct {
	ProfileImage string	`json:"profileImage"`
	BackgroundImage string	`json:"backgroundImage"`
 }

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
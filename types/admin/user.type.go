package types

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

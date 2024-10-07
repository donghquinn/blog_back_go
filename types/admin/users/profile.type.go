package types

type UserGetProfileRequest struct {
	UserId string `json:"userId"`
	BlogId string `json:"blogId"`
}

// 프로필 색 변경 요청
type UserUpdateProfileColorRequest struct {
	Color string `json:"color"`
}

// 블로그 타이틀 변경 요청
type UserUpdateBlogTitleRequest struct {
	Title string `json:"title"`
}

// 프로필 수정
type UserChangeProfileRequest struct {
	Name         string `json:"name,omitempty"`
	Color        string `json:"color,omitempty"`
	Title        string `json:"title,omitempty"`
	Instagram    string `json:"instagram,omitempty"`
	GithubUrls   string `json:"githubUrl,omitempty"`
	PersonalUrls string `json:"personalUrl,omitempty"`
	Memo         string `json:"memo,omitempty"`
	// BackgroundImage string `json:"backgroundImage" binding:"optional"`
	// ProfileImage string `json:"profileImage" binding:"optional"`
}

// 유저 프로필 쿼리 결과
type SelectUserProfileQueryResult struct {
	UserId      string
	UserEmail   string
	UserName    string
	Color       string
	Title       string
	Instagram   *string
	GithubUrl   *string
	PersonalUrl *string
	Memo        *string
}

// 파일 데이터 쿼리 결과
type SelectFileQueryResult struct {
	FileFormat    string
	FileType      string
	TargetPurpose string
	TargetId      string
	ObjectName    string
}

// 유저 프로필 이미지 파일 응답
type UserImageFileData struct {
	ProfileImage    string `json:"profileImage"`
	BackgroundImage string `json:"backgroundImage"`
}

// 우저 프로필 데이터 응답
type UserProfileDataResponseType struct {
	UserId      string            `json:"userId"`
	UserEmail   string            `json:"userEmail"`
	UserName    string            `json:"userName"`
	Color       string            `json:"color"`
	Title       string            `json:"title"`
	Instagram   *string           `json:"instagram"`
	GithubUrl   *string           `json:"githubUrl"`
	PersonalUrl *string           `json:"personalUrl"`
	Memo        *string           `json:"memo"`
	Images      UserImageFileData `json:"images"`
}

// 유저 프로필 응답 구조체
type ResponseProfileType struct {
	Status        int                         `json:"status"`
	Code          string                      `json:"code"`
	Result        bool                        `json:"result"`
	ProfileResult UserProfileDataResponseType `json:"profileResult"`
	Message       string                      `json:"message"`
}

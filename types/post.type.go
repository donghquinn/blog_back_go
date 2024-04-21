package types

type SelectAllPostData struct {
	PostTitle string `json:"postTitle"`
	PostContents string `json:"postContents"`
	UserId string `json:"user_id"`
	RegDate string `json:"regDate"`
	ModDate string `json:"modDate"`
}

type RegisterPostRequest struct {
	PostTitle string `json:"postTitle"`
	PostContents string `json:"postContents"`
}
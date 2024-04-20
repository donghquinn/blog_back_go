package types

type SelectAllPostData struct {
	Post_title string `json:"postTitle"`
	Post_contents string `json:"postContents"`
	User_id string `json:"user_id"`
	Reg_date string `json:"regDate"`
	Mod_date string `json:"modDate"`
}
package entity

type PostTableEntity struct {
	Post_seq int
	User_id string
	Post_title string
	Post_contents string
	Post_status int
	Viewed int
	Is_pinned int
	Is_secret int
	Blog_owner string
	Reg_date string
	Mod_date string
}

package types

type UserSignupRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
	Name string `json:"name"`
 }

 type UserLoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
 }

 type UserLoginQueryResult struct {
	User_id string
	User_password string
	User_status string
 }
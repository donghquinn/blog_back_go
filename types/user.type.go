package types

type UserSignupRequest struct {
	UserEmail string `json:"userEmail"`
	UserPassword string `json:"userPassword"`
	UserName string `json:"userName"`
 }

 type UserLoginRequest struct {
	UserEmail string `json:"userEmail"`
	UserPassword string `json:"userPassword"`
 }

 type UserLoginQueryResult struct {
	User_id string
	User_password string
	User_status string
 }
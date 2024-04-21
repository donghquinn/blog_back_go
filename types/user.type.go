package types

type UserSignupRequest struct {
	UserEmail string `json:"userEmail"`
	UserPassword string `json:"userPassword"`
	UserName string `json:"userName"`
 }
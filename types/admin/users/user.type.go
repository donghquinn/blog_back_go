package types

// 유저 패스워드 변경 요청
 type UserChangePasswordRequest struct {
	Password string `json:"password" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
 }


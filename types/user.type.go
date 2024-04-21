package types

type UserSignupRequest struct {
	Email string 		`json:"email"`
	Password string 	`json:"password"`
	Name string 		`json:"name"`
 }

 type UserLoginRequest struct {
	Email string 		`json:"email"`
	Password string 	`json:"password"`
 }

 type UserLoginQueryResult struct {
	UserId string			`json:"user_id"`
	UserPassword string	`json:"user_password"`
	UserStatus string		`json:"user_status"`
 }
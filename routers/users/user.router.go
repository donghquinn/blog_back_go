package routers

import (
	"net/http"

	"github.com/donghquinn/blog_back_go/controllers/users"
)


func UserRouter(server *http.ServeMux) {
	server.HandleFunc("POST /user/signup", users.SignupController)
	server.HandleFunc("POST /user/login", users.LoginController)

	server.HandleFunc("POST /user/search/email", users.SearchEmailController)
	server.HandleFunc("POST /user/search/password", users.SearchPasswordController)

	server.HandleFunc("POST /user/profile", users.GetUserProfileController)
	server.HandleFunc("POST /user/profile/update", users.UpdateProfileController)
	server.HandleFunc("POST /user/profile/title", users.UpdateTitleController)
	server.HandleFunc("POST /user/profile/color", users.UpdateColorController)
}
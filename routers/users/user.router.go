package routers

import (
	"net/http"

	controllers "github.com/donghquinn/blog_back_go/controllers/users"
)

func UserRouter(server *http.ServeMux) {
	server.HandleFunc("POST /user/signup", controllers.SignupController)
	server.HandleFunc("POST /user/login", controllers.LoginController)

	server.HandleFunc("POST /user/search/email", controllers.SearchEmailController)
	// server.HandleFunc("POST /user/search/password", controllers.SearchPasswordController)

	server.HandleFunc("POST /user/profile", controllers.GetUserProfileController)
}

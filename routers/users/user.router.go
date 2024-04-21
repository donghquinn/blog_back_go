package routers

import (
	"net/http"

	"github.com/donghquinn/blog_back_go/controllers/users"
)


func UserRouter(server *http.ServeMux) {
	server.HandleFunc("POST /user/signup", users.SignupController)
	server.HandleFunc("POST /user/login", users.LoginController)
}
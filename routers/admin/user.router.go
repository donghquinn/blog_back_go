package routers

import (
	"net/http"

	users "github.com/donghquinn/blog_back_go/controllers/admin/users"
)

func AdminUserRouter(server *http.ServeMux) {
	server.HandleFunc("POST /admin/user/profile/update", users.UpdateProfileController)
	server.HandleFunc("POST /admin/user/profile/title", users.UpdateTitleController)
	server.HandleFunc("POST /admin/user/profile/color", users.UpdateColorController)
}
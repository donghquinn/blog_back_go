package routers

import (
	"net/http"

	admincontrollers "github.com/donghquinn/blog_back_go/controllers/admin/users"
)

func AdminUserRouter(server *http.ServeMux) {
	server.HandleFunc("POST /admin/user/profile/update", admincontrollers.UpdateProfileController)
	server.HandleFunc("POST /admin/user/profile/title", admincontrollers.UpdateTitleController)
	server.HandleFunc("POST /admin/user/profile/color", admincontrollers.UpdateColorController)
}
package routers

import (
	"net/http"

	"github.com/donghquinn/blog_back_go/controllers/admin"
)

func AdminPostRouter(server *http.ServeMux) {
	server.HandleFunc("POST /admin/post/register", admin.RegisterPostController)
	server.HandleFunc("POST /admin/post/delete", admin.DeletePostController)

	server.HandleFunc("POST /admin/post/update/pin", admin.UpdatePinPostController)
	server.HandleFunc("POST /admin/post/update/unpin", admin.UpdateUnPinPostController)

}
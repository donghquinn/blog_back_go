package routers

import (
	"net/http"

	posts "github.com/donghquinn/blog_back_go/controllers/admin/posts"
)

func AdminPostRouter(server *http.ServeMux) {
	server.HandleFunc("POST /admin/post/register", posts.RegisterPostController)
	server.HandleFunc("POST /admin/post/edit", posts.EditPostController)
	
	server.HandleFunc("POST /admin/post/delete", posts.DeletePostController)

	server.HandleFunc("POST /admin/post/update/pin", posts.UpdatePinPostController)
	server.HandleFunc("POST /admin/post/update/unpin", posts.UpdateUnPinPostController)

}
package routers

import (
	"net/http"

	admincontrollers "github.com/donghquinn/blog_back_go/controllers/admin/posts"
)

func AdminPostRouter(server *http.ServeMux) {
	server.HandleFunc("POST /admin/post/register", admincontrollers.RegisterPostController)
	server.HandleFunc("POST /admin/post/edit", admincontrollers.EditPostController)
	
	server.HandleFunc("POST /admin/post/delete", admincontrollers.DeletePostController)

	server.HandleFunc("POST /admin/post/update/pin", admincontrollers.UpdatePinPostController)
	server.HandleFunc("POST /admin/post/update/unpin", admincontrollers.UpdateUnPinPostController)
 	server.HandleFunc("POST /admin/post/update/secret", admincontrollers.ChangeToSecretPostController)
	server.HandleFunc("POST /admin/post/update/unsecret", admincontrollers.ChangeToNotSecretPostController)
}
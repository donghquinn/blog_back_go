package routers

import (
	"net/http"

	admincontrollers "github.com/donghquinn/blog_back_go/controllers/admin/upload"
	controllers "github.com/donghquinn/blog_back_go/controllers/posts"
)

func UploadImageController(server *http.ServeMux) {
	server.HandleFunc("POST /admin/upload/image/profile", admincontrollers.UploadProfileImageController)
	server.HandleFunc("POST /admin/upload/image/background", admincontrollers.UploadBackgroundImageController)
	
	server.HandleFunc("POST /admin/upload/image/post", admincontrollers.UploadPostImageController)
	server.HandleFunc("POST /upload/image/post", controllers.UploadPostImageController )
}
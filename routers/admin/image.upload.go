package routers

import (
	"net/http"

	upload "github.com/donghquinn/blog_back_go/controllers/admin/upload"
)

func UploadImageController(server *http.ServeMux) {
	server.HandleFunc("POST /admin/upload/image/profile", upload.UploadProfileImageController)
	server.HandleFunc("POST /admin/upload/image/background", upload.UploadBackgroundImageController)
	
	server.HandleFunc("POST /admin/upload/image/post", upload.UploadPostImageController)
}
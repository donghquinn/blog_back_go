package routers

import (
	"net/http"

	"github.com/donghquinn/blog_back_go/controllers/upload"
)

func UploadImageController(server *http.ServeMux) {
	server.HandleFunc("POST /admin/upload/image/profile", upload.UploadProfileImageController)
	server.HandleFunc("POST /admin/upload/image/background", upload.UploadBackgroundImageController)
	
	server.HandleFunc("POST /admin/upload/image/post", upload.UploadPostImageController)
}
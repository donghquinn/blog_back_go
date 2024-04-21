package routers

import (
	"net/http"

	"github.com/donghquinn/blog_back_go/controllers/upload"
)

func UploadImageController(server *http.ServeMux) {
	server.HandleFunc("POST /upload/image/profile", upload.UploadProfileImageController)
}
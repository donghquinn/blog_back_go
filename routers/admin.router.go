package routers

import (
	"net/http"

	routers "github.com/donghquinn/blog_back_go/routers/admin"
)

func AdminRouter(server *http.ServeMux ) {
	routers.AdminPostRouter(server)
	routers.UploadImageController(server)
	routers.AdminUserRouter(server)
}
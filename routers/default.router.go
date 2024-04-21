package routers

import (
	"net/http"

	"github.com/donghquinn/blog_back_go/controllers"
	postRouters "github.com/donghquinn/blog_back_go/routers/posts"
	uploadRouters "github.com/donghquinn/blog_back_go/routers/upload"
	userRouters "github.com/donghquinn/blog_back_go/routers/users"
)

func DefaultRouter(server *http.ServeMux) {
	server.HandleFunc("GET /api", controllers.DefaultController)

	userRouters.UserRouter(server)
	uploadRouters.UploadImageController(server)
	postRouters.PostRouter(server)
}
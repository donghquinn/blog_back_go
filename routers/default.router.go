package routers

import (
	"net/http"

	"github.com/donghquinn/blog_back_go/controllers"
	routers "github.com/donghquinn/blog_back_go/routers/users"
)

func DefaultRouter(server *http.ServeMux) {
	server.HandleFunc("GET /", controllers.DefaultController)

	routers.UserRouter(server)
}
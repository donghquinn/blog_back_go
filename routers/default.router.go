package routers

import (
	"net/http"

	"github.com/donghquinn/blog_back_go/controllers"
	postRouters "github.com/donghquinn/blog_back_go/routers/posts"
	userRouters "github.com/donghquinn/blog_back_go/routers/users"
)

func DefaultRouter(server *http.ServeMux) {
	server.HandleFunc("GET /api", controllers.DefaultController)

	server.HandleFunc("GET /", controllers.CorsTestController)

	userRouters.UserRouter(server)
	postRouters.PostRouter(server)
}
package routers

import (
	"net/http"

	"github.com/donghquinn/blog_back_go/controllers/posts"
)

func PostRouter(server *http.ServeMux) {
	server.HandleFunc("POST /post/register", posts.RegisterPostController)
	server.HandleFunc("POST /post/list", posts.GetPostController)
}
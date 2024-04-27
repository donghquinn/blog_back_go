package routers

import (
	"net/http"

	"github.com/donghquinn/blog_back_go/controllers/posts"
)

func PostRouter(server *http.ServeMux) {
	server.HandleFunc("POST /post/contents", posts.PostContentsController)
	server.HandleFunc("POST /post/list", posts.GetPostController)

	server.HandleFunc("POST /post/list/tag", posts.GetPostsByTagController)
}
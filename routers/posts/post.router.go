package routers

import (
	"net/http"

	controllers "github.com/donghquinn/blog_back_go/controllers/posts"
)

func PostRouter(server *http.ServeMux) {
	server.HandleFunc("POST /post/contents", controllers.PostContentsController)
	server.HandleFunc("POST /post/list", controllers.GetPostController)
	server.HandleFunc("POST /post/list/unpinned", controllers.GetPinnedPostController)

	server.HandleFunc("POST /post/list/tag", controllers.GetPostsByTagController)
	server.HandleFunc("POST /post/list/category", controllers.GetPostsByCategoryController)

	server.HandleFunc("POST /post/url", controllers.GetImageUrl)
}
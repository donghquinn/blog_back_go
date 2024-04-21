package routers

import (
	"net/http"

	routers "github.com/donghquinn/blog_back_go/routers/users"
)

func DefaultRouter(server *http.ServeMux) {
	routers.UserRouter(server)
}
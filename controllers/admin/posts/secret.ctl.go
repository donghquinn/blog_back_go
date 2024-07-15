package admincontrollers

import (
	"net/http"

	"github.com/donghquinn/blog_back_go/auth"
	"github.com/donghquinn/blog_back_go/dto"
)

func ChangeToSecretPostController(response http.ResponseWriter, request *http.Request) {
	userId, _, _, blogId, err := auth.ValidateJwtToken(request)

	if (err != nil ) {
		dto.SetErrorResponse(response, 401, "01", "JWT Validate Error", err)
		return
	}

}
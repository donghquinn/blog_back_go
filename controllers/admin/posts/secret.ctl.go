package admincontrollers

import (
	"net/http"

	"github.com/donghquinn/blog_back_go/auth"
	"github.com/donghquinn/blog_back_go/dto"
	post "github.com/donghquinn/blog_back_go/libraries/post/admin"
	types "github.com/donghquinn/blog_back_go/types/admin/posts"
	"github.com/donghquinn/blog_back_go/utils"
)

func ChangeToSecretPostController(response http.ResponseWriter, request *http.Request) {
	_, _, _, _, err := auth.ValidateJwtToken(request)

	if (err != nil ) {
		dto.SetErrorResponse(response, 401, "01", "JWT Validate Error", err)
		return
	}

	var changeRequest types.ChangeToRequest

	parseErr := utils.DecodeBody(request, &changeRequest)

	if (parseErr != nil ) {
		dto.SetErrorResponse(response, 402, "02", "Invalid Request Body", parseErr)
		return
	}
	
	changeErr := post.ChangeToSecretPost(changeRequest.PostSeq)

	if (changeErr != nil ) {
		dto.SetErrorResponse(response, 403, "03", "Change Secret Failed", changeErr)
		return
	}

	dto.SetResponse(response, 200, "01")
}
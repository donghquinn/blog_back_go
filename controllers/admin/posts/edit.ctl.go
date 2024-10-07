package admincontrollers

import (
	"net/http"

	"github.com/donghquinn/blog_back_go/auth"
	"github.com/donghquinn/blog_back_go/dto"
	post "github.com/donghquinn/blog_back_go/libraries/post/admin"
	types "github.com/donghquinn/blog_back_go/types/admin/posts"
	"github.com/donghquinn/blog_back_go/utils"
)

func EditPostController(res http.ResponseWriter, req *http.Request) {
	userId, _, _, blogId, err := auth.ValidateJwtToken(req)

	if err != nil {
		dto.Response(res, dto.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "EDP001",
			Message: "JWT Validate Error",
		})

		return
	}

	var editPostRequest types.EditPostRequest

	parseErr := utils.DecodeBody(req, &editPostRequest)

	if parseErr != nil {
		dto.Response(res, dto.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "EDP002",
			Message: "Parse Error",
		})
		return
	}

	editErr := post.EditPost(editPostRequest, userId, blogId)

	if editErr != nil {
		dto.Response(res, dto.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "EDP003",
			Message: "Edit Post Error",
		})
		return
	}

	dto.Response(res, dto.CommonResponseWithMessage{
		Status:  http.StatusOK,
		Code:    "0000",
		Message: "Success",
	})
}

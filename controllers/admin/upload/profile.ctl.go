package admincontrollers

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/donghquinn/blog_back_go/auth"
	"github.com/donghquinn/blog_back_go/dto"
	"github.com/donghquinn/blog_back_go/libraries/database"
	upload "github.com/donghquinn/blog_back_go/libraries/upload/image"
	queries "github.com/donghquinn/blog_back_go/queries/upload"
)

// 프로필 이미지 업로드
func UploadProfileImageController(res http.ResponseWriter, req *http.Request) {
	userId, _, _, _, err := auth.ValidateJwtToken(req)

	if err != nil {
		dto.Response(res, dto.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "UPI001",
			Message: "JWT Token Validation Error",
		})

		return
	}

	// 요청으로부터 이미지 파일 가져오기
	file, handler, fileErr := upload.GetImagefileFromRequest(res, req)

	if fileErr != nil {
		dto.Response(res, dto.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "UPI002",
			Message: "Upload Image Error",
		})
		return
	}

	// 파일 생성
	tempFile, tempErr := upload.CreateFileImage(res, req, file, handler)

	if tempErr != nil {
		dto.Response(res, dto.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "UPI003",
			Message: "Upload Image Error",
		})
		return
	}

	contentType := handler.Header["Content-Type"][0]

	// 이미지 업로드 - minio
	_, uploadErr := database.UploadImage(handler.Filename, tempFile.Name(), contentType)

	if uploadErr != nil {
		dto.Response(res, dto.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "UPI004",
			Message: "Upload Image Error",
		})
		return
	}

	connect, _ := database.InitDatabaseConnection()

	// 데이터 입력 - DB
	_, insertErr := connect.InsertQuery(
		queries.InsertProfileImageData,
		// USER ID from JWT
		"1",
		userId,
		"user_table",
		"USER_PROFILE",
		strconv.Itoa(int(handler.Size)),
		handler.Filename,
		contentType)

	if insertErr != nil {
		dto.Response(res, dto.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "UPI005",
			Message: "Insert Image Error",
		})
		return
	}

	defer connect.Close()

	removeErr := os.Remove(tempFile.Name())

	if removeErr != nil {
		log.Printf("[UPLOAD] Remove Saved Image Error: %v", removeErr)

		dto.Response(res, dto.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "UPI006",
			Message: "Upload Image Error",
		})
		return
	}

	dto.Response(res, dto.CommonResponseWithMessage{
		Status:  http.StatusOK,
		Code:    "0000",
		Message: "Success",
	})
}

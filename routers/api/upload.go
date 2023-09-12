package api

//import (
//	"gin-demo/pkg/e"
//	"gin-demo/pkg/logging"
//	"gin-demo/pkg/upload"
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//func UploadImage (context *gin.Context) {
//	code := e.SUCCESS
//	data := make(map[string]string)
//
//	file, image, err := context.Request.FormFile("image")
//	//file , err:= context.FormFile("image")
//	if err != nil {
//		logging.Warn(err)
//		code = e.ERROR
//		context.JSON(http.StatusOK, gin.H{
//			"code" : code,
//			"msg" : e.GetMsg(code),
//			"data" : data,
//		})
//	}
//	if image == nil {
//		code = e.INVALID_PARAMS
//	} else {
//		imageName := upload.GetImageName(image.Filename)
//		fullPath := upload.GetImageFullPath()
//		savePath := upload.GetImagePath()
//
//		src := fullPath + imageName
//
//		if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
//			code = e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
//		} else {
//			err := upload.CheckImage(fullPath)
//			if err != nil {
//				logging.Warn(err)
//				code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL
//			} else if err := context.SaveUploadedFile(image, src); err != nil {
//				logging.Warn(err)
//				code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL
//			} else {
//				data["image_url"] = upload.GetImageFullUrl(imageName)
//				data["image_save_url"] = savePath + imageName
//			}
//		}
//	}
//	context.JSON(http.StatusOK, gin.H{
//		"code" : code,
//		"msg" : e.GetMsg(code),
//		"data" : data,
//	})
//}

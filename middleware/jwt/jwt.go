package jwt

import (
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		//var code int
		//var data interface{}
		//
		//code = e.SUCCESS
		//token := context.Query("token")
		//
		//if token == "" {
		//	code = e.INVALID_PARAMS
		//} else {
		//	claims, err := util.ParseToken(token)
		//	if err != nil {
		//		code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		//	} else if time.Now().Unix() > claims.ExpiresAt {
		//		code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
		//	}
		//}
		//
		//if code != e.SUCCESS {
		//	context.JSON(http.StatusUnauthorized, gin.H{
		//		"code" : code,
		//		"msg" : e.GetMsg(code),
		//		"data" : data,
		//	})
		//	context.Abort()
		//	return
		//}
		context.Next()
	}
}

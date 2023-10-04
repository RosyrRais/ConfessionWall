package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsonResponse(c *gin.Context, httpStatusCode int, code int, msg string, data interface{}) {
	c.JSON(httpStatusCode, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func JsonResponseSuccess(c *gin.Context, data interface{}) {
	JsonResponse(c, http.StatusOK, 200, "OK", data)
}

func JsonResponseError(c *gin.Context, code int, msg string) {
	JsonResponse(c, http.StatusInternalServerError, code, msg, nil)
}

func JsonResponseOtherError(c *gin.Context) {
	JsonResponse(c, 500, 777, "其他错误", nil)
}

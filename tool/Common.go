package tool

import "github.com/gin-gonic/gin"

const (
	SUCCESS = 1
	FAILED  = 0
)

func Success(context *gin.Context, v interface{}) {
	context.JSON(200, gin.H{
		"code":    SUCCESS,
		"message": "成功",
		"data":    v,
	})
}

func Failed(context *gin.Context, v interface{}) {
	context.JSON(400, gin.H{
		"code":    FAILED,
		"message": "解析失败",
		"data":    v,
	})
}

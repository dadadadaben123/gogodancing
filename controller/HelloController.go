package controller

import "github.com/gin-gonic/gin"

// 绑定方法
type HelloController struct {
}

func (hello *HelloController) Router(engine *gin.Engine) {
	//注册路由
	engine.GET("/hello", hello.Hello)

}

func (hello *HelloController) Hello(context *gin.Context) {
	//处理输入和输出
	context.JSON(200, gin.H{
		"code":    1,
		"message": "hello gin",
	})
}

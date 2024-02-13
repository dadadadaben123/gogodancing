package main

import (
	"github.com/gin-gonic/gin"
	"github.com/songzhonghuasongzhonghua/gogodancing/controller"
	"github.com/songzhonghuasongzhonghua/gogodancing/tool"
	"net/http"
)

func main() {
	config, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}
	_, err = tool.OrmEngine(&config.Database)
	if err != nil {
		panic(err.Error())
	}

	engine := gin.Default()
	//初始化Redis
	tool.InitRedisStore(config)
	//初始化sessions
	tool.InitSession(engine)
	//跨域中间件
	engine.Use(Cors())
	registerHelloController(engine)
	registerLoginController(engine)
	registerMemberController(engine)
	engine.Run(":" + config.AppPort)

}

func registerHelloController(engine *gin.Engine) {
	new(controller.HelloController).Router(engine)
}

// 注册登录路由
func registerLoginController(engine *gin.Engine) {
	new(controller.LoginController).Router(engine)
}

// 注册member路由
func registerMemberController(engine *gin.Engine) {
	new(controller.MemberController).Router(engine)
}

// 全局跨域中间件
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		context.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		method := context.Request.Method

		if method == "OPTIONS" {
			context.JSON(http.StatusOK, "options allow")
		}

		context.Next()
	}
}

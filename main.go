package main

import (
	"github.com/gin-gonic/gin"
	"github.com/songzhonghuasongzhonghua/gogodancing/controller"
	"github.com/songzhonghuasongzhonghua/gogodancing/tool"
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
	registerHelloController(engine)
	registerLoginController(engine)
	engine.Run(":" + config.AppPort)

}

func registerHelloController(engine *gin.Engine) {
	new(controller.HelloController).Router(engine)
}

// 注册登录路由
func registerLoginController(engine *gin.Engine) {
	new(controller.LoginController).Router(engine)
}

package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/songzhonghuasongzhonghua/gogodancing/controller"
	"github.com/songzhonghuasongzhonghua/gogodancing/tool"
)

type Person struct {
	Id   int
	Name string
	Age  int
}

func main() {
	config, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err)
	}

	app := gin.Default()
	registerHelloRouter(app)
	app.Run(config.AppHost + ":" + config.AppPort)
}

func registerHelloRouter(engin *gin.Engine) {
	new(controller.HelloController).Router(engin)
}

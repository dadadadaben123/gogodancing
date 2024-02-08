package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/songzhonghuasongzhonghua/gogodancing/param"
	"github.com/songzhonghuasongzhonghua/gogodancing/service"
	"github.com/songzhonghuasongzhonghua/gogodancing/tool"
)

type LoginController struct {
}

func (lc *LoginController) Router(engine *gin.Engine) {
	engine.POST("/login", lc.login)
}

func (lc *LoginController) login(context *gin.Context) {
	loginParam := new(param.LoginParam)
	err := context.ShouldBindJSON(loginParam)
	if err != nil {
		tool.Failed(context, err)
		return
	}

	member, err := service.LoginService(loginParam)
	if err != nil {
		tool.Failed(context, err)
		return
	}

	tool.Success(context, member)

}

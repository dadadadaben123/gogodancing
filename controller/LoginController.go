package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/songzhonghuasongzhonghua/gogodancing/param"
	"github.com/songzhonghuasongzhonghua/gogodancing/service"
	"github.com/songzhonghuasongzhonghua/gogodancing/tool"
	"strconv"
)

type LoginController struct {
}

func (lc *LoginController) Router(engine *gin.Engine) {
	engine.POST("/login", lc.login)

	engine.GET("/captcha", lc.captcha)

	engine.POST("/verify_captcha", lc.verifyCaptcha)

	engine.POST("/login_pwd", lc.loginPwd)
}

func (lc *LoginController) verifyCaptcha(context *gin.Context) {
	params := new(tool.CaptchaResult)
	err := context.Bind(&params)
	fmt.Println(params, "params==================")
	if err != nil {
		tool.Failed(context, "解析参数失败")
		return
	}

	verify := tool.VerifyCaptcha(params.CaptchaId, params.VerifyCode)
	if verify {
		tool.Success(context, "参数验证正确")
	} else {
		tool.Failed(context, "参数验证错误")
	}

}

func (lc *LoginController) captcha(context *gin.Context) {
	tool.Captcha(context)
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

// 用户+密码+验证码登录
func (lc *LoginController) loginPwd(context *gin.Context) {
	//解析参数
	var loginParams param.LoginPwdParams

	err := context.Bind(&loginParams)
	if err != nil {
		tool.Failed(context, "解析参数失败")
		return
	}
	fmt.Println("loginParams", loginParams)
	//验证验证码
	validate := tool.VerifyCaptcha(loginParams.CaptchaId, loginParams.CaptchaValue)
	if !validate {
		tool.Failed(context, "验证码验证错误")
		return
	}

	fmt.Println("validate", validate)

	//登录
	member := service.LoginPwdService(loginParams.Name, loginParams.Password)
	fmt.Println("member", member)
	if member.Id != 0 {
		//登录成功
		memberJson, _ := json.Marshal(member)
		//保存状态到session
		tool.SetSess(context, "user_"+strconv.FormatInt(member.Id, 10), memberJson)
		tool.Success(context, &member)
		return
	}

	tool.Failed(context, "登录失败")

}

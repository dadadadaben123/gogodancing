package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/songzhonghuasongzhonghua/gogodancing/dao"
	"github.com/songzhonghuasongzhonghua/gogodancing/model"
	"github.com/songzhonghuasongzhonghua/gogodancing/service"
	"github.com/songzhonghuasongzhonghua/gogodancing/tool"
	"log"
	"strconv"
	"time"
)

type MemberController struct {
}

func (mc *MemberController) Router(engine *gin.Engine) {
	engine.POST("/upload/avatar", mc.uploadAvatar)

	engine.GET("/userinfo", mc.userInfo)
}

func (mc *MemberController) userInfo(context *gin.Context) {
	cookie, err := tool.CookieAuth(context)
	if err != nil {
		tool.Failed(context, "用户未登录")

	}

	memberId, err := strconv.Atoi(cookie.Value)
	if err != nil {
		tool.Failed(context, "cookie解析失败")

	}

	member := dao.QueryUserInfo(int64(memberId))
	if member == nil {
		tool.Failed(context, "未查询到用户信息")
		return
	}

	tool.Success(context, member)

}

func (mc *MemberController) uploadAvatar(context *gin.Context) {
	//解析参数
	userId := context.PostForm("user_id")
	file, err := context.FormFile("avatar")
	if err != nil {
		log.Fatal(err.Error())
		tool.Failed(context, "文件获取失败")
		return
	}

	//验证是否登录
	session := tool.GetSess(context, "user_"+userId)
	fmt.Println(session)
	if session == nil {
		tool.Failed(context, "当前无权限上传")
		return
	}
	var member model.Member
	json.Unmarshal(session.([]byte), &member)

	//保存文件到本地
	fileName := "./uploadFile/" + strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	err = context.SaveUploadedFile(file, fileName)
	if err != nil {
		tool.Failed(context, "文件保存到本地失败")
		return
	}
	//保存至数据库
	memberService := service.MemberService{}
	path := memberService.UploadAvatarService(member.Id, fileName[1:])
	if path != "" {
		tool.Success(context, "http://localhost:8090"+path)
		return
	}
	tool.Failed(context, "文件保存至数据库失败")
}

package dao

import (
	"fmt"
	"github.com/songzhonghuasongzhonghua/gogodancing/model"
	"github.com/songzhonghuasongzhonghua/gogodancing/tool"
)

type MemberDao struct {
	*tool.Orm
}

func UpdateAvatar(userId int64, path string) int64 {
	md := MemberDao{tool.DbOrm}
	fmt.Println("user_id", userId)
	result, _ := md.Where("id = ?", userId).Update(&model.Member{Avatar: path})
	fmt.Println("result------", result)
	return result
}

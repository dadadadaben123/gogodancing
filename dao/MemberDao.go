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

func QueryUserInfo(userId int64) *model.Member {
	member := new(model.Member)
	md := MemberDao{tool.DbOrm}
	_, err := md.Where("id = ?", userId).Get(member)
	if err != nil {
		return nil
	}
	return member
}

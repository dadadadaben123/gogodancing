package service

import (
	"github.com/songzhonghuasongzhonghua/gogodancing/dao"
)

type MemberService struct {
}

func (ms *MemberService) UploadAvatarService(userId int64, fileName string) string {

	result := dao.UpdateAvatar(userId, fileName)

	if result == 0 {
		return ""
	}

	return fileName

}

func (ms *MemberService) UserInfoService(userId int64) {

}

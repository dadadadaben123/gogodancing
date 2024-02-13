package dao

import (
	"github.com/songzhonghuasongzhonghua/gogodancing/model"
	"github.com/songzhonghuasongzhonghua/gogodancing/tool"
)

type LoginEngine struct {
	*tool.Orm
}

func LoginDao(insertParam *model.Member) (int64, error) {

	db := LoginEngine{tool.DbOrm}

	result, err := db.Insert(insertParam)
	if err != nil {
		return 0, err
	}

	return result, nil

}

func QueryMember(name string, password string) *model.Member {
	db := LoginEngine{tool.DbOrm}

	member := new(model.Member)

	_, err := db.Where("user_name = ? and password = ?", name, password).Get(member)
	if err != nil {
		member.UserName = name
		member.Password = password

	}

	return member

}

func InsertMember(member *model.Member) error {
	db := LoginEngine{tool.DbOrm}
	_, err := db.Insert(member)
	if err != nil {
		return err
	}
	return nil
}

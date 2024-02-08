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

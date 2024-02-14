package dao

import (
	"github.com/songzhonghuasongzhonghua/gogodancing/model"
	"github.com/songzhonghuasongzhonghua/gogodancing/tool"
)

type FoodCategoryDao struct {
	*tool.Orm
}

func (fcd *FoodCategoryDao) QueryCategoryDao() ([]model.FoodCategory, error) {
	var categories []model.FoodCategory
	err := fcd.Engine.Find(&categories)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

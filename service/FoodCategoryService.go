package service

import (
	"github.com/songzhonghuasongzhonghua/gogodancing/dao"
	"github.com/songzhonghuasongzhonghua/gogodancing/model"
	"github.com/songzhonghuasongzhonghua/gogodancing/tool"
)

type FoodCategoryService struct {
}

func (fcs *FoodCategoryService) QueryFoodCategory() ([]model.FoodCategory, error) {
	foodCategoryDao := dao.FoodCategoryDao{Orm: tool.DbOrm}
	categories, err := foodCategoryDao.QueryCategoryDao()
	return categories, err
}

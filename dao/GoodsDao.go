package dao

import (
	"github.com/songzhonghuasongzhonghua/gogodancing/model"
	"github.com/songzhonghuasongzhonghua/gogodancing/tool"
)

type GoodsDao struct {
	*tool.Orm
}

func NewGoodsDao() *GoodsDao {
	return &GoodsDao{tool.DbOrm}
}

func (gd *GoodsDao) QueryGoods(shopId int) []model.Goods {
	var goods []model.Goods
	err := gd.Where("shop_id = ?", shopId).Find(&goods)
	if err != nil {
		return nil
	}
	return goods
}

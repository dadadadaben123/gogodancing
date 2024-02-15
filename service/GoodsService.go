package service

import (
	"github.com/songzhonghuasongzhonghua/gogodancing/dao"
	"github.com/songzhonghuasongzhonghua/gogodancing/model"
)

type GoodsService struct {
}

func (gs *GoodsService) QueryGoods(shopId int) []model.Goods {
	goodsDao := dao.NewGoodsDao()
	return goodsDao.QueryGoods(shopId)
}

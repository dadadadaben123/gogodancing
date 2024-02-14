package dao

import (
	"github.com/songzhonghuasongzhonghua/gogodancing/model"
	"github.com/songzhonghuasongzhonghua/gogodancing/tool"
)

type ShopDao struct {
	*tool.Orm
}

const DEFATLT_RANGE = 10

func (sd *ShopDao) GetShopList(long, lat float64, keywords string) []model.Shop {
	var shops []model.Shop
	if keywords == "" {
		err := sd.Engine.Where("longitude > ? and longitude < ? and latitude > ? and latitude < ?", long-DEFATLT_RANGE, long+DEFATLT_RANGE, lat-DEFATLT_RANGE, lat+DEFATLT_RANGE).Find(&shops)
		if err != nil {
			return nil
		}

	} else {
		err := sd.Engine.Where("longitude > ? and longitude < ? and latitude > ? and latitude < ? and name like ?", long-DEFATLT_RANGE, long+DEFATLT_RANGE, lat-DEFATLT_RANGE, lat+DEFATLT_RANGE, "%"+keywords+"%").Find(&shops)
		if err != nil {
			return nil
		}
	}

	return shops
}

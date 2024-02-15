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

func (sd *ShopDao) GetServiceByShopId(shopId int64) []model.Service {
	var services []model.Service
	err := sd.Orm.Table("service").Join("INNER", "shop_service", "service.id = shop_service.service_id and shop_service.shop_id = ?", shopId).Find(&services)
	if err != nil {
		return nil
	}
	return services

}

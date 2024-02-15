package service

import (
	"github.com/songzhonghuasongzhonghua/gogodancing/dao"
	"github.com/songzhonghuasongzhonghua/gogodancing/model"
	"github.com/songzhonghuasongzhonghua/gogodancing/tool"
	"strconv"
)

type ShopService struct {
}

func (sc *ShopService) QueryServiceByShopId(shopId int64) []model.Service {
	newShopDao := dao.ShopDao{tool.DbOrm}
	return newShopDao.GetServiceByShopId(shopId)
}

func (sc *ShopService) GetShopList(long, lat string) []model.Shop {
	longitude, err := strconv.ParseFloat(long, 10)
	if err != nil {
		return nil
	}
	latitude, err := strconv.ParseFloat(lat, 10)
	if err != nil {
		return nil
	}

	newShopDao := dao.ShopDao{Orm: tool.DbOrm}
	return newShopDao.GetShopList(longitude, latitude, "")

}

func (sc *ShopService) SearchShops(long, lat, keywords string) []model.Shop {
	longitude, err := strconv.ParseFloat(long, 10)
	if err != nil {
		return nil
	}
	latitude, err := strconv.ParseFloat(lat, 10)
	if err != nil {
		return nil
	}

	newShopDao := dao.ShopDao{Orm: tool.DbOrm}

	return newShopDao.GetShopList(longitude, latitude, keywords)
}

package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/songzhonghuasongzhonghua/gogodancing/service"
	"github.com/songzhonghuasongzhonghua/gogodancing/tool"
)

type ShopController struct {
}

type Location struct {
	Longitude string `form:"longitude"`
	Latitude  string `form:"latitude"`
}

func (sc *ShopController) Router(engine *gin.Engine) {
	engine.GET("/shop_list", sc.getShopList)
	engine.GET("/search_shops", sc.searchShops)
}

func (sc *ShopController) searchShops(context *gin.Context) {
	longitude := context.Query("longitude")
	latitude := context.Query("latitude")
	keywords := context.Query("keywords")
	if keywords == "" {
		tool.Failed(context, "搜索字段为空")
		return
	}
	fmt.Println(longitude, latitude, keywords)

	if longitude == "" || longitude == "undefined" || latitude == "" || latitude == "undefined" {
		longitude = "6.10"
		latitude = "7.10"
	}
	shopService := service.ShopService{}
	shops := shopService.SearchShops(longitude, latitude, keywords)

	if len(shops) == 0 {
		tool.Failed(context, "获取的商户列表为空")
		return
	}
	tool.Success(context, shops)
}

func (sc *ShopController) getShopList(context *gin.Context) {
	location := new(Location)
	err := context.ShouldBindQuery(location)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	fmt.Println(location)
	if location.Longitude == "" || location.Longitude == "undefined" || location.Latitude == "" || location.Latitude == "undefined" {
		location.Longitude = "6.10"
		location.Latitude = "7.10"
	}

	shopService := service.ShopService{}
	shops := shopService.GetShopList(location.Longitude, location.Latitude)
	if len(shops) == 0 {
		tool.Failed(context, "获取商店列表为空")
		return
	}

	tool.Success(context, shops)

}

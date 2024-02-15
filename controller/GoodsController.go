package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/songzhonghuasongzhonghua/gogodancing/service"
	"github.com/songzhonghuasongzhonghua/gogodancing/tool"
	"strconv"
)

type GoodsController struct {
}

func (gc *GoodsController) Router(engine *gin.Engine) {
	engine.GET("/goods", gc.goods)
}

func (gc *GoodsController) goods(context *gin.Context) {
	shopId := context.Query("shop_id")
	shopIdInt, err := strconv.Atoi(string(shopId))
	if err != nil {
		tool.Failed(context, "shop_id解析失败")
		return
	}

	goodsService := service.GoodsService{}
	goods := goodsService.QueryGoods(shopIdInt)
	if len(goods) == 0 {
		tool.Failed(context, "goods为空")
		return
	}

	tool.Success(context, goods)

}

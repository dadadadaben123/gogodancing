package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/songzhonghuasongzhonghua/gogodancing/service"
	"github.com/songzhonghuasongzhonghua/gogodancing/tool"
)

type FoodCategoryController struct {
}

func (fcc *FoodCategoryController) Router(engine *gin.Engine) {
	engine.GET("/food_category", fcc.foodCategory)
}

func (fcc *FoodCategoryController) foodCategory(context *gin.Context) {
	//service
	foodCategoryService := service.FoodCategoryService{}
	categories, err := foodCategoryService.QueryFoodCategory()
	if err != nil {
		tool.Failed(context, err)
		return
	}

	tool.Success(context, categories)

}

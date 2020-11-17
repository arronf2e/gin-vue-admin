package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitGarageRouter(Router *gin.RouterGroup) {
	GarageRouter := Router.Group("garage").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		GarageRouter.POST("createGarage", v1.CreateGarage)   // 新建Garage
		GarageRouter.DELETE("deleteGarage", v1.DeleteGarage) // 删除Garage
		GarageRouter.DELETE("deleteGarageByIds", v1.DeleteGarageByIds) // 批量删除Garage
		GarageRouter.PUT("updateGarage", v1.UpdateGarage)    // 更新Garage
		GarageRouter.GET("findGarage", v1.FindGarage)        // 根据ID获取Garage
		GarageRouter.GET("getGarageList", v1.GetGarageList)  // 获取Garage列表
	}
}

package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitServiceTypeRouter(Router *gin.RouterGroup) {
	ServiceTypeRouter := Router.Group("serviceType").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		ServiceTypeRouter.POST("createServiceType", v1.CreateServiceType)   // 新建ServiceType
		ServiceTypeRouter.DELETE("deleteServiceType", v1.DeleteServiceType) // 删除ServiceType
		ServiceTypeRouter.DELETE("deleteServiceTypeByIds", v1.DeleteServiceTypeByIds) // 批量删除ServiceType
		ServiceTypeRouter.PUT("updateServiceType", v1.UpdateServiceType)    // 更新ServiceType
		ServiceTypeRouter.GET("findServiceType", v1.FindServiceType)        // 根据ID获取ServiceType
		ServiceTypeRouter.GET("getServiceTypeList", v1.GetServiceTypeList)  // 获取ServiceType列表
	}
}

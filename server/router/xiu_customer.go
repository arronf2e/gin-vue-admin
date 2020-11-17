package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitWxCustomerRouter(Router *gin.RouterGroup) {
	CustomerRouter := Router.Group("customer").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		CustomerRouter.POST("createCustomer", v1.CreateCustomer)   // 新建Customer
		CustomerRouter.DELETE("deleteCustomer", v1.DeleteCustomer) // 删除Customer
		CustomerRouter.DELETE("deleteCustomerByIds", v1.DeleteCustomerByIds) // 批量删除Customer
		CustomerRouter.PUT("updateCustomer", v1.UpdateCustomer)    // 更新Customer
		CustomerRouter.GET("findCustomer", v1.FindCustomer)        // 根据ID获取Customer
		CustomerRouter.GET("getCustomerList", v1.GetCustomerList)  // 获取Customer列表
	}
}

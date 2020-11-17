package v1

import (
	"gin-vue-admin/global"
    "gin-vue-admin/model"
    "gin-vue-admin/model/request"
    "gin-vue-admin/model/response"
    "gin-vue-admin/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @Tags ServiceType
// @Summary 创建ServiceType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ServiceType true "创建ServiceType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /serviceType/createServiceType [post]
func CreateServiceType(c *gin.Context) {
	var serviceType model.ServiceType
	_ = c.ShouldBindJSON(&serviceType)
	if err := service.CreateServiceType(serviceType); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ServiceType
// @Summary 删除ServiceType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ServiceType true "删除ServiceType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /serviceType/deleteServiceType [delete]
func DeleteServiceType(c *gin.Context) {
	var serviceType model.ServiceType
	_ = c.ShouldBindJSON(&serviceType)
	if err := service.DeleteServiceType(serviceType); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ServiceType
// @Summary 批量删除ServiceType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ServiceType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /serviceType/deleteServiceTypeByIds [delete]
func DeleteServiceTypeByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteServiceTypeByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags ServiceType
// @Summary 更新ServiceType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ServiceType true "更新ServiceType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /serviceType/updateServiceType [put]
func UpdateServiceType(c *gin.Context) {
	var serviceType model.ServiceType
	_ = c.ShouldBindJSON(&serviceType)
	if err := service.UpdateServiceType(&serviceType); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags ServiceType
// @Summary 用id查询ServiceType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ServiceType true "用id查询ServiceType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /serviceType/findServiceType [get]
func FindServiceType(c *gin.Context) {
	var serviceType model.ServiceType
	_ = c.ShouldBindQuery(&serviceType)
	if err, reserviceType := service.GetServiceType(serviceType.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reserviceType": reserviceType}, c)
	}
}

// @Tags ServiceType
// @Summary 分页获取ServiceType列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ServiceTypeSearch true "分页获取ServiceType列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /serviceType/getServiceTypeList [get]
func GetServiceTypeList(c *gin.Context) {
	var pageInfo request.ServiceTypeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetServiceTypeInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}

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

// @Tags Garage
// @Summary 创建Garage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Garage true "创建Garage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /garage/createGarage [post]
func CreateGarage(c *gin.Context) {
	var garage model.Garage
	_ = c.ShouldBindJSON(&garage)
	if err := service.CreateGarage(garage); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Garage
// @Summary 删除Garage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Garage true "删除Garage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /garage/deleteGarage [delete]
func DeleteGarage(c *gin.Context) {
	var garage model.Garage
	_ = c.ShouldBindJSON(&garage)
	if err := service.DeleteGarage(garage); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Garage
// @Summary 批量删除Garage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Garage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /garage/deleteGarageByIds [delete]
func DeleteGarageByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteGarageByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Garage
// @Summary 更新Garage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Garage true "更新Garage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /garage/updateGarage [put]
func UpdateGarage(c *gin.Context) {
	var garage model.Garage
	_ = c.ShouldBindJSON(&garage)
	if err := service.UpdateGarage(&garage); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Garage
// @Summary 用id查询Garage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Garage true "用id查询Garage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /garage/findGarage [get]
func FindGarage(c *gin.Context) {
	var garage model.Garage
	_ = c.ShouldBindQuery(&garage)
	if err, regarage := service.GetGarage(garage.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"regarage": regarage}, c)
	}
}

// @Tags Garage
// @Summary 分页获取Garage列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GarageSearch true "分页获取Garage列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /garage/getGarageList [get]
func GetGarageList(c *gin.Context) {
	var pageInfo request.GarageSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetGarageInfoList(pageInfo); err != nil {
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

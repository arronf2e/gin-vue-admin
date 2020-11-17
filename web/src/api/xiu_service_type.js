import service from '@/utils/request'

// @Tags ServiceType
// @Summary 创建ServiceType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ServiceType true "创建ServiceType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /serviceType/createServiceType [post]
export const createServiceType = (data) => {
     return service({
         url: "/serviceType/createServiceType",
         method: 'post',
         data
     })
 }


// @Tags ServiceType
// @Summary 删除ServiceType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ServiceType true "删除ServiceType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /serviceType/deleteServiceType [delete]
 export const deleteServiceType = (data) => {
     return service({
         url: "/serviceType/deleteServiceType",
         method: 'delete',
         data
     })
 }

// @Tags ServiceType
// @Summary 删除ServiceType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ServiceType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /serviceType/deleteServiceType [delete]
 export const deleteServiceTypeByIds = (data) => {
     return service({
         url: "/serviceType/deleteServiceTypeByIds",
         method: 'delete',
         data
     })
 }

// @Tags ServiceType
// @Summary 更新ServiceType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ServiceType true "更新ServiceType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /serviceType/updateServiceType [put]
 export const updateServiceType = (data) => {
     return service({
         url: "/serviceType/updateServiceType",
         method: 'put',
         data
     })
 }


// @Tags ServiceType
// @Summary 用id查询ServiceType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ServiceType true "用id查询ServiceType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /serviceType/findServiceType [get]
 export const findServiceType = (params) => {
     return service({
         url: "/serviceType/findServiceType",
         method: 'get',
         params
     })
 }


// @Tags ServiceType
// @Summary 分页获取ServiceType列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取ServiceType列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /serviceType/getServiceTypeList [get]
 export const getServiceTypeList = (params) => {
     return service({
         url: "/serviceType/getServiceTypeList",
         method: 'get',
         params
     })
 }
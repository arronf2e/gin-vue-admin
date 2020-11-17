import service from '@/utils/request'

// @Tags Garage
// @Summary 创建Garage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Garage true "创建Garage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /garage/createGarage [post]
export const createGarage = (data) => {
     return service({
         url: "/garage/createGarage",
         method: 'post',
         data
     })
 }


// @Tags Garage
// @Summary 删除Garage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Garage true "删除Garage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /garage/deleteGarage [delete]
 export const deleteGarage = (data) => {
     return service({
         url: "/garage/deleteGarage",
         method: 'delete',
         data
     })
 }

// @Tags Garage
// @Summary 删除Garage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Garage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /garage/deleteGarage [delete]
 export const deleteGarageByIds = (data) => {
     return service({
         url: "/garage/deleteGarageByIds",
         method: 'delete',
         data
     })
 }

// @Tags Garage
// @Summary 更新Garage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Garage true "更新Garage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /garage/updateGarage [put]
 export const updateGarage = (data) => {
     return service({
         url: "/garage/updateGarage",
         method: 'put',
         data
     })
 }


// @Tags Garage
// @Summary 用id查询Garage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Garage true "用id查询Garage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /garage/findGarage [get]
 export const findGarage = (params) => {
     return service({
         url: "/garage/findGarage",
         method: 'get',
         params
     })
 }


// @Tags Garage
// @Summary 分页获取Garage列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Garage列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /garage/getGarageList [get]
 export const getGarageList = (params) => {
     return service({
         url: "/garage/getGarageList",
         method: 'get',
         params
     })
 }
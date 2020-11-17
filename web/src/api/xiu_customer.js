import service from '@/utils/request'

// @Tags Customer
// @Summary 创建Customer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Customer true "创建Customer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/createCustomer [post]
export const createCustomer = (data) => {
     return service({
         url: "/customer/createCustomer",
         method: 'post',
         data
     })
 }


// @Tags Customer
// @Summary 删除Customer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Customer true "删除Customer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /customer/deleteCustomer [delete]
 export const deleteCustomer = (data) => {
     return service({
         url: "/customer/deleteCustomer",
         method: 'delete',
         data
     })
 }

// @Tags Customer
// @Summary 删除Customer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Customer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /customer/deleteCustomer [delete]
 export const deleteCustomerByIds = (data) => {
     return service({
         url: "/customer/deleteCustomerByIds",
         method: 'delete',
         data
     })
 }

// @Tags Customer
// @Summary 更新Customer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Customer true "更新Customer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /customer/updateCustomer [put]
 export const updateCustomer = (data) => {
     return service({
         url: "/customer/updateCustomer",
         method: 'put',
         data
     })
 }


// @Tags Customer
// @Summary 用id查询Customer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Customer true "用id查询Customer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /customer/findCustomer [get]
 export const findCustomer = (params) => {
     return service({
         url: "/customer/findCustomer",
         method: 'get',
         params
     })
 }


// @Tags Customer
// @Summary 分页获取Customer列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Customer列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/getCustomerList [get]
 export const getCustomerList = (params) => {
     return service({
         url: "/customer/getCustomerList",
         method: 'get',
         params
     })
 }
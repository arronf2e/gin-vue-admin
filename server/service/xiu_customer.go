package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateCustomer
//@description: 创建Customer记录
//@param: customer model.Customer
//@return: err error

func CreateCustomer(customer model.Customer) (err error) {
	err = global.GVA_DB.Create(&customer).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteCustomer
//@description: 删除Customer记录
//@param: customer model.Customer
//@return: err error

func DeleteCustomer(customer model.Customer) (err error) {
	err = global.GVA_DB.Delete(customer).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteCustomerByIds
//@description: 批量删除Customer记录
//@param: ids request.IdsReq
//@return: err error

func DeleteCustomerByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Customer{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateCustomer
//@description: 更新Customer记录
//@param: customer *model.Customer
//@return: err error

func UpdateCustomer(customer *model.Customer) (err error) {
	err = global.GVA_DB.Save(customer).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetCustomer
//@description: 根据id获取Customer记录
//@param: id uint
//@return: err error, customer model.Customer

func GetCustomer(id uint) (err error, customer model.Customer) {
	err = global.GVA_DB.Where("id = ?", id).First(&customer).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetCustomerInfoList
//@description: 分页获取Customer记录
//@param: info request.CustomerSearch
//@return: err error, list interface{}, total int64

func GetWxCustomerInfoList(info request.CustomerSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Customer{})
    var customers []model.Customer
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&customers).Error
	return err, customers, total
}
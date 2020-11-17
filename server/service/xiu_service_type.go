package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateServiceType
//@description: 创建ServiceType记录
//@param: serviceType model.ServiceType
//@return: err error

func CreateServiceType(serviceType model.ServiceType) (err error) {
	err = global.GVA_DB.Create(&serviceType).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteServiceType
//@description: 删除ServiceType记录
//@param: serviceType model.ServiceType
//@return: err error

func DeleteServiceType(serviceType model.ServiceType) (err error) {
	err = global.GVA_DB.Delete(serviceType).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteServiceTypeByIds
//@description: 批量删除ServiceType记录
//@param: ids request.IdsReq
//@return: err error

func DeleteServiceTypeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.ServiceType{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateServiceType
//@description: 更新ServiceType记录
//@param: serviceType *model.ServiceType
//@return: err error

func UpdateServiceType(serviceType *model.ServiceType) (err error) {
	err = global.GVA_DB.Save(serviceType).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetServiceType
//@description: 根据id获取ServiceType记录
//@param: id uint
//@return: err error, serviceType model.ServiceType

func GetServiceType(id uint) (err error, serviceType model.ServiceType) {
	err = global.GVA_DB.Where("id = ?", id).First(&serviceType).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetServiceTypeInfoList
//@description: 分页获取ServiceType记录
//@param: info request.ServiceTypeSearch
//@return: err error, list interface{}, total int64

func GetServiceTypeInfoList(info request.ServiceTypeSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.ServiceType{})
    var serviceTypes []model.ServiceType
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&serviceTypes).Error
	return err, serviceTypes, total
}
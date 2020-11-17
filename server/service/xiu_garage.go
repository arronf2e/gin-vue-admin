package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateGarage
//@description: 创建Garage记录
//@param: garage model.Garage
//@return: err error

func CreateGarage(garage model.Garage) (err error) {
	err = global.GVA_DB.Create(&garage).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteGarage
//@description: 删除Garage记录
//@param: garage model.Garage
//@return: err error

func DeleteGarage(garage model.Garage) (err error) {
	err = global.GVA_DB.Delete(garage).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteGarageByIds
//@description: 批量删除Garage记录
//@param: ids request.IdsReq
//@return: err error

func DeleteGarageByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Garage{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateGarage
//@description: 更新Garage记录
//@param: garage *model.Garage
//@return: err error

func UpdateGarage(garage *model.Garage) (err error) {
	err = global.GVA_DB.Save(garage).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetGarage
//@description: 根据id获取Garage记录
//@param: id uint
//@return: err error, garage model.Garage

func GetGarage(id uint) (err error, garage model.Garage) {
	err = global.GVA_DB.Where("id = ?", id).First(&garage).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetGarageInfoList
//@description: 分页获取Garage记录
//@param: info request.GarageSearch
//@return: err error, list interface{}, total int64

func GetGarageInfoList(info request.GarageSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Garage{})
    var garages []model.Garage
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&garages).Error
	return err, garages, total
}
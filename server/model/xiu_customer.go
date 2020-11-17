// 自动生成模板Customer
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Customer struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:"`
      LicensePlate  string `json:"licensePlate" form:"licensePlate" gorm:"column:licensePlate;comment:"`
      Contact  string `json:"contact" form:"contact" gorm:"column:contact;comment:"`
}


func (Customer) TableName() string {
  return "xiu_customer"
}

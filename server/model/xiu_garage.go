// 自动生成模板Garage
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Garage struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:"`
      Location  string `json:"location" form:"location" gorm:"column:location;comment:"`
      OpeningHours  string `json:"openingHours" form:"openingHours" gorm:"column:openingHours;comment:"`
      BusinessScope  string `json:"businessScope" form:"businessScope" gorm:"column:businessScope;comment:"`
      Contacter  string `json:"contacter" form:"contacter" gorm:"column:contacter;comment:"`
      Contact  string `json:"contact" form:"contact" gorm:"column:contact;comment:"`
      Detail  string `json:"detail" form:"detail" gorm:"column:detail;comment:"`
}


func (Garage) TableName() string {
  return "xiu_garage"
}

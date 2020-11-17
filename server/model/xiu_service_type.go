// 自动生成模板ServiceType
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type ServiceType struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:"`
}


func (ServiceType) TableName() string {
  return "xiu_service_type"
}

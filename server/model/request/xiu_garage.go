package request

import "gin-vue-admin/model"

type GarageSearch struct{
    model.Garage
    PageInfo
}
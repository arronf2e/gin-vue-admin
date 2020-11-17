package request

import "gin-vue-admin/model"

type ServiceTypeSearch struct{
    model.ServiceType
    PageInfo
}
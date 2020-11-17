package request

import "gin-vue-admin/model"

type CustomerSearch struct{
    model.Customer
    PageInfo
}
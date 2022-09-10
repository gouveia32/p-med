package models

import (
	"github.com/beego/beego/v2/client/orm"
)

// AdminRole struct
type AdminRole struct {
	Id          int    `orm:"column(id);auto;size(11)" description:"表ID"`
	Name        string `orm:"column(name);size(50)" description:"Nome"`
	Description string `orm:"column(description);size(100)" description:"Descrição"`
	Url         string `orm:"column(url);size(1000)" description:"权限"`
	Status      int8   `orm:"column(status);size(1);default(1)" description:"Ativo 0：Não 1：Sim"`
}

// SearchField 定义模型的可搜索字段
func (*AdminRole) SearchField() []string {
	return []string{"name", "description"}
}

// NoDeletionId 禁止删除的数据id
func (*AdminRole) NoDeletionId() []int {
	return []int{1}
}

// WhereField 定义模型可作为条件的字段
func (*AdminRole) WhereField() []string {
	return []string{}
}

// TimeField 定义可做为时间范围查询的字段
func (*AdminRole) TimeField() []string {
	return []string{}
}

// TableName 自定义table Nome
func (*AdminRole) TableName() string {
	return "admin_role"
}

//在init中注册定义的model
func init() {
	orm.RegisterModel(new(AdminRole))
}

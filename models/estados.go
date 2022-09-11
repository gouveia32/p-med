package models

import (
	"github.com/beego/beego/v2/client/orm"
)

// UserLevel struct
type Estado struct {
	Uf   string `orm:"column(estado);size(2)" description:"UF" json:"uf"`
	Name string `orm:"column(nome);size(20)" description:"Nome" json:"nome"`
}

// TableName
func (*Estado) TableName() string {
	return "estados"
}

// SearchField
func (*Estado) SearchField() []string {
	return []string{"name", "description"}
}

// NoDeletionId
func (*Estado) NoDeletionUf() []string {
	return []string{}
}

// WhereField
func (*Estado) WhereField() []string {
	return []string{}
}

//init
func init() {
	orm.RegisterModel(new(Estado))
}

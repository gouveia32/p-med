/*
*
@author : CCB_f
@
*/
package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Item struct {
	Id int64 `form:"id"        		alias:"ID" `
	/* 	ListaId 	int    `orm:"column(lista_id);size(10);default(0)" description:"id da lista" json:"lista_id"` */
	Descricao string `form:"descricao"   			alias:"Descricao"`
}

type Lista struct {
	Id   int64   `form:"id"        		alias:"ID" `
	Nome string  `form:"nome"      		alias:"Nome"    		valid:"Required;MaxSize(20)"`
	Item []*Item `orm:"-"`
}

// TableName
func (*Lista) TableName() string {
	return "lista"
}

// SearchField
func (*Lista) SearchField() []string {
	return []string{"nome", "valor"}
}

// NoDeletionId
func (*Lista) NoDeletionId() []int {
	return []int{}
}

// WhereField
func (*Lista) WhereField() []string {
	return []string{}
}

// TimeField
func (*Lista) TimeField() []string {
	return []string{}
}

// Registre o campo definido no init
func init() {
	orm.RegisterModel(new(Lista))
}

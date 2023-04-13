/*
*
@author : CCB_f
@
*/
package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type ItemLista struct {
	Id        int64  `form:"id"        		alias:"ID" `
	ListaId   int64  `orm:"column(lista_id);size(10);default(0)" description:"id da lista" json:"lista_id"`
	Nome      string `form:"nome"      		alias:"Nome"    		valid:"Required;MaxSize(20)"`
	Descricao string `form:"descricao"   			alias:"Descricao"`
}

type ItemListaBd struct {
	Id      int64 `form:"id"        		alias:"ID" `
	ListaId int64 `orm:"column(lista_id);size(10);default(0)" description:"id da lista" json:"lista_id"`
	//Nome      string `form:"nome"      		alias:"Nome"    		valid:"Required;MaxSize(20)"`
	Descricao string `form:"descricao"   			alias:"Descricao"`
}

// TableName
func (*ItemListaBd) TableName() string {
	return "item_lista"
}

// SearchField
func (*ItemListaBd) SearchField() []string {
	return []string{"listaId", "descricao"}
}

// NoDeletionId
func (*ItemListaBd) NoDeletionId() []int {
	return []int{}
}

// WhereField
func (*ItemListaBd) WhereField() []string {
	return []string{}
}

// TimeField
func (*ItemListaBd) TimeField() []string {
	return []string{}
}

// Registre o campo definido no init
func init() {
	orm.RegisterModel(new(ItemListaBd))
}

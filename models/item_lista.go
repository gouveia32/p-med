/**
@author : CCB_f
@
*/
package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type ItemLista struct {
	Id          int64  `form:"id"        		alias:"ID" `
	ListaId 	int    `orm:"column(lista_id);size(10);default(0)" description:"id da lista" json:"lista_id"`
	Valor	 	string `form:"valor"   			alias:"Valor"`
}

// TableName
func (*ItemLista) TableName() string {
	return "item_lista"
}

// SearchField
func (*ItemLista) SearchField() []string {
	return []string{"listaId", "valor"}
}

// NoDeletionId
func (*ItemLista) NoDeletionId() []int {
	return []int{}
}

// WhereField
func (*ItemLista) WhereField() []string {
	return []string{}
}

// TimeField
func (*ItemLista) TimeField() []string {
	return []string{}
}

//Registre o campo definido no init
func init() {
	orm.RegisterModel(new(ItemLista))
}

/**
@author : CCB_f
@
*/
package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Cor struct {
	Id     int    `form:"id"        		alias:"ID" `
	Nome   string `form:"nome"      		alias:"Nome"    		valid:"Required;MaxSize(120)"`
	Cor    string `form:"cor"      		alias:"Cor"    		valid:"Required;MaxSize(6)"`
	Estado int8   `form:"estado"    		alias:"Estado"`
}

// TableName
func (*Cor) TableName() string {
	return "cor"
}

// SearchField
func (*Cor) SearchField() []string {
	return []string{"nome"}
}

// NoDeletionId
func (*Cor) NoDeletionId() []int {
	return []int{}
}

// WhereField
func (*Cor) WhereField() []string {
	return []string{}
}

// TimeField
func (*Cor) TimeField() []string {
	return []string{}
}

//Registre o cor definido no init
func init() {
	orm.RegisterModel(new(Cor))
}

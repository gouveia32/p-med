/**
@author : CCB_f
@
*/
package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Modelo struct {
	Id          int64  `form:"id"        		alias:"ID" `
	Nome        string `form:"nome"      		alias:"Nome"    		valid:"Required;MaxSize(120)"`
	Detalhe     string `form:"detalhe"   		alias:"Detalhe"`
	Estado      int8   `form:"estado"    		alias:"Estado"`
	CriadorId   int    `form:"criador_id"`
	AlteradorId int    `form:"alterador_id"`
	CriadoEm    int    `form:"criado_em"`
	AlteradoEm  int    `form:"alterado_em"`
}

// TableName
func (*Modelo) TableName() string {
	return "modelo"
}

// SearchField
func (*Modelo) SearchField() []string {
	return []string{"nome", "detalhe"}
}

// NoDeletionId
func (*Modelo) NoDeletionId() []int {
	return []int{}
}

// WhereField
func (*Modelo) WhereField() []string {
	return []string{}
}

// TimeField
func (*Modelo) TimeField() []string {
	return []string{}
}

//Registre o modelo definido no init
func init() {
	orm.RegisterModel(new(Modelo))
}

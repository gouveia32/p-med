/**
@author : CCB_f
@
*/
package models

import (
	"github.com/beego/beego/v2/client/orm"
)

//Respostas json
type Resposta struct {
	Tipo    string
	Resposta   string
}

type Campo struct {
	Id           int64  `form:"id"        		alias:"ID" `
	Original	 string `form:"original"   		alias:"Original"`
	Tipo		 string `form:"tipo"   			alias:"Tipo"`
	Nome         string `form:"nome"      		alias:"Nome"    		valid:"Required;MaxSize(120)"`
	Descricao    string `form:"descricao"   	alias:"Descricao"`
	Resposta     string `orm:"column(resposta);type(text)" description:"Resposta" json:"resposta"`
	RespostaStrut []*Resposta `orm:"-"`
	ValorInicial string `form:"valor_inicial"   alias:"ValorInicial"`
}

// TableName
func (*Campo) TableName() string {
	return "campo"
}

// SearchField
func (*Campo) SearchField() []string {
	return []string{"nome", "descricao", "resposta"}
}

// NoDeletionId
func (*Campo) NoDeletionId() []int {
	return []int{}
}

// WhereField
func (*Campo) WhereField() []string {
	return []string{}
}

// TimeField
func (*Campo) TimeField() []string {
	return []string{}
}

//Registre o campo definido no init
func init() {
	orm.RegisterModel(new(Campo))
}

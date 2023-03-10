/**
@author : CCB_f
@
*/
package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Nota struct {
	Id            int    `form:"id"        		alias:"ID" `
	Nome          string `form:"nome"      		alias:"Nome"    		valid:"Required;MaxSize(120)"`
	Acompanhante  string  `form:"acompanhante"    		alias:"Acompanhante"  		valid:"MaxSize(120)"`
	TipoNota      int    `form:"tipo_nota"`
	AtendimentoId int    `form:"atendimento_id" 	validate:"min:0"`
	Conteudo      string `form:"conteudo"   		alias:"Conteudo"`
	CorId         int    `form:"cor_id"`
	EtiquetaId    int    `form:"etiqueta_id" 		validate:"min:0"`
	CriadoEm      int    `form:"criado_em"`
	AlteradoEm    int    `form:"alterado_em"`
	Estado        int8   `form:"estado"    		alias:"Estado"`
}

// TableName
func (*Nota) TableName() string {
	return "nota"
}

// SearchField
func (*Nota) SearchField() []string {
	return []string{"nome"}
}

// NoDeletionId
func (*Nota) NoDeletionId() []int {
	return []int{}
}

// WhereField
func (*Nota) WhereField() []string {
	return []string{}
}

// TimeField
func (*Nota) TimeField() []string {
	return []string{}
}

//Registre o modelo definido no init
func init() {
	orm.RegisterModel(new(Nota))
}

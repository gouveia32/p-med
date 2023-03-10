/**
@author : CCB_f
@
*/
package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Atendimento struct {
	Id         int    `form:"id"			alias:"ID" `
	PacienteId int    `form:"paciente_id"	validate:"min:0"`
	Nome       string `form:"nome"      	alias:"Nome"		valid:"Required;MaxSize(120)"`
	Acompanhante  string  `form:"acompanhante"    		alias:"Acompanhante"  		valid:"MaxSize(120)"`
	Conteudo   string `form:"conteudo"   	alias:"Conteudo"`
	CorId      int    `form:"cor_id"`
	EtiquetaId int    `form:"etiqueta_id" 		validate:"min:0"`
	//CriadorId   int    `form:"criador_id"`
	//AlteradorId int    `form:"alterador_id"`
	CriadoEm   int  `form:"criado_em"`
	AlteradoEm int  `form:"alterado_em"`
	Estado     int8 `form:"estado"`
	//Paciente    *Paciente `orm:"reverse(one)"`
}

/* type AtendimentoPaciente struct {
	Id           int64  `form:"id"				alias:"ID" `
	NomePaciente string `form:"nome_paciente"	alias:"NomePaciente"`
	Email        string `form:"email"      		alias:"Email"`
	Telefone     string `form:"telefone"      	alias:"Telefone"`
	QtAtend      int    `form:"qt_atend"      	alias:"QtAtend"`
	Estado       int    `form:"estado"`
} */

// TableName
func (*Atendimento) TableName() string {
	return "atendimento"
}

// SearchField
func (*Atendimento) SearchField() []string {
	return []string{"nome"}
}

// SearchField
/* func (*AtendimentoPaciente) SearchFieldPaciente() []string {
	return []string{"nome", "telefone"}
} */

// NoDeletionId
func (*Atendimento) NoDeletionId() []int {
	return []int{}
}

// WhereField
func (*Atendimento) WhereField() []string {
	return []string{}
}

// TimeField
func (*Atendimento) TimeField() []string {
	return []string{}
}

//Registre o modelo definido no init
func init() {
	orm.RegisterModel(new(Atendimento))
}

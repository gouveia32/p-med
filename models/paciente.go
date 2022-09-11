/**
@author : CCB_f
@
*/
package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Paciente struct {
	Id         int64  `form:"id"        		alias:"ID" `
	Nome       string `form:"nome"      		alias:"Nome"    		valid:"Required;MaxSize(120)"`
	Sexo       string `form:"sexo"	alias:"Sexo"	valid:"MaxSize(1)"`
	Nascimento string `orm:"column(nascimento);size(10);default(0)" description:"Data de Nascimento" json:"nascimento"`
	//Nascimento    time.Time `orm:"type(datetime)"`
	Email         string  `form:"email"    			alias:"email"  			valid:"MaxSize(60)"`
	Telefone      string  `form:"telefone"    		alias:"Telefone"  		valid:"MaxSize(20)"`
	Logradoro     string  `form:"logradoro"    		alias:"logradoro"  		valid:"MaxSize(100)"`
	Numero        string  `form:"numero"    		alias:"Numero"  		valid:"MaxSize(9)"`
	Bairro        string  `form:"bairro"    	alias:"Bairro"  	valid:"MaxSize(40)"`
	Municipio     string  `form:"municipio"    		alias:"Municipio"  		valid:"MaxSize(40)"`
	Uf            string  `form:"uf"    			alias:"Uf"  			valid:"MaxSize(2)"`
	Cep           string  `form:"cep"    			alias:"Cep"  			valid:"MaxSize(9)"`
	Altura        int     `form:"altura"    			alias:"Altura"`
	Peso          float64 `form:"peso"    			alias:"Peso" orm:"digits(12);decimals(2)"`
	Imagem        string  `form:"imagem"    			alias:"Imagem"  			valid:"MaxSize(100)"`
	CriadoEm      int     `orm:"column(criado_em);size(10);default(0)" description:"Hora da criação" json:"criado_em"`
	AlteradoEm    int     `orm:"column(alterado_em);size(10);default(0)" description:"Hora da alteração" json:"alterado_em"`
	NoSelecionado int     `form:"no_selecionado"    			alias:"NoSelecionado"`
	Estado        int8    `form:"estado"    		alias:"Estado"`
	//Atendimentos  []*Atendimento `orm:"reverse(many)" json:"atendimentos"`
}

// TableName
func (*Paciente) TableName() string {
	return "paciente"
}

// SearchField
func (*Paciente) SearchField() []string {
	return []string{"nome", "telefone"}
}

// NoDeletionId
func (*Paciente) NoDeletionId() []int {
	return []int{}
}

// WhereField
func (*Paciente) WhereField() []string {
	return []string{}
}

// TimeField
func (*Paciente) TimeField() []string {
	return []string{}
}

//Registre o modelo definido no init
func init() {
	orm.RegisterModel(new(Paciente))
}

// GetEstados
func (*Paciente) GetEstados() []string {
	return []string{"--", "AL", "AM", "AP", "BA", "CE", "DF", "ES", "GO", "MA", "MT", "MS", "MG", "PR", "PB", "PE", "PI", "RJ", "RN", "RS", "RO", "RR", "SC", "SE", "SP", "PA", "TO"}
}

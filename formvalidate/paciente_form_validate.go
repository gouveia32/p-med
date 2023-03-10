package formvalidate

import "github.com/gookit/validate"

// PacienteForm paciente
type PacienteForm struct {
	Id            int64   `form:"id"        		alias:"ID" `
	Nome          string  `form:"nome"      		alias:"Nome"    		valid:"Required;MaxSize(120)"`
	Sexo          string  `form:"sexo"	alias:"Sexo"	valid:"MaxSize(1)"`
	Nascimento    string  `form:"nascimento" description:"Data de Nascimento" json:"nascimento"`
	Acompanhante  string  `form:"acompanhante"    		alias:"Acompanhante"  		valid:"MaxSize(120)"`
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
}

// Messages Validação personalizada
func (f PacienteForm) Messages() map[string]string {
	return validate.MS{
		"Nome.required": "Nome não pode ser vazio.",
	}
}

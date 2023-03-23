package formvalidate

import "github.com/gookit/validate"

// CampoForm campo
type CampoForm struct {
	Id           int64  `form:"id"        		alias:"ID" `
	Original	 string `form:"original"   		alias:"Original"`
	Tipo		 string `form:"tipo"   			alias:"Tipo"`
	Nome         string `form:"nome"      		alias:"Nome"    	valid:"Required;MaxSize(120)"`
	Descricao    string `form:"descricao"   	alias:"Descricao"`
	Resposta     string `form:"resposta"   		alias:"Resposta"	json:"resposta"`
	ValorInicial string `form:"valor_inicial"   alias:"ValorInicial"`
}

// Messages Validação personalizada
func (f CampoForm) Messages() map[string]string {
	return validate.MS{
		"Nome.required": "Nome não pode ser vazio.",
	}
}

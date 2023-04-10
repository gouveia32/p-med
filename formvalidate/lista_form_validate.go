package formvalidate

import "github.com/gookit/validate"

// CampoForm campo
type ListaForm struct {
	Id          int64  `form:"id"        		alias:"Id" `
	ListaId     int64  `form:"lista_id"       	alias:"ListaId" `
	Nome        string `form:"nome"      		alias:"Nome"    		valid:"Required;MaxSize(20)"`
	Descricao 	string `form:"descricao"   		alias:"Descricao"`
}

// Messages Validação personalizada
func (f ListaForm) Messages() map[string]string {
	return validate.MS{
		"Nome.required": "Nome não pode ser vazio.",
	}
}

package formvalidate

import "github.com/gookit/validate"

// CampoForm campo
type ListaForm struct {
	Id          int64  `form:"id"        		alias:"Id" `
	Nome        string `form:"nome"      		alias:"Nome"    		valid:"Required;MaxSize(20)"`
	Valor	 	string `form:"valor"   			alias:"Valor"`
}

// Messages Validação personalizada
func (f ListaForm) Messages() map[string]string {
	return validate.MS{
		"Nome.required": "Nome não pode ser vazio.",
	}
}

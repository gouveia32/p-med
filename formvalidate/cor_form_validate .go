package formvalidate

import "github.com/gookit/validate"

// CorForm cor
type CorForm struct {
	Id     int    `form:"id"        		alias:"ID" `
	Nome   string `form:"nome"      		alias:"Nome"    		valid:"Required;MaxSize(120)"`
	Cor    string `form:"cor"      		alias:"Cor"    		valid:"Required;MaxSize(6)"`
	Estado int8   `form:"estado"    		alias:"Estado"`
}

// Messages Validação personalizada
func (f CorForm) Messages() map[string]string {
	return validate.MS{
		"Nome.required": "Nome não pode ser vazio.",
	}
}

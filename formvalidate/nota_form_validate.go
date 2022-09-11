package formvalidate

import "github.com/gookit/validate"

// NotaForm nota
type NotaForm struct {
	Id            int    `form:"id"        		alias:"ID" `
	NotaId        int    `form:"NotaId"`
	AtendId       int    `form:"AtendId"`
	Nome          string `form:"nome"      		alias:"Nome"    		valid:"Required;MaxSize(120)"`
	TipoNota      int    `form:"tipo_nota"`
	AtendimentoId int    `form:"atendimento_id" 	validate:"min:0"`
	Conteudo      string `form:"conteudo"   		alias:"Conteudo"`
	CorId         int    `form:"cor_id"`
	EtiquetaId    int    `form:"etiqueta_id" 		validate:"min:0"`
	CriadoEm      int    `form:"criado_em"`
	AlteradoEm    int    `form:"alterado_em"`
	Estado        int8   `form:"estado"    		alias:"Estado"`
}

// Messages Validação personalizada
func (f NotaForm) Messages() map[string]string {
	return validate.MS{
		"Nome.required": "Nome não pode ser vazio.",
	}
}

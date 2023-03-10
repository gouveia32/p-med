package formvalidate

import "github.com/gookit/validate"

// AtendimentoForm atendimento
type AtendimentoForm struct {
	Id         int    `form:"id"        		alias:"ID" `
	AtendId    int    `form:"AtendId"`
	PacienteId int    `form:"paciente_id" validate:"min:0"`
	Nome       string `form:"nome"      		alias:"Nome"    		valid:"Required;MaxSize(120)"`
	Acompanhante  string  `form:"acompanhante"    		alias:"Acompanhante"  		valid:"MaxSize(120)"`
	Conteudo   string `form:"conteudo"   		alias:"Conteudo"`
	CorId      int    `form:"cor_id"`
	EtiquetaId int    `form:"etiqueta_id" 		validate:"min:0"`
	Estado     int    `form:"estado"`
	//CriadorId   int    `form:"criador_id"`
	//AlteradorId int    `form:"alterador_id"`
	CriadoEm   int `form:"criado_em"`
	AlteradoEm int `form:"alterado_em"`
}

// Messages Validação personalizada
func (f AtendimentoForm) Messages() map[string]string {
	return validate.MS{
		"Nome.required": "Nome não pode ser vazio.",
	}
}

package formvalidate

import "github.com/gookit/validate"

// ModeloForm modelo
type ModeloForm struct {
	Id          int64  `form:"id"        		alias:"ID" `
	Nome        string `form:"nome"      		alias:"Nome"    	valid:"Required;MaxSize(120)"`
	Detalhe     string `form:"detalhe"   		alias:"Detalhe"`
	Estado      int    `form:"estado"`
	CriadorId   int    `form:"criador_d"`
	AlteradorId int    `form:"alterador_id`
	CriadoEm    int    `form:"criado_em"`
	AlteradoEm  int    `form:"alterado_em`
}

// Messages Validação personalizada
func (f ModeloForm) Messages() map[string]string {
	return validate.MS{
		"Nome.required": "Nome não pode ser vazio.",
	}
}
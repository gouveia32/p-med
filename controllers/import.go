package controllers

import (
	"p-med/models"
	"p-med/services"

	"time"
)

// CampoController struct
type ImportController struct {
	baseController
}

// Import modelo
func (uc *ImportController) Receita() {
	nomeModelo := ""
	paciente_id, _ := uc.GetInt64("pacienteId", 0)
	paciente := new(models.Paciente)
	if (paciente_id) != 0 {
		var pacienteService services.PacienteService
		paciente = pacienteService.GetPacienteById(paciente_id)
		uc.Data["paciente"] = paciente
	}
	modeloId, _ := uc.GetInt64("modeloId", 0)
	if (modeloId) != 0 {
		var modeloService services.ModeloService
		var campoService services.CampoService

		modelo := modeloService.GetModeloById(modeloId)
		nomeModelo = modelo.Nome
		campos := campoService.CampoGetInConteudo(modelo.Detalhe)

		//modelo.Campos = campos
		//ajusta campos automáticos como data, paciente, etc...
		for _, v := range campos {
			switch v.Tipo {
			case "data":
				v.ValorInicial = time.Now().Format("02 de 01 de 2006")
				break
			case "paciente":
				v.ValorInicial = paciente.Nome
				break
			case "dia":
				v.ValorInicial = time.Now().Format("02")
				break
			case "marque_1":
				break
			case "escolha":
				break
			}
		}

		if nomeModelo == "Consulta" {
			modeloService.Ajustes(modelo, campos)
		}

		uc.Data["Campos"] = campos
		uc.Data["Modelo"] = modelo
	}

	if nomeModelo == "Consulta" {
		uc.TplName = "import/consulta.tpl"
	} else {
		uc.TplName = "import/receita.html"
	}

}

// Import modelo
func (uc *ImportController) Template() {

	modeloId, _ := uc.GetInt64("modeloId", 7)
	if (modeloId) != 0 {
		var modeloService services.ModeloService

		modelo := modeloService.GetModeloById(modeloId)

		uc.Data["Modelo"] = modelo

	}

	uc.TplName = "import/consulta.tpl"
}

package controllers

import (
	"p-med/models"
	"p-med/services"
	"strings"
	"fmt"

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
	data_consulta := uc.GetString("dataConsulta")
	//fmt.Println(data_consulta)
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
			//fmt.Println("Campo: ",v.Nome)
			switch v.Nome {
				
			case "hoje":
				if (v.Tipo == "data") {
					v.ValorInicial = time.Now().Format("2006-01-02")
				}
				if (v.Tipo == "data_extenso") { 
					v.ValorInicial = time.Now().Format("02 de 01 de 2006")
				}
				break
			case "data_consulta":
				if (data_consulta != "") {
					v.ValorInicial = data_consulta
				} else {
					v.ValorInicial = time.Now().Format("2006-01-02")
				}
				
				break				
			case "paciente":
				v.ValorInicial = paciente.Nome
				break
			case "nascimento":
				v.ValorInicial = paciente.Nascimento
				break				
			case "idade":

				
				//loc, _ := time.LoadLocation("Local")
				d1, err := time.Parse("2006-01-02", paciente.Nascimento)

				if (err != nil) {
					continue	
				}

				//fmt.Println(d1)
				

				// Subtrair as duas datas para obter uma duração
				dur := time.Now().Sub(d1)

				// Converter a duração em horas e dividir por 24 para obter o número de dias
				days := dur.Hours() / 24

				if (days > 730) { // dois anos
				// Dividir o número de dias por 365 para obter o número de anos
					years := days / 365
					v.ValorInicial = fmt.Sprintf("%.0f anos",years)
				} else {
					monthes := days / 30
					v.ValorInicial = fmt.Sprintf("%.0f meses",monthes)
				}
				
				break
			case "endereco":
				v.ValorInicial = paciente.Logradoro + ", " + paciente.Numero
				break
			case "acompanhante":
				v.ValorInicial = paciente.Acompanhante
				break
			case "dia":
				v.ValorInicial = time.Now().Format("02")
				break
			case "marque_1":
				break
			case "escolha":
				fmt.Println("escolha",v.Tipo)
				break
			}
			if (v.ValorInicial != "") {
				modelo.Detalhe = strings.ReplaceAll(modelo.Detalhe, "[[" + v.Original + "]]", v.ValorInicial)
			}




/* 			if (v.ValorInicial != "") {
				modelo.Detalhe += v.ValorInicial 
			} */
		}

		if nomeModelo == "Anaminese" {
			modeloService.Ajustes(modelo, campos)
		}

		//fmt.Println(campos)
		uc.Data["Campos"] = campos
		uc.Data["Modelo"] = modelo
	}

	if nomeModelo == "Anaminese" {

		uc.TplName = "import/consulta.html"
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

	uc.TplName = "import/consulta.html"
}

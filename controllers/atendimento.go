package controllers

import (
	"p-med/formvalidate"
	"p-med/global"
	"p-med/global/response"
	"p-med/models"
	"p-med/services"
	"p-med/utils"
	"p-med/utils/exceloffice"
	"strconv"
	"strings"
	"time"

	"github.com/adam-hanna/arrayOperations"
	"github.com/gookit/validate"
)

// AtendimentoController struct
type AtendimentoController struct {
	baseController
}

// Index
func (uc *AtendimentoController) Index() {
	var atendimentoService services.AtendimentoService

	Busca := ""
	for k, val := range gQueryParams {
		if k == "_keywords" {
			Busca = val[0]
		}
	}
	//Busca := gQueryParams
	estado, _ := uc.GetInt("estado", 1)
	BuscaEstado := ""
	switch estado {
	case 0:
		BuscaEstado = "estado=0"
	case 1:
		BuscaEstado = "estado=1"
	}

	cond := ""
	if BuscaEstado != "" {
		cond = "WHERE " + BuscaEstado
	}

	if Busca != "" {
		if cond == "" {
			cond = "WHERE nome LIKE '%" + Busca + "%'"
		} else {
			cond = cond + " AND nome LIKE '%" + Busca + "%'"
		}
	}

	//
	atendimentoLevelMap := make(map[int]string)

	//data, pagination := atendimentoService.GetPaginateDataAtend(admin["per_page"].(int), gQueryParams)
	data, pagination := atendimentoService.PacienteAtendGetList(1, admin["per_page"].(int), cond)

	//fmt.Println(data)
	uc.Data["data"] = data
	uc.Data["paginate"] = pagination
	uc.Data["atendimento_level_map"] = atendimentoLevelMap

	uc.Layout = "public/base.html"
	uc.TplName = "atendimento/index.html"
}

// Vizualizar
func (uc *AtendimentoController) Vizualizar() {

	data := uc.GetString("form_data")

	//
	//fmt.Println("Atend:", data)

	uc.Data["data"] = data
	/* 	uc.Layout = "public/base.html"
	   	uc.TplName = "atendimento/vizualizar.html" */
	uc.ServeJSON()

}

// Export
func (uc *AtendimentoController) Export() {
	exportData := uc.GetString("export_data")
	if exportData == "1" {
		var atendimentoService services.AtendimentoService

		data := atendimentoService.GetExportData(gQueryParams)
		header := []string{"ID", "Nome", "Conteudo"}
		body := [][]string{}
		for _, item := range data {
			record := []string{
				strconv.FormatInt(int64(item.Id), 10),
			}
			record = append(record, item.Nome)

			body = append(body, record)
		}
		uc.Ctx.ResponseWriter.Header().Set("a", "b")
		exceloffice.ExportData(header, body, "atendimento-"+time.Now().Format("2006-01-02"), "", "", uc.Ctx.ResponseWriter)
	}

	response.Error(uc.Ctx)
}

// Add
func (uc *AtendimentoController) Add() {

	uc.Layout = "public/base.html"
	uc.TplName = "atendimento/add.html"
}

// Create
func (uc *AtendimentoController) Create() {
	var atendimentoForm formvalidate.AtendimentoForm

	if err := uc.ParseForm(&atendimentoForm); err != nil {
		response.ErrorWithMessage(err.Error(), uc.Ctx)
	}

	v := validate.Struct(atendimentoForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), uc.Ctx)
	}

	var atendimentoService services.AtendimentoService
	insertID := atendimentoService.Create(&atendimentoForm)

	url := global.URL_BACK

	if insertID > 0 {
		paciente_id, _ := uc.GetInt64("paciente_id", 0)
		var pacienteService services.PacienteService
		paciente := pacienteService.GetPacienteById(paciente_id)
		var no_selecionado = paciente.NoSelecionado

		pacienteService.PacienteChangeNoSelecionado(paciente_id, int64(no_selecionado+1))

		response.SuccessWithMessageAndUrl("Adicionado com sucesso", url, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Atendimento
func (uc *AtendimentoController) Atendimento() {
	//uc.Data["zTree"] = true //Introduzir o ztre

	id, _ := uc.GetInt64("id", 0)
	if id <= 0 {
		response.ErrorWithMessage("Param is error.", uc.Ctx)
	}

	var pacienteService services.PacienteService
	var corService services.CorService
	var modeloService services.ModeloService

	corlists, _ := corService.GetPaginateData(100, gQueryParams)
	uc.Data["corlists"] = corlists

	modelolists, _ := modeloService.GetPaginateData(100, gQueryParams)
	uc.Data["modelolists"] = modelolists

	paciente := pacienteService.GetPacienteById(id)
	uc.Data["paciente"] = paciente

	//uc.GetNodes()

	uc.Layout = "public/base.html"
	uc.TplName = "atendimento/atendimento.html"

}

/*
-------------------------------------------------------------------------------

	GetNode
	obter todos os nós
*/
func (uc *AtendimentoController) GetNodes() {

	var atendimentoService services.AtendimentoService
	var notaService services.NotaService
	id := 2
	pid := 1

	paciente_id, _ := uc.GetInt("paciente_id", 0)

	//fmt.Println("no GetNodes")
	pId := strconv.Itoa(paciente_id)
	lstAtend, _ := atendimentoService.AtendimentoGetList(1, 100, " WHERE estado=1 AND paciente_id="+pId)
	list := make([]map[string]interface{}, 0)

	row0 := make(map[string]interface{})
	row0["id"] = 1
	row0["pId"] = 0
	row0["name"] = "Atendimentos"
	list = append(list, row0)

	for _, v := range lstAtend {
		row := make(map[string]interface{})
		row["id"] = id
		row["pId"] = 1
		row["AtendId"] = v.Id
		row["name"] = time.Unix(int64(v.CriadoEm), 0).Format("2006-01-02") + " " + v.Nome
		row["data_consulta"] = time.Unix(int64(v.CriadoEm), 0).Format("2006-01-02")
		pid = id //time.Unix(paciente.Nascimento, 0)

		list = append(list, row)
		lstNota, count2 := notaService.NotasGetList(1, 100, " WHERE estado=1 AND atendimento_id="+strconv.Itoa(v.Id))
		if count2 > 0 {
			for _, v2 := range lstNota {
				row2 := make(map[string]interface{})
				row2["pId"] = pid
				id++
				row2["id"] = id
				row2["name"] = v2.Nome
				row2["data_consulta"] = row["data_consulta"]
				row2["AtendId"] = v.Id
				row2["NotaId"] = v2.Id

				list = append(list, row2)
			}
		}
		id++
	}
	uc.Data["json"] = list
	uc.ServeJSON()
}

/*
-------------------------------------------------------------------------------

	GetNode
*/
func (uc *AtendimentoController) GetNode() {
	var notaService services.NotaService
	var pacienteService services.PacienteService
	var corService services.CorService
	var atendimentoService services.AtendimentoService

	row := make(map[string]interface{})
	id, _ := uc.GetInt("id", 0)
	paciente_id, _ := uc.GetInt("pacienteId", 0)
	if paciente_id > 0 {
		pacienteService.PacienteChangeNoSelecionado(int64(paciente_id), int64(id))
	}

	// Verifica primeiro se é nota
	n_id, _ := uc.GetInt("NotaId", 0)
	if (n_id) != 0 {
		nota := notaService.GetNotaById(n_id)
		/* 		if err == nil {
			uc.ajaxMsg(err.Error(), MSG_ERR)
		} */
		row["id"] = nota.Id
		row["TipoNota"] = nota.TipoNota
		row["NotaId"] = nota.Id

		row["AtendId"] = nota.AtendimentoId
		if nota.AtendimentoId > 0 {
			atendimento := atendimentoService.GetAtendimentoById(nota.AtendimentoId)
			row["data_consulta"] = time.Unix(int64(atendimento.CriadoEm), 0).Format("2006-01-02")
		}
		row["Nome"] = nota.Nome

		row["Conteudo"] = nota.Conteudo
		row["nota"] = nota

		cor := corService.GetCorById(nota.CorId)
		if cor != nil {
			row["Cor"] = "#" + cor.Cor
			row["CorId"] = cor.Id
		} else {
			row["Cor"] = "#fff00"
		}
		row["idAtual"] = id

	} else {
		atendimento_id, _ := uc.GetInt("AtendId", 0)
		atendimento := atendimentoService.GetAtendimentoById(atendimento_id)
		/* 		if err == nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		} */
		row["id"] = atendimento.Id
		row["AtendId"] = atendimento.Id
		row["NotaId"] = 0
		row["TipoNota"] = 0
		row["Nome"] = atendimento.Nome
		row["data_consulta"] = time.Unix(int64(atendimento.CriadoEm), 0).Format("2006-01-02")
		row["Conteudo"] = atendimento.Conteudo
		row["nota"] = atendimento

		row["CorId"] = atendimento.CorId
		cor := corService.GetCorById(atendimento.CorId)
		if cor != nil {
			row["Cor"] = "#" + cor.Cor
			row["CorId"] = cor.Id
		} else {
			row["Cor"] = "#fff00"
		}

		row["idAtual"] = id
	}

	uc.Data["json"] = row
	uc.ServeJSON()
}

/*
-------------------------------------------------------------------------------

	AjusteConteudo
*/
func (uc *AtendimentoController) AjusteConteudo() {
	var notaForm formvalidate.NotaForm

	if err := uc.ParseForm(&notaForm); err != nil {
		response.ErrorWithMessage(err.Error(), uc.Ctx)
	}

	v := validate.Struct(notaForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), uc.Ctx)
	}

	row := make(map[string]interface{})
	conteudo := uc.GetString("conteudo")
	//fmt.Println("Conteudo")
	//fmt.Println(notaForm)
	//> Sim <

	//result := strings.ReplaceAll(conteudo, "> Sim <", "> [Sim] <")

	row["conteudo"] = conteudo

	uc.Data["json"] = row
	uc.ServeJSON()
}

/*
-------------------------------------------------------------------------------

	Edit Paciente
*/
func (uc *AtendimentoController) EditPaciente() {
	var pacienteService services.PacienteService

	uc.Data["pageTitle"] = "Alterando Paciente"

	uc.Data["uf"] = new(models.Paciente).GetEstados()

	id, _ := uc.GetInt64("paciente_id", 0)

	paciente := pacienteService.GetPacienteById(id)

	row := make(map[string]interface{})
	row["id"] = paciente.Id
	row["nome"] = paciente.Nome
	row["acompanhante"] = paciente.Acompanhante
	row["sexo"] = paciente.Sexo
	row["imagem"] = paciente.Imagem
	row["logradoro"] = paciente.Logradoro
	row["municipio"] = paciente.Municipio
	row["numero"] = paciente.Numero
	row["uf"] = paciente.Uf
	row["cep"] = paciente.Cep
	row["altura"] = paciente.Altura
	row["peso"] = paciente.Peso
	row["telefone"] = paciente.Telefone
	row["email"] = paciente.Email
	row["nascimento"] = paciente.Nascimento
	row["estado"] = int8(paciente.Estado)
	//fmt.Println(paciente.Acompanhante)

	uc.Data["paciente"] = row
	uc.TplName = "atendimento/editpaciente.html"

}

// Update
func (uc *AtendimentoController) Update() {
	var atendimentoForm formvalidate.AtendimentoForm
	if err := uc.ParseForm(&atendimentoForm); err != nil {
		response.ErrorWithMessage(err.Error(), uc.Ctx)
	}

	if atendimentoForm.Id <= 0 {
		response.ErrorWithMessage("Parâmetros com erro.", uc.Ctx)
	}

	v := validate.Struct(atendimentoForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), uc.Ctx)
	}

	var atendimentoService services.AtendimentoService
	num := atendimentoService.Update(&atendimentoForm)

	if num > 0 {
		response.Success(uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Enable
func (uc *AtendimentoController) Enable() {
	idStr := uc.GetString("id")
	ids := make([]int, 0)
	var idArr []int

	if idStr == "" {
		uc.Ctx.Input.Bind(&ids, "id")
	} else {
		id, _ := strconv.Atoi(idStr)
		idArr = append(idArr, id)
	}

	if len(ids) > 0 {
		idArr = ids
	}

	if len(idArr) == 0 {
		response.ErrorWithMessage("Selecione um nível de usuário.", uc.Ctx)
	}

	var atendimentoService services.AtendimentoService
	num := atendimentoService.Enable(idArr)
	if num > 0 {
		response.SuccessWithMessageAndUrl("Operação bem-sucedida", global.URL_RELOAD, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Disable
func (uc *AtendimentoController) Disable() {
	idStr := uc.GetString("id")
	ids := make([]int, 0)
	var idArr []int

	if idStr == "" {
		uc.Ctx.Input.Bind(&ids, "id")
	} else {
		id, _ := strconv.Atoi(idStr)
		idArr = append(idArr, id)
	}

	if len(ids) > 0 {
		idArr = ids
	}

	if len(idArr) == 0 {
		response.ErrorWithMessage("Selecione usuários com deficiência.", uc.Ctx)
	}

	var atendimentoService services.AtendimentoService
	num := atendimentoService.Disable(idArr)
	if num > 0 {
		response.SuccessWithMessageAndUrl("Operação bem-sucedida", global.URL_RELOAD, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Del
func (uc *AtendimentoController) Del() {
	idStr := uc.GetString("a_Id")
	ids := make([]int, 0)
	var idArr []int

	if idStr == "" {
		uc.Ctx.Input.Bind(&ids, "id")
	} else {
		id, _ := strconv.Atoi(idStr)
		idArr = append(idArr, id)
	}

	if len(ids) > 0 {
		idArr = ids
	}

	if len(idArr) == 0 {
		response.ErrorWithMessage("ID de parâmetro errado.", uc.Ctx)
	}

	noDeletionID := new(models.Atendimento).NoDeletionId()

	m, b := arrayOperations.Intersect(noDeletionID, idArr)

	if len(noDeletionID) > 0 && len(m.Interface().([]int)) > 0 && b {
		response.ErrorWithMessage("ID é"+strings.Join(utils.IntArrToStringArr(noDeletionID), ",")+"os dados não podem ser excluídos!", uc.Ctx)
	}

	var atendimentoService services.AtendimentoService
	count := atendimentoService.Del(idArr)

	if count > 0 {
		paciente_id, _ := uc.GetInt64("paciente_Id", 0)
		var pacienteService services.PacienteService
		paciente := pacienteService.GetPacienteById(paciente_id)
		var no_selecionado = paciente.NoSelecionado

		pacienteService.PacienteChangeNoSelecionado(paciente_id, int64(no_selecionado-1))

		//excluir todas notas associadas a este atendimento
		var notaService services.NotaService
		lstNota, count := notaService.NotasGetList(1, 100, " WHERE estado=1 AND atendimento_id="+idStr)
		if count > 0 {
			var idArr []int
			for _, v := range lstNota {
				idArr = append(idArr, v.Id)
			}
			notaService.Del(idArr)
		}

		response.SuccessWithMessageAndUrl("Operação bem-sucedida", global.URL_RELOAD, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

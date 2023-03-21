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

// PacienteController struct
type PacienteController struct {
	baseController
}

// Index
func (uc *PacienteController) Index() {
	var pacienteService services.PacienteService

	//
	pacienteLevelMap := make(map[int]string)

	data, pagination := pacienteService.GetPaginateData(admin["per_page"].(int), gQueryParams)

	uc.Data["data"] = data
	uc.Data["paginate"] = pagination
	uc.Data["paciente_level_map"] = pacienteLevelMap

	uc.Layout = "public/base.html"
	uc.TplName = "paciente/index.html"
}

// Export
func (uc *PacienteController) Export() {
	exportData := uc.GetString("export_data")
	if exportData == "1" {
		var pacienteService services.PacienteService

		data := pacienteService.GetExportData(gQueryParams)
		header := []string{"ID", "Nome", "Contato", "Email", "Telefone 1", "Telefone 2", "Telefone 3"}
		body := [][]string{}
		for _, item := range data {
			record := []string{
				strconv.FormatInt(item.Id, 10),
			}
			record = append(record, item.Nome)
			record = append(record, item.Email)
			record = append(record, item.Telefone)

			body = append(body, record)
		}
		uc.Ctx.ResponseWriter.Header().Set("a", "b")
		exceloffice.ExportData(header, body, "paciente-"+time.Now().Format("2006-01-02-15-04-05"), "", "", uc.Ctx.ResponseWriter)
	}

	response.Error(uc.Ctx)
}

// Add
func (uc *PacienteController) Add() {

	uc.Data["uf"] = new(models.Paciente).GetEstados()
	uc.Layout = "public/base.html"
	uc.TplName = "paciente/add.html"
}

// Create
func (uc *PacienteController) Create() {
	var pacienteForm formvalidate.PacienteForm

	
	if err := uc.ParseForm(&pacienteForm); err != nil {
		response.ErrorWithMessage(err.Error(), uc.Ctx)
	}

	v := validate.Struct(pacienteForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), uc.Ctx)
	}

	var pacienteService services.PacienteService
	insertID := pacienteService.Create(&pacienteForm)

	url := global.URL_BACK

	if insertID > 0 {
		response.SuccessWithMessageAndUrl("Adicionado com sucesso", url, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Edit
func (uc *PacienteController) Edit() {
	id, _ := uc.GetInt64("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Param is error.", uc.Ctx)
	}

	var pacienteService services.PacienteService

	paciente := pacienteService.GetPacienteById(id)
	if paciente == nil {
		response.ErrorWithMessage("Não encontrado informações para o Id.", uc.Ctx)
	}

	//
	dn := paciente.Nascimento
	if dn == "0001-01-01" {
		paciente.Nascimento = ""
	}

	uc.Data["data"] = paciente
	uc.Data["uf"] = new(models.Paciente).GetEstados()

	uc.Layout = "public/base.html"
	uc.TplName = "paciente/edit.html"
}

// Update
func (uc *PacienteController) Update() {
	var pacienteForm formvalidate.PacienteForm
	if err := uc.ParseForm(&pacienteForm); err != nil {
		response.ErrorWithMessage(err.Error(), uc.Ctx)
	}

	if pacienteForm.Id <= 0 {
		response.ErrorWithMessage("Parâmetros com erro.", uc.Ctx)
	}

	v := validate.Struct(pacienteForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), uc.Ctx)
	}

	var pacienteService services.PacienteService
	num := pacienteService.Update(&pacienteForm)

	if num > 0 {
		response.Success(uc.Ctx)
		//response.SuccessWithMessageAndUrl("Atualizado!", global.URL_RELOAD, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Enable
func (uc *PacienteController) Enable() {
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

	var pacienteService services.PacienteService
	num := pacienteService.Enable(idArr)
	if num > 0 {
		response.SuccessWithMessageAndUrl("Operação bem-sucedida", global.URL_RELOAD, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Disable
func (uc *PacienteController) Disable() {
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

	var pacienteService services.PacienteService
	num := pacienteService.Disable(idArr)
	if num > 0 {
		response.SuccessWithMessageAndUrl("Operação bem-sucedida", global.URL_RELOAD, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Del
func (uc *PacienteController) Del() {
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
		response.ErrorWithMessage("ID de parâmetro errado.", uc.Ctx)
	}

	noDeletionID := new(models.Paciente).NoDeletionId()

	m, b := arrayOperations.Intersect(noDeletionID, idArr)

	if len(noDeletionID) > 0 && len(m.Interface().([]int)) > 0 && b {
		response.ErrorWithMessage("ID é"+strings.Join(utils.IntArrToStringArr(noDeletionID), ",")+"os dados não podem ser excluídos!", uc.Ctx)
	}

	var pacienteService services.PacienteService
	count := pacienteService.Del(idArr)

	if count > 0 {
		response.SuccessWithMessageAndUrl("Operação bem-sucedida", global.URL_RELOAD, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

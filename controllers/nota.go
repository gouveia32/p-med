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

// NotaController struct
type NotaController struct {
	baseController
}

// Index
func (uc *NotaController) Index() {
	var notaService services.NotaService

	//
	notaLevelMap := make(map[int]string)

	data, pagination := notaService.GetPaginateData(admin["per_page"].(int), gQueryParams)

	uc.Data["data"] = data
	uc.Data["paginate"] = pagination
	uc.Data["nota_level_map"] = notaLevelMap

	uc.Layout = "public/base.html"
	uc.TplName = "nota/index.html"
}

// Export
func (uc *NotaController) Export() {
	exportData := uc.GetString("export_data")
	if exportData == "1" {
		var notaService services.NotaService

		data := notaService.GetExportData(gQueryParams)
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
		exceloffice.ExportData(header, body, "nota-"+time.Now().Format("2006-01-02"), "", "", uc.Ctx.ResponseWriter)
	}

	response.Error(uc.Ctx)
}

// Add
func (uc *NotaController) Add() {

	uc.Layout = "public/base.html"
	uc.TplName = "nota/add.html"
}

// Create
func (uc *NotaController) Create() {
	var notaForm formvalidate.NotaForm

	if err := uc.ParseForm(&notaForm); err != nil {
		response.ErrorWithMessage(err.Error(), uc.Ctx)
	}

	v := validate.Struct(notaForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), uc.Ctx)
	}

	var notaService services.NotaService
	insertID := notaService.Create(&notaForm)

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

// Edit
func (uc *NotaController) Edit() {
	id, _ := uc.GetInt("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Param is error.", uc.Ctx)
	}

	var notaService services.NotaService

	nota := notaService.GetNotaById(id)
	if nota == nil {
		response.ErrorWithMessage("Não encontrado informações para o Id.", uc.Ctx)
	}

	//

	uc.Data["data"] = nota

	uc.Layout = "public/base.html"
	uc.TplName = "nota/edit.html"
}

// Update
func (uc *NotaController) Update() {
	var notaForm formvalidate.NotaForm
	if err := uc.ParseForm(&notaForm); err != nil {
		response.ErrorWithMessage(err.Error(), uc.Ctx)
	}

	if notaForm.Id <= 0 {
		response.ErrorWithMessage("Parâmetros com erro.", uc.Ctx)
	}

	v := validate.Struct(notaForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), uc.Ctx)
	}

	var notaService services.NotaService
	num := notaService.Update(&notaForm)

	if num > 0 {
		response.Success(uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Enable
func (uc *NotaController) Enable() {
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

	var notaService services.NotaService
	num := notaService.Enable(idArr)
	if num > 0 {
		response.SuccessWithMessageAndUrl("Operação bem-sucedida", global.URL_RELOAD, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Disable
func (uc *NotaController) Disable() {
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

	var notaService services.NotaService
	num := notaService.Disable(idArr)
	if num > 0 {
		response.SuccessWithMessageAndUrl("Operação bem-sucedida", global.URL_RELOAD, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Del
func (uc *NotaController) Del() {
	idStr := uc.GetString("n_Id")
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

	noDeletionID := new(models.Nota).NoDeletionId()

	m, b := arrayOperations.Intersect(noDeletionID, idArr)

	if len(noDeletionID) > 0 && len(m.Interface().([]int)) > 0 && b {
		response.ErrorWithMessage("ID é"+strings.Join(utils.IntArrToStringArr(noDeletionID), ",")+"os dados não podem ser excluídos!", uc.Ctx)
	}

	var notaService services.NotaService
	count := notaService.Del(idArr)

	if count > 0 {
		paciente_id, _ := uc.GetInt64("paciente_Id", 0)
		var pacienteService services.PacienteService
		paciente := pacienteService.GetPacienteById(paciente_id)
		var no_selecionado = paciente.NoSelecionado

		pacienteService.PacienteChangeNoSelecionado(paciente_id, int64(no_selecionado-1))

		response.SuccessWithMessageAndUrl("Operação bem-sucedida", global.URL_RELOAD, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

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

// ModeloController struct
type ModeloController struct {
	baseController
}

// Index
func (uc *ModeloController) Index() {
	var modeloService services.ModeloService

	//
	modeloLevelMap := make(map[int]string)

	data, pagination := modeloService.GetPaginateData(admin["per_page"].(int), gQueryParams)

	uc.Data["data"] = data
	uc.Data["paginate"] = pagination
	uc.Data["modelo_level_map"] = modeloLevelMap

	uc.Layout = "public/base.html"
	uc.TplName = "modelo/index.html"
}

// Export
func (uc *ModeloController) Export() {
	exportData := uc.GetString("export_data")
	if exportData == "1" {
		var modeloService services.ModeloService

		data := modeloService.GetExportData(gQueryParams)
		header := []string{"ID", "Nome", "Detalhe"}
		body := [][]string{}
		for _, item := range data {
			record := []string{
				strconv.FormatInt(item.Id, 10),
			}
			record = append(record, item.Nome)

			body = append(body, record)
		}
		uc.Ctx.ResponseWriter.Header().Set("a", "b")
		exceloffice.ExportData(header, body, "modelo-"+time.Now().Format("2006-01-02-15-04-05"), "", "", uc.Ctx.ResponseWriter)
	}

	response.Error(uc.Ctx)
}

// Add
func (uc *ModeloController) Add() {

	uc.Layout = "public/base.html"
	uc.TplName = "modelo/add.html"
}

// Create
func (uc *ModeloController) Create() {
	var modeloForm formvalidate.ModeloForm

	if err := uc.ParseForm(&modeloForm); err != nil {
		response.ErrorWithMessage(err.Error(), uc.Ctx)
	}

	v := validate.Struct(modeloForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), uc.Ctx)
	}

	var modeloService services.ModeloService
	insertID := modeloService.Create(&modeloForm)

	url := global.URL_BACK

	if insertID > 0 {
		response.SuccessWithMessageAndUrl("Adicionado com sucesso", url, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Edit
func (uc *ModeloController) Edit() {
	id, _ := uc.GetInt64("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Parâmetros com erro.", uc.Ctx)
	}

	var modeloService services.ModeloService

	modelo := modeloService.GetModeloById(id)
	if modelo == nil {
		response.ErrorWithMessage("Não encontrado informações para o Id.", uc.Ctx)
	}

	//

	uc.Data["data"] = modelo

	uc.Layout = "public/base.html"
	uc.TplName = "modelo/edit.html"
}

// Update
func (uc *ModeloController) Update() {
	var modeloForm formvalidate.ModeloForm
	if err := uc.ParseForm(&modeloForm); err != nil {
		response.ErrorWithMessage(err.Error(), uc.Ctx)
	}

	if modeloForm.Id <= 0 {
		response.ErrorWithMessage("Parâmetros com erro.", uc.Ctx)
	}

	v := validate.Struct(modeloForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), uc.Ctx)
	}

	var modeloService services.ModeloService
	num := modeloService.Update(&modeloForm)

	if num > 0 {
		response.Success(uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Enable
func (uc *ModeloController) Enable() {
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

	var modeloService services.ModeloService
	num := modeloService.Enable(idArr)
	if num > 0 {
		response.SuccessWithMessageAndUrl("Operação bem-sucedida", global.URL_RELOAD, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Disable
func (uc *ModeloController) Disable() {
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

	var modeloService services.ModeloService
	num := modeloService.Disable(idArr)
	if num > 0 {
		response.SuccessWithMessageAndUrl("Operação bem-sucedida", global.URL_RELOAD, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Del
func (uc *ModeloController) Del() {
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

	noDeletionID := new(models.Modelo).NoDeletionId()

	m, b := arrayOperations.Intersect(noDeletionID, idArr)

	if len(noDeletionID) > 0 && len(m.Interface().([]int)) > 0 && b {
		response.ErrorWithMessage("ID é"+strings.Join(utils.IntArrToStringArr(noDeletionID), ",")+"os dados não podem ser excluídos!", uc.Ctx)
	}

	var modeloService services.ModeloService
	count := modeloService.Del(idArr)

	if count > 0 {
		response.SuccessWithMessageAndUrl("Operação bem-sucedida", global.URL_RELOAD, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

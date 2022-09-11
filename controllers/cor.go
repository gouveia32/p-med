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

// CorController struct
type CorController struct {
	baseController
}

// Index
func (uc *CorController) Index() {
	var corService services.CorService

	//
	corLevelMap := make(map[int]string)

	data, pagination := corService.GetPaginateData(admin["per_page"].(int), gQueryParams)

	uc.Data["data"] = data
	uc.Data["paginate"] = pagination
	uc.Data["cor_level_map"] = corLevelMap

	uc.Layout = "public/base.html"
	uc.TplName = "cor/index.html"
}

// Export
func (uc *CorController) Export() {
	exportData := uc.GetString("export_data")
	if exportData == "1" {
		var corService services.CorService

		data := corService.GetExportData(gQueryParams)
		header := []string{"ID", "Nome"}
		body := [][]string{}
		for _, item := range data {
			record := []string{
				strconv.FormatInt(int64(item.Id), 10),
			}
			record = append(record, item.Nome)

			body = append(body, record)
		}
		uc.Ctx.ResponseWriter.Header().Set("a", "b")
		exceloffice.ExportData(header, body, "cor-"+time.Now().Format("2006-01-02-15-04-05"), "", "", uc.Ctx.ResponseWriter)
	}

	response.Error(uc.Ctx)
}

// Add
func (uc *CorController) Add() {

	uc.Layout = "public/base.html"
	uc.TplName = "cor/add.html"
}

// Create
func (uc *CorController) Create() {
	var corForm formvalidate.CorForm

	if err := uc.ParseForm(&corForm); err != nil {
		response.ErrorWithMessage(err.Error(), uc.Ctx)
	}

	v := validate.Struct(corForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), uc.Ctx)
	}

	var corService services.CorService
	insertID := corService.Create(&corForm)

	url := global.URL_BACK

	if insertID > 0 {
		response.SuccessWithMessageAndUrl("Adicionado com sucesso", url, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Edit
func (uc *CorController) Edit() {
	id, _ := uc.GetInt("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Parâmetros com erro.", uc.Ctx)
	}

	var corService services.CorService

	cor := corService.GetCorById(id)
	if cor == nil {
		response.ErrorWithMessage("Não encontrado informações para o Id.", uc.Ctx)
	}

	//
	uc.Data["data"] = cor

	uc.Layout = "public/base.html"
	uc.TplName = "cor/edit.html"
}

// Update
func (uc *CorController) Update() {
	var corForm formvalidate.CorForm
	if err := uc.ParseForm(&corForm); err != nil {
		response.ErrorWithMessage(err.Error(), uc.Ctx)
	}

	if corForm.Id <= 0 {
		response.ErrorWithMessage("Parâmetros com erro.", uc.Ctx)
	}

	v := validate.Struct(corForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), uc.Ctx)
	}

	var corService services.CorService
	num := corService.Update(&corForm)

	if num > 0 {
		response.Success(uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Enable
func (uc *CorController) Enable() {
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
		response.ErrorWithMessage("Selecione registro(s) antes.", uc.Ctx)
	}

	var corService services.CorService
	num := corService.Enable(idArr)
	if num > 0 {
		response.SuccessWithMessageAndUrl("Operação bem-sucedida", global.URL_RELOAD, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Disable
func (uc *CorController) Disable() {
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
		response.ErrorWithMessage("Selecione registro(s) antes.", uc.Ctx)
	}

	var corService services.CorService
	num := corService.Disable(idArr)
	if num > 0 {
		response.SuccessWithMessageAndUrl("Operação bem-sucedida", global.URL_RELOAD, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Del
func (uc *CorController) Del() {
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

	noDeletionID := new(models.Cor).NoDeletionId()

	m, b := arrayOperations.Intersect(noDeletionID, idArr)

	if len(noDeletionID) > 0 && len(m.Interface().([]int)) > 0 && b {
		response.ErrorWithMessage("ID é"+strings.Join(utils.IntArrToStringArr(noDeletionID), ",")+"os dados não podem ser excluídos!", uc.Ctx)
	}

	var corService services.CorService
	count := corService.Del(idArr)

	if count > 0 {
		response.SuccessWithMessageAndUrl("Operação bem-sucedida", global.URL_RELOAD, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

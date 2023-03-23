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
	//"fmt"

	"github.com/adam-hanna/arrayOperations"
	"github.com/gookit/validate"
)

// CampoController struct
type CampoController struct {
	baseController
}

// Index
func (uc *CampoController) Index() {
	var campoService services.CampoService

	//
	campoLevelMap := make(map[int]string)

	data, pagination := campoService.GetPaginateData(admin["per_page"].(int), gQueryParams)

	//fmt.Println ("data",data)
	uc.Data["data"] = data
	uc.Data["paginate"] = pagination
	uc.Data["campo_level_map"] = campoLevelMap

	uc.Layout = "public/base.html"
	uc.TplName = "campo/index.html"
}

// Export
func (uc *CampoController) Export() {
	exportData := uc.GetString("export_data")
	if exportData == "1" {
		var campoService services.CampoService

		data := campoService.GetExportData(gQueryParams)
		header := []string{"ID", "Nome", "Descrição"}
		body := [][]string{}
		for _, item := range data {
			record := []string{
				strconv.FormatInt(item.Id, 10),
			}
			record = append(record, item.Nome)

			body = append(body, record)
		}
		uc.Ctx.ResponseWriter.Header().Set("a", "b")
		exceloffice.ExportData(header, body, "campo-"+time.Now().Format("2006-01-02"), "", "", uc.Ctx.ResponseWriter)
	}

	response.Error(uc.Ctx)
}

// Add
func (uc *CampoController) Add() {

	uc.Layout = "public/base.html"
	uc.TplName = "campo/add.html"
}

// Create
func (uc *CampoController) Create() {
	var campoForm formvalidate.CampoForm

	if err := uc.ParseForm(&campoForm); err != nil {
		response.ErrorWithMessage(err.Error(), uc.Ctx)
	}

	v := validate.Struct(campoForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), uc.Ctx)
	}

	var campoService services.CampoService
	insertID := campoService.Create(&campoForm)

	url := global.URL_BACK

	if insertID > 0 {
		response.SuccessWithMessageAndUrl("Adicionado com sucesso", url, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Edit
func (uc *CampoController) Edit() {
	id, _ := uc.GetInt64("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Parâmetros com erro.", uc.Ctx)
	}

	var campoService services.CampoService

	campo := campoService.GetCampoById(id)
	if campo == nil {
		response.ErrorWithMessage("Não encontrado informações para o Id.", uc.Ctx)
	}

	//

	uc.Data["data"] = campo

	uc.Layout = "public/base.html"
	uc.TplName = "campo/edit.html"
}

// Update
func (uc *CampoController) Update() {
	var campoForm formvalidate.CampoForm
	if err := uc.ParseForm(&campoForm); err != nil {
		response.ErrorWithMessage(err.Error(), uc.Ctx)
	}

	if campoForm.Id <= 0 {
		response.ErrorWithMessage("Parâmetros com erro.", uc.Ctx)
	}

	v := validate.Struct(campoForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), uc.Ctx)
	}

	var campoService services.CampoService
	num := campoService.Update(&campoForm)

	if num > 0 {
		response.Success(uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Del
func (uc *CampoController) Del() {
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

	noDeletionID := new(models.Campo).NoDeletionId()

	m, b := arrayOperations.Intersect(noDeletionID, idArr)

	if len(noDeletionID) > 0 && len(m.Interface().([]int)) > 0 && b {
		response.ErrorWithMessage("ID é"+strings.Join(utils.IntArrToStringArr(noDeletionID), ",")+"os dados não podem ser excluídos!", uc.Ctx)
	}

	var campoService services.CampoService
	count := campoService.Del(idArr)

	if count > 0 {
		response.SuccessWithMessageAndUrl("Operação bem-sucedida", global.URL_RELOAD, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

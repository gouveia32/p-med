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

// ListaController struct
type ItemListaController struct {
	baseController
}

// Index
func (uc *ItemListaController) Index() {
	var listaService services.ListaService

	//
	listaLevelMap := make(map[int]string)

	data, pagination := listaService.GetPaginateData(admin["per_page"].(int), gQueryParams)

	//fmt.Println ("data",data)
	uc.Data["data"] = data
	uc.Data["paginate"] = pagination
	uc.Data["lista_level_map"] = listaLevelMap

	uc.Layout = "public/base.html"
	uc.TplName = "lista/index.html"
}

// Export
func (uc *ItemListaController) Export() {
	exportData := uc.GetString("export_data")
	if exportData == "1" {
		var listaService services.ListaService

		data := listaService.GetExportData(gQueryParams)
		header := []string{"Id", "Nome", "Valor"}
		body := [][]string{}
		for _, item := range data {
			record := []string{
				strconv.FormatInt(item.Id, 10),
			}
			record = append(record, item.Nome)

			body = append(body, record)
		}
		uc.Ctx.ResponseWriter.Header().Set("a", "b")
		exceloffice.ExportData(header, body, "lista-"+time.Now().Format("2006-01-02"), "", "", uc.Ctx.ResponseWriter)
	}

	response.Error(uc.Ctx)
}

// Add
func (uc *ItemListaController) Add() {

	uc.Layout = "public/base.html"
	uc.TplName = "lista/add.html"
}

// Create
func (uc *ItemListaController) Create() {
	var listaForm formvalidate.ListaForm

	if err := uc.ParseForm(&listaForm); err != nil {
		response.ErrorWithMessage(err.Error(), uc.Ctx)
	}

	v := validate.Struct(listaForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), uc.Ctx)
	}

	var listaService services.ListaService
	insertID := listaService.Create(&listaForm)

	url := global.URL_BACK

	if insertID > 0 {
		response.SuccessWithMessageAndUrl("Adicionado com sucesso", url, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Edit
func (uc *ItemListaController) Edit() {
	id, _ := uc.GetInt64("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Parâmetros com erro.", uc.Ctx)
	}

	var listaService services.ListaService

	lista := listaService.GetListaById(id)
	if lista == nil {
		response.ErrorWithMessage("Não encontrado informações para o Id.", uc.Ctx)
	}

	//

	uc.Data["data"] = lista

	uc.Layout = "public/base.html"
	uc.TplName = "lista/edit.html"
}

// Update
func (uc *ItemListaController) Update() {
	var listaForm formvalidate.ListaForm
	if err := uc.ParseForm(&listaForm); err != nil {
		response.ErrorWithMessage(err.Error(), uc.Ctx)
	}

	if listaForm.Id <= 0 {
		response.ErrorWithMessage("Parâmetros com erro.", uc.Ctx)
	}

	v := validate.Struct(listaForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), uc.Ctx)
	}

	var listaService services.ListaService
	num := listaService.Update(&listaForm)

	if num > 0 {
		response.Success(uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Del
func (uc *ItemListaController) Del() {
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
		response.ErrorWithMessage("Id de parâmetro errado.", uc.Ctx)
	}

	noDeletionID := new(models.Lista).NoDeletionId()

	m, b := arrayOperations.Intersect(noDeletionID, idArr)

	if len(noDeletionID) > 0 && len(m.Interface().([]int)) > 0 && b {
		response.ErrorWithMessage("ID é"+strings.Join(utils.IntArrToStringArr(noDeletionID), ",")+"os dados não podem ser excluídos!", uc.Ctx)
	}

	var listaService services.ListaService
	count := listaService.Del(idArr)

	if count > 0 {
		response.SuccessWithMessageAndUrl("Operação bem-sucedida", global.URL_RELOAD, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

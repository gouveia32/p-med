package controllers

import (
	"p-med/formvalidate"
	"p-med/global"
	"p-med/global/response"
	"p-med/models"
	"p-med/services"
	"p-med/utils"

	//"p-med/utils/exceloffice"
	"strconv"
	"strings"

	//"time"

	//"fmt"

	"github.com/adam-hanna/arrayOperations"
	"github.com/gookit/validate"
)

// ListaController struct
type ListaController struct {
	baseController
}

// Index
func (uc *ListaController) Index() {
	var listaService services.ListaService

	//listas := [2]string{"receita", "exame"}
	listas := listaService.GetAllLista()
	Busca := ""
	for k, val := range gQueryParams {
		if k == "_keywords" {
			Busca = val[0]
		}
	}

	cond := "WHERE lista.nome = '" + listas[0].Nome + "'"

	if Busca != "" {
		if cond == "" {
			cond = "WHERE nome LIKE '%" + Busca + "%'"
		} else {
			cond = cond + " AND nome LIKE '%" + Busca + "%'"
		}
	}

	//fmt.Println("cond:", cond)
	//
	listaLevelMap := make(map[int]string)

	//data, pagination := atendimentoService.GetPaginateDataAtend(admin["per_page"].(int), gQueryParams)
	data, pagination := listaService.GetPaginateDataItemLista(1, admin["per_page"].(int), cond)

	//fmt.Println ("data",data)
	uc.Data["lista"] = listas
	uc.Data["data"] = data
	uc.Data["paginate"] = pagination
	uc.Data["lista_level_map"] = listaLevelMap

	uc.Layout = "public/base.html"
	uc.TplName = "lista/index.html"
}

// Add
func (uc *ListaController) Add() {

	uc.Layout = "public/base.html"
	uc.TplName = "lista/add.html"
}

// Create
func (uc *ListaController) Create() {
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
func (uc *ListaController) Edit() {
	id, _ := uc.GetInt64("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Parâmetros com erro.", uc.Ctx)
	}

	var itemListaService services.ItemListaService

	item := itemListaService.GetListaById(id)
	if item == nil {
		response.ErrorWithMessage("Não encontrado informações para o Id.", uc.Ctx)
	}

	//

	//fmt.Println("lista_retornada:", lista)
	uc.Data["data"] = item

	uc.Layout = "public/base.html"
	uc.TplName = "lista/edit.html"
}

// Update
func (uc *ListaController) Update() {
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
func (uc *ListaController) Del() {
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

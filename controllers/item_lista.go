package controllers

import (
	"p-med/formvalidate"
	"p-med/global"
	"p-med/global/response"
	"p-med/models"
	"p-med/services"
	"p-med/utils"
	"strconv"
	"strings"

	"github.com/adam-hanna/arrayOperations"
	"github.com/gookit/validate"
)

// ListaController struct
type ItemListaController struct {
	baseController
}

// Add
func (uc *ItemListaController) Add() {
	listaNome := uc.GetString("nome")
	var listaService services.ListaService
	lista := new(models.Lista)

	//fmt.Println("listaNome: ", listaNome)

	lista = listaService.GetOneListaByNome(listaNome)

	//fmt.Println("lista: ", lista.Id)

	uc.Data["lista"] = lista

	uc.Layout = "public/base.html"
	uc.TplName = "lista/add.html"
}

// Create
func (uc *ItemListaController) Create() {
	var listaForm formvalidate.ListaForm

	//fmt.Println("AQUI")

	if err := uc.ParseForm(&listaForm); err != nil {
		response.ErrorWithMessage(err.Error(), uc.Ctx)
	}

	v := validate.Struct(listaForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), uc.Ctx)
	}

	var itemListaService services.ItemListaService
	insertID := itemListaService.Create(&listaForm)

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

	var itemListaService services.ItemListaService

	item := itemListaService.GetItemListaById(id)
	if item == nil {
		response.ErrorWithMessage("Não encontrado informações para o Id.", uc.Ctx)
	}

	//

	uc.Data["data"] = item

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

	var itemListaService services.ItemListaService
	num := itemListaService.Update(&listaForm)

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

	noDeletionID := new(models.ItemListaBd).NoDeletionId()

	m, b := arrayOperations.Intersect(noDeletionID, idArr)

	if len(noDeletionID) > 0 && len(m.Interface().([]int)) > 0 && b {
		response.ErrorWithMessage("ID é"+strings.Join(utils.IntArrToStringArr(noDeletionID), ",")+"os dados não podem ser excluídos!", uc.Ctx)
	}

	//fmt.Println("aqui", idArr)

	var itemListaService services.ItemListaService
	count := itemListaService.Del(idArr)

	//fmt.Println("aqui2", count)
	if count > 0 {
		response.SuccessWithMessageAndUrl("Operação bem-sucedida", global.URL_RELOAD, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

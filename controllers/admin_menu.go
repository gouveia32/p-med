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

// AdminMenuController struct.
type AdminMenuController struct {
	baseController
}

// Index 菜单首页.
func (amc *AdminMenuController) Index() {
	var adminTreeService services.AdminTreeService
	amc.Data["data"] = adminTreeService.AdminMenuTree()

	amc.Layout = "public/base.html"
	amc.TplName = "admin_menu/index.html"
}

// Add Novo菜单界面.
func (amc *AdminMenuController) Add() {

	var adminTreeService services.AdminTreeService
	parentID, _ := amc.GetInt("parent_id", 0)
	parents := adminTreeService.Menu(parentID, 0)

	amc.Data["parents"] = parents
	amc.Data["log_method"] = new(models.AdminMenu).GetLogMethod()

	amc.Layout = "public/base.html"
	amc.TplName = "admin_menu/add.html"
}

// Create Novo菜单.
func (amc *AdminMenuController) Create() {
	var adminMenuService services.AdminMenuService
	adminMenuForm := formvalidate.AdminMenuForm{}

	if err := amc.ParseForm(&adminMenuForm); err != nil {
		response.ErrorWithMessage(err.Error(), amc.Ctx)
	}

	//去除Url前后两侧的空格
	if adminMenuForm.Url != "" {
		adminMenuForm.Url = strings.TrimSpace(adminMenuForm.Url)
	}

	//数据校验
	v := validate.Struct(adminMenuForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), amc.Ctx)
	}

	//Novo Registro之前url验重
	if adminMenuService.IsExistUrl(adminMenuForm.Url, adminMenuForm.Id) {
		response.ErrorWithMessage("url["+adminMenuForm.Url+"] já existe.", amc.Ctx)
	}

	//创建
	_, err := adminMenuService.Create(&adminMenuForm)
	if err != nil {
		response.Error(amc.Ctx)
	}

	url := global.URL_BACK
	if adminMenuForm.IsCreate == 1 {
		url = global.URL_RELOAD
	}

	response.SuccessWithMessageAndUrl("Adicionado com sucesso", url, amc.Ctx)
}

// Edit 编辑菜单界面.
func (amc *AdminMenuController) Edit() {
	id, _ := amc.GetInt("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Param is error.", amc.Ctx)
	}

	var (
		adminMenuService services.AdminMenuService
		adminTreeService services.AdminTreeService
	)

	adminMenu := adminMenuService.GetAdminMenuById(id)
	if adminMenu == nil {
		response.ErrorWithMessage("Not Found Info By Id.", amc.Ctx)
	}

	parentID := adminMenu.ParentId
	parents := adminTreeService.Menu(parentID, 0)

	amc.Data["data"] = adminMenu
	amc.Data["parents"] = parents
	amc.Data["log_method"] = new(models.AdminMenu).GetLogMethod()

	amc.Layout = "public/base.html"
	amc.TplName = "admin_menu/edit.html"
}

// Update 菜单更新.
func (amc *AdminMenuController) Update() {
	var adminMenuService services.AdminMenuService
	adminMenuForm := formvalidate.AdminMenuForm{}

	if err := amc.ParseForm(&adminMenuForm); err != nil {
		response.ErrorWithMessage(err.Error(), amc.Ctx)
	}

	//去除Url前后两侧的空格
	if adminMenuForm.Url != "" {
		adminMenuForm.Url = strings.TrimSpace(adminMenuForm.Url)
	}

	//数据校验
	v := validate.Struct(adminMenuForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), amc.Ctx)
	}

	//Novo Registro之前url验重
	if adminMenuService.IsExistUrl(adminMenuForm.Url, adminMenuForm.Id) {
		response.ErrorWithMessage("url["+adminMenuForm.Url+"] Já existe.", amc.Ctx)
	}

	count := adminMenuService.Update(&adminMenuForm)

	if count > 0 {
		response.Success(amc.Ctx)
	} else {
		response.Error(amc.Ctx)
	}
}

// Del Exluir.
func (amc *AdminMenuController) Del() {
	idStr := amc.GetString("id")
	ids := make([]int, 0)
	var idArr []int

	if idStr == "" {
		amc.Ctx.Input.Bind(&ids, "id")
	} else {
		id, _ := strconv.Atoi(idStr)
		idArr = append(idArr, id)
	}

	if len(ids) > 0 {
		idArr = ids
	}

	if len(idArr) == 0 {
		response.ErrorWithMessage("Parâmetro errado.", amc.Ctx)
	}

	var adminMenuService services.AdminMenuService
	//判断SimNão有子菜单
	if adminMenuService.IsChildMenu(idArr) {
		response.ErrorWithMessage("Existem submenus que não podem ser excluídos!", amc.Ctx)
	}

	noDeletionID := new(models.AdminMenu).NoDeletionId()

	m, b := arrayOperations.Intersect(noDeletionID, idArr)

	if len(noDeletionID) > 0 && len(m.Interface().([]int)) > 0 && b {
		response.ErrorWithMessage("ID é "+strings.Join(utils.IntArrToStringArr(noDeletionID), ",")+"Não pode ser excluído!!", amc.Ctx)
	}

	count := adminMenuService.Del(idArr)

	if count > 0 {
		response.SuccessWithMessageAndUrl("Operação realizada com sucesso.", global.URL_RELOAD, amc.Ctx)
	} else {
		response.Error(amc.Ctx)
	}
}

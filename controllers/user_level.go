package controllers

import (
	"p-med/formvalidate"
	"p-med/global"
	"p-med/global/response"
	"p-med/models"
	"p-med/services"
	"p-med/utils"
	"p-med/utils/exceloffice"
	"p-med/utils/template"
	"strconv"
	"strings"
	"time"

	"github.com/adam-hanna/arrayOperations"
	"github.com/gookit/validate"
)

// UserLevelController struct
type UserLevelController struct {
	baseController
}

// Index 用户等级 列表页
func (ulc *UserLevelController) Index() {
	var userLevelService services.UserLevelService
	data, pagination := userLevelService.GetPaginateData(admin["per_page"].(int), gQueryParams)
	ulc.Data["data"] = data
	ulc.Data["paginate"] = pagination

	ulc.Layout = "public/base.html"
	ulc.TplName = "user_level/index.html"
}

// Export 导出
func (ulc *UserLevelController) Export() {
	exportData := ulc.GetString("export_data")
	if exportData == "1" {
		var userLevelService services.UserLevelService
		data := userLevelService.GetExportData(gQueryParams)
		header := []string{"ID", "Nome", "Descrição", "Ativo", "Criação"}
		body := [][]string{}
		for _, item := range data {
			record := []string{
				strconv.Itoa(item.Id),
				item.Name,
				item.Description,
			}
			if item.Status == 1 {
				record = append(record, "Sim")
			} else {
				record = append(record, "Não")
			}
			record = append(record, template.UnixTimeForFormat(item.CreateTime))
			body = append(body, record)
		}
		ulc.Ctx.ResponseWriter.Header().Set("a", "b")
		exceloffice.ExportData(header, body, "user_level-"+time.Now().Format("2006-01-02-15-04-05"), "", "", ulc.Ctx.ResponseWriter)
	}

	response.Error(ulc.Ctx)
}

// Add 用户等级-Novo Registro界面
func (ulc *UserLevelController) Add() {
	ulc.Layout = "public/base.html"
	ulc.TplName = "user_level/add.html"
}

// Create 用户等级-Novo Registro
func (ulc *UserLevelController) Create() {
	var userLevelForm formvalidate.UserLevelForm
	if err := ulc.ParseForm(&userLevelForm); err != nil {
		response.ErrorWithMessage(err.Error(), ulc.Ctx)
	}

	v := validate.Struct(userLevelForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), ulc.Ctx)
	}

	//处理图片上传
	_, _, err := ulc.GetFile("img")
	if err == nil {
		var attachmentService services.AttachmentService
		attachmentInfo, err := attachmentService.Upload(ulc.Ctx, "img", loginUser.Id, 0)
		if err != nil || attachmentInfo == nil {
			response.ErrorWithMessage(err.Error(), ulc.Ctx)
		} else {
			userLevelForm.Img = attachmentInfo.Url
		}
	}

	var userLevelService services.UserLevelService
	insertID := userLevelService.Create(&userLevelForm)

	url := global.URL_BACK
	if userLevelForm.IsCreate == 1 {
		url = global.URL_RELOAD
	}

	if insertID > 0 {
		response.SuccessWithMessageAndUrl("Adicionado com sucesso", url, ulc.Ctx)
	} else {
		response.Error(ulc.Ctx)
	}
}

// Edit 用户等级-修改界面
func (ulc *UserLevelController) Edit() {
	id, _ := ulc.GetInt("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Param is error.", ulc.Ctx)
	}

	var userLevelService services.UserLevelService

	userLevel := userLevelService.GetUserLevelById(id)
	if userLevel == nil {
		response.ErrorWithMessage("Not Found Info By Id.", ulc.Ctx)
	}

	ulc.Data["data"] = userLevel

	ulc.Layout = "public/base.html"
	ulc.TplName = "user_level/edit.html"
}

// Update 用户等级-修改
func (ulc *UserLevelController) Update() {
	var userLevelForm formvalidate.UserLevelForm
	if err := ulc.ParseForm(&userLevelForm); err != nil {
		response.ErrorWithMessage(err.Error(), ulc.Ctx)
	}

	if userLevelForm.Id <= 0 {
		response.ErrorWithMessage("Parâmetros com erro.", ulc.Ctx)
	}

	v := validate.Struct(userLevelForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), ulc.Ctx)
	}

	_, _, err := ulc.GetFile("img")
	if err == nil {
		//处理图片上传
		var attachmentService services.AttachmentService
		attachmentInfo, err := attachmentService.Upload(ulc.Ctx, "img", loginUser.Id, 0)
		if err != nil || attachmentInfo == nil {
			response.ErrorWithMessage(err.Error(), ulc.Ctx)
		} else {
			userLevelForm.Img = attachmentInfo.Url
		}
	}

	var userLevelService services.UserLevelService
	num := userLevelService.Update(&userLevelForm)

	if num > 0 {
		response.Success(ulc.Ctx)
	} else {
		response.Error(ulc.Ctx)
	}
}

// Enable Ativar
func (ulc *UserLevelController) Enable() {
	idStr := ulc.GetString("id")
	ids := make([]int, 0)
	var idArr []int

	if idStr == "" {
		ulc.Ctx.Input.Bind(&ids, "id")
	} else {
		id, _ := strconv.Atoi(idStr)
		idArr = append(idArr, id)
	}

	if len(ids) > 0 {
		idArr = ids
	}

	if len(idArr) == 0 {
		response.ErrorWithMessage("请选择用户等级.", ulc.Ctx)
	}

	var userLevelService services.UserLevelService
	num := userLevelService.Enable(idArr)
	if num > 0 {
		response.SuccessWithMessageAndUrl("Operação realizada com sucesso.", global.URL_RELOAD, ulc.Ctx)
	} else {
		response.Error(ulc.Ctx)
	}
}

// Disable Desativar
func (ulc *UserLevelController) Disable() {
	idStr := ulc.GetString("id")
	ids := make([]int, 0)
	var idArr []int

	if idStr == "" {
		ulc.Ctx.Input.Bind(&ids, "id")
	} else {
		id, _ := strconv.Atoi(idStr)
		idArr = append(idArr, id)
	}

	if len(ids) > 0 {
		idArr = ids
	}

	if len(idArr) == 0 {
		response.ErrorWithMessage("请选择Desativar的用户.", ulc.Ctx)
	}

	var userLevelService services.UserLevelService
	num := userLevelService.Disable(idArr)
	if num > 0 {
		response.SuccessWithMessageAndUrl("Operação realizada com sucesso.", global.URL_RELOAD, ulc.Ctx)
	} else {
		response.Error(ulc.Ctx)
	}
}

// Del Exluir
func (ulc *UserLevelController) Del() {
	idStr := ulc.GetString("id")
	ids := make([]int, 0)
	var idArr []int

	if idStr == "" {
		ulc.Ctx.Input.Bind(&ids, "id")
	} else {
		id, _ := strconv.Atoi(idStr)
		idArr = append(idArr, id)
	}

	if len(ids) > 0 {
		idArr = ids
	}

	if len(idArr) == 0 {
		response.ErrorWithMessage("Parâmetro errado.", ulc.Ctx)
	}

	noDeletionID := new(models.UserLevel).NoDeletionId()

	m, b := arrayOperations.Intersect(noDeletionID, idArr)

	if len(noDeletionID) > 0 && len(m.Interface().([]int)) > 0 && b {
		response.ErrorWithMessage("ID为"+strings.Join(utils.IntArrToStringArr(noDeletionID), ",")+"的数据无法删除!", ulc.Ctx)
	}

	var userLevelService services.UserLevelService
	count := userLevelService.Del(idArr)

	if count > 0 {
		response.SuccessWithMessageAndUrl("Operação realizada com sucesso.", global.URL_RELOAD, ulc.Ctx)
	} else {
		response.Error(ulc.Ctx)
	}
}

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

// UserController struct
type UserController struct {
	baseController
}

// Index 用户等级 列表页
func (uc *UserController) Index() {
	var userService services.UserService
	var userLevelService services.UserLevelService

	//获取用户等级
	userLevel := userLevelService.GetUserLevel()
	userLevelMap := make(map[int]string)
	for _, item := range userLevel {
		userLevelMap[item.Id] = item.Name
	}

	data, pagination := userService.GetPaginateData(admin["per_page"].(int), gQueryParams)

	uc.Data["data"] = data
	uc.Data["paginate"] = pagination
	uc.Data["user_level_map"] = userLevelMap

	uc.Layout = "public/base.html"
	uc.TplName = "user/index.html"
}

// Export 导出
func (uc *UserController) Export() {
	exportData := uc.GetString("export_data")
	if exportData == "1" {
		var userService services.UserService
		var userLevelService services.UserLevelService
		userLevel := userLevelService.GetUserLevel()
		userLevelMap := make(map[int]string)
		for _, item := range userLevel {
			userLevelMap[item.Id] = item.Name
		}

		data := userService.GetExportData(gQueryParams)
		header := []string{"ID", "头像", "用户等级", "用户名", "手机号", "昵称", "Ativo", "Criação"}
		body := [][]string{}
		for _, item := range data {
			record := []string{
				strconv.Itoa(item.Id),
				item.Avatar,
			}
			userLevelName, ok := userLevelMap[item.UserLevelId]
			if ok {
				record = append(record, userLevelName)
			}
			record = append(record, item.Username)
			record = append(record, item.Mobile)
			record = append(record, item.Nickname)

			if item.Status == 1 {
				record = append(record, "Sim")
			} else {
				record = append(record, "Não")
			}
			record = append(record, template.UnixTimeForFormat(item.CreateTime))
			body = append(body, record)
		}
		uc.Ctx.ResponseWriter.Header().Set("a", "b")
		exceloffice.ExportData(header, body, "user-"+time.Now().Format("2006-01-02-15-04-05"), "", "", uc.Ctx.ResponseWriter)
	}

	response.Error(uc.Ctx)
}

// Add 用户-Novo Registro界面
func (uc *UserController) Add() {
	var userLevelService services.UserLevelService

	//获取用户等级
	userLevel := userLevelService.GetUserLevel()

	uc.Data["user_level_list"] = userLevel
	uc.Layout = "public/base.html"
	uc.TplName = "user/add.html"
}

// Create Novo用户
func (uc *UserController) Create() {
	var userForm formvalidate.UserForm
	if err := uc.ParseForm(&userForm); err != nil {
		response.ErrorWithMessage(err.Error(), uc.Ctx)
	}

	v := validate.Struct(userForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), uc.Ctx)
	}

	//处理图片上传
	_, _, err := uc.GetFile("avatar")
	if err == nil {
		var attachmentService services.AttachmentService
		attachmentInfo, err := attachmentService.Upload(uc.Ctx, "avatar", loginUser.Id, 0)
		if err != nil || attachmentInfo == nil {
			response.ErrorWithMessage(err.Error(), uc.Ctx)
		} else {
			userForm.Avatar = attachmentInfo.Url
		}
	}

	var userService services.UserService
	insertID := userService.Create(&userForm)

	url := global.URL_BACK
	if userForm.IsCreate == 1 {
		url = global.URL_RELOAD
	}

	if insertID > 0 {
		response.SuccessWithMessageAndUrl("Adicionado com sucesso", url, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Edit 用户-修改界面
func (uc *UserController) Edit() {
	id, _ := uc.GetInt("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Param is error.", uc.Ctx)
	}

	var userService services.UserService

	user := userService.GetUserById(id)
	if user == nil {
		response.ErrorWithMessage("Not Found Info By Id.", uc.Ctx)
	}

	//获取用户等级
	var userLevelService services.UserLevelService
	userLevel := userLevelService.GetUserLevel()

	uc.Data["user_level_list"] = userLevel
	uc.Data["data"] = user

	uc.Layout = "public/base.html"
	uc.TplName = "user/edit.html"
}

// Update 用户-修改
func (uc *UserController) Update() {
	var userForm formvalidate.UserForm
	if err := uc.ParseForm(&userForm); err != nil {
		response.ErrorWithMessage(err.Error(), uc.Ctx)
	}

	if userForm.Id <= 0 {
		response.ErrorWithMessage("Parâmetros com erro.", uc.Ctx)
	}

	v := validate.Struct(userForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), uc.Ctx)
	}

	_, _, err := uc.GetFile("avatar")
	if err == nil {
		//处理图片上传
		var attachmentService services.AttachmentService
		attachmentInfo, err := attachmentService.Upload(uc.Ctx, "avatar", loginUser.Id, 0)
		if err != nil || attachmentInfo == nil {
			response.ErrorWithMessage(err.Error(), uc.Ctx)
		} else {
			userForm.Avatar = attachmentInfo.Url
		}
	}

	var userService services.UserService
	num := userService.Update(&userForm)

	if num > 0 {
		response.Success(uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Enable Ativar
func (uc *UserController) Enable() {
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
		response.ErrorWithMessage("请选择用户等级.", uc.Ctx)
	}

	var userService services.UserService
	num := userService.Enable(idArr)
	if num > 0 {
		response.SuccessWithMessageAndUrl("Operação realizada com sucesso.", global.URL_RELOAD, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Disable Desativar
func (uc *UserController) Disable() {
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
		response.ErrorWithMessage("请选择Desativar的用户.", uc.Ctx)
	}

	var userService services.UserService
	num := userService.Disable(idArr)
	if num > 0 {
		response.SuccessWithMessageAndUrl("Operação realizada com sucesso.", global.URL_RELOAD, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Del Exluir
func (uc *UserController) Del() {
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
		response.ErrorWithMessage("Parâmetro errado.", uc.Ctx)
	}

	noDeletionID := new(models.User).NoDeletionId()

	m, b := arrayOperations.Intersect(noDeletionID, idArr)

	if len(noDeletionID) > 0 && len(m.Interface().([]int)) > 0 && b {
		response.ErrorWithMessage("ID为"+strings.Join(utils.IntArrToStringArr(noDeletionID), ",")+"的数据无法删除!", uc.Ctx)
	}

	var userService services.UserService
	count := userService.Del(idArr)

	if count > 0 {
		response.SuccessWithMessageAndUrl("Operação realizada com sucesso.", global.URL_RELOAD, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

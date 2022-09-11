package controllers

import (
	"encoding/base64"
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

// AdminUserController struct
type AdminUserController struct {
	baseController
}

// Index 用户管理-首页
func (auc *AdminUserController) Index() {
	var adminUserService services.AdminUserService
	data, pagination := adminUserService.GetPaginateData(admin["per_page"].(int), gQueryParams)
	auc.Data["data"] = data
	auc.Data["paginate"] = pagination

	auc.Layout = "public/base.html"
	auc.TplName = "admin_user/index.html"
}

// Add 用户管理-Novo Registro界面
func (auc *AdminUserController) Add() {
	var adminRoleService services.AdminRoleService
	roles := adminRoleService.GetAllData()

	auc.Data["roles"] = roles
	auc.Layout = "public/base.html"
	auc.TplName = "admin_user/add.html"
}

// Create 用户管理-Novo Registro界面
func (auc *AdminUserController) Create() {
	var adminUserForm formvalidate.AdminUserForm
	if err := auc.ParseForm(&adminUserForm); err != nil {
		response.ErrorWithMessage(err.Error(), auc.Ctx)
	}
	roles := make([]string, 0)
	auc.Ctx.Input.Bind(&roles, "role")

	adminUserForm.Role = strings.Join(roles, ",")

	v := validate.Struct(adminUserForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), auc.Ctx)
	}

	//账号验重
	var adminUserService services.AdminUserService
	if adminUserService.IsExistName(strings.TrimSpace(adminUserForm.Username), 0) {
		response.ErrorWithMessage("账号已经存在", auc.Ctx)
	}
	//默认头像
	adminUserForm.Avatar = "/static/admin/images/avatar.png"

	insertID := adminUserService.Create(&adminUserForm)

	url := global.URL_BACK
	if adminUserForm.IsCreate == 1 {
		url = global.URL_RELOAD
	}

	if insertID > 0 {
		response.SuccessWithMessageAndUrl("Adicionado com sucesso", url, auc.Ctx)
	} else {
		response.Error(auc.Ctx)
	}
}

// Edit 系统管理-用户管理-修改界面
func (auc *AdminUserController) Edit() {
	id, _ := auc.GetInt("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Param is error.", auc.Ctx)
	}

	var (
		adminUserService services.AdminUserService
		adminRoleService services.AdminRoleService
	)

	adminUser := adminUserService.GetAdminUserById(id)
	if adminUser == nil {
		response.ErrorWithMessage("Not Found Info By Id.", auc.Ctx)
	}

	roles := adminRoleService.GetAllData()
	auc.Data["roles"] = roles
	auc.Data["data"] = adminUser
	auc.Data["role_arr"] = strings.Split(adminUser.Role, ",")

	auc.Layout = "public/base.html"
	auc.TplName = "admin_user/edit.html"
}

// Update 系统管理-用户管理-修改
func (auc *AdminUserController) Update() {
	var adminUserForm formvalidate.AdminUserForm
	if err := auc.ParseForm(&adminUserForm); err != nil {
		response.ErrorWithMessage(err.Error(), auc.Ctx)
	}

	if adminUserForm.Id <= 0 {
		response.ErrorWithMessage("Parâmetros com erro.", auc.Ctx)
	}

	roles := make([]string, 0)
	auc.Ctx.Input.Bind(&roles, "role")

	adminUserForm.Role = strings.Join(roles, ",")

	v := validate.Struct(adminUserForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), auc.Ctx)
	}

	//账号验重
	var adminUserService services.AdminUserService
	if adminUserService.IsExistName(strings.TrimSpace(adminUserForm.Username), adminUserForm.Id) {
		response.ErrorWithMessage("账号已经存在", auc.Ctx)
	}

	num := adminUserService.Update(&adminUserForm)

	if num > 0 {
		response.Success(auc.Ctx)
	} else {
		response.Error(auc.Ctx)
	}
}

// Enable Ativar
func (auc *AdminUserController) Enable() {
	idStr := auc.GetString("id")
	ids := make([]int, 0)
	var idArr []int

	if idStr == "" {
		auc.Ctx.Input.Bind(&ids, "id")
	} else {
		id, _ := strconv.Atoi(idStr)
		idArr = append(idArr, id)
	}

	if len(ids) > 0 {
		idArr = ids
	}

	if len(idArr) == 0 {
		response.ErrorWithMessage("请选择Ativar的用户.", auc.Ctx)
	}

	var adminUserService services.AdminUserService
	num := adminUserService.Enable(idArr)
	if num > 0 {
		response.SuccessWithMessageAndUrl("Operação realizada com sucesso.", global.URL_RELOAD, auc.Ctx)
	} else {
		response.Error(auc.Ctx)
	}
}

// Disable Desativar
func (auc *AdminUserController) Disable() {
	idStr := auc.GetString("id")
	ids := make([]int, 0)
	var idArr []int

	if idStr == "" {
		auc.Ctx.Input.Bind(&ids, "id")
	} else {
		id, _ := strconv.Atoi(idStr)
		idArr = append(idArr, id)
	}

	if len(ids) > 0 {
		idArr = ids
	}

	if len(idArr) == 0 {
		response.ErrorWithMessage("请选择Desativar的用户.", auc.Ctx)
	}

	var adminUserService services.AdminUserService
	num := adminUserService.Disable(idArr)
	if num > 0 {
		response.SuccessWithMessageAndUrl("Operação realizada com sucesso.", global.URL_RELOAD, auc.Ctx)
	} else {
		response.Error(auc.Ctx)
	}
}

// Del Exluir
func (auc *AdminUserController) Del() {
	idStr := auc.GetString("id")
	ids := make([]int, 0)
	var idArr []int

	if idStr == "" {
		auc.Ctx.Input.Bind(&ids, "id")
	} else {
		id, _ := strconv.Atoi(idStr)
		idArr = append(idArr, id)
	}

	if len(ids) > 0 {
		idArr = ids
	}

	if len(idArr) == 0 {
		response.ErrorWithMessage("Parâmetro errado.", auc.Ctx)
	}

	noDeletionID := new(models.AdminUser).NoDeletionId()

	m, b := arrayOperations.Intersect(noDeletionID, idArr)

	if len(noDeletionID) > 0 && len(m.Interface().([]int)) > 0 && b {
		response.ErrorWithMessage("ID为"+strings.Join(utils.IntArrToStringArr(noDeletionID), ",")+"的数据无法删除!", auc.Ctx)
	}

	var adminUserService services.AdminUserService
	count := adminUserService.Del(idArr)

	if count > 0 {
		response.SuccessWithMessageAndUrl("Operação realizada com sucesso.", global.URL_RELOAD, auc.Ctx)
	} else {
		response.Error(auc.Ctx)
	}
}

// Profile 系统管理-Perfil
func (auc *AdminUserController) Profile() {
	auc.Layout = "public/base.html"
	auc.TplName = "admin_user/profile.html"
}

// UpdateNickName 系统管理-Perfil-修改昵称
func (auc *AdminUserController) UpdateNickName() {
	id, err := auc.GetInt("id")
	nickname := strings.TrimSpace(auc.GetString("nickname"))

	if nickname == "" || err != nil {
		response.ErrorWithMessage("参数错误", auc.Ctx)
	}

	// 验证SimNãoSim登陆用户，这里也可不用提供的id，使用登陆的id即可
	if loginUser.Id != id {
		response.ErrorWithMessage("数据非法", auc.Ctx)
	}

	var adminUserService services.AdminUserService
	num := adminUserService.UpdateNickName(id, nickname)

	if num > 0 {
		//修改成功后，更新session的Conecte-se用户信息
		loginAdminUser := adminUserService.GetAdminUserById(id)
		auc.SetSession(global.LOGIN_USER, *loginAdminUser)
		response.SuccessWithMessageAndUrl("修改成功", global.URL_RELOAD, auc.Ctx)
	} else {
		response.Error(auc.Ctx)
	}
}

// UpdatePassword 系统管理-Perfil-修改密码
func (auc *AdminUserController) UpdatePassword() {
	id, err := auc.GetInt("id")
	password := auc.GetString("password")
	newPassword := auc.GetString("new_password")
	reNewPassword := auc.GetString("renew_password")

	// 验证SimNãoSim登陆用户，这里也可不用提供的id，使用登陆的id即可
	if loginUser.Id != id {
		response.ErrorWithMessage("数据非法", auc.Ctx)
	}

	if err != nil || password == "" || newPassword == "" || reNewPassword == "" {
		response.ErrorWithMessage("Bad Parameter.", auc.Ctx)
	}

	if newPassword != reNewPassword {
		response.ErrorWithMessage("两次输入的密码不一致.", auc.Ctx)
	}

	if password == newPassword {
		response.ErrorWithMessage("新密码与旧密码一致，无需修改", auc.Ctx)
	}

	loginUserPassword, err := base64.StdEncoding.DecodeString(loginUser.Password)

	if err != nil {
		response.ErrorWithMessage("err:"+err.Error(), auc.Ctx)
	}

	if !utils.PasswordVerify(password, string(loginUserPassword)) {
		response.ErrorWithMessage("当前密码不正确", auc.Ctx)
	}

	var adminUserService services.AdminUserService
	num := adminUserService.UpdatePassword(id, newPassword)
	if num > 0 {
		response.SuccessWithMessageAndUrl("修改成功", global.URL_RELOAD, auc.Ctx)
	} else {
		response.Error(auc.Ctx)
	}
}

// UpdateAvatar 系统管理-Perfil-修改头像
func (auc *AdminUserController) UpdateAvatar() {
	_, _, err := auc.GetFile("avatar")
	if err != nil {
		response.ErrorWithMessage("上传头像错误"+err.Error(), auc.Ctx)
	}

	var (
		attachmentService services.AttachmentService
		adminUserService  services.AdminUserService
	)
	attachmentInfo, err := attachmentService.Upload(auc.Ctx, "avatar", loginUser.Id, 0)
	if err != nil || attachmentInfo == nil {
		response.ErrorWithMessage(err.Error(), auc.Ctx)
	} else {
		//头像上传成功，更新用户的avatar头像信息
		num := adminUserService.UpdateAvatar(loginUser.Id, attachmentInfo.Url)
		if num > 0 {
			//修改成功后，更新session的Conecte-se用户信息
			loginAdminUser := adminUserService.GetAdminUserById(loginUser.Id)
			auc.SetSession(global.LOGIN_USER, *loginAdminUser)
			response.SuccessWithMessageAndUrl("修改成功", global.URL_RELOAD, auc.Ctx)
		} else {
			response.Error(auc.Ctx)
		}
	}

}

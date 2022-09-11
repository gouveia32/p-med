package controllers

import (
	"net/http"
	"os"
	"p-med/formvalidate"
	"p-med/global"
	"p-med/global/response"
	"p-med/services"
	"p-med/utils"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/adapter/validation"
	"github.com/dchest/captcha"
	"github.com/gookit/validate"
)

var adminLogService services.AdminLogService

// AuthController struct
type AuthController struct {
	baseController
}

// Login Conecte-se界面
func (ac *AuthController) Login() {
	//加载Conecte-se配置信息
	var settingService services.SettingService
	data := settingService.Show(1)
	for _, setting := range data {
		settingService.LoadOrUpdateGlobalBaseConfig(setting)
	}

	//获取Conecte-se配置信息
	loginConfig := struct {
		Token      string
		Captcha    string
		Background string
	}{
		Token:   global.BA_CONFIG.Login.Token,
		Captcha: global.BA_CONFIG.Login.Captcha,
	}
	//Conecte-se背景图片
	if _, err := os.Stat(strings.TrimLeft(global.BA_CONFIG.Login.Background, "/")); err != nil {
		global.BA_CONFIG.Login.Background = "/static/admin/images/login-default-bg.jpg"
	}
	loginConfig.Background = global.BA_CONFIG.Login.Background

	//login界面只需要name字段
	admin := map[string]interface{}{
		"name":  global.BA_CONFIG.Base.Name,
		"title": "Conecte-se",
	}

	ac.Data["login_config"] = loginConfig
	//Login
	ac.Data["captcha"] = utils.GetCaptcha()
	ac.Data["admin"] = admin

	ac.TplName = "auth/login.html"
}

// Logout 退出Conecte-se
func (ac *AuthController) Logout() {
	ac.DelSession(global.LOGIN_USER)
	ac.Ctx.SetCookie(global.LOGIN_USER_ID, "", -1)
	ac.Ctx.SetCookie(global.LOGIN_USER_ID_SIGN, "", -1)
	ac.Redirect("/admin/auth/login", http.StatusFound)
}

// CheckLogin Conecte-se认证
func (ac *AuthController) CheckLogin() {
	//数据校验
	valid := validation.Validation{}
	loginForm := formvalidate.LoginForm{}

	if err := ac.ParseForm(&loginForm); err != nil {
		response.ErrorWithMessage(err.Error(), ac.Ctx)
	}

	v := validate.Struct(loginForm)

	//看SimNão需要校验验证码
	isCaptcha, _ := strconv.Atoi(global.BA_CONFIG.Login.Captcha)
	if isCaptcha > 0 {
		valid.Required(loginForm.Captcha, "captcha").Message("por favor insira o código de verificação.")
		if ok := captcha.VerifyString(loginForm.CaptchaId, loginForm.Captcha); !ok {
			response.ErrorWithMessage("Erro no código de verificação", ac.Ctx)
		}
	}

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), ac.Ctx)
	}

	//基础验证通过后，进行用户验证
	var adminUserService services.AdminUserService
	loginUser, err := adminUserService.CheckLogin(loginForm, ac.Ctx)
	if err != nil {
		response.ErrorWithMessage(err.Error(), ac.Ctx)
	}

	//Conecte-se日志记录
	adminLogService.LoginLog(loginUser.Id, ac.Ctx)

	redirect, _ := ac.GetSession("redirect").(string)
	if redirect != "" {
		response.SuccessWithMessageAndUrl("Logado com sucesso", redirect, ac.Ctx)
	} else {
		response.SuccessWithMessageAndUrl("Logado com scesso", "/admin/index/index", ac.Ctx)
	}
}

// RefreshCaptcha Atualizar验证码
func (ac *AuthController) RefreshCaptcha() {
	captchaID := ac.GetString("captchaId")
	res := map[string]interface{}{
		"isNew": false,
	}
	if captchaID == "" {
		res["msg"] = "Erro de parâmetro."
	}

	isReload := captcha.Reload(captchaID)
	if isReload {
		res["captchaId"] = captchaID
	} else {
		res["isNew"] = true
		res["captcha"] = utils.GetCaptcha()
	}

	ac.Data["json"] = res

	ac.ServeJSON()
}

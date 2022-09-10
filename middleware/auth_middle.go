package middleware

import (
	"fmt"
	"p-med/global"
	"p-med/global/response"
	"p-med/models"
	"p-med/services"
	"strconv"
	"strings"

	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/adapter/context"
	context2 "github.com/beego/beego/v2/server/web/context"
)

// AuthMiddle 中间件
func AuthMiddle() {

	//不需要验证的url
	authExcept := map[string]interface{}{
		"admin/auth/login":           0,
		"admin/auth/check_login":     1,
		"admin/auth/logout":          2,
		"admin/auth/captcha":         3,
		"admin/editor/server":        4,
		"admin/auth/refresh_captcha": 5,
	}

	//Conecte-se认证中间件过滤器
	var filterLogin = func(ctx *context.Context) {
		url := strings.TrimLeft(ctx.Input.URL(), "/")

		//需要进行Conecte-se验证
		if !isAuthExceptUrl(strings.ToLower(url), authExcept) {
			//验证SimNãoConecte-se
			loginUser, isLogin := isLogin(ctx)
			if !isLogin {
				response.ErrorWithMessageAndUrl("Não logado", "/admin/auth/login", (*context2.Context)(ctx))
				return
			}

			//验证，SimNão有权限访问
			var adminUserService services.AdminUserService
			if loginUser.Id != 1 && !adminUserService.AuthCheck(url, authExcept, loginUser) {
				errorBackURL := global.URL_CURRENT
				if ctx.Request.Method == "GET" {
					errorBackURL = ""
				}
				response.ErrorWithMessageAndUrl("Sem permissão", errorBackURL, (*context2.Context)(ctx))
				return
			}
		}

		checkAuth, _ := strconv.Atoi(ctx.Request.PostForm.Get("check_auth"))

		if checkAuth == 1 {
			response.Success((*context2.Context)(ctx))
			return
		}

	}

	beego.InsertFilter("/admin/*", beego.BeforeRouter, filterLogin)
}

//判断SimNãoSim不需要验证Conecte-se的url,只针对admin模块路由的判断
func isAuthExceptUrl(url string, m map[string]interface{}) bool {
	urlArr := strings.Split(url, "/")
	if len(urlArr) > 3 {
		url = fmt.Sprintf("%s/%s/%s", urlArr[0], urlArr[1], urlArr[2])
	}
	_, ok := m[url]
	if ok {
		return true
	}
	return false
}

//SimNãoConecte-se
func isLogin(ctx *context.Context) (*models.AdminUser, bool) {
	loginUser, ok := ctx.Input.Session(global.LOGIN_USER).(models.AdminUser)
	if !ok {
		loginUserIDStr := ctx.GetCookie(global.LOGIN_USER_ID)
		loginUserIDSign := ctx.GetCookie(global.LOGIN_USER_ID_SIGN)

		if loginUserIDStr != "" && loginUserIDSign != "" {
			loginUserID, _ := strconv.Atoi(loginUserIDStr)
			var adminUserService services.AdminUserService
			loginUserPointer := adminUserService.GetAdminUserById(loginUserID)

			if loginUserPointer != nil && loginUserPointer.GetSignStrByAdminUser((*context2.Context)(ctx)) == loginUserIDSign {
				ctx.Output.Session(global.LOGIN_USER, *loginUserPointer)
				return loginUserPointer, true
			}
		}
		return nil, false
	}

	return &loginUser, true
}

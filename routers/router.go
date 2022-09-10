package routers

import (
	"net/http"
	"p-med/controllers"
	"p-med/middleware"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/dchest/captcha"
)

func init() {
	//授权Conecte-se中间件
	middleware.AuthMiddle()

	web.Get("/", func(ctx *context.Context) {
		ctx.Redirect(http.StatusFound, "/admin/index/index")
	})

	//admin模块路由
	admin := web.NewNamespace("/admin",
		//UEditor控制器
		web.NSRouter("/editor/server", &controllers.EditorController{}, "get,post:Server"),

		//Conecte-se页
		web.NSRouter("/auth/login", &controllers.AuthController{}, "get:Login"),
		//退出Conecte-se
		web.NSRouter("/auth/logout", &controllers.AuthController{}, "get:Logout"),
		//二维码图片输出
		web.NSHandler("/auth/captcha/*.png", captcha.Server(240, 80)),
		//Conecte-se认证
		web.NSRouter("/auth/check_login", &controllers.AuthController{}, "post:CheckLogin"),
		//atualizar验证码
		web.NSRouter("/auth/refresh_captcha", &controllers.AuthController{}, "post:RefreshCaptcha"),

		//首页
		web.NSRouter("/index/index", &controllers.IndexController{}, "get:Index"),

		//用户管理
		web.NSRouter("/admin_user/index", &controllers.AdminUserController{}, "get:Index"),

		//操作日志
		web.NSRouter("/admin_log/index", &controllers.AdminLogController{}, "get:Index"),

		//菜单管理
		web.NSRouter("/admin_menu/index", &controllers.AdminMenuController{}, "get:Index"),
		//菜单管理-Novo Registro菜单-界面
		web.NSRouter("/admin_menu/add", &controllers.AdminMenuController{}, "get:Add"),
		//菜单管理-Novo Registro菜单-创建
		web.NSRouter("/admin_menu/create", &controllers.AdminMenuController{}, "post:Create"),
		//菜单管理-修改菜单-界面
		web.NSRouter("/admin_menu/edit", &controllers.AdminMenuController{}, "get:Edit"),
		//菜单管理-更新菜单
		web.NSRouter("/admin_menu/update", &controllers.AdminMenuController{}, "post:Update"),
		//菜单管理-Excluir Menu
		web.NSRouter("/admin_menu/del", &controllers.AdminMenuController{}, "post:Del"),

		//系统管理-Perfil
		web.NSRouter("/admin_user/profile", &controllers.AdminUserController{}, "get:Profile"),
		//系统管理-Perfil-修改昵称
		web.NSRouter("/admin_user/update_nickname", &controllers.AdminUserController{}, "post:UpdateNickName"),
		//系统管理-Perfil-修改密码
		web.NSRouter("/admin_user/update_password", &controllers.AdminUserController{}, "post:UpdatePassword"),
		//系统管理-Perfil-修改头像
		web.NSRouter("/admin_user/update_avatar", &controllers.AdminUserController{}, "post:UpdateAvatar"),
		//系统管理-用户管理-Novo Registro界面
		web.NSRouter("/admin_user/add", &controllers.AdminUserController{}, "get:Add"),
		//系统管理-用户管理-Novo Registro
		web.NSRouter("/admin_user/create", &controllers.AdminUserController{}, "post:Create"),
		//系统管理-用户管理-修改界面
		web.NSRouter("/admin_user/edit", &controllers.AdminUserController{}, "get:Edit"),
		//系统管理-用户管理-修改
		web.NSRouter("/admin_user/update", &controllers.AdminUserController{}, "post:Update"),
		//系统管理-用户管理-Ativar
		web.NSRouter("/admin_user/enable", &controllers.AdminUserController{}, "post:Enable"),
		//系统管理-用户管理-Desativar
		web.NSRouter("/admin_user/disable", &controllers.AdminUserController{}, "post:Disable"),
		//系统管理-用户管理-删除
		web.NSRouter("/admin_user/del", &controllers.AdminUserController{}, "post:Del"),

		//系统管理-角色管理
		web.NSRouter("/admin_role/index", &controllers.AdminRoleController{}, "get:Index"),
		//系统管理-角色管理-Novo Registro界面
		web.NSRouter("/admin_role/add", &controllers.AdminRoleController{}, "get:Add"),
		//系统管理-角色管理-Novo Registro
		web.NSRouter("/admin_role/create", &controllers.AdminRoleController{}, "post:Create"),
		//菜单管理-角色管理-修改界面
		web.NSRouter("/admin_role/edit", &controllers.AdminRoleController{}, "get:Edit"),
		//菜单管理-角色管理-修改
		web.NSRouter("/admin_role/update", &controllers.AdminRoleController{}, "post:Update"),
		//菜单管理-角色管理-删除
		web.NSRouter("/admin_role/del", &controllers.AdminRoleController{}, "post:Del"),
		//菜单管理-角色管理-Ativar角色
		web.NSRouter("/admin_role/enable", &controllers.AdminRoleController{}, "post:Enable"),
		//菜单管理-角色管理-Desativar角色
		web.NSRouter("/admin_role/disable", &controllers.AdminRoleController{}, "post:Disable"),
		//菜单管理-角色管理-角色授权界面
		web.NSRouter("/admin_role/access", &controllers.AdminRoleController{}, "get:Access"),
		//菜单管理-角色管理-角色授权
		web.NSRouter("/admin_role/access_operate", &controllers.AdminRoleController{}, "post:AccessOperate"),

		//设置中心-后台设置
		web.NSRouter("/setting/admin", &controllers.SettingController{}, "get:Admin"),
		//设置中心-更新设置
		web.NSRouter("/setting/update", &controllers.SettingController{}, "post:Update"),

		//系统管理-开发管理-数据维护
		web.NSRouter("/database/table", &controllers.DatabaseController{}, "get:Table"),
		//系统管理-开发管理-数据维护-优化表
		web.NSRouter("/database/optimize", &controllers.DatabaseController{}, "post:Optimize"),
		//系统管理-开发管理-数据维护-修复表
		web.NSRouter("/database/repair", &controllers.DatabaseController{}, "post:Repair"),
		//系统管理-开发管理-数据维护-查看详情
		web.NSRouter("/database/view", &controllers.DatabaseController{}, "get,post:View"),

		//用户等级管理
		web.NSRouter("/user_level/index", &controllers.UserLevelController{}, "get:Index"),
		//用户等级管理-Novo Registro界面
		web.NSRouter("/user_level/add", &controllers.UserLevelController{}, "get:Add"),
		//用户等级管理-Novo Registro
		web.NSRouter("/user_level/create", &controllers.UserLevelController{}, "post:Create"),
		//用户等级管理-修改界面
		web.NSRouter("/user_level/edit", &controllers.UserLevelController{}, "get:Edit"),
		//用户等级管理-修改
		web.NSRouter("/user_level/update", &controllers.UserLevelController{}, "post:Update"),
		//用户等级管理-Ativar
		web.NSRouter("/user_level/enable", &controllers.UserLevelController{}, "post:Enable"),
		//用户等级管理-Desativar
		web.NSRouter("/user_level/disable", &controllers.UserLevelController{}, "post:Disable"),
		//用户等级管理-删除
		web.NSRouter("/user_level/del", &controllers.UserLevelController{}, "post:Del"),
		//用户等级管理-导出
		web.NSRouter("/user_level/export", &controllers.UserLevelController{}, "get:Export"),

		//用户管理
		web.NSRouter("/user/index", &controllers.UserController{}, "get:Index"),
		//用户管理-Novo Registro界面
		web.NSRouter("/user/add", &controllers.UserController{}, "get:Add"),
		//用户管理-Novo Registro
		web.NSRouter("/user/create", &controllers.UserController{}, "post:Create"),
		//用户管理-修改界面
		web.NSRouter("/user/edit", &controllers.UserController{}, "get:Edit"),
		//用户管理-修改
		web.NSRouter("/user/update", &controllers.UserController{}, "post:Update"),
		//用户管理-Ativar
		web.NSRouter("/user/enable", &controllers.UserController{}, "post:Enable"),
		//用户管理-Desativar
		web.NSRouter("/user/disable", &controllers.UserController{}, "post:Disable"),
		//用户管理-删除
		web.NSRouter("/user/del", &controllers.UserController{}, "post:Del"),
		//用户管理-导出
		web.NSRouter("/user/export", &controllers.UserController{}, "get:Export"),

		//Ger pacientes
		web.NSRouter("/paciente/index", &controllers.PacienteController{}, "get:Index"),
		web.NSRouter("/paciente/add", &controllers.PacienteController{}, "get:Add"),
		web.NSRouter("/paciente/create", &controllers.PacienteController{}, "post:Create"),
		web.NSRouter("/paciente/edit", &controllers.PacienteController{}, "get:Edit"),
		web.NSRouter("/paciente/update", &controllers.PacienteController{}, "post:Update"),
		web.NSRouter("/paciente/enable", &controllers.PacienteController{}, "post:Enable"),
		web.NSRouter("/paciente/disable", &controllers.PacienteController{}, "post:Disable"),
		web.NSRouter("/paciente/del", &controllers.PacienteController{}, "post:Del"),
		web.NSRouter("/paciente/export", &controllers.PacienteController{}, "get:Export"),

		//Ger atendimento
		web.NSRouter("/atendimento/index", &controllers.AtendimentoController{}, "get:Index"),
		web.NSRouter("/atendimento/add", &controllers.AtendimentoController{}, "get:Add"),
		web.NSRouter("/atendimento/create", &controllers.AtendimentoController{}, "post:Create"),
		web.NSRouter("/atendimento/atendimento", &controllers.AtendimentoController{}, "get:Atendimento"),
		web.NSRouter("/atendimento/ajusteconteudo", &controllers.AtendimentoController{}, "post:AjusteConteudo"),
		web.NSRouter("/atendimento/getnode", &controllers.AtendimentoController{}, "post:GetNode"),
		web.NSRouter("/atendimento/getnodes", &controllers.AtendimentoController{}, "post:GetNodes"),
		web.NSRouter("/atendimento/editpaciente", &controllers.AtendimentoController{}, "get:EditPaciente"),
		web.NSRouter("/atendimento/update", &controllers.AtendimentoController{}, "post:Update"),
		web.NSRouter("/atendimento/enable", &controllers.AtendimentoController{}, "post:Enable"),
		web.NSRouter("/atendimento/disable", &controllers.AtendimentoController{}, "post:Disable"),
		web.NSRouter("/atendimento/del", &controllers.AtendimentoController{}, "post:Del"),
		web.NSRouter("/atendimento/export", &controllers.AtendimentoController{}, "get:Export"),

		web.NSRouter("/atendimento/getnodes", &controllers.AtendimentoController{}, "get:GetNodes"),

		//Ger notas
		web.NSRouter("/nota/index", &controllers.NotaController{}, "get:Index"),
		web.NSRouter("/nota/add", &controllers.NotaController{}, "get:Add"),
		web.NSRouter("/nota/create", &controllers.NotaController{}, "post:Create"),
		web.NSRouter("/nota/update", &controllers.NotaController{}, "post:Update"),
		web.NSRouter("/nota/enable", &controllers.NotaController{}, "post:Enable"),
		web.NSRouter("/nota/disable", &controllers.NotaController{}, "post:Disable"),
		web.NSRouter("/nota/del", &controllers.NotaController{}, "post:Del"),
		web.NSRouter("/nota/export", &controllers.NotaController{}, "get:Export"),

		//Ger modelo
		web.NSRouter("/modelo/index", &controllers.ModeloController{}, "get:Index"),
		web.NSRouter("/modelo/add", &controllers.ModeloController{}, "get:Add"),
		web.NSRouter("/modelo/create", &controllers.ModeloController{}, "post:Create"),
		web.NSRouter("/modelo/edit", &controllers.ModeloController{}, "get:Edit"),
		web.NSRouter("/modelo/update", &controllers.ModeloController{}, "post:Update"),
		web.NSRouter("/modelo/enable", &controllers.ModeloController{}, "post:Enable"),
		web.NSRouter("/modelo/disable", &controllers.ModeloController{}, "post:Disable"),
		web.NSRouter("/modelo/del", &controllers.ModeloController{}, "post:Del"),
		web.NSRouter("/modelo/export", &controllers.ModeloController{}, "get:Export"),

		//Ger campo
		web.NSRouter("/campo/index", &controllers.CampoController{}, "get:Index"),
		web.NSRouter("/campo/add", &controllers.CampoController{}, "get:Add"),
		web.NSRouter("/campo/create", &controllers.CampoController{}, "post:Create"),
		web.NSRouter("/campo/edit", &controllers.CampoController{}, "get:Edit"),
		web.NSRouter("/campo/update", &controllers.CampoController{}, "post:Update"),
		web.NSRouter("/campo/del", &controllers.CampoController{}, "post:Del"),
		web.NSRouter("/campo/export", &controllers.CampoController{}, "get:Export"),

		//Ger cor
		web.NSRouter("/cor/index", &controllers.CorController{}, "get:Index"),
		web.NSRouter("/cor/add", &controllers.CorController{}, "get:Add"),
		web.NSRouter("/cor/create", &controllers.CorController{}, "post:Create"),
		web.NSRouter("/cor/edit", &controllers.CorController{}, "get:Edit"),
		web.NSRouter("/cor/update", &controllers.CorController{}, "post:Update"),
		web.NSRouter("/cor/enable", &controllers.CorController{}, "post:Enable"),
		web.NSRouter("/cor/disable", &controllers.CorController{}, "post:Disable"),
		web.NSRouter("/cor/del", &controllers.CorController{}, "post:Del"),
		web.NSRouter("/cor/export", &controllers.CorController{}, "get:Export"),

		web.NSRouter("/import/receita", &controllers.ImportController{}, "get:Receita"),
		web.NSRouter("/import/template", &controllers.ImportController{}, "get:Template"),
		web.NSRouter("/print/atendimento", &controllers.PrintController{}, "get:Atendimento"),
		web.NSRouter("/print/nota", &controllers.PrintController{}, "get:Nota"),
	)

	web.AddNamespace(admin)
}

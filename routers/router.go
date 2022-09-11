package routers

import (
	"pm/app/AppController"

	"github.com/astaxie/beego"
)

func init() {
	//Roteamento de back-end
	//beego.Include(&AppController.LoginController{})
	//beego.Include(&AppController.IndexController{})
	//beego.Include(&AppController.MenuController{})
	//beego.Include(&AppController.UserController{})

	//Roteamento de front-end
	//beego.Include(&ApiController.IndexController{})
	beego.SetStaticPath("/upload", "upload")                                        //上传目录
	beego.Router("/", &AppController.IndexController{}, "get:Index")                //首页
	beego.Router("/layout", &AppController.IndexController{}, "get:Layout")         //layout页面
	beego.Router("/main", &AppController.IndexController{}, "get:Main")             //página home
	beego.Router("/client", &AppController.ClientController{}, "get:Lists")         //página home
	beego.Router("/client/search", &AppController.ClientController{}, "get:Search") //página home
	beego.Router("/login", &AppController.LoginController{}, "*:Login")             //Entrar
	beego.Router("/logout", &AppController.LoginController{}, "*:Logout")           //Sair

	beego.AutoRouter(&AppController.MenuController{})
	beego.AutoRouter(&AppController.ClientController{})
	beego.AutoRouter(&AppController.IndexController{})
	beego.AutoRouter(&AppController.SysController{})

}

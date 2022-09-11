package main

import (
	_ "p-med/initialize/conf"
	_ "p-med/initialize/mysql"
	_ "p-med/initialize/session"
	_ "p-med/routers"
	_ "p-med/utils/template"
	beego "github.com/beego/beego/v2/adapter"
)

func main() {

	//输出文件名和行号
	beego.SetLogFuncCall(true)

	//启动beego
	beego.Run()
}

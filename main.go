package main

import (
	_ "pm/pm/init"
	_ "pm/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

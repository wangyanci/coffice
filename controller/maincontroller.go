package controller

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func(this *MainController) Home() {
	fmt.Println("xxxx")
	this.Data["Title"] = "json在线解析神器"
	this.TplName = "index.html"
}

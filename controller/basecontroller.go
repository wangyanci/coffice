package controller

import (
	"fmt"
	"github.com/astaxie/beego"
	e "github.com/wangyanci/coffice/exception"
	"github.com/wangyanci/coffice/utils"
	"github.com/wangyanci/coffice/version"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Home() {
	fmt.Println("xxxx")
	c.Data["Title"] = "json在线解析神器"
	c.TplName = "index.html"
}

func (c *BaseController)Version() {
	versionInfo := map[string]string{
		"release": version.Version,
		"commit": version.GitCommit,
		"buildDate": version.BuildDate,
	}

	c.OutputSuccess(versionInfo)
}

func (c *BaseController)Health() {
	c.OutputSuccess(nil)
}

func (c *BaseController)NotFound() {
	c.OutputErrorV4Code(e.GLOBAL_ROUTE_NOT_FOUND)
}


//如果不传响应码，则默认使用200，否则只有第一个有效
func (c *BaseController) OutputSuccess(data interface{}, code ...int) {
	utils.OutputSuccess(c.Ctx, data, code...)
}

//从coffice错误码返回错误信息
func (c *BaseController) OutputErrorV4Code(code e.ErrorCode, errors ...error) {
	utils.OutputErrorV4Code(c.Ctx, code, errors...)
}

//从coffice错误返回错误信息
func (c *BaseController) OutputV4Error(k4SErr *e.K4SError, errors ...error) {
	utils.OutputV4Error(c.Ctx, k4SErr, errors...)
}

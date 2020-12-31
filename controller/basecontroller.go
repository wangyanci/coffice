package controller

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
	e "github.com/wangyanci/coffice/exception"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Home() {
	fmt.Println("xxxx")
	c.Data["Title"] = "json在线解析神器"
	c.TplName = "index.html"
}

//如果不传响应码，则默认使用200，否则只有第一个有效
func (c *BaseController) OutputSuccess(data interface{}, code ...int) {
	c.Data["json"] = data
	c.Ctx.Output.SetStatus(http.StatusOK)
	if len(code) != 0 {
		c.Ctx.Output.SetStatus(code[0])
	}

	c.ServeJSON()
}

//从coffice错误码返回错误信息
func (c *BaseController) OutputErrorV4Code(code e.ErrorCode, msg ...string) {
	c.Data["json"] = code.Code2K4SERROR(msg...)
	c.Ctx.Output.SetStatus(code.CodeInfo(e.STATUSCODE).(int))
	c.ServeJSON()
}

//从coffice错误返回错误信息
func (c *BaseController) OutputError(k4SErr *e.K4SError, msg ...string) {

	c.Data["json"] = k4SErr.AppendMsg(msg...)
	c.Ctx.Output.SetStatus(k4SErr.Code.CodeInfo(e.STATUSCODE).(int))
	c.ServeJSON()
}

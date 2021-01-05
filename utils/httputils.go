package utils

import (
	"net/http"

	e "github.com/wangyanci/coffice/exception"

	"github.com/astaxie/beego/context"
)

//如果不传响应码，则默认使用200，否则只有第一个传入的响应码有效
func OutputSuccess(ctx *context.Context, data interface{}, code ...int) {
	ctx.Output.SetStatus(http.StatusOK)
	if len(code) != 0 {
		ctx.Output.SetStatus(code[0])
	}
	if data != nil {
		ctx.Output.JSON(data, false, true)
	}
}

//从coffice错误码返回错误信息
func OutputErrorV4Code(ctx *context.Context, code e.ErrorCode, errors ...error) {
	ctx.Output.SetStatus(code.CodeInfo(e.STATUSCODE).(int))
	ctx.Output.JSON(code.Code2K4SERROR(errors...), false, true)
}

//从coffice错误返回错误信息
func OutputV4Error(ctx *context.Context, k4SErr *e.K4SError, errors ...error) {
	ctx.Output.SetStatus(k4SErr.Code.CodeInfo(e.STATUSCODE).(int))
	ctx.Output.JSON(k4SErr.AppendMsg(errors...), false, true)
}

package utils

import (
	"encoding/json"

	e6rror "vueApp/error"

	"github.com/astaxie/beego/context"
)

func ResponseWithSuccess(ctx *context.Context, code int, payload interface{}) {
	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	ctx.ResponseWriter.WriteHeader(code)
	if payload != nil {
		response, _ := json.Marshal(payload)
		ctx.ResponseWriter.Write(response)
	}
}

func ResponseWithError(ctx *context.Context, code int, err error) {
	e6rr := e6rror.E609rror{
		Code: code,
		Msg: e6rror.ErrorMessage[code].Msg,
	}
	response, _ := json.Marshal(e6rr)
	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	ctx.ResponseWriter.WriteHeader(e6rror.ErrorMessage[code].Code)
	ctx.ResponseWriter.Write(response)
}
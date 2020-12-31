package utils

//"encoding/json"
//
//
//"github.com/astaxie/beego/context"

//func ResponseWithSuccess(ctx *context.Context, code int, payload interface{}) {
//	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
//	ctx.ResponseWriter.WriteHeader(code)
//	if payload != nil {
//		response, _ := json.Marshal(payload)
//		ctx.ResponseWriter.Write(response)
//	}
//}
//
//func ResponseWithError(ctx *context.Context, code int, err error) {
//	e6rr := e6rror.E609rror{
//		Code: code,
//		Msg: e6rror.ErrorMessage[code].Msg,
//	}
//	response, _ := json.Marshal(e6rr)
//	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
//	ctx.ResponseWriter.WriteHeader(e6rror.ErrorMessage[code].Code)
//	ctx.ResponseWriter.Write(response)
//}
//
//func ResponseWithError111(ctx *context.Context, code string, err error) {
//	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
//	ctx.ResponseWriter.WriteHeader(500)
//	ctx.ResponseWriter.Write([]byte(`{"code":"S2I.4042003301","error":"unmarshal request body failed"}`))
//}

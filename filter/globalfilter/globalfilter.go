package globalfilter

import (
	"errors"

	"github.com/wangyanci/coffice/utils"

	"github.com/astaxie/beego/context"
)

var skipMediaTypeCheckRouter = map[string]map[string]bool{
	"/": {
		"Any": true,
	},
}

func PreDeal(ctx *context.Context) {
	ctx.Input.RequestBody = ctx.Input.CopyBody(1024)

	if utils.IsSkipFilterRouter(ctx.Input.URL(), ctx.Input.Method(), skipMediaTypeCheckRouter) {
		return
	}

	contentType := ctx.Input.Header("Content-Type")
	if contentType == "application/json" {
		return
	}

	err := errors.New("content type is not application/json")
	utils.ResponseWithError(ctx, error.GLOBAL_ALL_MEDIA_TYPE_ERROE, err)
}

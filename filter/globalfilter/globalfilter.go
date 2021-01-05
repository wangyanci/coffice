package globalfilter

import (
	"errors"
	e "github.com/wangyanci/coffice/exception"
	"strings"

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

	if utils.IsSkipFilterRouter(ctx, ctx.Input.Method(), skipMediaTypeCheckRouter) {
		return
	}

	contentType := ctx.Input.Header("Content-Type")
	if strings.Index(contentType, "application/json") != -1 {
		return
	}

	err := errors.New("content type is not application/json")
	utils.OutputErrorV4Code(ctx, e.GLOBAL_ALL_MEDIA_TYPE_ERROE, err)
}

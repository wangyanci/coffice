package identfyfilter

import (
	"errors"
	"net/http"

	"github.com/wangyanci/coffice/auth"
	e "github.com/wangyanci/coffice/exception"
	"github.com/wangyanci/coffice/logs"
	"github.com/wangyanci/coffice/utils"

	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
)

var skipAuthRouter = map[string]map[string]bool{
	"/": {
		"Any": true,
	},
	"/v1/auth": {
		http.MethodPost: true,
	},
	"/v1/users": {
		http.MethodPost: true,
	},
}

func Identfy(ctx *context.Context) {
	if utils.IsSkipFilterRouter(ctx.Input.URL(), ctx.Input.Method(), skipAuthRouter) {
		return
	}

	tokenStr := ctx.Input.Header("X-Auth-Token")
	if tokenStr == "" {
		err := errors.New("the token is empty")
		logs.Logger.Info("authentication failed err: %v", err)
		utils.OutputErrorV4Code(ctx, e.AUTH_GET_TOKEN_FAIL, err)
		return
	}

	token, err := jwt.Parse(tokenStr, auth.Keyfunc)
	if err != nil {
		utils.OutputErrorV4Code(ctx, e.AUTH_GET_DECRYPT_FAIL, err)
		return
	}

	if !token.Valid {
		err := errors.New("validate token failed")
		utils.OutputErrorV4Code(ctx, e.AUTH_GET_VALIDATE_FAIL, err)
	}
}

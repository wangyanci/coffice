package authcontroller

import (
	"fmt"
	"encoding/json"
	"net/http"

	"vueApp/auth"
	"vueApp/error"
	"vueApp/logs"
	"vueApp/model"
	"vueApp/utils"

	"github.com/astaxie/beego"
)

type AuthController struct {
	beego.Controller
}

func(this *AuthController) Auth() {
	user := model.User{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	if err != nil {
		logs.Logger.Error("authentication failed err: %v", err)
		utils.ResponseWithError(this.Ctx, error.AUTH_POST_UNMARSHAL_FAIL, err)
		return
	}

	toke, err := auth.GetAuthToken(user)
	if err != nil {
		fmt.Println(err)
		utils.ResponseWithError(this.Ctx, error.AUTH_POST_ENCRYPT_FAIL, err)
		return
	}

	this.Ctx.ResponseWriter.Header().Set("X-Auth-Token", toke)
	utils.ResponseWithSuccess(this.Ctx, http.StatusCreated, nil)
}

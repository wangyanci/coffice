package authcontroller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/wangyanci/coffice/auth"
	"github.com/wangyanci/coffice/controller"
	e "github.com/wangyanci/coffice/exception"
	"github.com/wangyanci/coffice/logs"
	"github.com/wangyanci/coffice/model"
)

type AuthController struct {
	controller.BaseController
}

func (c *AuthController) Auth() {
	user := model.User{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if err != nil {
		logs.Logger.Error("authentication failed err: %v", err)
		c.OutputErrorV4Code(e.AUTH_POST_UNMARSHAL_FAIL, err)
		return
	}

	if user.DomainName == "" || user.Secret == "" {
		err := errors.New("username or password is empty")
		c.OutputErrorV4Code(e.AUTH_POST_UNMARSHAL_FAIL, err)
		return
	}

	if user.DomainName != "wangyanci" || user.Secret != "123456" {
		c.OutputErrorV4Code(e.AUTH_PASSWORD_INVAILD, err)
		return
	}

	token, err := auth.GetAuthToken(user)
	if err != nil {
		fmt.Println(err)
		c.OutputErrorV4Code(e.AUTH_POST_ENCRYPT_FAIL, err)
		return
	}

	c.Ctx.ResponseWriter.Header().Set("X-Auth-Token", token)
	c.OutputSuccess(nil, http.StatusCreated)
}

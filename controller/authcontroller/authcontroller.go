package authcontroller

import (
	"encoding/json"
	"fmt"
	"github.com/wangyanci/coffice/service/userservice"
	v "github.com/wangyanci/coffice/utils/validation"
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

func (c *AuthController) Auth()  {
	user := &model.User{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, user)
	if err != nil {
		logs.Logger.Error("authentication failed err: %v", err)
		c.OutputErrorV4Code(e.AUTH_POST_UNMARSHAL_FAIL, err)
		return
	}


	fieldErr := v.ValidateUser(user)
	if fieldErr != nil {
		c.OutputErrorV4Code(e.GLOBAL_REQUEST_PARAM_INVALID, fieldErr)
		return
	}

	ok, k4sErr := userservice.UserService.ValidateUser(user)
	if k4sErr != nil {
		c.OutputErrorV4Code(e.AUTH_VALIDATE_INTERNAL_ERROR, err)
		return
	}

	if !ok {
		c.OutputErrorV4Code(e.AUTH_PASSWORD_INVAILD)
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

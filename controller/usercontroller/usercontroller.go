package usercontroller

import (
	"encoding/json"
	"fmt"
	"github.com/wangyanci/coffice/model"
	service "github.com/wangyanci/coffice/service/userservice"
	"net/http"

	"github.com/wangyanci/coffice/controller"
	e "github.com/wangyanci/coffice/exception"
	v "github.com/wangyanci/coffice/utils/validation"
)

type UserController struct {
	controller.BaseController
}

func (c *UserController) CreateUser() {
	user := &model.User{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, user)
	if err != nil {
		c.OutputErrorV4Code(e.GLOBAL_REQUEST_UNMARSHAL_ERROR, err)
		return
	}

	fieldErr := v.ValidateUser(user)
	if fieldErr != nil {
		c.OutputErrorV4Code(e.GLOBAL_REQUEST_PARAM_INVALID, fieldErr)
		return
	}

	k4sErr := service.UserService.CreateUser(user)
	if k4sErr != nil {
		c.OutputV4Error(k4sErr)
		return
	}

	c.OutputSuccess(nil, http.StatusCreated)
}

func (c *UserController) ListUser() {
	c.Ctx.WriteString("get user list!")
}

func (c *UserController) GetUserById() {
	id := c.Ctx.Input.Param(":id")
	resp := fmt.Sprintf("get user by id %s!", id)
	c.Ctx.WriteString(resp)
}

func (c *UserController) DeleteUserById() {
	id := c.Ctx.Input.Param(":id")
	resp := fmt.Sprintf("delete user by id %s!", id)
	c.Ctx.WriteString(resp)
}

func (c *UserController) UpdateUserById() {
	id := c.Ctx.Input.Param(":id")
	resp := fmt.Sprintf("update user by id %s!", id)
	c.Ctx.WriteString(resp)
}

func (c *UserController) IsUserExist() {
	name := c.Ctx.Input.Param(":name")

	fieldErr := v.ValidateUserName(name)
	if fieldErr != nil {
		c.OutputErrorV4Code(e.GLOBAL_REQUEST_PARAM_INVALID, fieldErr)
		return
	}

	exist := service.UserService.IsUserExist(name)
	if !exist {
		c.OutputErrorV4Code(e.USER_NOT_FOUND)
		return
	}

	c.OutputSuccess(nil)
}

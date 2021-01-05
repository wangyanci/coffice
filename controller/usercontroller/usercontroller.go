package usercontroller

import (
	"fmt"
	service "github.com/wangyanci/coffice/service/userservice"

	"github.com/wangyanci/coffice/controller"
	e "github.com/wangyanci/coffice/exception"
)

type UserController struct {
	controller.BaseController
}

func (this *UserController) CreateUser() {
	service.UserService.CreateUser()
	this.Ctx.WriteString("create user!")
}

func (this *UserController) ListUser() {
	this.Ctx.WriteString("get user list!")
}

func (this *UserController) GetUserById() {
	id := this.Ctx.Input.Param(":id")
	resp := fmt.Sprintf("get user by id %s!", id)
	this.Ctx.WriteString(resp)
}

func (this *UserController) DeleteUserById() {
	id := this.Ctx.Input.Param(":id")
	resp := fmt.Sprintf("delete user by id %s!", id)
	this.Ctx.WriteString(resp)
}

func (this *UserController) UpdateUserById() {
	id := this.Ctx.Input.Param(":id")
	resp := fmt.Sprintf("update user by id %s!", id)
	this.Ctx.WriteString(resp)
}

func (this *UserController) IsUserExist() {
	name := this.Ctx.Input.Param(":name")
	if name == "wangyanci" {
		this.OutputErrorV4Code(e.USER_NAME_USED)
		return
	}

	this.OutputSuccess(nil)
}

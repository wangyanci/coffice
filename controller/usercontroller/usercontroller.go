package usercontroller

import (
	"fmt"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func(this *UserController) CreateUser() {
	this.Ctx.WriteString("create user!")
}

func(this *UserController) ListUser() {
	this.Ctx.WriteString("get user list!")
}

func(this *UserController) GetUserById() {
	id := this.Ctx.Input.Param(":id")
	resp := fmt.Sprintf("get user by id %s!", id)
	this.Ctx.WriteString(resp)
}

func(this *UserController) DeleteUserById() {
	id := this.Ctx.Input.Param(":id")
	resp := fmt.Sprintf("delete user by id %s!", id)
	this.Ctx.WriteString(resp)
}

func(this *UserController) UpdateUserById() {
	id := this.Ctx.Input.Param(":id")
	resp := fmt.Sprintf("update user by id %s!", id)
	this.Ctx.WriteString(resp)
}
package router

import (
	"github.com/wangyanci/coffice/controller"
	"github.com/wangyanci/coffice/controller/authcontroller"
	"github.com/wangyanci/coffice/controller/usercontroller"

	"github.com/astaxie/beego"
)

func InitRouter() {
	mainController := &controller.BaseController{}
	authController := &authcontroller.AuthController{}
	userController := &usercontroller.UserController{}

	beego.Router("/", mainController, "*:Home")

	v1 := beego.NewNamespace("v1",
		beego.NSNamespace("/auth",
			beego.NSRouter("/", authController, "post:Auth"),
		),

		beego.NSNamespace("/users",
			beego.NSRouter("/", userController, "post:CreateUser"),
			beego.NSRouter("/", userController, "get:ListUser"),
			beego.NSRouter("/:id", userController, "get:GetUserById"),
			beego.NSRouter("/:id", userController, "delete:DeleteUserById"),
			beego.NSRouter("/:id", userController, "put:UpdateUserById"),
		),
	)

	beego.AddNamespace(v1)
}

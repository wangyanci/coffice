package router

import (
	"github.com/wangyanci/coffice/controller"
	"github.com/wangyanci/coffice/controller/authcontroller"
	"github.com/wangyanci/coffice/controller/usercontroller"

	"github.com/astaxie/beego"
)

func InitRouter() {
	baseController := &controller.BaseController{}
	authController := &authcontroller.AuthController{}
	userController := &usercontroller.UserController{}

	beego.Router("/", baseController, "*:Version")
	beego.Router("/healthy", baseController, "*:Health")

	v1 := beego.NewNamespace("v1",
		beego.NSNamespace("/auth",
			beego.NSRouter("/", authController, "post:Auth"),
		),

		beego.NSNamespace("/users",
			beego.NSRouter("/", userController, "post:CreateUser"),
			beego.NSRouter("/", userController, "get:ListUser"),
			beego.NSRouter("/:id:int", userController, "get:GetUserById"),
			beego.NSRouter("/:name:string", userController, "get:IsUserExist"),
			beego.NSRouter("/:id:int", userController, "delete:DeleteUserById"),
			beego.NSRouter("/:id:int", userController, "put:UpdateUserById"),
		),
	)


	beego.AddNamespace(v1)
	beego.Router("*", baseController, "*:NotFound" )
}

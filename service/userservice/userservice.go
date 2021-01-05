package userservice

//import "github.com/astaxie/beego"

var UserService UserServiceInterface = new(userService)
type UserServiceInterface interface {
	CreateUser()
}
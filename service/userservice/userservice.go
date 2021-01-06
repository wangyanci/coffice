package userservice

import (
	e "github.com/wangyanci/coffice/exception"
	"github.com/wangyanci/coffice/model"
)

//import "github.com/astaxie/beego"

var UserService UserServiceInterface = new(userService)
type UserServiceInterface interface {
	CreateUser(user *model.User) *e.K4SError
	IsUserExist(userName string) bool
	GetUsersByFilter()([]model.User, *e.K4SError)
	ValidateUser(*model.User) (bool, *e.K4SError)
}
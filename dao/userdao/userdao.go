package userdao

import (
	"github.com/wangyanci/coffice/dao"
	e "github.com/wangyanci/coffice/exception"
	"github.com/wangyanci/coffice/model"
)

var UserDao UserDaoInterface = new(userDao)
type UserDaoInterface interface {
	InsertUser(*model.User)*e.K4SError
	GetUserByName(string)(*model.User, *e.K4SError)
	IsUserExist(string)bool
	GetUsersByFilter(dao.Filters)([]model.User, *e.K4SError)
}

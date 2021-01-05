package userdao

import "github.com/wangyanci/coffice/model"

var UserDao UserDaoInterface = new(userDao)
type UserDaoInterface interface {
	InsertUser(string)(*model.User, error)
}

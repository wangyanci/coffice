package userservice

import (
	"github.com/wangyanci/coffice/dao"
	"github.com/wangyanci/coffice/dao/userdao"
	"github.com/wangyanci/coffice/model"
)

type userService struct {

}

func (s *userService) CreateUser() {
	var user *model.User
	if err := dao.UseTransaction(func(userDao userdao.UserDaoInterface,
		userDao2 userdao.UserDaoInterface) (err error) {
		user, err = userDao.InsertUser("")
		return err
	}, userdao.UserDao, userdao.UserDao); err != nil {

	}
}
package userservice

import (
	"github.com/wangyanci/coffice/dao"
	"github.com/wangyanci/coffice/dao/userdao"
	e "github.com/wangyanci/coffice/exception"
	"github.com/wangyanci/coffice/model"
)

type userService struct {

}

func (s *userService) CreateUser(user *model.User) *e.K4SError {
	//var user *model.User
	//if err := dao.UseTransaction(func(userDao userdao.UserDaoInterface,
	//	userDao2 userdao.UserDaoInterface) (err error) {
	//	user, err = userDao.InsertUser("")
	//	return err
	//}, userdao.UserDao, userdao.UserDao); err != nil {
	//
	//}
	exist := userdao.UserDao.IsUserExist(user.DomainName)
	if exist {
		return e.USER_IS_EXIST.Code2K4SERROR()
	}

	return userdao.UserDao.InsertUser(user)
}

func (s *userService) IsUserExist(userName string) bool {
	return userdao.UserDao.IsUserExist(userName)
}

func (s *userService) GetUsersByFilter()([]model.User, *e.K4SError) {
	filters := dao.Filters{}
	return userdao.UserDao.GetUsersByFilter(filters)
}

func (s *userService) ValidateUser(user *model.User) (bool, *e.K4SError) {
	filters := dao.Filters{
		"DomainName": []interface{}{user.DomainName},
		"Secret": []interface{}{user.Secret},
	}

	users, k4sErr := userdao.UserDao.GetUsersByFilter(filters)
	if k4sErr != nil {
		return false, k4sErr
	}

	if len(users) == 0 {
		return false, nil
	}

	return true, nil
}
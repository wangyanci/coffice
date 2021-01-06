package userdao

import (
	"github.com/astaxie/beego/orm"
	"github.com/wangyanci/coffice/dao"
	e "github.com/wangyanci/coffice/exception"
	"github.com/wangyanci/coffice/model"
)

type userDao struct {
	orm orm.Ormer
}

func (m *userDao) InsertUser(user *model.User) *e.K4SError {
	o := m.getOrm()

	//当使用了非递增类型主键（如字符串），会报错LastInsertId is not supported by this driver
	_, err := o.Insert(user)
	if err != nil {
		return e.USER_DB_INSERT_FAIL.Code2K4SERROR(err)
	}

	return nil
}

func (m *userDao) GetUserByName(userName string)(*model.User, *e.K4SError) {
	o := m.getOrm()
	user := new(model.User)
	err := o.QueryTable(user).Filter("DomainName", userName).One(user)
	if err != nil {
		return nil, e.USER_DB_GET_BY_NAME_FAIL.Code2K4SERROR(err)
	}

	return user, nil
}

func (m *userDao) IsUserExist(userName string)bool {
	o := m.getOrm()
	return o.QueryTable(new(model.User)).Filter("DomainName", userName).Exist()
}

func (m *userDao) GetUsersByFilter(filters  dao.Filters)([]model.User, *e.K4SError){
	o := m.getOrm()
	users := make([]model.User,0)
	qs := o.QueryTable(new(model.User))
	for i, v := range filters {
		qs = qs.Filter(i, v)
	}

	_, err := qs.All(&users)
	if err != nil {
		return nil, e.USER_DB_LIST_BY_FILTER_FAIL.Code2K4SERROR(err)
	}

	return users, nil
}


func (m *userDao) getOrm() (orm.Ormer) {
	if m.orm == nil {
		return orm.NewOrm()
	}

	return m.orm
}



package userdao

import (
	"github.com/astaxie/beego/orm"
	"github.com/wangyanci/coffice/model"
)

type userDao struct {
	orm orm.Ormer
}

func (m *userDao) InsertUser(aaa string) (*model.User, error) {
	//o := orm.NewOrm()
	return nil, nil
}



func (m *userDao) getOrm() (orm.Ormer) {
	if m.orm == nil {
		return orm.NewOrm()
	}

	return m.orm
}



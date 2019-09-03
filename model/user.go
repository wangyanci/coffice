package model

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Age         int     `json:"age" orm:"column(age)"`
	Secret      string  `json:"secret" orm:"column(secret)"`
	DomainId    string	`json:"domain_id" orm:"column(domain_id);pk"`
	DomainName  string  `json:"domain_name" orm:"column(domain_name)"`
	Telephone   string  `json:"telephone" orm:"column(telephone);null"`
	Email       string  `json:"email" orm:"column(email); null"`
	IsValidate  string  `json:"is_validate" orm:"column(is_validate)"`
	AccountType string  `json:"account_type" orm:"column(account_type)"`
}

func init() {
	orm.RegisterModel(new(User))
}

func TableName() string {
	return "user"
}

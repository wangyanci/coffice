package validation

import (
	"regexp"
	"strings"

	"github.com/wangyanci/coffice/model"
)
const (
	UserNameRex = "^[^_.*?$0-9][a-zA-Z0-9_\u4e00-\u9fa5]+$"
)

//type aaa struct {
//
//}
//
//func(a *aaa)Error()string{
//	return ""
//}

func ValidateUser(user *model.User) (err *FieldError) {
	err = err.Also(ValidateUserName(user.DomainName)).
		Also(ValidateUserPassword(user.Secret)).ViaField("user")
	//fmt.Println("xxxx err: ", err)
	//fmt.Println("xxxx err: ", err==nil)
	//var kkkk *FieldError
	//fmt.Println("yyyy err: ", kkkk==nil)
	//fmt.Println("yyyy err: ", error(kkkk)==nil)
	//var mmm *e.K4SError
	//fmt.Println("zzzz err: ", mmm==nil)
	//fmt.Println("zzzz err: ", error(mmm)==nil)
	//
	//var nnn *e.K4SError
	//fmt.Println("pppp err: ", nnn==nil)
	//fmt.Println("pppp err: ", error(nnn)==nil)

	return
}

func ValidateUserName(userName string) (err *FieldError) {
	if !regexp.MustCompile(UserNameRex).MatchString(userName) {
		err = err.Also(ErrInvalidValue(userName, CurrentField).SetDetails("user name must match %s", UserNameRex)).ViaField("userName")
	}

	return err
}

func ValidateUserPassword(password string)(err *FieldError) {
	if strings.TrimSpace(password) == "" {
		err = err.Also(ErrInvalidValue(password, CurrentField).SetDetails("user password must not empty")).ViaField("password")
	}

	return err
}
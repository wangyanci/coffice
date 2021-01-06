package dao

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/astaxie/beego/orm"
)

func UseTransaction(fn interface{}, daos ...interface{}) error {
	return (&transactionManager{
		orm: orm.NewOrm(),
	}).invoke(fn, daos...).do()
}

type transactionManager struct {
	orm    orm.Ormer
	errors []error
	fn     reflect.Value
	params []reflect.Value
}

type Filters map[string][]interface{}

//func Model2Filters(model interface{}, columns ...string) Filters {
//	tp := reflect.TypeOf(model)
//	v := reflect.ValueOf(model)
//
//	if
//}

func (t *transactionManager) invoke(fn interface{}, daos ...interface{}) *transactionManager {
	tp := reflect.TypeOf(fn)
	t.fn = reflect.ValueOf(fn)

	//v := reflect.ValueOf(fn)

	if tp.Kind() != reflect.Func {
		t.errors = append(t.errors, errors.New("TransactionManager invoke param type must be func"))
		return t
	}

	daosT := []reflect.Type{}
	daosV := []reflect.Value{}
	for _, v := range daos {
		fmt.Println("xxx: ", reflect.TypeOf(v).Elem())
		daosT = append(daosT, reflect.TypeOf(v))
		daosV = append(daosV, reflect.ValueOf(v))
	}

	fmt.Println("daosT: ", daosT)
	fmt.Println("daosV: ", daosV)

	if len(daos) != tp.NumIn() {
		t.errors = append(t.errors, errors.New("TransactionManager invoke param count must be equal"))
		return t
	}

	for i := 0; i < tp.NumIn(); i++ {
		fmt.Printf("fn param [%+v] ====== [%+v]\n", daosT[i], tp.In(i))
		if !daosT[i].Implements(tp.In(i)) {
			t.errors = append(t.errors, errors.New("TransactionManager invoke param type must be consistent"))
			return t
		}

		fmt.Printf("fn param [%+v] Implements type [%+v]\n", daosT[i], tp.In(i))

		dao := reflect.New(daosT[i].Elem())
		o := dao.Elem().FieldByName("orm")
		if !o.IsValid() {
			t.errors = append(t.errors, errors.New("TransactionManager invoke param is not a dao"))
			return t
		}

		fmt.Println("oooo: ", o)

		o.Set(reflect.ValueOf(t.orm))

		t.params = append(t.params, dao)

	}

	fmt.Printf("fn param fn return  type [%+v]\n", tp.Out(0))

	if tp.NumOut() != 1 && tp.Out(0) != reflect.TypeOf(errors.New("")) {
		t.errors = append(t.errors, errors.New("TransactionManager invoke fn return type must be error"))
		return t
	}

	return t
}

func (t *transactionManager) do() error {
	var err error
	if len(t.errors) != 0 {
		return getErrors(t.errors)
	}

	err = t.orm.Begin()
	if err != nil {
		return err
	}

	results := t.fn.Call(t.params)
	if len(results) != 1 {
		err = errors.New("TransactionManager invoke fn return type must be error")
		return err
	}

	if results[0].IsNil() {
		err = t.orm.Commit()
		if err == nil {
			return nil
		}
	}

	return getErrors(append(t.errors, err, t.orm.Rollback()))
}

//func Test() {
//	err := TransactionManage().Invoke(func(userdao UserDaoInterface,productdao ProductDaoInterface)error {
//		return  nil
//	}, UserDao, ProductDao).Do()
//
//	fmt.Println("error", err)
//}

func getErrors(errs []error) error {
	var (
		hasErr bool
		errMsg string
	)

	for _, e := range errs {
		if e != nil {
			hasErr = true
			errMsg += e.Error()
		}
	}

	if hasErr {
		return errors.New(errMsg)
	}

	return nil
}

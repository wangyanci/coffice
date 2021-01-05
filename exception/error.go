package exception

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var _ error = new(K4SError)

type ErrorCode string
type K4SError struct {
	Code   ErrorCode `json:"code"`
	Msg    string    `json:"message"`
	Detail string    `json:"detail,omitempty"`
}

const _CONCAT = "; "
const _SEPERATE = "."
const _SERVICE = "K4S"
const _CONCATERROR = ":"

const (
	SVC = iota + 1
	STATUSCODE
	APIVERSION
	MODULENAME
	SEQUENCE
)
const (
	_MODULEL_AUTH   = (iota + 1) * 1000
	_MODULEL_USER   = (iota + 1) * 1000
	_MODULEL_GLOBAL = (iota + 1) * 1000
)

const (
	_APIVERSION_V1 = (iota + 1) * 100000
)

const (
	_K4S_NOTFUND        = http.StatusNotFound * 10000000
	_K4S_CONFLICT       = http.StatusConflict * 10000000
	_K4S_BADREQUEST     = http.StatusBadRequest * 10000000
	_K4S_UNAUTHORIZED   = http.StatusUnauthorized * 10000000
	_K4S_INTERNALERROR  = http.StatusInternalServerError * 10000000
	_K4S_MEDIATYPEERROR = http.StatusUnsupportedMediaType * 10000000
)

//K4S.4042003301
//服务名.STATUSCODE+APIVERSION+MODULENAME+SEQUENCE
//(?<SVC>[A-Z0-9]{3}).(?<STATUSCODE>[0-9]{3})(?<APIVERSION>[0-9]{2})(?<MODULENAME>[0-9]{2})(?<SEQUENCE>[0-9]{3})

//router: /v1/auth
var (
	AUTH_GET_TOKEN_FAIL        = _V1(_K4S_UNAUTHORIZED + _MODULEL_AUTH + 001)
	AUTH_GET_DECRYPT_FAIL      = _V1(_K4S_UNAUTHORIZED + _MODULEL_AUTH + 002)
	AUTH_POST_ENCRYPT_FAIL     = _V1(_K4S_INTERNALERROR + _MODULEL_AUTH + 003)
	AUTH_GET_VALIDATE_FAIL     = _V1(_K4S_UNAUTHORIZED + _MODULEL_AUTH + 004)
	AUTH_PASSWORD_INVAILD      = _V1(_K4S_BADREQUEST + _MODULEL_AUTH + 005)
	AUTH_POST_UNMARSHAL_FAIL   = _V1(_K4S_BADREQUEST + _MODULEL_AUTH + 006)
	AUTH_POST_MEDIA_TYPE_ERROE = _V1(_K4S_MEDIATYPEERROR + _MODULEL_AUTH + 007)
)
var (
	USER_NAME_USED = _V1(_K4S_CONFLICT + _MODULEL_USER + 001)
	USER_NOT_FOUND = _V1(_K4S_NOTFUND + _MODULEL_USER + 002)
)

//router: /*
var (
	GLOBAL_UNKNOWN_ERROE        = _V1(_K4S_INTERNALERROR + _MODULEL_GLOBAL + 000)
	GLOBAL_ROUTE_NOT_FOUND      = _V1(_K4S_NOTFUND + _MODULEL_GLOBAL + 001)
	GLOBAL_ALL_MEDIA_TYPE_ERROE = _V1(_K4S_MEDIATYPEERROR + _MODULEL_GLOBAL + 002)
)

var k4SERRORS = map[ErrorCode]*K4SError{
	AUTH_POST_MEDIA_TYPE_ERROE: {Msg: "unsupport media type."},
	AUTH_POST_UNMARSHAL_FAIL:   {Msg: "unmarshal request body failed."},
	AUTH_POST_ENCRYPT_FAIL:     {Msg: "encrypt token failed."},
	AUTH_GET_TOKEN_FAIL:        {Msg: "authentication failed."},
	AUTH_GET_DECRYPT_FAIL:      {Msg: "authentication failed."},
	AUTH_GET_VALIDATE_FAIL:     {Msg: "authentication failed."},
	AUTH_PASSWORD_INVAILD:      {Msg: "authentication failed, userName or passwd is correct."},

	USER_NAME_USED: {Msg: "user is exist."},
	USER_NOT_FOUND: {Msg: "user not found."},

	GLOBAL_UNKNOWN_ERROE:        {Msg: "internal error, unknown err occur."},
	GLOBAL_ROUTE_NOT_FOUND:      {Msg: "resource not found."},
	GLOBAL_ALL_MEDIA_TYPE_ERROE: {Msg: "unsupport media type."},
}

func _V1(s int) ErrorCode {
	return ErrorCode(_SERVICE + _SEPERATE + strconv.Itoa(s+_APIVERSION_V1))
}

func (code ErrorCode) CodeInfo(name int) interface{} {
	rex := fmt.Sprintf("(?P<SVC>[A-Z0-9]{3}).(?P<STATUSCODE>[0-9]{3})(?P<APIVERSION>[0-9]{2})(?P<MODULENAME>[0-9]{2})(?P<SEQUENCE>[0-9]{3})")
	str := regexp.MustCompile(rex).FindStringSubmatch(string(code))[name]
	if num, err := strconv.Atoi(str); err == nil {
		return num
	} else {
		return str
	}
}

func (code ErrorCode) Code2K4SERROR(errors ...error) *K4SError {
	k4SERROR := &K4SError{Msg: k4SERRORS[code].Msg}
	if k4SERROR == nil {
		k4SERROR = k4SERRORS[GLOBAL_UNKNOWN_ERROE]
	}
	k4SERROR.Code = code
	k4SERROR.AppendMsg(errors...)
	return k4SERROR
}

func (k *K4SError) Error() string {
	return k.Msg + _CONCATERROR + k.Detail
}

func (k *K4SError) AppendMsg(errors ...error) *K4SError {
	detail := strings.TrimSpace(k.Detail)
	if detail != "" && detail != _CONCAT {
		k.Detail = strings.Join(append([]string{detail}, errs2string(errors...)...), _CONCAT)
		return k
	}

	k.Detail = strings.Join(errs2string(errors...), _CONCAT)
	return k
}

func errs2string(errors ...error) []string {
	var msgs []string
	for _, v := range errors {
		if v != nil {
			msgs = append(msgs, v.Error())
		}
	}

	return msgs
}

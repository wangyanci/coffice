package exception

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type ErrorCode string
type K4SError struct {
	Code   ErrorCode `json:"code"`
	Msg    string    `json:"message"`
	Detail string    `json:"detail,omitempty"`
}

const _CONCAT = "; "
const _SEPERATE = "."
const _SERVICE = "K4S"

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
	_K4S_BADREQUEST     = http.StatusBadRequest * 1000000
	_K4S_UNAUTHORIZED   = http.StatusUnauthorized * 1000000
	_K4S_INTERNALERROR  = http.StatusInternalServerError * 1000000
	_K4S_MEDIATYPEERROR = http.StatusUnsupportedMediaType * 1000000
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
	AUTH_POST_UNMARSHAL_FAIL   = _V1(_K4S_BADREQUEST + _MODULEL_AUTH + 005)
	AUTH_POST_MEDIA_TYPE_ERROE = _V1(_K4S_MEDIATYPEERROR + _MODULEL_AUTH + 006)
)

//router: /*
var (
	GLOBAL_ALL_MEDIA_TYPE_ERROE = _V1(_K4S_MEDIATYPEERROR + _MODULEL_GLOBAL + 001)
)

var k4SERRORS = map[ErrorCode]*K4SError{
	AUTH_POST_MEDIA_TYPE_ERROE:  {Msg: "unsupport media type"},
	AUTH_POST_UNMARSHAL_FAIL:    {Msg: "unmarshal request body failed"},
	AUTH_POST_ENCRYPT_FAIL:      {Msg: "encrypt token failed"},
	AUTH_GET_TOKEN_FAIL:         {Msg: "authentication failed"},
	AUTH_GET_DECRYPT_FAIL:       {Msg: "authentication failed"},
	AUTH_GET_VALIDATE_FAIL:      {Msg: "authentication failed"},
	GLOBAL_ALL_MEDIA_TYPE_ERROE: {Msg: "unsupport media type"},
}

func _V1(s int) ErrorCode {
	return ErrorCode(_SERVICE + _SEPERATE + strconv.Itoa(s+_APIVERSION_V1))
}

func (code ErrorCode) CodeInfo(name int) interface{} {
	rex := fmt.Sprintf("(?P<SVC>[A-Z0-9]{3}).(?P<STATUSCODE>[0-9]{3})(?P<APIVERSION>[0-9]{2})(?P<MODULENAME>[0-9]{2})(?P<SEQUENCE>[0-9]{3})")
	str := regexp.MustCompile(rex).FindStringSubmatch(string(code))[name]
	if num, err := strconv.Atoi(str); err != nil {
		return num
	} else {
		return str
	}
}

func (code ErrorCode) Code2K4SERROR(msg ...string) *K4SError {
	k4SERROR := k4SERRORS[code]
	k4SERROR.Code = code
	k4SERROR.AppendMsg(msg...)
	return k4SERROR
}

func (error *K4SError) AppendMsg(msg ...string) *K4SError {
	detail := strings.TrimSpace(error.Detail)
	if detail != "" && detail != _CONCAT {
		error.Detail = strings.Join(append([]string{detail}, msg...), _CONCAT)
		return error
	}

	error.Detail = strings.Join(msg, _CONCAT)
	return error
}

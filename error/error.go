package error

import (
	"net/http"
)

type E609rror struct {
	Code int         `json:"code"`
	Msg  string      `json:"error"`
}

const (
	_MODEL_AUTH = (iota + 1) * 100
	_MODEL_USER = (iota + 1) * 100
	_MODEL_GLOBAL = (iota + 1) * 100
)

const (
	_STATUS_BADREQUEST = http.StatusBadRequest * 10000
	_STATUS_UNAUTHORIZED = http.StatusUnauthorized * 10000
	_STATUS_INTERNALERROR = http.StatusInternalServerError * 10000
	_STATUS_MEDIATYPEERROR = http.StatusUnsupportedMediaType * 10000
)


//router: /v1/auth
const (
	AUTH_POST_MEDIA_TYPE_ERROE = _STATUS_MEDIATYPEERROR + _MODEL_AUTH + iota +1
	AUTH_POST_UNMARSHAL_FAIL = _STATUS_BADREQUEST + _MODEL_AUTH + iota
	AUTH_POST_ENCRYPT_FAIL = _STATUS_INTERNALERROR + _MODEL_AUTH + iota
	AUTH_GET_TOKEN_FAIL = _STATUS_UNAUTHORIZED + _MODEL_AUTH + iota
	AUTH_GET_DECRYPT_FAIL = _STATUS_UNAUTHORIZED + _MODEL_AUTH + iota
	AUTH_GET_VALIDATE_FAIL = _STATUS_UNAUTHORIZED + _MODEL_AUTH + iota
)

//router: /*
const (
	GLOBAL_ALL_MEDIA_TYPE_ERROE = _STATUS_MEDIATYPEERROR + _MODEL_GLOBAL + iota +1
)

var ErrorMessage = map[int]E609rror {
	AUTH_POST_MEDIA_TYPE_ERROE: {http.StatusUnsupportedMediaType, "unsupport media type"},
	AUTH_POST_UNMARSHAL_FAIL: {http.StatusBadRequest, "unmarshal request body failed"},
	AUTH_POST_ENCRYPT_FAIL: {http.StatusInternalServerError, "encrypt token failed"},
	AUTH_GET_TOKEN_FAIL: {http.StatusUnauthorized, "authentication failed"},
	AUTH_GET_DECRYPT_FAIL: {http.StatusUnauthorized, "authentication failed"},
	AUTH_GET_VALIDATE_FAIL: {http.StatusUnauthorized, "authentication failed"},
	GLOBAL_ALL_MEDIA_TYPE_ERROE: {http.StatusUnsupportedMediaType, "unsupport media type"},
}
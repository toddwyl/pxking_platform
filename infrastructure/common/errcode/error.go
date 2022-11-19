package errcode

import (
	"fmt"
	"net/http"
	"sync"
)

// Error 返回错误码和消息的结构体
// nolint: govet
type Error struct {
	ErrCode int    `json:"retcode"`
	Message string `json:"message"`
}

var errorCodes = map[int]struct{}{}
var toStatus sync.Map

// NewError create a error
func NewError(code int, msg string) *Error {
	if _, ok := errorCodes[code]; ok {
		err := fmt.Errorf("code %d is exsit, please change one", code)
		fmt.Printf("err:%s", err)
	}
	errorCodes[code] = struct{}{}
	return &Error{ErrCode: code, Message: msg}
}

// CloneErrorWithDetail create an error with detail
func CloneErrorWithDetail(e *Error, msg string) *Error {
	err := cloneError(e)
	err.Message = err.Message + "|| detail:{ " + msg + "}"
	return err
}

func cloneError(e *Error) *Error {
	return &Error{
		ErrCode: e.Code(),
		Message: e.Msg(),
	}

}

// Error return a error string
func (e Error) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", e.Code(), e.Msg())
}

// Code return error code
func (e *Error) Code() int {
	return e.ErrCode
}

// Msg return error msg
func (e *Error) Msg() string {
	return e.Message
}

// Msgf format error string
func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.Message, args...)
}

// SetHTTPStatusCode set a specific http status code to err
func SetHTTPStatusCode(err *Error, status int) {
	toStatus.Store(err.Code(), status)
}

// ToHTTPStatusCode convert custom error code to http status code and avoid return unknown status code.
func ToHTTPStatusCode(code int) int {
	if status, ok := toStatus.Load(code); ok {
		return status.(int)
	}

	return http.StatusBadRequest
}

// DecodeErr 对错误进行解码，返回错误code和错误提示
func DecodeErr(err error) (int, string) {
	if err == nil {
		return Success.Code(), Success.Msg()
	}

	switch typed := err.(type) {
	case *Error:
		return typed.Code(), typed.Msg()
	default:
	}

	return ErrInternalServer.Code(), err.Error()
}

func initToStatus() {
	for code, status := range map[int]int{
		Success.Code():               http.StatusOK,
		ErrInternalServer.Code():     http.StatusInternalServerError,
		ErrNotFound.Code():           http.StatusNotFound,
		ErrInvalidParam.Code():       http.StatusBadRequest,
		ErrToken.Code():              http.StatusUnauthorized,
		ErrInvalidToken.Code():       http.StatusUnauthorized,
		ErrTokenTimeout.Code():       http.StatusUnauthorized,
		ErrTooManyRequests.Code():    http.StatusTooManyRequests,
		ErrServiceUnavailable.Code(): http.StatusServiceUnavailable,
	} {
		toStatus.Store(code, status)
	}
}

func init() {
	initToStatus()
}

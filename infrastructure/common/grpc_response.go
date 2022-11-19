package common

import (
	"github.com/go-eagle/eagle/infrastructure/common/errcode"
	"reflect"
)

// grpc resp must have retcode and message field

type GrpcResp interface {
	GetRetcode() int32
	GetMessage() string
}

func ErrGrpcRespWrap(resp GrpcResp, err *errcode.Error) GrpcResp {
	val := reflect.ValueOf(resp).Elem()
	retcode := val.FieldByName("Retcode")
	if retcode.CanSet() {
		retcode.SetInt(int64(err.Code()))
	}

	message := val.FieldByName("Message")
	if message.CanSet() {
		message.SetString(err.Message)
	}
	return resp
}

func SuccessGrpcRespWrap(resp GrpcResp) GrpcResp {
	val := reflect.ValueOf(resp).Elem()
	retcode := val.FieldByName("Retcode")
	if retcode.CanSet() {
		retcode.SetInt(0)
	}

	message := val.FieldByName("Message")
	if message.CanSet() {
		message.SetString("Success")
	}
	return resp
}

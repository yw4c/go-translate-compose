package grpc

import "fmt"

const(
	LoginFail = 1
	ParamRequire = 2
)

var statusText = map[int32]string {
	LoginFail: "帳號或密碼錯誤",
	ParamRequire: "參數 %s 不存在",
}

//impl Error
type GrpcError struct {
	Code int32
	Message string
}

func (e *GrpcError) Error() string{
	return e.Message
}

func NewGrpcError(code int32, param... string) *GrpcError {

	convered := make([]interface{}, len(param))
	for i, v := range param {
		convered[i] = v
	}
	message := fmt.Sprintf(statusText[code], convered...)

	return &GrpcError{
		Code: code,
		Message: message,
	}
}
package errors

import "errors"

// 设计风格：matrix-batch 的发生的任何一个错误都被全局映射为一个 Go error
// 如果需要传递到客户端，则会再进行 error -> statusCode 的映射
var (
	ErrUnknown  = errors.New("unknown internal error")
	ErrOperated = errors.New("operated")
	ErrUnknownState = errors.New("unknown state")
	ErrUserNotFound = errors.New("user not found")
)

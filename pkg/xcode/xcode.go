package xcode

import "strconv"

type XCode interface {
	Error() string
	Code() int
	Message() string
	Details() []interface{}
}

type Code struct {
	code int
	msg  string
}

func (c Code) Error() string {
	if len(c.msg) > 0 {
		return c.msg //返回具体错误信息
	}

	return strconv.Itoa(c.code) //没有信息就返回状态码
}

func (c Code) Code() int {
	return c.code
}

func (c Code) Message() string {
	return c.Error()
}

func (c Code) Details() []interface{} {
	return nil
}

func add(code int, msg string) Code {
	return Code{code: code, msg: msg}
}
func New(code int, msg string) Code {
	return Code{code: code, msg: msg}
}
func String(s string) Code {
	if len(s) == 0 {
		return OK
	}
	code, err := strconv.Atoi(s)
	if err != nil {
		return ServerErr //不能转化直接返回服务错误(500)
	}

	return Code{code: code}
}

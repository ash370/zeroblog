package code

import "zeroblog/pkg/xcode"

var (
	RegisterMobileEmpty     = xcode.New(10001, "注册手机号不能为空")
	VerificationCodeEmpty   = xcode.New(10002, "验证码不能为空")
	MobileHasRegistered     = xcode.New(10003, "手机号已经注册")
	LoginMobileEmpty        = xcode.New(10004, "手机号不能为空")
	RegisterPasswdEmpty     = xcode.New(10005, "密码不能为空")
	VerificationCodeExpired = xcode.New(10006, "验证码已过期")
	VerificationCodeFailed  = xcode.New(10007, "验证码错误")
)

package errcode

var (
	Success                   = NewError(200, "成功")
	ServerError               = NewError(1000, "服务内部错误")
	InvalidParams             = NewError(1001, "入参错误")
	NotFound                  = NewError(1002, "找不到")
	UnauthorizedAuthNotExist  = NewError(1003, "鉴权失败，账号验证失败")
	UnauthorizedTokenError    = NewError(1004, "鉴权失败，令牌错误")
	UnauthorizedTokenTimeout  = NewError(1005, "鉴权失败，令牌超时")
	UnauthorizedTokenGenerate = NewError(1006, "鉴权失败，令牌生成失败")
	TooManyRequests           = NewError(1007, "请求过多")
	IsSignup                  = NewError(1008, "账号已注册")
	NotSignup                 = NewError(1009, "账号未注册")
)

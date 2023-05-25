package errorx

// 参数错误
const paramsError = 10001
const cacheError = 10002

// 用户不存在
const userNotFound = 20001

// 用户已存在
const userExist = 20002

// 密码错误
const passwordError = 20003

var (
	ParamsError  = NewCodeError(paramsError, "参数错误")
	CacheError   = NewCodeError(cacheError, "缓存错误")
	UserNotFound = NewCodeError(userNotFound, "用户不存在")
	UserExist    = NewCodeError(userExist, "用户已存在")

	PasswordError = NewCodeError(passwordError, "密码错误")
)

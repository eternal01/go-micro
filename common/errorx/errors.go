package errorx

// 参数错误
const paramsError = 10001

// 缓存错误
const cacheError = 10002

// 记录不存在
const notFoundError = 10003

// 存在子级记录
const existChildrenError = 10004

// 用户不存在
const userNotFoundError = 20001

// 用户已存在
const userExistError = 20002

// 密码错误
const passwordError = 20003

var (
	ParamsError        = NewCodeError(paramsError, "参数错误")
	CacheError         = NewCodeError(cacheError, "缓存错误")
	NotFoundError      = NewCodeError(notFoundError, "记录不存在")
	ExistChildrenError = NewCodeError(existChildrenError, "存在子级记录")

	UserNotFoundError = NewCodeError(userNotFoundError, "用户不存在")
	UserExistError    = NewCodeError(userExistError, "用户已存在")

	PasswordError = NewCodeError(passwordError, "密码错误")
)

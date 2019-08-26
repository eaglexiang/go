package user

// 账户类型，PrivateUser的同时登录地限制为1，而SharedUser则没有限制
const (
	SharedUser  = iota // 0 表示不限制登录
	PrivateUser        // 1 表示只允许一个登录地
)

type user struct {
	ID       string
	Password string
}

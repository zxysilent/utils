package utils

// NewSucc 返回一个成功标识的结果格式
// 推荐使用 Succ
func NewSucc(msg string, data ...interface{}) (int, Reply) {
	return newReply(stSucc, msg, data...)
}

// NewFail 返回一个失败标识的结果格式
// 推荐使用 Fail
func NewFail(msg string, data ...interface{}) (int, Reply) {
	return newReply(stFail, msg, data...)
}

// NewPage 返回一个带有分页数据的结果格式
// 推荐使用 Page
func NewPage(msg string, items interface{}, count int) (int, Reply) {
	return 200, Reply{
		Code: stSucc,
		Msg:  msg,
		Data: page{
			Items: items,
			Count: count,
		},
	}
}

// NewErrIpt 返回一个输入错误的结果格式
// 推荐使用 ErrIpt
func NewErrIpt(msg string, data ...interface{}) (int, Reply) {
	return newReply(stErrIpt, msg, data...)
}

// NewErrOpt 返回一个输出错误的结果格式
// 推荐使用 ErrOpt
func NewErrOpt(msg string, data ...interface{}) (int, Reply) {
	return newReply(stErrOpt, msg, data...)
}

// NewErrDeny 返回一个没有权限的结果格式
// 推荐使用 ErrDeny
func NewErrDeny(msg string, data ...interface{}) (int, Reply) {
	return newReply(stErrDeny, msg, data...)
}

// NewErrJwt 返回一个通过验证的结果格式
// 推荐使用 ErrJwt
func NewErrJwt(msg string, data ...interface{}) (int, Reply) {
	return newReply(stErrJwt, msg, data...)
}

// NewErrSvr 返回一个服务端错误的结果格式
func NewErrSvr(msg string, data ...interface{}) (int, Reply) {
	return newReply(stErrSvr, msg, data...)
}

// NewExt 返回一个其他约定的结果格式
func NewExt(msg string, data ...interface{}) (int, Reply) {
	return newReply(stExt, msg, data...)
}

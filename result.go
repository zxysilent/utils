package utils

// //Reply  format
// type Reply struct {
// 	Code int         `json:"code"`
// 	Msg  string      `json:"msg"`
// 	Data interface{} `json:"data,omitempty"`
// }

// //page format
// //Message
// type page struct {
// 	Count int         `json:"count"`
// 	Items interface{} `json:"items"`
// }

// const (
// 	stSucc    int = 200 //正常
// 	stFail    int = 300 //失败
// 	stErrIpt  int = 310 //输入数据有误
// 	stErrOpt  int = 320 //无数据返回
// 	stErrDeny int = 330 //没有权限
// 	stErrJwt  int = 340 //jwt未通过验证
// 	stErrSvr  int = 350 //服务端错误
// 	stExt     int = 400 //其他约定 //eg 更新 token
// )

// func newRes(code int, msg string, data ...interface{}) (int, Reply) {
// 	if len(data) > 0 {
// 		return 200, Reply{
// 			Code: code,
// 			Msg:  msg,
// 			Data: data[0],
// 		}
// 	}
// 	return 200, Reply{
// 		Code: code,
// 		Msg:  msg,
// 	}
// }

//NewSucc 返回一个成功标识的结果格式
func NewSucc(msg string, data ...interface{}) (int, Reply) {
	return newRes(stSucc, msg, data...)
}

//NewFail 返回一个失败标识的结果格式
func NewFail(msg string, data ...interface{}) (int, Reply) {
	return newRes(stFail, msg, data...)
}

//NewPage 返回一个带有分页数据的结果格式
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

//NewErrIpt 返回一个输入错误的结果格式
func NewErrIpt(msg string, data ...interface{}) (int, Reply) {
	return newRes(stErrIpt, msg, data...)
}

//NewErrOpt 返回一个输出错误的结果格式
func NewErrOpt(msg string, data ...interface{}) (int, Reply) {
	return newRes(stErrOpt, msg, data...)
}

//NewErrDeny 返回一个没有权限的结果格式
func NewErrDeny(msg string, data ...interface{}) (int, Reply) {
	return newRes(stErrDeny, msg, data...)
}

//NewErrJwt 返回一个通过验证的结果格式
func NewErrJwt(msg string, data ...interface{}) (int, Reply) {
	return newRes(stErrJwt, msg, data...)
}

//NewErrSvr 返回一个服务端错误的结果格式
func NewErrSvr(msg string, data ...interface{}) (int, Reply) {
	return newRes(stErrSvr, msg, data...)
}

//NewExt 返回一个其他约定的结果格式
func NewExt(msg string, data ...interface{}) (int, Reply) {
	return newRes(stExt, msg, data...)
}

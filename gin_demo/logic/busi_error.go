package logic

import "gin_demo/gin_util"

var (
	ErrorGetInfo = &ginUtil.Error{UserMsg: "input param is nil", UserErrCode: 30001}
)

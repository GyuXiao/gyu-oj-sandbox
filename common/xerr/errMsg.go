package xerr

var mapCodMsg map[uint32]string

func init() {
	mapCodMsg = make(map[uint32]string)

	// 全局错误码
	mapCodMsg[SUCCESS] = "success"
	mapCodMsg[ERROR] = "error"
	mapCodMsg[UnknownError] = "未知错误"
	mapCodMsg[ServerCommonError] = "服务器开小差啦,稍后再来试一试"
	mapCodMsg[ParamFormatError] = "参数格式错误"
	mapCodMsg[RequestParamError] = "参数缺失或不规范"
	mapCodMsg[UnauthorizedError] = "鉴权失败错误"

	// sandbox
	mapCodMsg[CompileFailError] = "编译代码错误"
	mapCodMsg[RunFailError] = "运行代码错误"
	mapCodMsg[RunTimeoutError] = "运行代码时间超出限制错误"
	mapCodMsg[RunOutOfMemoryError] = "运行代码内存存储限制错误"
	mapCodMsg[SandboxError] = "代码沙箱系统错误"
}

func GetMsgByCode(errCode uint32) string {
	if msg, ok := mapCodMsg[errCode]; ok {
		return msg
	}
	return "服务器开小差啦,稍后再来试一试"
}

func IsCodeErr(errCode uint32) bool {
	if _, ok := mapCodMsg[errCode]; ok {
		return true
	}
	return false
}

package xerr

const (
	SUCCESS           uint32 = 0
	ERROR             uint32 = 1
	UnknownError      uint32 = 100000
	ServerCommonError uint32 = 100001
	ParamFormatError  uint32 = 100002
	RequestParamError uint32 = 100003
	UnauthorizedError uint32 = 100004
)

// sandbox
const (
	CompileFailError    uint32 = 110001
	RunFailError        uint32 = 110002
	RunTimeoutError     uint32 = 110003
	RunOutOfMemoryError uint32 = 110004
	SandboxError        uint32 = 110005
)

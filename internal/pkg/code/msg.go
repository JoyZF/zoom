package code

const (
	defaultMsg = "unknown error"
	ZH_CN      = "zh-cn"
	EN_US      = "en-us"
)

type CodeMsg struct {
	Code int
	Msg  string
}

// GetMsg return error msg
func GetMsg(code int, lang string) string {
	if codeMsg, ok := CodeMap[code]; ok {
		if msg, ok := codeMsg[lang]; ok {
			return msg.Msg
		} else {
			return defaultMsg
		}
	} else {
		return defaultMsg
	}
}

var CodeMap = map[int]map[string]CodeMsg{
	Success:     {"zh-cn": CodeMsg{Code: Success, Msg: "成功"}, "en-us": CodeMsg{Code: Success, Msg: "Success"}},
	SystemError: {"zh-cn": CodeMsg{Code: SystemError, Msg: "系统错误"}, "en-us": CodeMsg{Code: Success, Msg: "Success"}},
	NotFound:    {"zh-cn": CodeMsg{Code: NotFound, Msg: "未找到"}, "en-us": CodeMsg{Code: Success, Msg: "Success"}},
}

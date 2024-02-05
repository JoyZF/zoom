// Copyright 2024 Joy <joyssss94@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package code

import "github.com/JoyZF/zoom/global"

const (
	defaultMsg = "unknown error"
)

type CodeMsg struct {
	C   int
	Msg string
}

// GetMsg return error msg
func GetMsg(code int, lang string) string {
	if codeMsg, ok := CodeMap[code]; ok {
		if msg, ok := codeMsg[lang]; ok {
			return msg.Msg
		} else {
			return codeMsg[global.EN_US].Msg
		}
	} else {
		return defaultMsg
	}
}

func (c CodeMsg) HTTPStatus() int {
	if c.C > 500 {
		return 500
	}
	return c.C
}

func (c CodeMsg) String() string {
	return c.Msg
}

func (c CodeMsg) Reference() string {
	return "https://github.com/JoyZF/zoom"
}

func (c CodeMsg) Code() int {
	return c.C
}

var CodeMap = map[int]map[string]CodeMsg{
	Success:                 {global.ZH_CN: CodeMsg{C: Success, Msg: "成功"}, global.EN_US: CodeMsg{C: Success, Msg: "success"}},
	SystemError:             {global.ZH_CN: CodeMsg{C: SystemError, Msg: "系统错误"}, global.EN_US: CodeMsg{C: SystemError, Msg: "system error"}},
	NotFound:                {global.ZH_CN: CodeMsg{C: NotFound, Msg: "未找到"}, global.EN_US: CodeMsg{C: NotFound, Msg: "not found"}},
	ParamsError:             {global.ZH_CN: CodeMsg{C: ParamsError, Msg: "参数错误"}, global.EN_US: CodeMsg{C: ParamsError, Msg: "params error"}},
	GenericServiceErrorCode: {global.ZH_CN: CodeMsg{C: GenericServiceErrorCode, Msg: "Generic service error code"}, global.EN_US: CodeMsg{C: GenericServiceErrorCode, Msg: "Generic service error code"}},
}

// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package response

import (
	"net/http"

	"github.com/JoyZF/errors"
	"github.com/JoyZF/zlog"

	"github.com/JoyZF/zoom/global"
	"github.com/JoyZF/zoom/internal/pkg/code"

	"github.com/gin-gonic/gin"
)

// ErrResponse defines the return messages when an error occurred.
// Reference will be omitted if it does not exist.
// swagger:model
type ErrResponse struct {
	// Code defines the business error code.
	Code int `json:"code"`

	// Message contains the detail of this message.
	// This message is suitable to be exposed to external
	Message string `json:"message"`

	// Reference returns the reference document which maybe useful to solve this error.
	Reference string `json:"reference,omitempty"`
}

type SuccessResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type errDecode struct {
	Caller  string `json:"caller"`
	Code    int    `json:"code"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

// WriteResponse write an error or the response data into http response body.
// It use errors.ParseCoder to parse any error into errors.Coder
// errors.Coder contains error code, user-safe error message and http status code.
func WriteResponse(c *gin.Context, err error, data interface{}) {
	if err != nil {
		zlog.Errorf("%#+v", err)
		var message string
		coder := errors.ParseCoder(err)
		if err.Error() != code.GetMsg(coder.Code(), c.GetHeader(global.LangKey)) {
			message = err.Error()
		} else {
			message = code.GetMsg(coder.Code(), c.GetHeader(global.LangKey))
		}
		c.JSON(coder.HTTPStatus(), ErrResponse{
			Code:      coder.Code(),
			Message:   message,
			Reference: coder.Reference(),
		})

		return
	}
	if data == nil {
		data = ""
	}
	c.JSON(http.StatusOK, SuccessResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

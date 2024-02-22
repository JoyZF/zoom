// Copyright 2024 Joy <joyssss94@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package code

import (
	"fmt"
	"github.com/JoyZF/zoom/global"
	"testing"

	"github.com/JoyZF/errors"
)

func TestGetMsg(t *testing.T) {
	err := errors.WithCode(SystemError, GetMsg(SystemError, global.EN_US))
	fmt.Println(fmt.Sprintf("%+v", err))
}

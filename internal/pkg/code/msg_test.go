package code

import (
	"fmt"
	"github.com/JoyZF/errors"
	"testing"
)

func TestGetMsg(t *testing.T) {
	err := errors.WithCode(SystemError, GetMsg(SystemError, EN_US))
	fmt.Println(fmt.Sprintf("%+v", err))
}

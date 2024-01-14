package code

import (
	"context"
	"github.com/JoyZF/errors"
)

const lang = "lang"

// ErrorWithCode return error with code
func ErrorWithCode(ctx context.Context, code int) error {
	return errors.WithCode(code, GetMsg(code, getLangFromCtx(ctx)))
}

func getLangFromCtx(ctx context.Context) string {
	value := ctx.Value(lang)
	if value == nil {
		return ZH_CN
	}
	if l, ok := value.(string); ok {
		return l
	}
	return ZH_CN

}

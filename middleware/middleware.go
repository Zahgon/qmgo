package middleware

import (
	"context"

	"github.com/qiniu/qmgo/field"
	"github.com/qiniu/qmgo/hook"
	"github.com/qiniu/qmgo/operator"
	"github.com/qiniu/qmgo/validator"
)

// callback define the callback function type
type callback func(ctx context.Context, doc interface{}, opType operator.OpType, opts ...interface{}) error

// middlewareCallback the register callback slice
// some callbacks initial here without Register() for order
var middlewareCallback = []callback{
	hook.Do,
	field.Do,
	validator.Do,
}

// Register register callback into middleware
func Register(cb callback) { _ = "STUB: not implemented"; return }

// Do call every registers
// The doc is always the document to operate
func Do(ctx context.Context, content interface{}, opType operator.OpType, opts ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

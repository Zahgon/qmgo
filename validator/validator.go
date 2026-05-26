package validator

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/qiniu/qmgo/operator"
)

// use a single instance of Validate, it caches struct info
var validate = validator.New()

// SetValidate let validate can use custom rules
func SetValidate(v *validator.Validate) {
	_ = "STUB: not implemented"

	// validatorNeeded checks if the validator is needed to opType
	return
}

func validatorNeeded(opType operator.OpType) bool { _ = "STUB: not implemented"; return false }

// Do calls validator check
// Don't use opts here
func Do(ctx context.Context, doc interface{}, opType operator.OpType, opts ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// sliceHandle handles the slice docs
func sliceHandle(docs interface{}, opType operator.OpType) error {
	_ = "STUB: not implemented"
	// []interface{}{UserType{}...}
	return nil
}

// []UserType{}

// do check if opType is supported and call fieldHandler
func do(doc interface{}) error { _ = "STUB: not implemented"; return nil }

// validatorStruct check if kind of doc is validator supported struct
// same implement as validator
func validatorStruct(doc interface{}) bool { _ = "STUB: not implemented"; return false }

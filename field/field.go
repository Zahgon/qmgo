/*
 Copyright 2020 The Qmgo Authors.
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
     http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package field

import (
	"context"
	"time"

	"github.com/qiniu/qmgo/operator"
)

var nilTime time.Time

// filedHandler defines the relations between field type and handler
var fieldHandler = map[operator.OpType]func(doc interface{}) error{
	operator.BeforeInsert:  beforeInsert,
	operator.BeforeUpdate:  beforeUpdate,
	operator.BeforeReplace: beforeUpdate,
	operator.BeforeUpsert:  beforeUpsert,
}

//func init() {
//	middleware.Register(Do)
//}

// Do call the specific method to handle field based on fType
// Don't use opts here
func Do(ctx context.Context, doc interface{}, opType operator.OpType, opts ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

//fmt.Println("not support type")

// sliceHandle handles the slice docs
func sliceHandle(docs interface{}, opType operator.OpType) error {
	_ = "STUB: not implemented"
	// []interface{}{UserType{}...}
	return nil
}

// []UserType{}

// beforeInsert handles field before insert
// If value of field createAt is valid in doc, upsert doesn't change it
// If value of field id is valid in doc, upsert doesn't change it
// Change the value of field updateAt anyway
func beforeInsert(doc interface{}) error { _ = "STUB: not implemented"; return nil }

// beforeUpdate handles field before update
func beforeUpdate(doc interface{}) error { _ = "STUB: not implemented"; return nil }

// beforeUpsert handles field before upsert
// If value of field createAt is valid in doc, upsert doesn't change it
// If value of field id is valid in doc, upsert doesn't change it
// Change the value of field updateAt anyway
func beforeUpsert(doc interface{}) error { _ = "STUB: not implemented"; return nil }

// do check if opType is supported and call fieldHandler
func do(doc interface{}, opType operator.OpType) error { _ = "STUB: not implemented"; return nil }

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

package hook

import (
	"context"

	"github.com/qiniu/qmgo/operator"
)

// hookHandler defines the relations between hook type and handler
var hookHandler = map[operator.OpType]func(ctx context.Context, hook interface{}) error{
	operator.BeforeInsert:  beforeInsert,
	operator.AfterInsert:   afterInsert,
	operator.BeforeUpdate:  beforeUpdate,
	operator.AfterUpdate:   afterUpdate,
	operator.BeforeQuery:   beforeQuery,
	operator.AfterQuery:    afterQuery,
	operator.BeforeRemove:  beforeRemove,
	operator.AfterRemove:   afterRemove,
	operator.BeforeUpsert:  beforeUpsert,
	operator.AfterUpsert:   afterUpsert,
	operator.BeforeReplace: beforeUpdate,
	operator.AfterReplace:  afterUpdate,
}

//
//func init() {
//	middleware.Register(Do)
//}

// Do call the specific method to handle hook based on hType
// If opts has valid value, use it instead of original hook
func Do(ctx context.Context, hook interface{}, opType operator.OpType, opts ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// sliceHandle handles the slice hooks
func sliceHandle(ctx context.Context, hook interface{}, opType operator.OpType) error {
	_ = "STUB: not implemented"
	// []interface{}{UserType{}...}
	return nil
}

// []UserType{}

// BeforeInsertHook InsertHook defines the insert hook interface
type BeforeInsertHook interface {
	BeforeInsert(ctx context.Context) error
}
type AfterInsertHook interface {
	AfterInsert(ctx context.Context) error
}

// beforeInsert calls custom BeforeInsert
func beforeInsert(ctx context.Context, hook interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// afterInsert calls custom AfterInsert
func afterInsert(ctx context.Context, hook interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// BeforeUpdateHook defines the Update hook interface
type BeforeUpdateHook interface {
	BeforeUpdate(ctx context.Context) error
}
type AfterUpdateHook interface {
	AfterUpdate(ctx context.Context) error
}

// beforeUpdate calls custom BeforeUpdate
func beforeUpdate(ctx context.Context, hook interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// afterUpdate calls custom AfterUpdate
func afterUpdate(ctx context.Context, hook interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// BeforeQueryHook QueryHook defines the query hook interface
type BeforeQueryHook interface {
	BeforeQuery(ctx context.Context) error
}
type AfterQueryHook interface {
	AfterQuery(ctx context.Context) error
}

// beforeQuery calls custom BeforeQuery
func beforeQuery(ctx context.Context, hook interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// afterQuery calls custom AfterQuery
func afterQuery(ctx context.Context, hook interface{}) error { _ = "STUB: not implemented"; return nil }

// BeforeRemoveHook RemoveHook defines the remove hook interface
type BeforeRemoveHook interface {
	BeforeRemove(ctx context.Context) error
}
type AfterRemoveHook interface {
	AfterRemove(ctx context.Context) error
}

// beforeRemove calls custom BeforeRemove
func beforeRemove(ctx context.Context, hook interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// afterRemove calls custom AfterRemove
func afterRemove(ctx context.Context, hook interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// BeforeUpsertHook UpsertHook defines the upsert hook interface
type BeforeUpsertHook interface {
	BeforeUpsert(ctx context.Context) error
}
type AfterUpsertHook interface {
	AfterUpsert(ctx context.Context) error
}

// beforeUpsert calls custom BeforeUpsert
func beforeUpsert(ctx context.Context, hook interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// afterUpsert calls custom AfterUpsert
func afterUpsert(ctx context.Context, hook interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// do check if opType is supported and call hookHandler
func do(ctx context.Context, hook interface{}, opType operator.OpType) error {
	_ = "STUB: not implemented"
	return nil
}

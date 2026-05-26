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

package qmgo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// BulkResult is the result type returned by Bulk.Run operation.
type BulkResult struct {
	// The number of documents inserted.
	InsertedCount int64

	// The number of documents matched by filters in update and replace operations.
	MatchedCount int64

	// The number of documents modified by update and replace operations.
	ModifiedCount int64

	// The number of documents deleted.
	DeletedCount int64

	// The number of documents upserted by update and replace operations.
	UpsertedCount int64

	// A map of operation index to the _id of each upserted document.
	UpsertedIDs map[int64]interface{}
}

// Bulk is context for batching operations to be sent to database in a single
// bulk write.
//
// Bulk is not safe for concurrent use.
//
// Notes:
//
// Individual operations inside a bulk do not trigger middlewares or hooks
// at present.
//
// Different from original mgo, the qmgo implementation of Bulk does not emulate
// bulk operations individually on old versions of MongoDB servers that do not
// natively support bulk operations.
//
// Only operations supported by the official driver are exposed, that is why
// InsertMany is missing from the methods.
type Bulk struct {
	coll *Collection

	queue   []mongo.WriteModel
	ordered *bool
}

// Bulk returns a new context for preparing bulk execution of operations.
func (c *Collection) Bulk() *Bulk { _ = "STUB: not implemented"; return nil }

// SetOrdered marks the bulk as ordered or unordered.
//
// If ordered, writes does not continue after one individual write fails.
// Default is ordered.
func (b *Bulk) SetOrdered(ordered bool) *Bulk { _ = "STUB: not implemented"; return nil }

// InsertOne queues an InsertOne operation for bulk execution.
func (b *Bulk) InsertOne(doc interface{}) *Bulk { _ = "STUB: not implemented"; return nil }

// Remove queues a Remove operation for bulk execution.
func (b *Bulk) Remove(filter interface{}) *Bulk { _ = "STUB: not implemented"; return nil }

// RemoveId queues a RemoveId operation for bulk execution.
func (b *Bulk) RemoveId(id interface{}) *Bulk { _ = "STUB: not implemented"; return nil }

// RemoveAll queues a RemoveAll operation for bulk execution.
func (b *Bulk) RemoveAll(filter interface{}) *Bulk { _ = "STUB: not implemented"; return nil }

// Upsert queues an Upsert operation for bulk execution.
// The replacement should be document without operator
func (b *Bulk) Upsert(filter interface{}, replacement interface{}) *Bulk {
	_ = "STUB: not implemented"
	return nil
}

// UpsertOne queues an UpsertOne operation for bulk execution.
// The update should contain operator
func (b *Bulk) UpsertOne(filter interface{}, update interface{}) *Bulk {
	_ = "STUB: not implemented"
	return nil
}

// UpsertId queues an UpsertId operation for bulk execution.
// The replacement should be document without operator
func (b *Bulk) UpsertId(id interface{}, replacement interface{}) *Bulk {
	_ = "STUB: not implemented"
	return nil
}

// UpdateOne queues an UpdateOne operation for bulk execution.
// The update should contain operator
func (b *Bulk) UpdateOne(filter interface{}, update interface{}) *Bulk {
	_ = "STUB: not implemented"
	return nil
}

// UpdateId queues an UpdateId operation for bulk execution.
// The update should contain operator
func (b *Bulk) UpdateId(id interface{}, update interface{}) *Bulk {
	_ = "STUB: not implemented"
	return nil
}

// UpdateAll queues an UpdateAll operation for bulk execution.
// The update should contain operator
func (b *Bulk) UpdateAll(filter interface{}, update interface{}) *Bulk {
	_ = "STUB: not implemented"
	return nil
}

// Run executes the collected operations in a single bulk operation.
//
// A successful call resets the Bulk. If an error is returned, the internal
// queue of operations is unchanged, containing both successful and failed
// operations.
func (b *Bulk) Run(ctx context.Context) (*BulkResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// In original mgo, queue is not reset in case of error.

// Empty the queue for possible reuse, as per mgo's behavior.

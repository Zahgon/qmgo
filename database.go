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

	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/mongo"
	officialOpts "go.mongodb.org/mongo-driver/mongo/options"
)

// Database is a handle to a MongoDB database
type Database struct {
	database *mongo.Database

	registry *bsoncodec.Registry
}

// Collection gets collection from database
func (d *Database) Collection(name string, opts ...*options.CollectionOptions) *Collection {
	_ = "STUB: not implemented"
	return nil
}

// ListCollections lists all collections in the database.
func (d *Database) ListCollections(ctx context.Context, filter interface{}, opts ...*officialOpts.ListCollectionsOptions) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetDatabaseName returns the name of database
func (d *Database) GetDatabaseName() string { _ = "STUB: not implemented"; return "" }

// DropDatabase drops database
func (d *Database) DropDatabase(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// RunCommand executes the given command against the database.
//
// The runCommand parameter must be a document for the command to be executed. It cannot be nil.
// This must be an order-preserving type such as bson.D. Map types such as bson.M are not valid.
// If the command document contains a session ID or any transaction-specific fields, the behavior is undefined.
//
// The opts parameter can be used to specify options for this operation (see the options.RunCmdOptions documentation).
func (d *Database) RunCommand(ctx context.Context, runCommand interface{}, opts ...options.RunCommandOptions) *mongo.SingleResult {
	_ = "STUB: not implemented"
	return nil
}

// CreateCollection executes a create command to explicitly create a new collection with the specified name on the
// server. If the collection being created already exists, this method will return a mongo.CommandError. This method
// requires driver version 1.4.0 or higher.
//
// The opts parameter can be used to specify options for the operation (see the options.CreateCollectionOptions
// documentation).
func (db *Database) CreateCollection(ctx context.Context, name string, opts ...options.CreateCollectionOptions) error {
	_ = "STUB: not implemented"
	return nil
}

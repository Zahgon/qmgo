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
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Config for initial mongodb instance
type Config struct {
	// URI example: [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
	// URI Reference: https://docs.mongodb.com/manual/reference/connection-string/
	Uri      string `json:"uri"`
	Database string `json:"database"`
	Coll     string `json:"coll"`
	// ConnectTimeoutMS specifies a timeout that is used for creating connections to the server.
	//	If set to 0, no timeout will be used.
	//	The default is 30 seconds.
	ConnectTimeoutMS *int64 `json:"connectTimeoutMS"`
	// MaxPoolSize specifies that maximum number of connections allowed in the driver's connection pool to each server.
	// If this is 0, it will be set to math.MaxInt64,
	// The default is 100.
	MaxPoolSize *uint64 `json:"maxPoolSize"`
	// MinPoolSize specifies the minimum number of connections allowed in the driver's connection pool to each server. If
	// this is non-zero, each server's pool will be maintained in the background to ensure that the size does not fall below
	// the minimum. This can also be set through the "minPoolSize" URI option (e.g. "minPoolSize=100"). The default is 0.
	MinPoolSize *uint64 `json:"minPoolSize"`
	// SocketTimeoutMS specifies how long the driver will wait for a socket read or write to return before returning a
	// network error. If this is 0 meaning no timeout is used and socket operations can block indefinitely.
	// The default is 300,000 ms.
	SocketTimeoutMS *int64 `json:"socketTimeoutMS"`
	// ReadPreference determines which servers are considered suitable for read operations.
	// default is PrimaryMode
	ReadPreference *ReadPref `json:"readPreference"`
	// can be used to provide authentication options when configuring a Client.
	Auth *Credential `json:"auth"`
}

// Credential can be used to provide authentication options when configuring a Client.
//
// AuthMechanism: the mechanism to use for authentication. Supported values include "SCRAM-SHA-256", "SCRAM-SHA-1",
// "MONGODB-CR", "PLAIN", "GSSAPI", "MONGODB-X509", and "MONGODB-AWS". This can also be set through the "authMechanism"
// URI option. (e.g. "authMechanism=PLAIN"). For more information, see
// https://docs.mongodb.com/manual/core/authentication-mechanisms/.
// AuthSource: the name of the database to use for authentication. This defaults to "$external" for MONGODB-X509,
// GSSAPI, and PLAIN and "admin" for all other mechanisms. This can also be set through the "authSource" URI option
// (e.g. "authSource=otherDb").
//
// Username: the username for authentication. This can also be set through the URI as a username:password pair before
// the first @ character. For example, a URI for user "user", password "pwd", and host "localhost:27017" would be
// "mongodb://user:pwd@localhost:27017". This is optional for X509 authentication and will be extracted from the
// client certificate if not specified.
//
// Password: the password for authentication. This must not be specified for X509 and is optional for GSSAPI
// authentication.
//
// PasswordSet: For GSSAPI, this must be true if a password is specified, even if the password is the empty string, and
// false if no password is specified, indicating that the password should be taken from the context of the running
// process. For other mechanisms, this field is ignored.
type Credential struct {
	AuthMechanism string `json:"authMechanism"`
	AuthSource    string `json:"authSource"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	PasswordSet   bool   `json:"passwordSet"`
}

// ReadPref determines which servers are considered suitable for read operations.
type ReadPref struct {
	// MaxStaleness is the maximum amount of time to allow a server to be considered eligible for selection.
	// Supported from version 3.4.
	MaxStalenessMS int64 `json:"maxStalenessMS"`
	// indicates the user's preference on reads.
	// PrimaryMode as default
	Mode readpref.Mode `json:"mode"`
}

// QmgoClient specifies the instance to operate mongoDB
type QmgoClient struct {
	*Collection
	*Database
	*Client
}

// Open creates client instance according to config
// QmgoClient can operates all qmgo.client 、qmgo.database and qmgo.collection
func Open(ctx context.Context, conf *Config, o ...options.ClientOptions) (cli *QmgoClient, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Client creates client to mongo
type Client struct {
	client *mongo.Client
	conf   Config

	registry *bsoncodec.Registry
}

// NewClient creates Qmgo MongoDB client
func NewClient(ctx context.Context, conf *Config, o ...options.ClientOptions) (cli *Client, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// client creates connection to MongoDB
func client(ctx context.Context, opt *officialOpts.ClientOptions) (client *mongo.Client, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// half of default connect timeout

// newConnectOpts creates client options from conf
// Qmgo will follow this way official mongodb driver do：
// - the configuration in uri takes precedence over the configuration in the setter
// - Check the validity of the configuration in the uri, while the configuration in the setter is basically not checked
func newConnectOpts(conf *Config, o ...options.ClientOptions) (*officialOpts.ClientOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// newAuth create options.Credential from conf.Auth
func newAuth(auth Credential) (credential officialOpts.Credential, err error) {
	_ = "STUB: not implemented"
	return *new(officialOpts.Credential), nil
}

// Validate and process the username.

// newReadPref create readpref.ReadPref from config
func newReadPref(pref ReadPref) (*readpref.ReadPref, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close closes sockets to the topology referenced by this Client.
func (c *Client) Close(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Ping confirm connection is alive
func (c *Client) Ping(timeout int64) error { _ = "STUB: not implemented"; return nil }

// Database create connection to database
func (c *Client) Database(name string, options ...*options.DatabaseOptions) *Database {
	_ = "STUB: not implemented"
	return nil
}

// Session create one session on client
// Watch out, close session after operation done
func (c *Client) Session(opt ...*options.SessionOptions) (*Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DoTransaction do whole transaction in one function
// precondition：
// - version of mongoDB server >= v4.0
// - Topology of mongoDB server is not Single
// At the same time, please pay attention to the following
//   - make sure all operations in callback use the sessCtx as context parameter
//   - if operations in callback takes more than(include equal) 120s, the operations will not take effect,
//   - if operation in callback return qmgo.ErrTransactionRetry,
//     the whole transaction will retry, so this transaction must be idempotent
//   - if operations in callback return qmgo.ErrTransactionNotSupported,
//   - If the ctx parameter already has a Session attached to it, it will be replaced by this session.
func (c *Client) DoTransaction(ctx context.Context, callback func(sessCtx context.Context) (interface{}, error), opts ...*options.TransactionOptions) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ServerVersion get the version of mongoDB server, like 4.4.0
func (c *Client) ServerVersion() string { _ = "STUB: not implemented"; return "" }

// transactionAllowed check if transaction is allowed
func (c *Client) transactionAllowed() bool { _ = "STUB: not implemented"; return false }

// TODO dont know why need to do `cli, err := Open(ctx, &c.conf)` in topology() to get topo,
// Before figure it out, we only use this function in UT
//topo, err := c.topology()
//if topo == description.Single {
//	fmt.Println("transaction is not supported because mongo server topology is single")
//	return false
//}

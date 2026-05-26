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
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Now return Millisecond current time
func Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// NewObjectID generates a new ObjectID.
// Watch out: the way it generates objectID is different from mgo
func NewObjectID() primitive.ObjectID { _ = "STUB: not implemented"; return *new(primitive.ObjectID) }

// SplitSortField handle sort symbol: "+"/"-" in front of field
// if "+"， return sort as 1
// if "-"， return sort as -1
func SplitSortField(field string) (key string, sort int32) { _ = "STUB: not implemented"; return "", 0 }

// CompareVersions compares two version number strings (i.e. positive integers separated by
// periods). Comparisons are done to the lesser precision of the two versions. For example, 3.2 is
// considered equal to 3.2.11, whereas 3.2.0 is considered less than 3.2.11.
//
// Returns a positive int if version1 is greater than version2, a negative int if version1 is less
// than version2, and 0 if version1 is equal to version2.
func CompareVersions(v1 string, v2 string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

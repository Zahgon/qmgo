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

// CustomFields defines struct of supported custom fields
type CustomFields struct {
	createAt string
	updateAt string
	id       string
}

// CustomFieldsHook defines the interface, CustomFields return custom field user want to change
type CustomFieldsHook interface {
	CustomFields() CustomFieldsBuilder
}

// CustomFieldsBuilder defines the interface which user use to set custom fields
type CustomFieldsBuilder interface {
	SetUpdateAt(fieldName string) CustomFieldsBuilder
	SetCreateAt(fieldName string) CustomFieldsBuilder
	SetId(fieldName string) CustomFieldsBuilder
}

// NewCustom creates new Builder which is used to set the custom fields
func NewCustom() CustomFieldsBuilder {
	_ = "STUB: not implemented"
	return *

	// SetUpdateAt set the custom UpdateAt field
	new(CustomFieldsBuilder)
}

func (c *CustomFields) SetUpdateAt(fieldName string) CustomFieldsBuilder {
	_ = "STUB: not implemented"
	return *new(CustomFieldsBuilder)
}

// SetCreateAt set the custom CreateAt field
func (c *CustomFields) SetCreateAt(fieldName string) CustomFieldsBuilder {
	_ = "STUB: not implemented"
	return *new(CustomFieldsBuilder)
}

// SetId set the custom Id field
func (c *CustomFields) SetId(fieldName string) CustomFieldsBuilder {
	_ = "STUB: not implemented"
	return *new(CustomFieldsBuilder)
}

// CustomCreateTime changes the custom create time
func (c CustomFields) CustomCreateTime(doc interface{}) { _ = "STUB: not implemented"; return }

// CustomUpdateTime changes the custom update time
func (c CustomFields) CustomUpdateTime(doc interface{}) { _ = "STUB: not implemented"; return }

// CustomUpdateTime changes the custom update time
func (c CustomFields) CustomId(doc interface{}) { _ = "STUB: not implemented"; return }

// setTime changes the custom time fields
// The overWrite defines if change value when the filed has valid value
func setTime(doc interface{}, fieldName string, overWrite bool) { _ = "STUB: not implemented"; return }

// setId changes the custom Id fields
func setId(doc interface{}, fieldName string) { _ = "STUB: not implemented"; return }

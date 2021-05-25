/*
 * Koyeb Rest API
 *
 * The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. 
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"encoding/json"
)

// ListFunctionsReply struct for ListFunctionsReply
type ListFunctionsReply struct {
	Functions *[]FunctionListItem `json:"functions,omitempty"`
}

// NewListFunctionsReply instantiates a new ListFunctionsReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListFunctionsReply() *ListFunctionsReply {
	this := ListFunctionsReply{}
	return &this
}

// NewListFunctionsReplyWithDefaults instantiates a new ListFunctionsReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListFunctionsReplyWithDefaults() *ListFunctionsReply {
	this := ListFunctionsReply{}
	return &this
}

// GetFunctions returns the Functions field value if set, zero value otherwise.
func (o *ListFunctionsReply) GetFunctions() []FunctionListItem {
	if o == nil || o.Functions == nil {
		var ret []FunctionListItem
		return ret
	}
	return *o.Functions
}

// GetFunctionsOk returns a tuple with the Functions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFunctionsReply) GetFunctionsOk() (*[]FunctionListItem, bool) {
	if o == nil || o.Functions == nil {
		return nil, false
	}
	return o.Functions, true
}

// HasFunctions returns a boolean if a field has been set.
func (o *ListFunctionsReply) HasFunctions() bool {
	if o != nil && o.Functions != nil {
		return true
	}

	return false
}

// SetFunctions gets a reference to the given []FunctionListItem and assigns it to the Functions field.
func (o *ListFunctionsReply) SetFunctions(v []FunctionListItem) {
	o.Functions = &v
}

func (o ListFunctionsReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Functions != nil {
		toSerialize["functions"] = o.Functions
	}
	return json.Marshal(toSerialize)
}

type NullableListFunctionsReply struct {
	value *ListFunctionsReply
	isSet bool
}

func (v NullableListFunctionsReply) Get() *ListFunctionsReply {
	return v.value
}

func (v *NullableListFunctionsReply) Set(val *ListFunctionsReply) {
	v.value = val
	v.isSet = true
}

func (v NullableListFunctionsReply) IsSet() bool {
	return v.isSet
}

func (v *NullableListFunctionsReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFunctionsReply(val *ListFunctionsReply) *NullableListFunctionsReply {
	return &NullableListFunctionsReply{value: val, isSet: true}
}

func (v NullableListFunctionsReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListFunctionsReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



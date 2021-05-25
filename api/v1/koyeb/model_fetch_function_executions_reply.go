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

// FetchFunctionExecutionsReply struct for FetchFunctionExecutionsReply
type FetchFunctionExecutionsReply struct {
	Executions *[]FunctionRunInfoListItem `json:"executions,omitempty"`
	Limit *int64 `json:"limit,omitempty"`
	Offset *int64 `json:"offset,omitempty"`
	Count *int64 `json:"count,omitempty"`
}

// NewFetchFunctionExecutionsReply instantiates a new FetchFunctionExecutionsReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchFunctionExecutionsReply() *FetchFunctionExecutionsReply {
	this := FetchFunctionExecutionsReply{}
	return &this
}

// NewFetchFunctionExecutionsReplyWithDefaults instantiates a new FetchFunctionExecutionsReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchFunctionExecutionsReplyWithDefaults() *FetchFunctionExecutionsReply {
	this := FetchFunctionExecutionsReply{}
	return &this
}

// GetExecutions returns the Executions field value if set, zero value otherwise.
func (o *FetchFunctionExecutionsReply) GetExecutions() []FunctionRunInfoListItem {
	if o == nil || o.Executions == nil {
		var ret []FunctionRunInfoListItem
		return ret
	}
	return *o.Executions
}

// GetExecutionsOk returns a tuple with the Executions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchFunctionExecutionsReply) GetExecutionsOk() (*[]FunctionRunInfoListItem, bool) {
	if o == nil || o.Executions == nil {
		return nil, false
	}
	return o.Executions, true
}

// HasExecutions returns a boolean if a field has been set.
func (o *FetchFunctionExecutionsReply) HasExecutions() bool {
	if o != nil && o.Executions != nil {
		return true
	}

	return false
}

// SetExecutions gets a reference to the given []FunctionRunInfoListItem and assigns it to the Executions field.
func (o *FetchFunctionExecutionsReply) SetExecutions(v []FunctionRunInfoListItem) {
	o.Executions = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *FetchFunctionExecutionsReply) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchFunctionExecutionsReply) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *FetchFunctionExecutionsReply) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *FetchFunctionExecutionsReply) SetLimit(v int64) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *FetchFunctionExecutionsReply) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchFunctionExecutionsReply) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *FetchFunctionExecutionsReply) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *FetchFunctionExecutionsReply) SetOffset(v int64) {
	o.Offset = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *FetchFunctionExecutionsReply) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchFunctionExecutionsReply) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *FetchFunctionExecutionsReply) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *FetchFunctionExecutionsReply) SetCount(v int64) {
	o.Count = &v
}

func (o FetchFunctionExecutionsReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Executions != nil {
		toSerialize["executions"] = o.Executions
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableFetchFunctionExecutionsReply struct {
	value *FetchFunctionExecutionsReply
	isSet bool
}

func (v NullableFetchFunctionExecutionsReply) Get() *FetchFunctionExecutionsReply {
	return v.value
}

func (v *NullableFetchFunctionExecutionsReply) Set(val *FetchFunctionExecutionsReply) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchFunctionExecutionsReply) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchFunctionExecutionsReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchFunctionExecutionsReply(val *FetchFunctionExecutionsReply) *NullableFetchFunctionExecutionsReply {
	return &NullableFetchFunctionExecutionsReply{value: val, isSet: true}
}

func (v NullableFetchFunctionExecutionsReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchFunctionExecutionsReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



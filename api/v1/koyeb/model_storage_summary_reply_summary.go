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

// StorageSummaryReplySummary struct for StorageSummaryReplySummary
type StorageSummaryReplySummary struct {
	StoreCount *float64 `json:"store_count,omitempty"`
	StackCount *float64 `json:"stack_count,omitempty"`
}

// NewStorageSummaryReplySummary instantiates a new StorageSummaryReplySummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageSummaryReplySummary() *StorageSummaryReplySummary {
	this := StorageSummaryReplySummary{}
	return &this
}

// NewStorageSummaryReplySummaryWithDefaults instantiates a new StorageSummaryReplySummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageSummaryReplySummaryWithDefaults() *StorageSummaryReplySummary {
	this := StorageSummaryReplySummary{}
	return &this
}

// GetStoreCount returns the StoreCount field value if set, zero value otherwise.
func (o *StorageSummaryReplySummary) GetStoreCount() float64 {
	if o == nil || o.StoreCount == nil {
		var ret float64
		return ret
	}
	return *o.StoreCount
}

// GetStoreCountOk returns a tuple with the StoreCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSummaryReplySummary) GetStoreCountOk() (*float64, bool) {
	if o == nil || o.StoreCount == nil {
		return nil, false
	}
	return o.StoreCount, true
}

// HasStoreCount returns a boolean if a field has been set.
func (o *StorageSummaryReplySummary) HasStoreCount() bool {
	if o != nil && o.StoreCount != nil {
		return true
	}

	return false
}

// SetStoreCount gets a reference to the given float64 and assigns it to the StoreCount field.
func (o *StorageSummaryReplySummary) SetStoreCount(v float64) {
	o.StoreCount = &v
}

// GetStackCount returns the StackCount field value if set, zero value otherwise.
func (o *StorageSummaryReplySummary) GetStackCount() float64 {
	if o == nil || o.StackCount == nil {
		var ret float64
		return ret
	}
	return *o.StackCount
}

// GetStackCountOk returns a tuple with the StackCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSummaryReplySummary) GetStackCountOk() (*float64, bool) {
	if o == nil || o.StackCount == nil {
		return nil, false
	}
	return o.StackCount, true
}

// HasStackCount returns a boolean if a field has been set.
func (o *StorageSummaryReplySummary) HasStackCount() bool {
	if o != nil && o.StackCount != nil {
		return true
	}

	return false
}

// SetStackCount gets a reference to the given float64 and assigns it to the StackCount field.
func (o *StorageSummaryReplySummary) SetStackCount(v float64) {
	o.StackCount = &v
}

func (o StorageSummaryReplySummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StoreCount != nil {
		toSerialize["store_count"] = o.StoreCount
	}
	if o.StackCount != nil {
		toSerialize["stack_count"] = o.StackCount
	}
	return json.Marshal(toSerialize)
}

type NullableStorageSummaryReplySummary struct {
	value *StorageSummaryReplySummary
	isSet bool
}

func (v NullableStorageSummaryReplySummary) Get() *StorageSummaryReplySummary {
	return v.value
}

func (v *NullableStorageSummaryReplySummary) Set(val *StorageSummaryReplySummary) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageSummaryReplySummary) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageSummaryReplySummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageSummaryReplySummary(val *StorageSummaryReplySummary) *NullableStorageSummaryReplySummary {
	return &NullableStorageSummaryReplySummary{value: val, isSet: true}
}

func (v NullableStorageSummaryReplySummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageSummaryReplySummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


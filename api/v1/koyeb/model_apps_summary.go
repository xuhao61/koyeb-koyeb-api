/*
Koyeb Rest API

The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"encoding/json"
)

// AppsSummary struct for AppsSummary
type AppsSummary struct {
	Total *string `json:"total,omitempty"`
	ByStatus *map[string]string `json:"by_status,omitempty"`
}

// NewAppsSummary instantiates a new AppsSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppsSummary() *AppsSummary {
	this := AppsSummary{}
	return &this
}

// NewAppsSummaryWithDefaults instantiates a new AppsSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppsSummaryWithDefaults() *AppsSummary {
	this := AppsSummary{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *AppsSummary) GetTotal() string {
	if o == nil || isNil(o.Total) {
		var ret string
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppsSummary) GetTotalOk() (*string, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *AppsSummary) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given string and assigns it to the Total field.
func (o *AppsSummary) SetTotal(v string) {
	o.Total = &v
}

// GetByStatus returns the ByStatus field value if set, zero value otherwise.
func (o *AppsSummary) GetByStatus() map[string]string {
	if o == nil || isNil(o.ByStatus) {
		var ret map[string]string
		return ret
	}
	return *o.ByStatus
}

// GetByStatusOk returns a tuple with the ByStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppsSummary) GetByStatusOk() (*map[string]string, bool) {
	if o == nil || isNil(o.ByStatus) {
    return nil, false
	}
	return o.ByStatus, true
}

// HasByStatus returns a boolean if a field has been set.
func (o *AppsSummary) HasByStatus() bool {
	if o != nil && !isNil(o.ByStatus) {
		return true
	}

	return false
}

// SetByStatus gets a reference to the given map[string]string and assigns it to the ByStatus field.
func (o *AppsSummary) SetByStatus(v map[string]string) {
	o.ByStatus = &v
}

func (o AppsSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.ByStatus) {
		toSerialize["by_status"] = o.ByStatus
	}
	return json.Marshal(toSerialize)
}

type NullableAppsSummary struct {
	value *AppsSummary
	isSet bool
}

func (v NullableAppsSummary) Get() *AppsSummary {
	return v.value
}

func (v *NullableAppsSummary) Set(val *AppsSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableAppsSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableAppsSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppsSummary(val *AppsSummary) *NullableAppsSummary {
	return &NullableAppsSummary{value: val, isSet: true}
}

func (v NullableAppsSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppsSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



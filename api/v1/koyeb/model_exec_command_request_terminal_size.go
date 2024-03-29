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

// ExecCommandRequestTerminalSize struct for ExecCommandRequestTerminalSize
type ExecCommandRequestTerminalSize struct {
	Height *int32 `json:"height,omitempty"`
	Width *int32 `json:"width,omitempty"`
}

// NewExecCommandRequestTerminalSize instantiates a new ExecCommandRequestTerminalSize object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecCommandRequestTerminalSize() *ExecCommandRequestTerminalSize {
	this := ExecCommandRequestTerminalSize{}
	return &this
}

// NewExecCommandRequestTerminalSizeWithDefaults instantiates a new ExecCommandRequestTerminalSize object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecCommandRequestTerminalSizeWithDefaults() *ExecCommandRequestTerminalSize {
	this := ExecCommandRequestTerminalSize{}
	return &this
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *ExecCommandRequestTerminalSize) GetHeight() int32 {
	if o == nil || isNil(o.Height) {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecCommandRequestTerminalSize) GetHeightOk() (*int32, bool) {
	if o == nil || isNil(o.Height) {
    return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *ExecCommandRequestTerminalSize) HasHeight() bool {
	if o != nil && !isNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *ExecCommandRequestTerminalSize) SetHeight(v int32) {
	o.Height = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *ExecCommandRequestTerminalSize) GetWidth() int32 {
	if o == nil || isNil(o.Width) {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecCommandRequestTerminalSize) GetWidthOk() (*int32, bool) {
	if o == nil || isNil(o.Width) {
    return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *ExecCommandRequestTerminalSize) HasWidth() bool {
	if o != nil && !isNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *ExecCommandRequestTerminalSize) SetWidth(v int32) {
	o.Width = &v
}

func (o ExecCommandRequestTerminalSize) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !isNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	return json.Marshal(toSerialize)
}

type NullableExecCommandRequestTerminalSize struct {
	value *ExecCommandRequestTerminalSize
	isSet bool
}

func (v NullableExecCommandRequestTerminalSize) Get() *ExecCommandRequestTerminalSize {
	return v.value
}

func (v *NullableExecCommandRequestTerminalSize) Set(val *ExecCommandRequestTerminalSize) {
	v.value = val
	v.isSet = true
}

func (v NullableExecCommandRequestTerminalSize) IsSet() bool {
	return v.isSet
}

func (v *NullableExecCommandRequestTerminalSize) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecCommandRequestTerminalSize(val *ExecCommandRequestTerminalSize) *NullableExecCommandRequestTerminalSize {
	return &NullableExecCommandRequestTerminalSize{value: val, isSet: true}
}

func (v NullableExecCommandRequestTerminalSize) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecCommandRequestTerminalSize) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



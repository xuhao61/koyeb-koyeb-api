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

// ExecCommandReply struct for ExecCommandReply
type ExecCommandReply struct {
	Stdout *ExecCommandIO `json:"stdout,omitempty"`
	Stderr *ExecCommandIO `json:"stderr,omitempty"`
	Exited *bool `json:"exited,omitempty"`
	ExitCode *int32 `json:"exit_code,omitempty"`
}

// NewExecCommandReply instantiates a new ExecCommandReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecCommandReply() *ExecCommandReply {
	this := ExecCommandReply{}
	return &this
}

// NewExecCommandReplyWithDefaults instantiates a new ExecCommandReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecCommandReplyWithDefaults() *ExecCommandReply {
	this := ExecCommandReply{}
	return &this
}

// GetStdout returns the Stdout field value if set, zero value otherwise.
func (o *ExecCommandReply) GetStdout() ExecCommandIO {
	if o == nil || isNil(o.Stdout) {
		var ret ExecCommandIO
		return ret
	}
	return *o.Stdout
}

// GetStdoutOk returns a tuple with the Stdout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecCommandReply) GetStdoutOk() (*ExecCommandIO, bool) {
	if o == nil || isNil(o.Stdout) {
    return nil, false
	}
	return o.Stdout, true
}

// HasStdout returns a boolean if a field has been set.
func (o *ExecCommandReply) HasStdout() bool {
	if o != nil && !isNil(o.Stdout) {
		return true
	}

	return false
}

// SetStdout gets a reference to the given ExecCommandIO and assigns it to the Stdout field.
func (o *ExecCommandReply) SetStdout(v ExecCommandIO) {
	o.Stdout = &v
}

// GetStderr returns the Stderr field value if set, zero value otherwise.
func (o *ExecCommandReply) GetStderr() ExecCommandIO {
	if o == nil || isNil(o.Stderr) {
		var ret ExecCommandIO
		return ret
	}
	return *o.Stderr
}

// GetStderrOk returns a tuple with the Stderr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecCommandReply) GetStderrOk() (*ExecCommandIO, bool) {
	if o == nil || isNil(o.Stderr) {
    return nil, false
	}
	return o.Stderr, true
}

// HasStderr returns a boolean if a field has been set.
func (o *ExecCommandReply) HasStderr() bool {
	if o != nil && !isNil(o.Stderr) {
		return true
	}

	return false
}

// SetStderr gets a reference to the given ExecCommandIO and assigns it to the Stderr field.
func (o *ExecCommandReply) SetStderr(v ExecCommandIO) {
	o.Stderr = &v
}

// GetExited returns the Exited field value if set, zero value otherwise.
func (o *ExecCommandReply) GetExited() bool {
	if o == nil || isNil(o.Exited) {
		var ret bool
		return ret
	}
	return *o.Exited
}

// GetExitedOk returns a tuple with the Exited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecCommandReply) GetExitedOk() (*bool, bool) {
	if o == nil || isNil(o.Exited) {
    return nil, false
	}
	return o.Exited, true
}

// HasExited returns a boolean if a field has been set.
func (o *ExecCommandReply) HasExited() bool {
	if o != nil && !isNil(o.Exited) {
		return true
	}

	return false
}

// SetExited gets a reference to the given bool and assigns it to the Exited field.
func (o *ExecCommandReply) SetExited(v bool) {
	o.Exited = &v
}

// GetExitCode returns the ExitCode field value if set, zero value otherwise.
func (o *ExecCommandReply) GetExitCode() int32 {
	if o == nil || isNil(o.ExitCode) {
		var ret int32
		return ret
	}
	return *o.ExitCode
}

// GetExitCodeOk returns a tuple with the ExitCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecCommandReply) GetExitCodeOk() (*int32, bool) {
	if o == nil || isNil(o.ExitCode) {
    return nil, false
	}
	return o.ExitCode, true
}

// HasExitCode returns a boolean if a field has been set.
func (o *ExecCommandReply) HasExitCode() bool {
	if o != nil && !isNil(o.ExitCode) {
		return true
	}

	return false
}

// SetExitCode gets a reference to the given int32 and assigns it to the ExitCode field.
func (o *ExecCommandReply) SetExitCode(v int32) {
	o.ExitCode = &v
}

func (o ExecCommandReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Stdout) {
		toSerialize["stdout"] = o.Stdout
	}
	if !isNil(o.Stderr) {
		toSerialize["stderr"] = o.Stderr
	}
	if !isNil(o.Exited) {
		toSerialize["exited"] = o.Exited
	}
	if !isNil(o.ExitCode) {
		toSerialize["exit_code"] = o.ExitCode
	}
	return json.Marshal(toSerialize)
}

type NullableExecCommandReply struct {
	value *ExecCommandReply
	isSet bool
}

func (v NullableExecCommandReply) Get() *ExecCommandReply {
	return v.value
}

func (v *NullableExecCommandReply) Set(val *ExecCommandReply) {
	v.value = val
	v.isSet = true
}

func (v NullableExecCommandReply) IsSet() bool {
	return v.isSet
}

func (v *NullableExecCommandReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecCommandReply(val *ExecCommandReply) *NullableExecCommandReply {
	return &NullableExecCommandReply{value: val, isSet: true}
}

func (v NullableExecCommandReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecCommandReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



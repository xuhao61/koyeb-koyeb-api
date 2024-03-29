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

// ResetPasswordRequest struct for ResetPasswordRequest
type ResetPasswordRequest struct {
	Email *string `json:"email,omitempty"`
}

// NewResetPasswordRequest instantiates a new ResetPasswordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResetPasswordRequest() *ResetPasswordRequest {
	this := ResetPasswordRequest{}
	return &this
}

// NewResetPasswordRequestWithDefaults instantiates a new ResetPasswordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResetPasswordRequestWithDefaults() *ResetPasswordRequest {
	this := ResetPasswordRequest{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ResetPasswordRequest) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetPasswordRequest) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ResetPasswordRequest) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ResetPasswordRequest) SetEmail(v string) {
	o.Email = &v
}

func (o ResetPasswordRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableResetPasswordRequest struct {
	value *ResetPasswordRequest
	isSet bool
}

func (v NullableResetPasswordRequest) Get() *ResetPasswordRequest {
	return v.value
}

func (v *NullableResetPasswordRequest) Set(val *ResetPasswordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableResetPasswordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableResetPasswordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResetPasswordRequest(val *ResetPasswordRequest) *NullableResetPasswordRequest {
	return &NullableResetPasswordRequest{value: val, isSet: true}
}

func (v NullableResetPasswordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResetPasswordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



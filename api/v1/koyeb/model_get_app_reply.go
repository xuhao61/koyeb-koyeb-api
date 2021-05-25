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

// GetAppReply struct for GetAppReply
type GetAppReply struct {
	App *App `json:"app,omitempty"`
}

// NewGetAppReply instantiates a new GetAppReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAppReply() *GetAppReply {
	this := GetAppReply{}
	return &this
}

// NewGetAppReplyWithDefaults instantiates a new GetAppReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAppReplyWithDefaults() *GetAppReply {
	this := GetAppReply{}
	return &this
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *GetAppReply) GetApp() App {
	if o == nil || o.App == nil {
		var ret App
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAppReply) GetAppOk() (*App, bool) {
	if o == nil || o.App == nil {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *GetAppReply) HasApp() bool {
	if o != nil && o.App != nil {
		return true
	}

	return false
}

// SetApp gets a reference to the given App and assigns it to the App field.
func (o *GetAppReply) SetApp(v App) {
	o.App = &v
}

func (o GetAppReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.App != nil {
		toSerialize["app"] = o.App
	}
	return json.Marshal(toSerialize)
}

type NullableGetAppReply struct {
	value *GetAppReply
	isSet bool
}

func (v NullableGetAppReply) Get() *GetAppReply {
	return v.value
}

func (v *NullableGetAppReply) Set(val *GetAppReply) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAppReply) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAppReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAppReply(val *GetAppReply) *NullableGetAppReply {
	return &NullableGetAppReply{value: val, isSet: true}
}

func (v NullableGetAppReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAppReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



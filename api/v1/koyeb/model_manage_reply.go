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

// ManageReply struct for ManageReply
type ManageReply struct {
	Url *string `json:"url,omitempty"`
}

// NewManageReply instantiates a new ManageReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageReply() *ManageReply {
	this := ManageReply{}
	return &this
}

// NewManageReplyWithDefaults instantiates a new ManageReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageReplyWithDefaults() *ManageReply {
	this := ManageReply{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ManageReply) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageReply) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ManageReply) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ManageReply) SetUrl(v string) {
	o.Url = &v
}

func (o ManageReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableManageReply struct {
	value *ManageReply
	isSet bool
}

func (v NullableManageReply) Get() *ManageReply {
	return v.value
}

func (v *NullableManageReply) Set(val *ManageReply) {
	v.value = val
	v.isSet = true
}

func (v NullableManageReply) IsSet() bool {
	return v.isSet
}

func (v *NullableManageReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageReply(val *ManageReply) *NullableManageReply {
	return &NullableManageReply{value: val, isSet: true}
}

func (v NullableManageReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



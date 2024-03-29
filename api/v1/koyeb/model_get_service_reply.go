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

// GetServiceReply struct for GetServiceReply
type GetServiceReply struct {
	Service *Service `json:"service,omitempty"`
}

// NewGetServiceReply instantiates a new GetServiceReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetServiceReply() *GetServiceReply {
	this := GetServiceReply{}
	return &this
}

// NewGetServiceReplyWithDefaults instantiates a new GetServiceReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetServiceReplyWithDefaults() *GetServiceReply {
	this := GetServiceReply{}
	return &this
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *GetServiceReply) GetService() Service {
	if o == nil || isNil(o.Service) {
		var ret Service
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServiceReply) GetServiceOk() (*Service, bool) {
	if o == nil || isNil(o.Service) {
    return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *GetServiceReply) HasService() bool {
	if o != nil && !isNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given Service and assigns it to the Service field.
func (o *GetServiceReply) SetService(v Service) {
	o.Service = &v
}

func (o GetServiceReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	return json.Marshal(toSerialize)
}

type NullableGetServiceReply struct {
	value *GetServiceReply
	isSet bool
}

func (v NullableGetServiceReply) Get() *GetServiceReply {
	return v.value
}

func (v *NullableGetServiceReply) Set(val *GetServiceReply) {
	v.value = val
	v.isSet = true
}

func (v NullableGetServiceReply) IsSet() bool {
	return v.isSet
}

func (v *NullableGetServiceReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetServiceReply(val *GetServiceReply) *NullableGetServiceReply {
	return &NullableGetServiceReply{value: val, isSet: true}
}

func (v NullableGetServiceReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetServiceReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



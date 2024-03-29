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

// GetRegionalDeploymentReply struct for GetRegionalDeploymentReply
type GetRegionalDeploymentReply struct {
	RegionalDeployment *RegionalDeployment `json:"regional_deployment,omitempty"`
}

// NewGetRegionalDeploymentReply instantiates a new GetRegionalDeploymentReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRegionalDeploymentReply() *GetRegionalDeploymentReply {
	this := GetRegionalDeploymentReply{}
	return &this
}

// NewGetRegionalDeploymentReplyWithDefaults instantiates a new GetRegionalDeploymentReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRegionalDeploymentReplyWithDefaults() *GetRegionalDeploymentReply {
	this := GetRegionalDeploymentReply{}
	return &this
}

// GetRegionalDeployment returns the RegionalDeployment field value if set, zero value otherwise.
func (o *GetRegionalDeploymentReply) GetRegionalDeployment() RegionalDeployment {
	if o == nil || isNil(o.RegionalDeployment) {
		var ret RegionalDeployment
		return ret
	}
	return *o.RegionalDeployment
}

// GetRegionalDeploymentOk returns a tuple with the RegionalDeployment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRegionalDeploymentReply) GetRegionalDeploymentOk() (*RegionalDeployment, bool) {
	if o == nil || isNil(o.RegionalDeployment) {
    return nil, false
	}
	return o.RegionalDeployment, true
}

// HasRegionalDeployment returns a boolean if a field has been set.
func (o *GetRegionalDeploymentReply) HasRegionalDeployment() bool {
	if o != nil && !isNil(o.RegionalDeployment) {
		return true
	}

	return false
}

// SetRegionalDeployment gets a reference to the given RegionalDeployment and assigns it to the RegionalDeployment field.
func (o *GetRegionalDeploymentReply) SetRegionalDeployment(v RegionalDeployment) {
	o.RegionalDeployment = &v
}

func (o GetRegionalDeploymentReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RegionalDeployment) {
		toSerialize["regional_deployment"] = o.RegionalDeployment
	}
	return json.Marshal(toSerialize)
}

type NullableGetRegionalDeploymentReply struct {
	value *GetRegionalDeploymentReply
	isSet bool
}

func (v NullableGetRegionalDeploymentReply) Get() *GetRegionalDeploymentReply {
	return v.value
}

func (v *NullableGetRegionalDeploymentReply) Set(val *GetRegionalDeploymentReply) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRegionalDeploymentReply) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRegionalDeploymentReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRegionalDeploymentReply(val *GetRegionalDeploymentReply) *NullableGetRegionalDeploymentReply {
	return &NullableGetRegionalDeploymentReply{value: val, isSet: true}
}

func (v NullableGetRegionalDeploymentReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRegionalDeploymentReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



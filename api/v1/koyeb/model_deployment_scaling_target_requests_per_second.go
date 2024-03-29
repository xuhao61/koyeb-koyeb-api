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

// DeploymentScalingTargetRequestsPerSecond struct for DeploymentScalingTargetRequestsPerSecond
type DeploymentScalingTargetRequestsPerSecond struct {
	Value *int64 `json:"value,omitempty"`
}

// NewDeploymentScalingTargetRequestsPerSecond instantiates a new DeploymentScalingTargetRequestsPerSecond object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentScalingTargetRequestsPerSecond() *DeploymentScalingTargetRequestsPerSecond {
	this := DeploymentScalingTargetRequestsPerSecond{}
	return &this
}

// NewDeploymentScalingTargetRequestsPerSecondWithDefaults instantiates a new DeploymentScalingTargetRequestsPerSecond object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentScalingTargetRequestsPerSecondWithDefaults() *DeploymentScalingTargetRequestsPerSecond {
	this := DeploymentScalingTargetRequestsPerSecond{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DeploymentScalingTargetRequestsPerSecond) GetValue() int64 {
	if o == nil || isNil(o.Value) {
		var ret int64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentScalingTargetRequestsPerSecond) GetValueOk() (*int64, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DeploymentScalingTargetRequestsPerSecond) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int64 and assigns it to the Value field.
func (o *DeploymentScalingTargetRequestsPerSecond) SetValue(v int64) {
	o.Value = &v
}

func (o DeploymentScalingTargetRequestsPerSecond) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentScalingTargetRequestsPerSecond struct {
	value *DeploymentScalingTargetRequestsPerSecond
	isSet bool
}

func (v NullableDeploymentScalingTargetRequestsPerSecond) Get() *DeploymentScalingTargetRequestsPerSecond {
	return v.value
}

func (v *NullableDeploymentScalingTargetRequestsPerSecond) Set(val *DeploymentScalingTargetRequestsPerSecond) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentScalingTargetRequestsPerSecond) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentScalingTargetRequestsPerSecond) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentScalingTargetRequestsPerSecond(val *DeploymentScalingTargetRequestsPerSecond) *NullableDeploymentScalingTargetRequestsPerSecond {
	return &NullableDeploymentScalingTargetRequestsPerSecond{value: val, isSet: true}
}

func (v NullableDeploymentScalingTargetRequestsPerSecond) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentScalingTargetRequestsPerSecond) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



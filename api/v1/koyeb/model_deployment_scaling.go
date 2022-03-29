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

// DeploymentScaling struct for DeploymentScaling
type DeploymentScaling struct {
	Scopes *[]string `json:"scopes,omitempty"`
	Min *int64 `json:"min,omitempty"`
	Max *int64 `json:"max,omitempty"`
}

// NewDeploymentScaling instantiates a new DeploymentScaling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentScaling() *DeploymentScaling {
	this := DeploymentScaling{}
	return &this
}

// NewDeploymentScalingWithDefaults instantiates a new DeploymentScaling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentScalingWithDefaults() *DeploymentScaling {
	this := DeploymentScaling{}
	return &this
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *DeploymentScaling) GetScopes() []string {
	if o == nil || o.Scopes == nil {
		var ret []string
		return ret
	}
	return *o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentScaling) GetScopesOk() (*[]string, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *DeploymentScaling) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *DeploymentScaling) SetScopes(v []string) {
	o.Scopes = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *DeploymentScaling) GetMin() int64 {
	if o == nil || o.Min == nil {
		var ret int64
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentScaling) GetMinOk() (*int64, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *DeploymentScaling) HasMin() bool {
	if o != nil && o.Min != nil {
		return true
	}

	return false
}

// SetMin gets a reference to the given int64 and assigns it to the Min field.
func (o *DeploymentScaling) SetMin(v int64) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *DeploymentScaling) GetMax() int64 {
	if o == nil || o.Max == nil {
		var ret int64
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentScaling) GetMaxOk() (*int64, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *DeploymentScaling) HasMax() bool {
	if o != nil && o.Max != nil {
		return true
	}

	return false
}

// SetMax gets a reference to the given int64 and assigns it to the Max field.
func (o *DeploymentScaling) SetMax(v int64) {
	o.Max = &v
}

func (o DeploymentScaling) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	if o.Min != nil {
		toSerialize["min"] = o.Min
	}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentScaling struct {
	value *DeploymentScaling
	isSet bool
}

func (v NullableDeploymentScaling) Get() *DeploymentScaling {
	return v.value
}

func (v *NullableDeploymentScaling) Set(val *DeploymentScaling) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentScaling) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentScaling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentScaling(val *DeploymentScaling) *NullableDeploymentScaling {
	return &NullableDeploymentScaling{value: val, isSet: true}
}

func (v NullableDeploymentScaling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentScaling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



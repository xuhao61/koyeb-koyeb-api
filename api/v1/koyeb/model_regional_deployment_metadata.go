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

// RegionalDeploymentMetadata struct for RegionalDeploymentMetadata
type RegionalDeploymentMetadata struct {
	RuntimeJobId *string `json:"runtime_job_id,omitempty"`
}

// NewRegionalDeploymentMetadata instantiates a new RegionalDeploymentMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionalDeploymentMetadata() *RegionalDeploymentMetadata {
	this := RegionalDeploymentMetadata{}
	return &this
}

// NewRegionalDeploymentMetadataWithDefaults instantiates a new RegionalDeploymentMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionalDeploymentMetadataWithDefaults() *RegionalDeploymentMetadata {
	this := RegionalDeploymentMetadata{}
	return &this
}

// GetRuntimeJobId returns the RuntimeJobId field value if set, zero value otherwise.
func (o *RegionalDeploymentMetadata) GetRuntimeJobId() string {
	if o == nil || isNil(o.RuntimeJobId) {
		var ret string
		return ret
	}
	return *o.RuntimeJobId
}

// GetRuntimeJobIdOk returns a tuple with the RuntimeJobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionalDeploymentMetadata) GetRuntimeJobIdOk() (*string, bool) {
	if o == nil || isNil(o.RuntimeJobId) {
    return nil, false
	}
	return o.RuntimeJobId, true
}

// HasRuntimeJobId returns a boolean if a field has been set.
func (o *RegionalDeploymentMetadata) HasRuntimeJobId() bool {
	if o != nil && !isNil(o.RuntimeJobId) {
		return true
	}

	return false
}

// SetRuntimeJobId gets a reference to the given string and assigns it to the RuntimeJobId field.
func (o *RegionalDeploymentMetadata) SetRuntimeJobId(v string) {
	o.RuntimeJobId = &v
}

func (o RegionalDeploymentMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RuntimeJobId) {
		toSerialize["runtime_job_id"] = o.RuntimeJobId
	}
	return json.Marshal(toSerialize)
}

type NullableRegionalDeploymentMetadata struct {
	value *RegionalDeploymentMetadata
	isSet bool
}

func (v NullableRegionalDeploymentMetadata) Get() *RegionalDeploymentMetadata {
	return v.value
}

func (v *NullableRegionalDeploymentMetadata) Set(val *RegionalDeploymentMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionalDeploymentMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionalDeploymentMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionalDeploymentMetadata(val *RegionalDeploymentMetadata) *NullableRegionalDeploymentMetadata {
	return &NullableRegionalDeploymentMetadata{value: val, isSet: true}
}

func (v NullableRegionalDeploymentMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionalDeploymentMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// DeploymentPort struct for DeploymentPort
type DeploymentPort struct {
	Port *int64 `json:"port,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
}

// NewDeploymentPort instantiates a new DeploymentPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentPort() *DeploymentPort {
	this := DeploymentPort{}
	return &this
}

// NewDeploymentPortWithDefaults instantiates a new DeploymentPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentPortWithDefaults() *DeploymentPort {
	this := DeploymentPort{}
	return &this
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *DeploymentPort) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentPort) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *DeploymentPort) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *DeploymentPort) SetPort(v int64) {
	o.Port = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *DeploymentPort) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentPort) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *DeploymentPort) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *DeploymentPort) SetProtocol(v string) {
	o.Protocol = &v
}

func (o DeploymentPort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentPort struct {
	value *DeploymentPort
	isSet bool
}

func (v NullableDeploymentPort) Get() *DeploymentPort {
	return v.value
}

func (v *NullableDeploymentPort) Set(val *DeploymentPort) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentPort) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentPort(val *DeploymentPort) *NullableDeploymentPort {
	return &NullableDeploymentPort{value: val, isSet: true}
}

func (v NullableDeploymentPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



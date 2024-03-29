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

// DeploymentDatabaseInfo struct for DeploymentDatabaseInfo
type DeploymentDatabaseInfo struct {
	NeonPostgres *DeploymentNeonPostgresDatabaseInfo `json:"neon_postgres,omitempty"`
}

// NewDeploymentDatabaseInfo instantiates a new DeploymentDatabaseInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentDatabaseInfo() *DeploymentDatabaseInfo {
	this := DeploymentDatabaseInfo{}
	return &this
}

// NewDeploymentDatabaseInfoWithDefaults instantiates a new DeploymentDatabaseInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentDatabaseInfoWithDefaults() *DeploymentDatabaseInfo {
	this := DeploymentDatabaseInfo{}
	return &this
}

// GetNeonPostgres returns the NeonPostgres field value if set, zero value otherwise.
func (o *DeploymentDatabaseInfo) GetNeonPostgres() DeploymentNeonPostgresDatabaseInfo {
	if o == nil || isNil(o.NeonPostgres) {
		var ret DeploymentNeonPostgresDatabaseInfo
		return ret
	}
	return *o.NeonPostgres
}

// GetNeonPostgresOk returns a tuple with the NeonPostgres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentDatabaseInfo) GetNeonPostgresOk() (*DeploymentNeonPostgresDatabaseInfo, bool) {
	if o == nil || isNil(o.NeonPostgres) {
    return nil, false
	}
	return o.NeonPostgres, true
}

// HasNeonPostgres returns a boolean if a field has been set.
func (o *DeploymentDatabaseInfo) HasNeonPostgres() bool {
	if o != nil && !isNil(o.NeonPostgres) {
		return true
	}

	return false
}

// SetNeonPostgres gets a reference to the given DeploymentNeonPostgresDatabaseInfo and assigns it to the NeonPostgres field.
func (o *DeploymentDatabaseInfo) SetNeonPostgres(v DeploymentNeonPostgresDatabaseInfo) {
	o.NeonPostgres = &v
}

func (o DeploymentDatabaseInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NeonPostgres) {
		toSerialize["neon_postgres"] = o.NeonPostgres
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentDatabaseInfo struct {
	value *DeploymentDatabaseInfo
	isSet bool
}

func (v NullableDeploymentDatabaseInfo) Get() *DeploymentDatabaseInfo {
	return v.value
}

func (v *NullableDeploymentDatabaseInfo) Set(val *DeploymentDatabaseInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentDatabaseInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentDatabaseInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentDatabaseInfo(val *DeploymentDatabaseInfo) *NullableDeploymentDatabaseInfo {
	return &NullableDeploymentDatabaseInfo{value: val, isSet: true}
}

func (v NullableDeploymentDatabaseInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentDatabaseInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



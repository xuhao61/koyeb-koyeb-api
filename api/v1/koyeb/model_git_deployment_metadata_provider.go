/*
Koyeb Rest API

The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"encoding/json"
	"fmt"
)

// GitDeploymentMetadataProvider the model 'GitDeploymentMetadataProvider'
type GitDeploymentMetadataProvider string

// List of GitDeploymentMetadata.Provider
const (
	GITDEPLOYMENTMETADATAPROVIDER_UNKNOWN GitDeploymentMetadataProvider = "UNKNOWN"
	GITDEPLOYMENTMETADATAPROVIDER_GITHUB GitDeploymentMetadataProvider = "GITHUB"
)

// All allowed values of GitDeploymentMetadataProvider enum
var AllowedGitDeploymentMetadataProviderEnumValues = []GitDeploymentMetadataProvider{
	"UNKNOWN",
	"GITHUB",
}

func (v *GitDeploymentMetadataProvider) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GitDeploymentMetadataProvider(value)
	for _, existing := range AllowedGitDeploymentMetadataProviderEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GitDeploymentMetadataProvider", value)
}

// NewGitDeploymentMetadataProviderFromValue returns a pointer to a valid GitDeploymentMetadataProvider
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGitDeploymentMetadataProviderFromValue(v string) (*GitDeploymentMetadataProvider, error) {
	ev := GitDeploymentMetadataProvider(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GitDeploymentMetadataProvider: valid values are %v", v, AllowedGitDeploymentMetadataProviderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GitDeploymentMetadataProvider) IsValid() bool {
	for _, existing := range AllowedGitDeploymentMetadataProviderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GitDeploymentMetadata.Provider value
func (v GitDeploymentMetadataProvider) Ptr() *GitDeploymentMetadataProvider {
	return &v
}

type NullableGitDeploymentMetadataProvider struct {
	value *GitDeploymentMetadataProvider
	isSet bool
}

func (v NullableGitDeploymentMetadataProvider) Get() *GitDeploymentMetadataProvider {
	return v.value
}

func (v *NullableGitDeploymentMetadataProvider) Set(val *GitDeploymentMetadataProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableGitDeploymentMetadataProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableGitDeploymentMetadataProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitDeploymentMetadataProvider(val *GitDeploymentMetadataProvider) *NullableGitDeploymentMetadataProvider {
	return &NullableGitDeploymentMetadataProvider{value: val, isSet: true}
}

func (v NullableGitDeploymentMetadataProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitDeploymentMetadataProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


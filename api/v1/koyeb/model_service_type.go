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

// ServiceType the model 'ServiceType'
type ServiceType string

// List of Service.Type
const (
	SERVICETYPE_INVALID_TYPE ServiceType = "INVALID_TYPE"
	SERVICETYPE_WEB ServiceType = "WEB"
	SERVICETYPE_WORKER ServiceType = "WORKER"
	SERVICETYPE_DATABASE ServiceType = "DATABASE"
)

// All allowed values of ServiceType enum
var AllowedServiceTypeEnumValues = []ServiceType{
	"INVALID_TYPE",
	"WEB",
	"WORKER",
	"DATABASE",
}

func (v *ServiceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServiceType(value)
	for _, existing := range AllowedServiceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServiceType", value)
}

// NewServiceTypeFromValue returns a pointer to a valid ServiceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServiceTypeFromValue(v string) (*ServiceType, error) {
	ev := ServiceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServiceType: valid values are %v", v, AllowedServiceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServiceType) IsValid() bool {
	for _, existing := range AllowedServiceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Service.Type value
func (v ServiceType) Ptr() *ServiceType {
	return &v
}

type NullableServiceType struct {
	value *ServiceType
	isSet bool
}

func (v NullableServiceType) Get() *ServiceType {
	return v.value
}

func (v *NullableServiceType) Set(val *ServiceType) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceType) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceType(val *ServiceType) *NullableServiceType {
	return &NullableServiceType{value: val, isSet: true}
}

func (v NullableServiceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


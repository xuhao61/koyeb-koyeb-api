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

// OrganizationMemberStatus the model 'OrganizationMemberStatus'
type OrganizationMemberStatus string

// List of OrganizationMember.Status
const (
	ORGANIZATIONMEMBERSTATUS_INVALID OrganizationMemberStatus = "INVALID"
	ORGANIZATIONMEMBERSTATUS_ACTIVE OrganizationMemberStatus = "ACTIVE"
	ORGANIZATIONMEMBERSTATUS_DELETED OrganizationMemberStatus = "DELETED"
)

// All allowed values of OrganizationMemberStatus enum
var AllowedOrganizationMemberStatusEnumValues = []OrganizationMemberStatus{
	"INVALID",
	"ACTIVE",
	"DELETED",
}

func (v *OrganizationMemberStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrganizationMemberStatus(value)
	for _, existing := range AllowedOrganizationMemberStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrganizationMemberStatus", value)
}

// NewOrganizationMemberStatusFromValue returns a pointer to a valid OrganizationMemberStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrganizationMemberStatusFromValue(v string) (*OrganizationMemberStatus, error) {
	ev := OrganizationMemberStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrganizationMemberStatus: valid values are %v", v, AllowedOrganizationMemberStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrganizationMemberStatus) IsValid() bool {
	for _, existing := range AllowedOrganizationMemberStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrganizationMember.Status value
func (v OrganizationMemberStatus) Ptr() *OrganizationMemberStatus {
	return &v
}

type NullableOrganizationMemberStatus struct {
	value *OrganizationMemberStatus
	isSet bool
}

func (v NullableOrganizationMemberStatus) Get() *OrganizationMemberStatus {
	return v.value
}

func (v *NullableOrganizationMemberStatus) Set(val *OrganizationMemberStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationMemberStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationMemberStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationMemberStatus(val *OrganizationMemberStatus) *NullableOrganizationMemberStatus {
	return &NullableOrganizationMemberStatus{value: val, isSet: true}
}

func (v NullableOrganizationMemberStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationMemberStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


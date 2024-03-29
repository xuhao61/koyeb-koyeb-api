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

// CreateOrganizationInvitationReply struct for CreateOrganizationInvitationReply
type CreateOrganizationInvitationReply struct {
	Invitation *OrganizationInvitation `json:"invitation,omitempty"`
}

// NewCreateOrganizationInvitationReply instantiates a new CreateOrganizationInvitationReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationInvitationReply() *CreateOrganizationInvitationReply {
	this := CreateOrganizationInvitationReply{}
	return &this
}

// NewCreateOrganizationInvitationReplyWithDefaults instantiates a new CreateOrganizationInvitationReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationInvitationReplyWithDefaults() *CreateOrganizationInvitationReply {
	this := CreateOrganizationInvitationReply{}
	return &this
}

// GetInvitation returns the Invitation field value if set, zero value otherwise.
func (o *CreateOrganizationInvitationReply) GetInvitation() OrganizationInvitation {
	if o == nil || isNil(o.Invitation) {
		var ret OrganizationInvitation
		return ret
	}
	return *o.Invitation
}

// GetInvitationOk returns a tuple with the Invitation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInvitationReply) GetInvitationOk() (*OrganizationInvitation, bool) {
	if o == nil || isNil(o.Invitation) {
    return nil, false
	}
	return o.Invitation, true
}

// HasInvitation returns a boolean if a field has been set.
func (o *CreateOrganizationInvitationReply) HasInvitation() bool {
	if o != nil && !isNil(o.Invitation) {
		return true
	}

	return false
}

// SetInvitation gets a reference to the given OrganizationInvitation and assigns it to the Invitation field.
func (o *CreateOrganizationInvitationReply) SetInvitation(v OrganizationInvitation) {
	o.Invitation = &v
}

func (o CreateOrganizationInvitationReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Invitation) {
		toSerialize["invitation"] = o.Invitation
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrganizationInvitationReply struct {
	value *CreateOrganizationInvitationReply
	isSet bool
}

func (v NullableCreateOrganizationInvitationReply) Get() *CreateOrganizationInvitationReply {
	return v.value
}

func (v *NullableCreateOrganizationInvitationReply) Set(val *CreateOrganizationInvitationReply) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationInvitationReply) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationInvitationReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationInvitationReply(val *CreateOrganizationInvitationReply) *NullableCreateOrganizationInvitationReply {
	return &NullableCreateOrganizationInvitationReply{value: val, isSet: true}
}

func (v NullableCreateOrganizationInvitationReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationInvitationReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



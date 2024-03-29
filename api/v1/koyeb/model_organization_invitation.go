/*
Koyeb Rest API

The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"encoding/json"
	"time"
)

// OrganizationInvitation struct for OrganizationInvitation
type OrganizationInvitation struct {
	Id *string `json:"id,omitempty"`
	Email *string `json:"email,omitempty"`
	Role *UserRoleRole `json:"role,omitempty"`
	Status *OrganizationInvitationStatus `json:"status,omitempty"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	OrganizationId *string `json:"organization_id,omitempty"`
	Organization *PublicOrganization `json:"organization,omitempty"`
	InviteeId *string `json:"invitee_id,omitempty"`
	Invitee *PublicUser `json:"invitee,omitempty"`
	InviterId *string `json:"inviter_id,omitempty"`
	Inviter *PublicUser `json:"inviter,omitempty"`
}

// NewOrganizationInvitation instantiates a new OrganizationInvitation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationInvitation() *OrganizationInvitation {
	this := OrganizationInvitation{}
	var role UserRoleRole = USERROLEROLE_INVALID
	this.Role = &role
	var status OrganizationInvitationStatus = ORGANIZATIONINVITATIONSTATUS_INVALID
	this.Status = &status
	return &this
}

// NewOrganizationInvitationWithDefaults instantiates a new OrganizationInvitation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationInvitationWithDefaults() *OrganizationInvitation {
	this := OrganizationInvitation{}
	var role UserRoleRole = USERROLEROLE_INVALID
	this.Role = &role
	var status OrganizationInvitationStatus = ORGANIZATIONINVITATIONSTATUS_INVALID
	this.Status = &status
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationInvitation) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitation) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationInvitation) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationInvitation) SetId(v string) {
	o.Id = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *OrganizationInvitation) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitation) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *OrganizationInvitation) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *OrganizationInvitation) SetEmail(v string) {
	o.Email = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *OrganizationInvitation) GetRole() UserRoleRole {
	if o == nil || isNil(o.Role) {
		var ret UserRoleRole
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitation) GetRoleOk() (*UserRoleRole, bool) {
	if o == nil || isNil(o.Role) {
    return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *OrganizationInvitation) HasRole() bool {
	if o != nil && !isNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given UserRoleRole and assigns it to the Role field.
func (o *OrganizationInvitation) SetRole(v UserRoleRole) {
	o.Role = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrganizationInvitation) GetStatus() OrganizationInvitationStatus {
	if o == nil || isNil(o.Status) {
		var ret OrganizationInvitationStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitation) GetStatusOk() (*OrganizationInvitationStatus, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrganizationInvitation) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given OrganizationInvitationStatus and assigns it to the Status field.
func (o *OrganizationInvitation) SetStatus(v OrganizationInvitationStatus) {
	o.Status = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *OrganizationInvitation) GetExpiresAt() time.Time {
	if o == nil || isNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitation) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ExpiresAt) {
    return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *OrganizationInvitation) HasExpiresAt() bool {
	if o != nil && !isNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *OrganizationInvitation) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *OrganizationInvitation) GetOrganizationId() string {
	if o == nil || isNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitation) GetOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationId) {
    return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *OrganizationInvitation) HasOrganizationId() bool {
	if o != nil && !isNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *OrganizationInvitation) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *OrganizationInvitation) GetOrganization() PublicOrganization {
	if o == nil || isNil(o.Organization) {
		var ret PublicOrganization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitation) GetOrganizationOk() (*PublicOrganization, bool) {
	if o == nil || isNil(o.Organization) {
    return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *OrganizationInvitation) HasOrganization() bool {
	if o != nil && !isNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given PublicOrganization and assigns it to the Organization field.
func (o *OrganizationInvitation) SetOrganization(v PublicOrganization) {
	o.Organization = &v
}

// GetInviteeId returns the InviteeId field value if set, zero value otherwise.
func (o *OrganizationInvitation) GetInviteeId() string {
	if o == nil || isNil(o.InviteeId) {
		var ret string
		return ret
	}
	return *o.InviteeId
}

// GetInviteeIdOk returns a tuple with the InviteeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitation) GetInviteeIdOk() (*string, bool) {
	if o == nil || isNil(o.InviteeId) {
    return nil, false
	}
	return o.InviteeId, true
}

// HasInviteeId returns a boolean if a field has been set.
func (o *OrganizationInvitation) HasInviteeId() bool {
	if o != nil && !isNil(o.InviteeId) {
		return true
	}

	return false
}

// SetInviteeId gets a reference to the given string and assigns it to the InviteeId field.
func (o *OrganizationInvitation) SetInviteeId(v string) {
	o.InviteeId = &v
}

// GetInvitee returns the Invitee field value if set, zero value otherwise.
func (o *OrganizationInvitation) GetInvitee() PublicUser {
	if o == nil || isNil(o.Invitee) {
		var ret PublicUser
		return ret
	}
	return *o.Invitee
}

// GetInviteeOk returns a tuple with the Invitee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitation) GetInviteeOk() (*PublicUser, bool) {
	if o == nil || isNil(o.Invitee) {
    return nil, false
	}
	return o.Invitee, true
}

// HasInvitee returns a boolean if a field has been set.
func (o *OrganizationInvitation) HasInvitee() bool {
	if o != nil && !isNil(o.Invitee) {
		return true
	}

	return false
}

// SetInvitee gets a reference to the given PublicUser and assigns it to the Invitee field.
func (o *OrganizationInvitation) SetInvitee(v PublicUser) {
	o.Invitee = &v
}

// GetInviterId returns the InviterId field value if set, zero value otherwise.
func (o *OrganizationInvitation) GetInviterId() string {
	if o == nil || isNil(o.InviterId) {
		var ret string
		return ret
	}
	return *o.InviterId
}

// GetInviterIdOk returns a tuple with the InviterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitation) GetInviterIdOk() (*string, bool) {
	if o == nil || isNil(o.InviterId) {
    return nil, false
	}
	return o.InviterId, true
}

// HasInviterId returns a boolean if a field has been set.
func (o *OrganizationInvitation) HasInviterId() bool {
	if o != nil && !isNil(o.InviterId) {
		return true
	}

	return false
}

// SetInviterId gets a reference to the given string and assigns it to the InviterId field.
func (o *OrganizationInvitation) SetInviterId(v string) {
	o.InviterId = &v
}

// GetInviter returns the Inviter field value if set, zero value otherwise.
func (o *OrganizationInvitation) GetInviter() PublicUser {
	if o == nil || isNil(o.Inviter) {
		var ret PublicUser
		return ret
	}
	return *o.Inviter
}

// GetInviterOk returns a tuple with the Inviter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitation) GetInviterOk() (*PublicUser, bool) {
	if o == nil || isNil(o.Inviter) {
    return nil, false
	}
	return o.Inviter, true
}

// HasInviter returns a boolean if a field has been set.
func (o *OrganizationInvitation) HasInviter() bool {
	if o != nil && !isNil(o.Inviter) {
		return true
	}

	return false
}

// SetInviter gets a reference to the given PublicUser and assigns it to the Inviter field.
func (o *OrganizationInvitation) SetInviter(v PublicUser) {
	o.Inviter = &v
}

func (o OrganizationInvitation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if !isNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if !isNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !isNil(o.InviteeId) {
		toSerialize["invitee_id"] = o.InviteeId
	}
	if !isNil(o.Invitee) {
		toSerialize["invitee"] = o.Invitee
	}
	if !isNil(o.InviterId) {
		toSerialize["inviter_id"] = o.InviterId
	}
	if !isNil(o.Inviter) {
		toSerialize["inviter"] = o.Inviter
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationInvitation struct {
	value *OrganizationInvitation
	isSet bool
}

func (v NullableOrganizationInvitation) Get() *OrganizationInvitation {
	return v.value
}

func (v *NullableOrganizationInvitation) Set(val *OrganizationInvitation) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationInvitation) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationInvitation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationInvitation(val *OrganizationInvitation) *NullableOrganizationInvitation {
	return &NullableOrganizationInvitation{value: val, isSet: true}
}

func (v NullableOrganizationInvitation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationInvitation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



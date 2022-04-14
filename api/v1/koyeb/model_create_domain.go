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

// CreateDomain struct for CreateDomain
type CreateDomain struct {
	Name *string `json:"name,omitempty"`
	Type *DomainType `json:"type,omitempty"`
	AppId *string `json:"app_id,omitempty"`
}

// NewCreateDomain instantiates a new CreateDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDomain() *CreateDomain {
	this := CreateDomain{}
	var type_ DomainType = DOMAINTYPE_AUTOASSIGNED
	this.Type = &type_
	return &this
}

// NewCreateDomainWithDefaults instantiates a new CreateDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDomainWithDefaults() *CreateDomain {
	this := CreateDomain{}
	var type_ DomainType = DOMAINTYPE_AUTOASSIGNED
	this.Type = &type_
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateDomain) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDomain) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateDomain) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateDomain) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateDomain) GetType() DomainType {
	if o == nil || o.Type == nil {
		var ret DomainType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDomain) GetTypeOk() (*DomainType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateDomain) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given DomainType and assigns it to the Type field.
func (o *CreateDomain) SetType(v DomainType) {
	o.Type = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *CreateDomain) GetAppId() string {
	if o == nil || o.AppId == nil {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDomain) GetAppIdOk() (*string, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *CreateDomain) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *CreateDomain) SetAppId(v string) {
	o.AppId = &v
}

func (o CreateDomain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.AppId != nil {
		toSerialize["app_id"] = o.AppId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateDomain struct {
	value *CreateDomain
	isSet bool
}

func (v NullableCreateDomain) Get() *CreateDomain {
	return v.value
}

func (v *NullableCreateDomain) Set(val *CreateDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDomain(val *CreateDomain) *NullableCreateDomain {
	return &NullableCreateDomain{value: val, isSet: true}
}

func (v NullableCreateDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



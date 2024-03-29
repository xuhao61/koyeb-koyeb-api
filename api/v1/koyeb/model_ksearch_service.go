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

// KsearchService struct for KsearchService
type KsearchService struct {
	Id *string `json:"id,omitempty"`
	OrganizationId *string `json:"organization_id,omitempty"`
	AppId *string `json:"app_id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewKsearchService instantiates a new KsearchService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKsearchService() *KsearchService {
	this := KsearchService{}
	return &this
}

// NewKsearchServiceWithDefaults instantiates a new KsearchService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKsearchServiceWithDefaults() *KsearchService {
	this := KsearchService{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KsearchService) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KsearchService) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KsearchService) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *KsearchService) SetId(v string) {
	o.Id = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *KsearchService) GetOrganizationId() string {
	if o == nil || isNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KsearchService) GetOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationId) {
    return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *KsearchService) HasOrganizationId() bool {
	if o != nil && !isNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *KsearchService) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *KsearchService) GetAppId() string {
	if o == nil || isNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KsearchService) GetAppIdOk() (*string, bool) {
	if o == nil || isNil(o.AppId) {
    return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *KsearchService) HasAppId() bool {
	if o != nil && !isNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *KsearchService) SetAppId(v string) {
	o.AppId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KsearchService) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KsearchService) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KsearchService) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KsearchService) SetName(v string) {
	o.Name = &v
}

func (o KsearchService) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if !isNil(o.AppId) {
		toSerialize["app_id"] = o.AppId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableKsearchService struct {
	value *KsearchService
	isSet bool
}

func (v NullableKsearchService) Get() *KsearchService {
	return v.value
}

func (v *NullableKsearchService) Set(val *KsearchService) {
	v.value = val
	v.isSet = true
}

func (v NullableKsearchService) IsSet() bool {
	return v.isSet
}

func (v *NullableKsearchService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKsearchService(val *KsearchService) *NullableKsearchService {
	return &NullableKsearchService{value: val, isSet: true}
}

func (v NullableKsearchService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKsearchService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



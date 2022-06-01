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

// KsearchInstance struct for KsearchInstance
type KsearchInstance struct {
	Id *string `json:"id,omitempty"`
	OrganizationId *string `json:"organization_id,omitempty"`
	AppId *string `json:"app_id,omitempty"`
	ServiceId *string `json:"service_id,omitempty"`
	AllocationId *string `json:"allocation_id,omitempty"`
}

// NewKsearchInstance instantiates a new KsearchInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKsearchInstance() *KsearchInstance {
	this := KsearchInstance{}
	return &this
}

// NewKsearchInstanceWithDefaults instantiates a new KsearchInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKsearchInstanceWithDefaults() *KsearchInstance {
	this := KsearchInstance{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KsearchInstance) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KsearchInstance) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KsearchInstance) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *KsearchInstance) SetId(v string) {
	o.Id = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *KsearchInstance) GetOrganizationId() string {
	if o == nil || o.OrganizationId == nil {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KsearchInstance) GetOrganizationIdOk() (*string, bool) {
	if o == nil || o.OrganizationId == nil {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *KsearchInstance) HasOrganizationId() bool {
	if o != nil && o.OrganizationId != nil {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *KsearchInstance) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *KsearchInstance) GetAppId() string {
	if o == nil || o.AppId == nil {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KsearchInstance) GetAppIdOk() (*string, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *KsearchInstance) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *KsearchInstance) SetAppId(v string) {
	o.AppId = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *KsearchInstance) GetServiceId() string {
	if o == nil || o.ServiceId == nil {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KsearchInstance) GetServiceIdOk() (*string, bool) {
	if o == nil || o.ServiceId == nil {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *KsearchInstance) HasServiceId() bool {
	if o != nil && o.ServiceId != nil {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *KsearchInstance) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetAllocationId returns the AllocationId field value if set, zero value otherwise.
func (o *KsearchInstance) GetAllocationId() string {
	if o == nil || o.AllocationId == nil {
		var ret string
		return ret
	}
	return *o.AllocationId
}

// GetAllocationIdOk returns a tuple with the AllocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KsearchInstance) GetAllocationIdOk() (*string, bool) {
	if o == nil || o.AllocationId == nil {
		return nil, false
	}
	return o.AllocationId, true
}

// HasAllocationId returns a boolean if a field has been set.
func (o *KsearchInstance) HasAllocationId() bool {
	if o != nil && o.AllocationId != nil {
		return true
	}

	return false
}

// SetAllocationId gets a reference to the given string and assigns it to the AllocationId field.
func (o *KsearchInstance) SetAllocationId(v string) {
	o.AllocationId = &v
}

func (o KsearchInstance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.OrganizationId != nil {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if o.AppId != nil {
		toSerialize["app_id"] = o.AppId
	}
	if o.ServiceId != nil {
		toSerialize["service_id"] = o.ServiceId
	}
	if o.AllocationId != nil {
		toSerialize["allocation_id"] = o.AllocationId
	}
	return json.Marshal(toSerialize)
}

type NullableKsearchInstance struct {
	value *KsearchInstance
	isSet bool
}

func (v NullableKsearchInstance) Get() *KsearchInstance {
	return v.value
}

func (v *NullableKsearchInstance) Set(val *KsearchInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableKsearchInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableKsearchInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKsearchInstance(val *KsearchInstance) *NullableKsearchInstance {
	return &NullableKsearchInstance{value: val, isSet: true}
}

func (v NullableKsearchInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKsearchInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


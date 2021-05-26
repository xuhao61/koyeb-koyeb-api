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

// GetCatalogStackReply struct for GetCatalogStackReply
type GetCatalogStackReply struct {
	CatalogStack *CatalogStack `json:"catalog_stack,omitempty"`
}

// NewGetCatalogStackReply instantiates a new GetCatalogStackReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCatalogStackReply() *GetCatalogStackReply {
	this := GetCatalogStackReply{}
	return &this
}

// NewGetCatalogStackReplyWithDefaults instantiates a new GetCatalogStackReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCatalogStackReplyWithDefaults() *GetCatalogStackReply {
	this := GetCatalogStackReply{}
	return &this
}

// GetCatalogStack returns the CatalogStack field value if set, zero value otherwise.
func (o *GetCatalogStackReply) GetCatalogStack() CatalogStack {
	if o == nil || o.CatalogStack == nil {
		var ret CatalogStack
		return ret
	}
	return *o.CatalogStack
}

// GetCatalogStackOk returns a tuple with the CatalogStack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCatalogStackReply) GetCatalogStackOk() (*CatalogStack, bool) {
	if o == nil || o.CatalogStack == nil {
		return nil, false
	}
	return o.CatalogStack, true
}

// HasCatalogStack returns a boolean if a field has been set.
func (o *GetCatalogStackReply) HasCatalogStack() bool {
	if o != nil && o.CatalogStack != nil {
		return true
	}

	return false
}

// SetCatalogStack gets a reference to the given CatalogStack and assigns it to the CatalogStack field.
func (o *GetCatalogStackReply) SetCatalogStack(v CatalogStack) {
	o.CatalogStack = &v
}

func (o GetCatalogStackReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CatalogStack != nil {
		toSerialize["catalog_stack"] = o.CatalogStack
	}
	return json.Marshal(toSerialize)
}

type NullableGetCatalogStackReply struct {
	value *GetCatalogStackReply
	isSet bool
}

func (v NullableGetCatalogStackReply) Get() *GetCatalogStackReply {
	return v.value
}

func (v *NullableGetCatalogStackReply) Set(val *GetCatalogStackReply) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCatalogStackReply) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCatalogStackReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCatalogStackReply(val *GetCatalogStackReply) *NullableGetCatalogStackReply {
	return &NullableGetCatalogStackReply{value: val, isSet: true}
}

func (v NullableGetCatalogStackReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCatalogStackReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


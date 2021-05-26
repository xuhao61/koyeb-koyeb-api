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

// CatalogStoreReply struct for CatalogStoreReply
type CatalogStoreReply struct {
	CatalogStore *CatalogStore `json:"catalog_store,omitempty"`
}

// NewCatalogStoreReply instantiates a new CatalogStoreReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogStoreReply() *CatalogStoreReply {
	this := CatalogStoreReply{}
	return &this
}

// NewCatalogStoreReplyWithDefaults instantiates a new CatalogStoreReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogStoreReplyWithDefaults() *CatalogStoreReply {
	this := CatalogStoreReply{}
	return &this
}

// GetCatalogStore returns the CatalogStore field value if set, zero value otherwise.
func (o *CatalogStoreReply) GetCatalogStore() CatalogStore {
	if o == nil || o.CatalogStore == nil {
		var ret CatalogStore
		return ret
	}
	return *o.CatalogStore
}

// GetCatalogStoreOk returns a tuple with the CatalogStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogStoreReply) GetCatalogStoreOk() (*CatalogStore, bool) {
	if o == nil || o.CatalogStore == nil {
		return nil, false
	}
	return o.CatalogStore, true
}

// HasCatalogStore returns a boolean if a field has been set.
func (o *CatalogStoreReply) HasCatalogStore() bool {
	if o != nil && o.CatalogStore != nil {
		return true
	}

	return false
}

// SetCatalogStore gets a reference to the given CatalogStore and assigns it to the CatalogStore field.
func (o *CatalogStoreReply) SetCatalogStore(v CatalogStore) {
	o.CatalogStore = &v
}

func (o CatalogStoreReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CatalogStore != nil {
		toSerialize["catalog_store"] = o.CatalogStore
	}
	return json.Marshal(toSerialize)
}

type NullableCatalogStoreReply struct {
	value *CatalogStoreReply
	isSet bool
}

func (v NullableCatalogStoreReply) Get() *CatalogStoreReply {
	return v.value
}

func (v *NullableCatalogStoreReply) Set(val *CatalogStoreReply) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogStoreReply) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogStoreReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogStoreReply(val *CatalogStoreReply) *NullableCatalogStoreReply {
	return &NullableCatalogStoreReply{value: val, isSet: true}
}

func (v NullableCatalogStoreReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogStoreReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


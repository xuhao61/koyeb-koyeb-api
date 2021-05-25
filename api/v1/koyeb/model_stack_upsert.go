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

// StackUpsert struct for StackUpsert
type StackUpsert struct {
	Name *string `json:"name,omitempty"`
	Repository *SCMRepository `json:"repository,omitempty"`
}

// NewStackUpsert instantiates a new StackUpsert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStackUpsert() *StackUpsert {
	this := StackUpsert{}
	return &this
}

// NewStackUpsertWithDefaults instantiates a new StackUpsert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStackUpsertWithDefaults() *StackUpsert {
	this := StackUpsert{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StackUpsert) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackUpsert) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StackUpsert) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StackUpsert) SetName(v string) {
	o.Name = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *StackUpsert) GetRepository() SCMRepository {
	if o == nil || o.Repository == nil {
		var ret SCMRepository
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StackUpsert) GetRepositoryOk() (*SCMRepository, bool) {
	if o == nil || o.Repository == nil {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *StackUpsert) HasRepository() bool {
	if o != nil && o.Repository != nil {
		return true
	}

	return false
}

// SetRepository gets a reference to the given SCMRepository and assigns it to the Repository field.
func (o *StackUpsert) SetRepository(v SCMRepository) {
	o.Repository = &v
}

func (o StackUpsert) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Repository != nil {
		toSerialize["repository"] = o.Repository
	}
	return json.Marshal(toSerialize)
}

type NullableStackUpsert struct {
	value *StackUpsert
	isSet bool
}

func (v NullableStackUpsert) Get() *StackUpsert {
	return v.value
}

func (v *NullableStackUpsert) Set(val *StackUpsert) {
	v.value = val
	v.isSet = true
}

func (v NullableStackUpsert) IsSet() bool {
	return v.isSet
}

func (v *NullableStackUpsert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStackUpsert(val *StackUpsert) *NullableStackUpsert {
	return &NullableStackUpsert{value: val, isSet: true}
}

func (v NullableStackUpsert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStackUpsert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// DatabaseRolePassword struct for DatabaseRolePassword
type DatabaseRolePassword struct {
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
}

// NewDatabaseRolePassword instantiates a new DatabaseRolePassword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseRolePassword() *DatabaseRolePassword {
	this := DatabaseRolePassword{}
	return &this
}

// NewDatabaseRolePasswordWithDefaults instantiates a new DatabaseRolePassword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseRolePasswordWithDefaults() *DatabaseRolePassword {
	this := DatabaseRolePassword{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *DatabaseRolePassword) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseRolePassword) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *DatabaseRolePassword) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *DatabaseRolePassword) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *DatabaseRolePassword) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseRolePassword) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
    return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *DatabaseRolePassword) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *DatabaseRolePassword) SetPassword(v string) {
	o.Password = &v
}

func (o DatabaseRolePassword) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableDatabaseRolePassword struct {
	value *DatabaseRolePassword
	isSet bool
}

func (v NullableDatabaseRolePassword) Get() *DatabaseRolePassword {
	return v.value
}

func (v *NullableDatabaseRolePassword) Set(val *DatabaseRolePassword) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseRolePassword) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseRolePassword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseRolePassword(val *DatabaseRolePassword) *NullableDatabaseRolePassword {
	return &NullableDatabaseRolePassword{value: val, isSet: true}
}

func (v NullableDatabaseRolePassword) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseRolePassword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



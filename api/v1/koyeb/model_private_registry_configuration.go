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

// PrivateRegistryConfiguration struct for PrivateRegistryConfiguration
type PrivateRegistryConfiguration struct {
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewPrivateRegistryConfiguration instantiates a new PrivateRegistryConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateRegistryConfiguration() *PrivateRegistryConfiguration {
	this := PrivateRegistryConfiguration{}
	return &this
}

// NewPrivateRegistryConfigurationWithDefaults instantiates a new PrivateRegistryConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateRegistryConfigurationWithDefaults() *PrivateRegistryConfiguration {
	this := PrivateRegistryConfiguration{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *PrivateRegistryConfiguration) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateRegistryConfiguration) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *PrivateRegistryConfiguration) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *PrivateRegistryConfiguration) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *PrivateRegistryConfiguration) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateRegistryConfiguration) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *PrivateRegistryConfiguration) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *PrivateRegistryConfiguration) SetPassword(v string) {
	o.Password = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PrivateRegistryConfiguration) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateRegistryConfiguration) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PrivateRegistryConfiguration) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PrivateRegistryConfiguration) SetUrl(v string) {
	o.Url = &v
}

func (o PrivateRegistryConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullablePrivateRegistryConfiguration struct {
	value *PrivateRegistryConfiguration
	isSet bool
}

func (v NullablePrivateRegistryConfiguration) Get() *PrivateRegistryConfiguration {
	return v.value
}

func (v *NullablePrivateRegistryConfiguration) Set(val *PrivateRegistryConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateRegistryConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateRegistryConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateRegistryConfiguration(val *PrivateRegistryConfiguration) *NullablePrivateRegistryConfiguration {
	return &NullablePrivateRegistryConfiguration{value: val, isSet: true}
}

func (v NullablePrivateRegistryConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateRegistryConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



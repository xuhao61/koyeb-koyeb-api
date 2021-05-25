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

// FunctionParameter struct for FunctionParameter
type FunctionParameter struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Required *bool `json:"required,omitempty"`
	Type *FunctionParameterType `json:"type,omitempty"`
	Default *interface{} `json:"default,omitempty"`
}

// NewFunctionParameter instantiates a new FunctionParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunctionParameter() *FunctionParameter {
	this := FunctionParameter{}
	var type_ FunctionParameterType = FUNCTIONPARAMETERTYPE_UNKNOWN
	this.Type = &type_
	return &this
}

// NewFunctionParameterWithDefaults instantiates a new FunctionParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionParameterWithDefaults() *FunctionParameter {
	this := FunctionParameter{}
	var type_ FunctionParameterType = FUNCTIONPARAMETERTYPE_UNKNOWN
	this.Type = &type_
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FunctionParameter) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionParameter) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FunctionParameter) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FunctionParameter) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FunctionParameter) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionParameter) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FunctionParameter) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FunctionParameter) SetDescription(v string) {
	o.Description = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *FunctionParameter) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionParameter) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *FunctionParameter) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *FunctionParameter) SetRequired(v bool) {
	o.Required = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FunctionParameter) GetType() FunctionParameterType {
	if o == nil || o.Type == nil {
		var ret FunctionParameterType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionParameter) GetTypeOk() (*FunctionParameterType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FunctionParameter) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given FunctionParameterType and assigns it to the Type field.
func (o *FunctionParameter) SetType(v FunctionParameterType) {
	o.Type = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *FunctionParameter) GetDefault() interface{} {
	if o == nil || o.Default == nil {
		var ret interface{}
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionParameter) GetDefaultOk() (*interface{}, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *FunctionParameter) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given interface{} and assigns it to the Default field.
func (o *FunctionParameter) SetDefault(v interface{}) {
	o.Default = &v
}

func (o FunctionParameter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	return json.Marshal(toSerialize)
}

type NullableFunctionParameter struct {
	value *FunctionParameter
	isSet bool
}

func (v NullableFunctionParameter) Get() *FunctionParameter {
	return v.value
}

func (v *NullableFunctionParameter) Set(val *FunctionParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionParameter(val *FunctionParameter) *NullableFunctionParameter {
	return &NullableFunctionParameter{value: val, isSet: true}
}

func (v NullableFunctionParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



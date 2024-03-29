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

// ReviewOrganizationCapacityReply struct for ReviewOrganizationCapacityReply
type ReviewOrganizationCapacityReply struct {
	HasCapacity *bool `json:"has_capacity,omitempty"`
}

// NewReviewOrganizationCapacityReply instantiates a new ReviewOrganizationCapacityReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewOrganizationCapacityReply() *ReviewOrganizationCapacityReply {
	this := ReviewOrganizationCapacityReply{}
	return &this
}

// NewReviewOrganizationCapacityReplyWithDefaults instantiates a new ReviewOrganizationCapacityReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewOrganizationCapacityReplyWithDefaults() *ReviewOrganizationCapacityReply {
	this := ReviewOrganizationCapacityReply{}
	return &this
}

// GetHasCapacity returns the HasCapacity field value if set, zero value otherwise.
func (o *ReviewOrganizationCapacityReply) GetHasCapacity() bool {
	if o == nil || isNil(o.HasCapacity) {
		var ret bool
		return ret
	}
	return *o.HasCapacity
}

// GetHasCapacityOk returns a tuple with the HasCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewOrganizationCapacityReply) GetHasCapacityOk() (*bool, bool) {
	if o == nil || isNil(o.HasCapacity) {
    return nil, false
	}
	return o.HasCapacity, true
}

// HasHasCapacity returns a boolean if a field has been set.
func (o *ReviewOrganizationCapacityReply) HasHasCapacity() bool {
	if o != nil && !isNil(o.HasCapacity) {
		return true
	}

	return false
}

// SetHasCapacity gets a reference to the given bool and assigns it to the HasCapacity field.
func (o *ReviewOrganizationCapacityReply) SetHasCapacity(v bool) {
	o.HasCapacity = &v
}

func (o ReviewOrganizationCapacityReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.HasCapacity) {
		toSerialize["has_capacity"] = o.HasCapacity
	}
	return json.Marshal(toSerialize)
}

type NullableReviewOrganizationCapacityReply struct {
	value *ReviewOrganizationCapacityReply
	isSet bool
}

func (v NullableReviewOrganizationCapacityReply) Get() *ReviewOrganizationCapacityReply {
	return v.value
}

func (v *NullableReviewOrganizationCapacityReply) Set(val *ReviewOrganizationCapacityReply) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewOrganizationCapacityReply) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewOrganizationCapacityReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewOrganizationCapacityReply(val *ReviewOrganizationCapacityReply) *NullableReviewOrganizationCapacityReply {
	return &NullableReviewOrganizationCapacityReply{value: val, isSet: true}
}

func (v NullableReviewOrganizationCapacityReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewOrganizationCapacityReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



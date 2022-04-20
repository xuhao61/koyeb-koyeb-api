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

// GetMetricsReply struct for GetMetricsReply
type GetMetricsReply struct {
	Metrics *[]GetMetricsReplyMetric `json:"metrics,omitempty"`
}

// NewGetMetricsReply instantiates a new GetMetricsReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMetricsReply() *GetMetricsReply {
	this := GetMetricsReply{}
	return &this
}

// NewGetMetricsReplyWithDefaults instantiates a new GetMetricsReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMetricsReplyWithDefaults() *GetMetricsReply {
	this := GetMetricsReply{}
	return &this
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *GetMetricsReply) GetMetrics() []GetMetricsReplyMetric {
	if o == nil || o.Metrics == nil {
		var ret []GetMetricsReplyMetric
		return ret
	}
	return *o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMetricsReply) GetMetricsOk() (*[]GetMetricsReplyMetric, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *GetMetricsReply) HasMetrics() bool {
	if o != nil && o.Metrics != nil {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given []GetMetricsReplyMetric and assigns it to the Metrics field.
func (o *GetMetricsReply) SetMetrics(v []GetMetricsReplyMetric) {
	o.Metrics = &v
}

func (o GetMetricsReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Metrics != nil {
		toSerialize["metrics"] = o.Metrics
	}
	return json.Marshal(toSerialize)
}

type NullableGetMetricsReply struct {
	value *GetMetricsReply
	isSet bool
}

func (v NullableGetMetricsReply) Get() *GetMetricsReply {
	return v.value
}

func (v *NullableGetMetricsReply) Set(val *GetMetricsReply) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMetricsReply) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMetricsReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMetricsReply(val *GetMetricsReply) *NullableGetMetricsReply {
	return &NullableGetMetricsReply{value: val, isSet: true}
}

func (v NullableGetMetricsReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMetricsReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



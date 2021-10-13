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
	"time"
)

// LogEntry struct for LogEntry
type LogEntry struct {
	Msg *string `json:"msg,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Labels *interface{} `json:"labels,omitempty"`
}

// NewLogEntry instantiates a new LogEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogEntry() *LogEntry {
	this := LogEntry{}
	return &this
}

// NewLogEntryWithDefaults instantiates a new LogEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogEntryWithDefaults() *LogEntry {
	this := LogEntry{}
	return &this
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *LogEntry) GetMsg() string {
	if o == nil || o.Msg == nil {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogEntry) GetMsgOk() (*string, bool) {
	if o == nil || o.Msg == nil {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *LogEntry) HasMsg() bool {
	if o != nil && o.Msg != nil {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *LogEntry) SetMsg(v string) {
	o.Msg = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *LogEntry) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogEntry) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LogEntry) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *LogEntry) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *LogEntry) GetLabels() interface{} {
	if o == nil || o.Labels == nil {
		var ret interface{}
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogEntry) GetLabelsOk() (*interface{}, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *LogEntry) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given interface{} and assigns it to the Labels field.
func (o *LogEntry) SetLabels(v interface{}) {
	o.Labels = &v
}

func (o LogEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Msg != nil {
		toSerialize["msg"] = o.Msg
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	return json.Marshal(toSerialize)
}

type NullableLogEntry struct {
	value *LogEntry
	isSet bool
}

func (v NullableLogEntry) Get() *LogEntry {
	return v.value
}

func (v *NullableLogEntry) Set(val *LogEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableLogEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableLogEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogEntry(val *LogEntry) *NullableLogEntry {
	return &NullableLogEntry{value: val, isSet: true}
}

func (v NullableLogEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



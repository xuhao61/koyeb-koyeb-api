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
	Message *string `json:"message,omitempty"`
	Level *string `json:"level,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Stream *string `json:"stream,omitempty"`
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

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *LogEntry) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogEntry) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *LogEntry) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *LogEntry) SetMessage(v string) {
	o.Message = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *LogEntry) GetLevel() string {
	if o == nil || o.Level == nil {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogEntry) GetLevelOk() (*string, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *LogEntry) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *LogEntry) SetLevel(v string) {
	o.Level = &v
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

// GetStream returns the Stream field value if set, zero value otherwise.
func (o *LogEntry) GetStream() string {
	if o == nil || o.Stream == nil {
		var ret string
		return ret
	}
	return *o.Stream
}

// GetStreamOk returns a tuple with the Stream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogEntry) GetStreamOk() (*string, bool) {
	if o == nil || o.Stream == nil {
		return nil, false
	}
	return o.Stream, true
}

// HasStream returns a boolean if a field has been set.
func (o *LogEntry) HasStream() bool {
	if o != nil && o.Stream != nil {
		return true
	}

	return false
}

// SetStream gets a reference to the given string and assigns it to the Stream field.
func (o *LogEntry) SetStream(v string) {
	o.Stream = &v
}

func (o LogEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Stream != nil {
		toSerialize["stream"] = o.Stream
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


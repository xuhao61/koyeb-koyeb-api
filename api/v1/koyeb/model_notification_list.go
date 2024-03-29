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

// NotificationList struct for NotificationList
type NotificationList struct {
	Notifications []Notification `json:"notifications,omitempty"`
	Limit *int64 `json:"limit,omitempty"`
	Offset *int64 `json:"offset,omitempty"`
	Count *int64 `json:"count,omitempty"`
	IsRead *bool `json:"is_read,omitempty"`
	IsSeen *bool `json:"is_seen,omitempty"`
	Unread *int64 `json:"unread,omitempty"`
	Unseen *int64 `json:"unseen,omitempty"`
}

// NewNotificationList instantiates a new NotificationList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationList() *NotificationList {
	this := NotificationList{}
	return &this
}

// NewNotificationListWithDefaults instantiates a new NotificationList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationListWithDefaults() *NotificationList {
	this := NotificationList{}
	return &this
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *NotificationList) GetNotifications() []Notification {
	if o == nil || isNil(o.Notifications) {
		var ret []Notification
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationList) GetNotificationsOk() ([]Notification, bool) {
	if o == nil || isNil(o.Notifications) {
    return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *NotificationList) HasNotifications() bool {
	if o != nil && !isNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given []Notification and assigns it to the Notifications field.
func (o *NotificationList) SetNotifications(v []Notification) {
	o.Notifications = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *NotificationList) GetLimit() int64 {
	if o == nil || isNil(o.Limit) {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationList) GetLimitOk() (*int64, bool) {
	if o == nil || isNil(o.Limit) {
    return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *NotificationList) HasLimit() bool {
	if o != nil && !isNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *NotificationList) SetLimit(v int64) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *NotificationList) GetOffset() int64 {
	if o == nil || isNil(o.Offset) {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationList) GetOffsetOk() (*int64, bool) {
	if o == nil || isNil(o.Offset) {
    return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *NotificationList) HasOffset() bool {
	if o != nil && !isNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *NotificationList) SetOffset(v int64) {
	o.Offset = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *NotificationList) GetCount() int64 {
	if o == nil || isNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationList) GetCountOk() (*int64, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *NotificationList) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *NotificationList) SetCount(v int64) {
	o.Count = &v
}

// GetIsRead returns the IsRead field value if set, zero value otherwise.
func (o *NotificationList) GetIsRead() bool {
	if o == nil || isNil(o.IsRead) {
		var ret bool
		return ret
	}
	return *o.IsRead
}

// GetIsReadOk returns a tuple with the IsRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationList) GetIsReadOk() (*bool, bool) {
	if o == nil || isNil(o.IsRead) {
    return nil, false
	}
	return o.IsRead, true
}

// HasIsRead returns a boolean if a field has been set.
func (o *NotificationList) HasIsRead() bool {
	if o != nil && !isNil(o.IsRead) {
		return true
	}

	return false
}

// SetIsRead gets a reference to the given bool and assigns it to the IsRead field.
func (o *NotificationList) SetIsRead(v bool) {
	o.IsRead = &v
}

// GetIsSeen returns the IsSeen field value if set, zero value otherwise.
func (o *NotificationList) GetIsSeen() bool {
	if o == nil || isNil(o.IsSeen) {
		var ret bool
		return ret
	}
	return *o.IsSeen
}

// GetIsSeenOk returns a tuple with the IsSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationList) GetIsSeenOk() (*bool, bool) {
	if o == nil || isNil(o.IsSeen) {
    return nil, false
	}
	return o.IsSeen, true
}

// HasIsSeen returns a boolean if a field has been set.
func (o *NotificationList) HasIsSeen() bool {
	if o != nil && !isNil(o.IsSeen) {
		return true
	}

	return false
}

// SetIsSeen gets a reference to the given bool and assigns it to the IsSeen field.
func (o *NotificationList) SetIsSeen(v bool) {
	o.IsSeen = &v
}

// GetUnread returns the Unread field value if set, zero value otherwise.
func (o *NotificationList) GetUnread() int64 {
	if o == nil || isNil(o.Unread) {
		var ret int64
		return ret
	}
	return *o.Unread
}

// GetUnreadOk returns a tuple with the Unread field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationList) GetUnreadOk() (*int64, bool) {
	if o == nil || isNil(o.Unread) {
    return nil, false
	}
	return o.Unread, true
}

// HasUnread returns a boolean if a field has been set.
func (o *NotificationList) HasUnread() bool {
	if o != nil && !isNil(o.Unread) {
		return true
	}

	return false
}

// SetUnread gets a reference to the given int64 and assigns it to the Unread field.
func (o *NotificationList) SetUnread(v int64) {
	o.Unread = &v
}

// GetUnseen returns the Unseen field value if set, zero value otherwise.
func (o *NotificationList) GetUnseen() int64 {
	if o == nil || isNil(o.Unseen) {
		var ret int64
		return ret
	}
	return *o.Unseen
}

// GetUnseenOk returns a tuple with the Unseen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationList) GetUnseenOk() (*int64, bool) {
	if o == nil || isNil(o.Unseen) {
    return nil, false
	}
	return o.Unseen, true
}

// HasUnseen returns a boolean if a field has been set.
func (o *NotificationList) HasUnseen() bool {
	if o != nil && !isNil(o.Unseen) {
		return true
	}

	return false
}

// SetUnseen gets a reference to the given int64 and assigns it to the Unseen field.
func (o *NotificationList) SetUnseen(v int64) {
	o.Unseen = &v
}

func (o NotificationList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Notifications) {
		toSerialize["notifications"] = o.Notifications
	}
	if !isNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !isNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !isNil(o.IsRead) {
		toSerialize["is_read"] = o.IsRead
	}
	if !isNil(o.IsSeen) {
		toSerialize["is_seen"] = o.IsSeen
	}
	if !isNil(o.Unread) {
		toSerialize["unread"] = o.Unread
	}
	if !isNil(o.Unseen) {
		toSerialize["unseen"] = o.Unseen
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationList struct {
	value *NotificationList
	isSet bool
}

func (v NullableNotificationList) Get() *NotificationList {
	return v.value
}

func (v *NullableNotificationList) Set(val *NotificationList) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationList) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationList(val *NotificationList) *NullableNotificationList {
	return &NullableNotificationList{value: val, isSet: true}
}

func (v NullableNotificationList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



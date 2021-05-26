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

// User Represent a User
type User struct {
	Id *string `json:"id,omitempty"`
	Email *string `json:"email,omitempty"`
	Status *UserAccountStatus `json:"status,omitempty"`
	StatusMessage *string `json:"status_message,omitempty"`
	AvatarUrl *string `json:"avatar_url,omitempty"`
	Verified *bool `json:"verified,omitempty"`
	TwoFactorAuthentication *bool `json:"two_factor_authentication,omitempty"`
	LastLogin *time.Time `json:"last_login,omitempty"`
	LastLoginIp *string `json:"last_login_ip,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	DefaultOrganizationId *string `json:"default_organization_id,omitempty"`
	NewsletterSubscribed *bool `json:"newsletter_subscribed,omitempty"`
	GithubId *string `json:"github_id,omitempty"`
	GithubUser *string `json:"github_user,omitempty"`
	Flags *[]UserFlags `json:"flags,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser() *User {
	this := User{}
	var status UserAccountStatus = USERACCOUNTSTATUS_WARNING
	this.Status = &status
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	var status UserAccountStatus = USERACCOUNTSTATUS_WARNING
	this.Status = &status
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *User) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *User) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *User) SetId(v string) {
	o.Id = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *User) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *User) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *User) SetEmail(v string) {
	o.Email = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *User) GetStatus() UserAccountStatus {
	if o == nil || o.Status == nil {
		var ret UserAccountStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetStatusOk() (*UserAccountStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *User) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given UserAccountStatus and assigns it to the Status field.
func (o *User) SetStatus(v UserAccountStatus) {
	o.Status = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *User) GetStatusMessage() string {
	if o == nil || o.StatusMessage == nil {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetStatusMessageOk() (*string, bool) {
	if o == nil || o.StatusMessage == nil {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *User) HasStatusMessage() bool {
	if o != nil && o.StatusMessage != nil {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *User) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise.
func (o *User) GetAvatarUrl() string {
	if o == nil || o.AvatarUrl == nil {
		var ret string
		return ret
	}
	return *o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAvatarUrlOk() (*string, bool) {
	if o == nil || o.AvatarUrl == nil {
		return nil, false
	}
	return o.AvatarUrl, true
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *User) HasAvatarUrl() bool {
	if o != nil && o.AvatarUrl != nil {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given string and assigns it to the AvatarUrl field.
func (o *User) SetAvatarUrl(v string) {
	o.AvatarUrl = &v
}

// GetVerified returns the Verified field value if set, zero value otherwise.
func (o *User) GetVerified() bool {
	if o == nil || o.Verified == nil {
		var ret bool
		return ret
	}
	return *o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetVerifiedOk() (*bool, bool) {
	if o == nil || o.Verified == nil {
		return nil, false
	}
	return o.Verified, true
}

// HasVerified returns a boolean if a field has been set.
func (o *User) HasVerified() bool {
	if o != nil && o.Verified != nil {
		return true
	}

	return false
}

// SetVerified gets a reference to the given bool and assigns it to the Verified field.
func (o *User) SetVerified(v bool) {
	o.Verified = &v
}

// GetTwoFactorAuthentication returns the TwoFactorAuthentication field value if set, zero value otherwise.
func (o *User) GetTwoFactorAuthentication() bool {
	if o == nil || o.TwoFactorAuthentication == nil {
		var ret bool
		return ret
	}
	return *o.TwoFactorAuthentication
}

// GetTwoFactorAuthenticationOk returns a tuple with the TwoFactorAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetTwoFactorAuthenticationOk() (*bool, bool) {
	if o == nil || o.TwoFactorAuthentication == nil {
		return nil, false
	}
	return o.TwoFactorAuthentication, true
}

// HasTwoFactorAuthentication returns a boolean if a field has been set.
func (o *User) HasTwoFactorAuthentication() bool {
	if o != nil && o.TwoFactorAuthentication != nil {
		return true
	}

	return false
}

// SetTwoFactorAuthentication gets a reference to the given bool and assigns it to the TwoFactorAuthentication field.
func (o *User) SetTwoFactorAuthentication(v bool) {
	o.TwoFactorAuthentication = &v
}

// GetLastLogin returns the LastLogin field value if set, zero value otherwise.
func (o *User) GetLastLogin() time.Time {
	if o == nil || o.LastLogin == nil {
		var ret time.Time
		return ret
	}
	return *o.LastLogin
}

// GetLastLoginOk returns a tuple with the LastLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLastLoginOk() (*time.Time, bool) {
	if o == nil || o.LastLogin == nil {
		return nil, false
	}
	return o.LastLogin, true
}

// HasLastLogin returns a boolean if a field has been set.
func (o *User) HasLastLogin() bool {
	if o != nil && o.LastLogin != nil {
		return true
	}

	return false
}

// SetLastLogin gets a reference to the given time.Time and assigns it to the LastLogin field.
func (o *User) SetLastLogin(v time.Time) {
	o.LastLogin = &v
}

// GetLastLoginIp returns the LastLoginIp field value if set, zero value otherwise.
func (o *User) GetLastLoginIp() string {
	if o == nil || o.LastLoginIp == nil {
		var ret string
		return ret
	}
	return *o.LastLoginIp
}

// GetLastLoginIpOk returns a tuple with the LastLoginIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLastLoginIpOk() (*string, bool) {
	if o == nil || o.LastLoginIp == nil {
		return nil, false
	}
	return o.LastLoginIp, true
}

// HasLastLoginIp returns a boolean if a field has been set.
func (o *User) HasLastLoginIp() bool {
	if o != nil && o.LastLoginIp != nil {
		return true
	}

	return false
}

// SetLastLoginIp gets a reference to the given string and assigns it to the LastLoginIp field.
func (o *User) SetLastLoginIp(v string) {
	o.LastLoginIp = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *User) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *User) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *User) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *User) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *User) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *User) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDefaultOrganizationId returns the DefaultOrganizationId field value if set, zero value otherwise.
func (o *User) GetDefaultOrganizationId() string {
	if o == nil || o.DefaultOrganizationId == nil {
		var ret string
		return ret
	}
	return *o.DefaultOrganizationId
}

// GetDefaultOrganizationIdOk returns a tuple with the DefaultOrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetDefaultOrganizationIdOk() (*string, bool) {
	if o == nil || o.DefaultOrganizationId == nil {
		return nil, false
	}
	return o.DefaultOrganizationId, true
}

// HasDefaultOrganizationId returns a boolean if a field has been set.
func (o *User) HasDefaultOrganizationId() bool {
	if o != nil && o.DefaultOrganizationId != nil {
		return true
	}

	return false
}

// SetDefaultOrganizationId gets a reference to the given string and assigns it to the DefaultOrganizationId field.
func (o *User) SetDefaultOrganizationId(v string) {
	o.DefaultOrganizationId = &v
}

// GetNewsletterSubscribed returns the NewsletterSubscribed field value if set, zero value otherwise.
func (o *User) GetNewsletterSubscribed() bool {
	if o == nil || o.NewsletterSubscribed == nil {
		var ret bool
		return ret
	}
	return *o.NewsletterSubscribed
}

// GetNewsletterSubscribedOk returns a tuple with the NewsletterSubscribed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetNewsletterSubscribedOk() (*bool, bool) {
	if o == nil || o.NewsletterSubscribed == nil {
		return nil, false
	}
	return o.NewsletterSubscribed, true
}

// HasNewsletterSubscribed returns a boolean if a field has been set.
func (o *User) HasNewsletterSubscribed() bool {
	if o != nil && o.NewsletterSubscribed != nil {
		return true
	}

	return false
}

// SetNewsletterSubscribed gets a reference to the given bool and assigns it to the NewsletterSubscribed field.
func (o *User) SetNewsletterSubscribed(v bool) {
	o.NewsletterSubscribed = &v
}

// GetGithubId returns the GithubId field value if set, zero value otherwise.
func (o *User) GetGithubId() string {
	if o == nil || o.GithubId == nil {
		var ret string
		return ret
	}
	return *o.GithubId
}

// GetGithubIdOk returns a tuple with the GithubId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetGithubIdOk() (*string, bool) {
	if o == nil || o.GithubId == nil {
		return nil, false
	}
	return o.GithubId, true
}

// HasGithubId returns a boolean if a field has been set.
func (o *User) HasGithubId() bool {
	if o != nil && o.GithubId != nil {
		return true
	}

	return false
}

// SetGithubId gets a reference to the given string and assigns it to the GithubId field.
func (o *User) SetGithubId(v string) {
	o.GithubId = &v
}

// GetGithubUser returns the GithubUser field value if set, zero value otherwise.
func (o *User) GetGithubUser() string {
	if o == nil || o.GithubUser == nil {
		var ret string
		return ret
	}
	return *o.GithubUser
}

// GetGithubUserOk returns a tuple with the GithubUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetGithubUserOk() (*string, bool) {
	if o == nil || o.GithubUser == nil {
		return nil, false
	}
	return o.GithubUser, true
}

// HasGithubUser returns a boolean if a field has been set.
func (o *User) HasGithubUser() bool {
	if o != nil && o.GithubUser != nil {
		return true
	}

	return false
}

// SetGithubUser gets a reference to the given string and assigns it to the GithubUser field.
func (o *User) SetGithubUser(v string) {
	o.GithubUser = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *User) GetFlags() []UserFlags {
	if o == nil || o.Flags == nil {
		var ret []UserFlags
		return ret
	}
	return *o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetFlagsOk() (*[]UserFlags, bool) {
	if o == nil || o.Flags == nil {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *User) HasFlags() bool {
	if o != nil && o.Flags != nil {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []UserFlags and assigns it to the Flags field.
func (o *User) SetFlags(v []UserFlags) {
	o.Flags = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *User) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *User) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *User) SetName(v string) {
	o.Name = &v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StatusMessage != nil {
		toSerialize["status_message"] = o.StatusMessage
	}
	if o.AvatarUrl != nil {
		toSerialize["avatar_url"] = o.AvatarUrl
	}
	if o.Verified != nil {
		toSerialize["verified"] = o.Verified
	}
	if o.TwoFactorAuthentication != nil {
		toSerialize["two_factor_authentication"] = o.TwoFactorAuthentication
	}
	if o.LastLogin != nil {
		toSerialize["last_login"] = o.LastLogin
	}
	if o.LastLoginIp != nil {
		toSerialize["last_login_ip"] = o.LastLoginIp
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.DefaultOrganizationId != nil {
		toSerialize["default_organization_id"] = o.DefaultOrganizationId
	}
	if o.NewsletterSubscribed != nil {
		toSerialize["newsletter_subscribed"] = o.NewsletterSubscribed
	}
	if o.GithubId != nil {
		toSerialize["github_id"] = o.GithubId
	}
	if o.GithubUser != nil {
		toSerialize["github_user"] = o.GithubUser
	}
	if o.Flags != nil {
		toSerialize["flags"] = o.Flags
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// CatalogStack struct for CatalogStack
type CatalogStack struct {
	Name *string `json:"name,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	ShortDescription *string `json:"short_description,omitempty"`
	Description *string `json:"description,omitempty"`
	Icon *string `json:"icon,omitempty"`
	Status *BaseCatalogStatus `json:"status,omitempty"`
	Tags *[]string `json:"tags,omitempty"`
	SourceControlRef *SCMReference `json:"source_control_ref,omitempty"`
	// Whether this stack is only a yaml or not (in which case it can be created as a non git managed stack).
	IsYamlOnly *bool `json:"is_yaml_only,omitempty"`
	// The yaml that is represented the stack.
	Yaml *string `json:"yaml,omitempty"`
}

// NewCatalogStack instantiates a new CatalogStack object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogStack() *CatalogStack {
	this := CatalogStack{}
	var status BaseCatalogStatus = BASECATALOGSTATUS_COMING_SOON
	this.Status = &status
	return &this
}

// NewCatalogStackWithDefaults instantiates a new CatalogStack object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogStackWithDefaults() *CatalogStack {
	this := CatalogStack{}
	var status BaseCatalogStatus = BASECATALOGSTATUS_COMING_SOON
	this.Status = &status
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CatalogStack) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogStack) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CatalogStack) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CatalogStack) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CatalogStack) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogStack) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CatalogStack) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CatalogStack) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetShortDescription returns the ShortDescription field value if set, zero value otherwise.
func (o *CatalogStack) GetShortDescription() string {
	if o == nil || o.ShortDescription == nil {
		var ret string
		return ret
	}
	return *o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogStack) GetShortDescriptionOk() (*string, bool) {
	if o == nil || o.ShortDescription == nil {
		return nil, false
	}
	return o.ShortDescription, true
}

// HasShortDescription returns a boolean if a field has been set.
func (o *CatalogStack) HasShortDescription() bool {
	if o != nil && o.ShortDescription != nil {
		return true
	}

	return false
}

// SetShortDescription gets a reference to the given string and assigns it to the ShortDescription field.
func (o *CatalogStack) SetShortDescription(v string) {
	o.ShortDescription = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CatalogStack) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogStack) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CatalogStack) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CatalogStack) SetDescription(v string) {
	o.Description = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *CatalogStack) GetIcon() string {
	if o == nil || o.Icon == nil {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogStack) GetIconOk() (*string, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *CatalogStack) HasIcon() bool {
	if o != nil && o.Icon != nil {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *CatalogStack) SetIcon(v string) {
	o.Icon = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CatalogStack) GetStatus() BaseCatalogStatus {
	if o == nil || o.Status == nil {
		var ret BaseCatalogStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogStack) GetStatusOk() (*BaseCatalogStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CatalogStack) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given BaseCatalogStatus and assigns it to the Status field.
func (o *CatalogStack) SetStatus(v BaseCatalogStatus) {
	o.Status = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CatalogStack) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogStack) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CatalogStack) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CatalogStack) SetTags(v []string) {
	o.Tags = &v
}

// GetSourceControlRef returns the SourceControlRef field value if set, zero value otherwise.
func (o *CatalogStack) GetSourceControlRef() SCMReference {
	if o == nil || o.SourceControlRef == nil {
		var ret SCMReference
		return ret
	}
	return *o.SourceControlRef
}

// GetSourceControlRefOk returns a tuple with the SourceControlRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogStack) GetSourceControlRefOk() (*SCMReference, bool) {
	if o == nil || o.SourceControlRef == nil {
		return nil, false
	}
	return o.SourceControlRef, true
}

// HasSourceControlRef returns a boolean if a field has been set.
func (o *CatalogStack) HasSourceControlRef() bool {
	if o != nil && o.SourceControlRef != nil {
		return true
	}

	return false
}

// SetSourceControlRef gets a reference to the given SCMReference and assigns it to the SourceControlRef field.
func (o *CatalogStack) SetSourceControlRef(v SCMReference) {
	o.SourceControlRef = &v
}

// GetIsYamlOnly returns the IsYamlOnly field value if set, zero value otherwise.
func (o *CatalogStack) GetIsYamlOnly() bool {
	if o == nil || o.IsYamlOnly == nil {
		var ret bool
		return ret
	}
	return *o.IsYamlOnly
}

// GetIsYamlOnlyOk returns a tuple with the IsYamlOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogStack) GetIsYamlOnlyOk() (*bool, bool) {
	if o == nil || o.IsYamlOnly == nil {
		return nil, false
	}
	return o.IsYamlOnly, true
}

// HasIsYamlOnly returns a boolean if a field has been set.
func (o *CatalogStack) HasIsYamlOnly() bool {
	if o != nil && o.IsYamlOnly != nil {
		return true
	}

	return false
}

// SetIsYamlOnly gets a reference to the given bool and assigns it to the IsYamlOnly field.
func (o *CatalogStack) SetIsYamlOnly(v bool) {
	o.IsYamlOnly = &v
}

// GetYaml returns the Yaml field value if set, zero value otherwise.
func (o *CatalogStack) GetYaml() string {
	if o == nil || o.Yaml == nil {
		var ret string
		return ret
	}
	return *o.Yaml
}

// GetYamlOk returns a tuple with the Yaml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogStack) GetYamlOk() (*string, bool) {
	if o == nil || o.Yaml == nil {
		return nil, false
	}
	return o.Yaml, true
}

// HasYaml returns a boolean if a field has been set.
func (o *CatalogStack) HasYaml() bool {
	if o != nil && o.Yaml != nil {
		return true
	}

	return false
}

// SetYaml gets a reference to the given string and assigns it to the Yaml field.
func (o *CatalogStack) SetYaml(v string) {
	o.Yaml = &v
}

func (o CatalogStack) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.ShortDescription != nil {
		toSerialize["short_description"] = o.ShortDescription
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Icon != nil {
		toSerialize["icon"] = o.Icon
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.SourceControlRef != nil {
		toSerialize["source_control_ref"] = o.SourceControlRef
	}
	if o.IsYamlOnly != nil {
		toSerialize["is_yaml_only"] = o.IsYamlOnly
	}
	if o.Yaml != nil {
		toSerialize["yaml"] = o.Yaml
	}
	return json.Marshal(toSerialize)
}

type NullableCatalogStack struct {
	value *CatalogStack
	isSet bool
}

func (v NullableCatalogStack) Get() *CatalogStack {
	return v.value
}

func (v *NullableCatalogStack) Set(val *CatalogStack) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogStack) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogStack) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogStack(val *CatalogStack) *NullableCatalogStack {
	return &NullableCatalogStack{value: val, isSet: true}
}

func (v NullableCatalogStack) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogStack) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



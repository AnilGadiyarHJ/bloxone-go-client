/*
Discovery Configuration API V2

The Discovery configuration service is a BloxOne Service that provides configuration for accessing and syncing the Cloud assets   Base Paths:  1. provider: **_/api/cloud_discovery/v2/_**

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clouddiscovery

import (
	"encoding/json"
)

// checks if the DdidnsrickettsSubAccountProvCredConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DdidnsrickettsSubAccountProvCredConfig{}

// DdidnsrickettsSubAccountProvCredConfig struct for DdidnsrickettsSubAccountProvCredConfig
type DdidnsrickettsSubAccountProvCredConfig struct {
	ProjectId            *string `json:"project_id,omitempty"`
	RoleArn              *string `json:"role_arn,omitempty"`
	TenantId             *string `json:"tenant_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DdidnsrickettsSubAccountProvCredConfig DdidnsrickettsSubAccountProvCredConfig

// NewDdidnsrickettsSubAccountProvCredConfig instantiates a new DdidnsrickettsSubAccountProvCredConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdidnsrickettsSubAccountProvCredConfig() *DdidnsrickettsSubAccountProvCredConfig {
	this := DdidnsrickettsSubAccountProvCredConfig{}
	return &this
}

// NewDdidnsrickettsSubAccountProvCredConfigWithDefaults instantiates a new DdidnsrickettsSubAccountProvCredConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdidnsrickettsSubAccountProvCredConfigWithDefaults() *DdidnsrickettsSubAccountProvCredConfig {
	this := DdidnsrickettsSubAccountProvCredConfig{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *DdidnsrickettsSubAccountProvCredConfig) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdidnsrickettsSubAccountProvCredConfig) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *DdidnsrickettsSubAccountProvCredConfig) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *DdidnsrickettsSubAccountProvCredConfig) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetRoleArn returns the RoleArn field value if set, zero value otherwise.
func (o *DdidnsrickettsSubAccountProvCredConfig) GetRoleArn() string {
	if o == nil || IsNil(o.RoleArn) {
		var ret string
		return ret
	}
	return *o.RoleArn
}

// GetRoleArnOk returns a tuple with the RoleArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdidnsrickettsSubAccountProvCredConfig) GetRoleArnOk() (*string, bool) {
	if o == nil || IsNil(o.RoleArn) {
		return nil, false
	}
	return o.RoleArn, true
}

// HasRoleArn returns a boolean if a field has been set.
func (o *DdidnsrickettsSubAccountProvCredConfig) HasRoleArn() bool {
	if o != nil && !IsNil(o.RoleArn) {
		return true
	}

	return false
}

// SetRoleArn gets a reference to the given string and assigns it to the RoleArn field.
func (o *DdidnsrickettsSubAccountProvCredConfig) SetRoleArn(v string) {
	o.RoleArn = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *DdidnsrickettsSubAccountProvCredConfig) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdidnsrickettsSubAccountProvCredConfig) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *DdidnsrickettsSubAccountProvCredConfig) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *DdidnsrickettsSubAccountProvCredConfig) SetTenantId(v string) {
	o.TenantId = &v
}

func (o DdidnsrickettsSubAccountProvCredConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DdidnsrickettsSubAccountProvCredConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	if !IsNil(o.RoleArn) {
		toSerialize["role_arn"] = o.RoleArn
	}
	if !IsNil(o.TenantId) {
		toSerialize["tenant_id"] = o.TenantId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DdidnsrickettsSubAccountProvCredConfig) UnmarshalJSON(data []byte) (err error) {
	varDdidnsrickettsSubAccountProvCredConfig := _DdidnsrickettsSubAccountProvCredConfig{}

	err = json.Unmarshal(data, &varDdidnsrickettsSubAccountProvCredConfig)

	if err != nil {
		return err
	}

	*o = DdidnsrickettsSubAccountProvCredConfig(varDdidnsrickettsSubAccountProvCredConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "project_id")
		delete(additionalProperties, "role_arn")
		delete(additionalProperties, "tenant_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDdidnsrickettsSubAccountProvCredConfig struct {
	value *DdidnsrickettsSubAccountProvCredConfig
	isSet bool
}

func (v NullableDdidnsrickettsSubAccountProvCredConfig) Get() *DdidnsrickettsSubAccountProvCredConfig {
	return v.value
}

func (v *NullableDdidnsrickettsSubAccountProvCredConfig) Set(val *DdidnsrickettsSubAccountProvCredConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDdidnsrickettsSubAccountProvCredConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDdidnsrickettsSubAccountProvCredConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDdidnsrickettsSubAccountProvCredConfig(val *DdidnsrickettsSubAccountProvCredConfig) *NullableDdidnsrickettsSubAccountProvCredConfig {
	return &NullableDdidnsrickettsSubAccountProvCredConfig{value: val, isSet: true}
}

func (v NullableDdidnsrickettsSubAccountProvCredConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDdidnsrickettsSubAccountProvCredConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
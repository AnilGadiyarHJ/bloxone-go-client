/*
Discovery Configuration API V2

The Discovery configuration service is a BloxOne Service that provides configuration for accessing and syncing the Cloud assets   Base Paths:  1. provider: **_/api/cloud_discovery/v2/_**

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clouddiscovery

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the DdidnsrickettsDestination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DdidnsrickettsDestination{}

// DdidnsrickettsDestination Destination information
type DdidnsrickettsDestination struct {
	// Destination configuration. Ex.: '{  \"dns\": {    \"view_name\": \"view 1\",    \"view_id\": \"dns/view/v1\",    \"consolidated_zone_data_enabled\": false,    \"sync_type\": \"read_only/read_write\"    \"split_view_enabled\": false  },  \"ipam\": {    \"ip_space\": \"\",  },  \"account\": {},  }'.
	Config *DdidnsrickettsDestinationConfig `json:"config,omitempty"`
	// Timestamp when the object has been created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Timestamp when the object has been deleted.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Destination type: DNS / IPAM / ACCOUNT.
	DestinationType string `json:"destination_type"`
	// Auto-generated unique destination ID. Format BloxID.
	Id *string `json:"id,omitempty"`
	// Timestamp when the object has been updated.
	UpdatedAt            *time.Time `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DdidnsrickettsDestination DdidnsrickettsDestination

// NewDdidnsrickettsDestination instantiates a new DdidnsrickettsDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdidnsrickettsDestination(destinationType string) *DdidnsrickettsDestination {
	this := DdidnsrickettsDestination{}
	this.DestinationType = destinationType
	return &this
}

// NewDdidnsrickettsDestinationWithDefaults instantiates a new DdidnsrickettsDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdidnsrickettsDestinationWithDefaults() *DdidnsrickettsDestination {
	this := DdidnsrickettsDestination{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *DdidnsrickettsDestination) GetConfig() DdidnsrickettsDestinationConfig {
	if o == nil || IsNil(o.Config) {
		var ret DdidnsrickettsDestinationConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdidnsrickettsDestination) GetConfigOk() (*DdidnsrickettsDestinationConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *DdidnsrickettsDestination) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given DdidnsrickettsDestinationConfig and assigns it to the Config field.
func (o *DdidnsrickettsDestination) SetConfig(v DdidnsrickettsDestinationConfig) {
	o.Config = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DdidnsrickettsDestination) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdidnsrickettsDestination) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DdidnsrickettsDestination) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DdidnsrickettsDestination) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *DdidnsrickettsDestination) GetDeletedAt() time.Time {
	if o == nil || IsNil(o.DeletedAt) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdidnsrickettsDestination) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *DdidnsrickettsDestination) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *DdidnsrickettsDestination) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetDestinationType returns the DestinationType field value
func (o *DdidnsrickettsDestination) GetDestinationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationType
}

// GetDestinationTypeOk returns a tuple with the DestinationType field value
// and a boolean to check if the value has been set.
func (o *DdidnsrickettsDestination) GetDestinationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationType, true
}

// SetDestinationType sets field value
func (o *DdidnsrickettsDestination) SetDestinationType(v string) {
	o.DestinationType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DdidnsrickettsDestination) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdidnsrickettsDestination) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DdidnsrickettsDestination) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DdidnsrickettsDestination) SetId(v string) {
	o.Id = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DdidnsrickettsDestination) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdidnsrickettsDestination) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DdidnsrickettsDestination) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DdidnsrickettsDestination) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o DdidnsrickettsDestination) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DdidnsrickettsDestination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	toSerialize["destination_type"] = o.DestinationType
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DdidnsrickettsDestination) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"destination_type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDdidnsrickettsDestination := _DdidnsrickettsDestination{}

	err = json.Unmarshal(data, &varDdidnsrickettsDestination)

	if err != nil {
		return err
	}

	*o = DdidnsrickettsDestination(varDdidnsrickettsDestination)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "config")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "destination_type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDdidnsrickettsDestination struct {
	value *DdidnsrickettsDestination
	isSet bool
}

func (v NullableDdidnsrickettsDestination) Get() *DdidnsrickettsDestination {
	return v.value
}

func (v *NullableDdidnsrickettsDestination) Set(val *DdidnsrickettsDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableDdidnsrickettsDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableDdidnsrickettsDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDdidnsrickettsDestination(val *DdidnsrickettsDestination) *NullableDdidnsrickettsDestination {
	return &NullableDdidnsrickettsDestination{value: val, isSet: true}
}

func (v NullableDdidnsrickettsDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDdidnsrickettsDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

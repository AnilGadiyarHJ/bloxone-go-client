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

// checks if the DdidnsrickettsSourceConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DdidnsrickettsSourceConfig{}

// DdidnsrickettsSourceConfig Source configuration
type DdidnsrickettsSourceConfig struct {
	// Account Schedule ID.
	AccountScheduleId *string                 `json:"account_schedule_id,omitempty"`
	Accounts          []DdidnsrickettsAccount `json:"accounts,omitempty"`
	// Cloud Credential ID.
	CloudCredentialId string `json:"cloud_credential_id"`
	// Timestamp when the object has been created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Credential configuration. Ex.: '{    \"access_identifier\": \"arn:aws:iam::1234:role/access_for_discovery\",    \"region\": \"us-east-1\",    \"enclave\": \"commercial/gov\"  }'.
	CredentialConfig *DdidnsrickettsCredentialConfig `json:"credential_config,omitempty"`
	// Timestamp when the object has been deleted.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Auto-generated unique source config ID. Format BloxID.
	Id *string `json:"id,omitempty"`
	// Provider account IDs such as accountID/ SubscriptionID to be restricted for a given source_config.
	RestrictedToAccounts []string `json:"restricted_to_accounts,omitempty"`
	// Timestamp when the object has been updated.
	UpdatedAt            *time.Time `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DdidnsrickettsSourceConfig DdidnsrickettsSourceConfig

// NewDdidnsrickettsSourceConfig instantiates a new DdidnsrickettsSourceConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdidnsrickettsSourceConfig(cloudCredentialId string) *DdidnsrickettsSourceConfig {
	this := DdidnsrickettsSourceConfig{}
	this.CloudCredentialId = cloudCredentialId
	return &this
}

// NewDdidnsrickettsSourceConfigWithDefaults instantiates a new DdidnsrickettsSourceConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdidnsrickettsSourceConfigWithDefaults() *DdidnsrickettsSourceConfig {
	this := DdidnsrickettsSourceConfig{}
	return &this
}

// GetAccountScheduleId returns the AccountScheduleId field value if set, zero value otherwise.
func (o *DdidnsrickettsSourceConfig) GetAccountScheduleId() string {
	if o == nil || IsNil(o.AccountScheduleId) {
		var ret string
		return ret
	}
	return *o.AccountScheduleId
}

// GetAccountScheduleIdOk returns a tuple with the AccountScheduleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdidnsrickettsSourceConfig) GetAccountScheduleIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountScheduleId) {
		return nil, false
	}
	return o.AccountScheduleId, true
}

// HasAccountScheduleId returns a boolean if a field has been set.
func (o *DdidnsrickettsSourceConfig) HasAccountScheduleId() bool {
	if o != nil && !IsNil(o.AccountScheduleId) {
		return true
	}

	return false
}

// SetAccountScheduleId gets a reference to the given string and assigns it to the AccountScheduleId field.
func (o *DdidnsrickettsSourceConfig) SetAccountScheduleId(v string) {
	o.AccountScheduleId = &v
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *DdidnsrickettsSourceConfig) GetAccounts() []DdidnsrickettsAccount {
	if o == nil || IsNil(o.Accounts) {
		var ret []DdidnsrickettsAccount
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdidnsrickettsSourceConfig) GetAccountsOk() ([]DdidnsrickettsAccount, bool) {
	if o == nil || IsNil(o.Accounts) {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *DdidnsrickettsSourceConfig) HasAccounts() bool {
	if o != nil && !IsNil(o.Accounts) {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []DdidnsrickettsAccount and assigns it to the Accounts field.
func (o *DdidnsrickettsSourceConfig) SetAccounts(v []DdidnsrickettsAccount) {
	o.Accounts = v
}

// GetCloudCredentialId returns the CloudCredentialId field value
func (o *DdidnsrickettsSourceConfig) GetCloudCredentialId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudCredentialId
}

// GetCloudCredentialIdOk returns a tuple with the CloudCredentialId field value
// and a boolean to check if the value has been set.
func (o *DdidnsrickettsSourceConfig) GetCloudCredentialIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudCredentialId, true
}

// SetCloudCredentialId sets field value
func (o *DdidnsrickettsSourceConfig) SetCloudCredentialId(v string) {
	o.CloudCredentialId = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DdidnsrickettsSourceConfig) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdidnsrickettsSourceConfig) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DdidnsrickettsSourceConfig) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DdidnsrickettsSourceConfig) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCredentialConfig returns the CredentialConfig field value if set, zero value otherwise.
func (o *DdidnsrickettsSourceConfig) GetCredentialConfig() DdidnsrickettsCredentialConfig {
	if o == nil || IsNil(o.CredentialConfig) {
		var ret DdidnsrickettsCredentialConfig
		return ret
	}
	return *o.CredentialConfig
}

// GetCredentialConfigOk returns a tuple with the CredentialConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdidnsrickettsSourceConfig) GetCredentialConfigOk() (*DdidnsrickettsCredentialConfig, bool) {
	if o == nil || IsNil(o.CredentialConfig) {
		return nil, false
	}
	return o.CredentialConfig, true
}

// HasCredentialConfig returns a boolean if a field has been set.
func (o *DdidnsrickettsSourceConfig) HasCredentialConfig() bool {
	if o != nil && !IsNil(o.CredentialConfig) {
		return true
	}

	return false
}

// SetCredentialConfig gets a reference to the given DdidnsrickettsCredentialConfig and assigns it to the CredentialConfig field.
func (o *DdidnsrickettsSourceConfig) SetCredentialConfig(v DdidnsrickettsCredentialConfig) {
	o.CredentialConfig = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *DdidnsrickettsSourceConfig) GetDeletedAt() time.Time {
	if o == nil || IsNil(o.DeletedAt) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdidnsrickettsSourceConfig) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *DdidnsrickettsSourceConfig) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *DdidnsrickettsSourceConfig) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DdidnsrickettsSourceConfig) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdidnsrickettsSourceConfig) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DdidnsrickettsSourceConfig) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DdidnsrickettsSourceConfig) SetId(v string) {
	o.Id = &v
}

// GetRestrictedToAccounts returns the RestrictedToAccounts field value if set, zero value otherwise.
func (o *DdidnsrickettsSourceConfig) GetRestrictedToAccounts() []string {
	if o == nil || IsNil(o.RestrictedToAccounts) {
		var ret []string
		return ret
	}
	return o.RestrictedToAccounts
}

// GetRestrictedToAccountsOk returns a tuple with the RestrictedToAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdidnsrickettsSourceConfig) GetRestrictedToAccountsOk() ([]string, bool) {
	if o == nil || IsNil(o.RestrictedToAccounts) {
		return nil, false
	}
	return o.RestrictedToAccounts, true
}

// HasRestrictedToAccounts returns a boolean if a field has been set.
func (o *DdidnsrickettsSourceConfig) HasRestrictedToAccounts() bool {
	if o != nil && !IsNil(o.RestrictedToAccounts) {
		return true
	}

	return false
}

// SetRestrictedToAccounts gets a reference to the given []string and assigns it to the RestrictedToAccounts field.
func (o *DdidnsrickettsSourceConfig) SetRestrictedToAccounts(v []string) {
	o.RestrictedToAccounts = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DdidnsrickettsSourceConfig) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdidnsrickettsSourceConfig) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DdidnsrickettsSourceConfig) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DdidnsrickettsSourceConfig) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o DdidnsrickettsSourceConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DdidnsrickettsSourceConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountScheduleId) {
		toSerialize["account_schedule_id"] = o.AccountScheduleId
	}
	if !IsNil(o.Accounts) {
		toSerialize["accounts"] = o.Accounts
	}
	toSerialize["cloud_credential_id"] = o.CloudCredentialId
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.CredentialConfig) {
		toSerialize["credential_config"] = o.CredentialConfig
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.RestrictedToAccounts) {
		toSerialize["restricted_to_accounts"] = o.RestrictedToAccounts
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DdidnsrickettsSourceConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloud_credential_id",
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

	varDdidnsrickettsSourceConfig := _DdidnsrickettsSourceConfig{}

	err = json.Unmarshal(data, &varDdidnsrickettsSourceConfig)

	if err != nil {
		return err
	}

	*o = DdidnsrickettsSourceConfig(varDdidnsrickettsSourceConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "account_schedule_id")
		delete(additionalProperties, "accounts")
		delete(additionalProperties, "cloud_credential_id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "credential_config")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "restricted_to_accounts")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDdidnsrickettsSourceConfig struct {
	value *DdidnsrickettsSourceConfig
	isSet bool
}

func (v NullableDdidnsrickettsSourceConfig) Get() *DdidnsrickettsSourceConfig {
	return v.value
}

func (v *NullableDdidnsrickettsSourceConfig) Set(val *DdidnsrickettsSourceConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDdidnsrickettsSourceConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDdidnsrickettsSourceConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDdidnsrickettsSourceConfig(val *DdidnsrickettsSourceConfig) *NullableDdidnsrickettsSourceConfig {
	return &NullableDdidnsrickettsSourceConfig{value: val, isSet: true}
}

func (v NullableDdidnsrickettsSourceConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDdidnsrickettsSourceConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

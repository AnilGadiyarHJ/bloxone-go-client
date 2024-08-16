/*
IPAM Federation API

The DDI IPAM Federation application enables a SaaS administrator to manage multiple IPAM systems from one central control point CSP.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipamfederation

import (
	"encoding/json"
)

// checks if the FederationHealthConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FederationHealthConfig{}

// FederationHealthConfig struct for FederationHealthConfig
type FederationHealthConfig struct {
	AppName              *string                 `json:"appName,omitempty"`
	Message              *string                 `json:"message,omitempty"`
	Ophid                *string                 `json:"ophid,omitempty"`
	StatusCode           *HealthConfigStatusCode `json:"statusCode,omitempty"`
	Version              *string                 `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FederationHealthConfig FederationHealthConfig

// NewFederationHealthConfig instantiates a new FederationHealthConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFederationHealthConfig() *FederationHealthConfig {
	this := FederationHealthConfig{}
	var statusCode HealthConfigStatusCode = HEALTHCONFIGSTATUSCODE_SUCCESS
	this.StatusCode = &statusCode
	return &this
}

// NewFederationHealthConfigWithDefaults instantiates a new FederationHealthConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFederationHealthConfigWithDefaults() *FederationHealthConfig {
	this := FederationHealthConfig{}
	var statusCode HealthConfigStatusCode = HEALTHCONFIGSTATUSCODE_SUCCESS
	this.StatusCode = &statusCode
	return &this
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *FederationHealthConfig) GetAppName() string {
	if o == nil || IsNil(o.AppName) {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationHealthConfig) GetAppNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppName) {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *FederationHealthConfig) HasAppName() bool {
	if o != nil && !IsNil(o.AppName) {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *FederationHealthConfig) SetAppName(v string) {
	o.AppName = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *FederationHealthConfig) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationHealthConfig) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *FederationHealthConfig) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *FederationHealthConfig) SetMessage(v string) {
	o.Message = &v
}

// GetOphid returns the Ophid field value if set, zero value otherwise.
func (o *FederationHealthConfig) GetOphid() string {
	if o == nil || IsNil(o.Ophid) {
		var ret string
		return ret
	}
	return *o.Ophid
}

// GetOphidOk returns a tuple with the Ophid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationHealthConfig) GetOphidOk() (*string, bool) {
	if o == nil || IsNil(o.Ophid) {
		return nil, false
	}
	return o.Ophid, true
}

// HasOphid returns a boolean if a field has been set.
func (o *FederationHealthConfig) HasOphid() bool {
	if o != nil && !IsNil(o.Ophid) {
		return true
	}

	return false
}

// SetOphid gets a reference to the given string and assigns it to the Ophid field.
func (o *FederationHealthConfig) SetOphid(v string) {
	o.Ophid = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *FederationHealthConfig) GetStatusCode() HealthConfigStatusCode {
	if o == nil || IsNil(o.StatusCode) {
		var ret HealthConfigStatusCode
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationHealthConfig) GetStatusCodeOk() (*HealthConfigStatusCode, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *FederationHealthConfig) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given HealthConfigStatusCode and assigns it to the StatusCode field.
func (o *FederationHealthConfig) SetStatusCode(v HealthConfigStatusCode) {
	o.StatusCode = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *FederationHealthConfig) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationHealthConfig) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *FederationHealthConfig) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *FederationHealthConfig) SetVersion(v string) {
	o.Version = &v
}

func (o FederationHealthConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FederationHealthConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppName) {
		toSerialize["appName"] = o.AppName
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Ophid) {
		toSerialize["ophid"] = o.Ophid
	}
	if !IsNil(o.StatusCode) {
		toSerialize["statusCode"] = o.StatusCode
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FederationHealthConfig) UnmarshalJSON(data []byte) (err error) {
	varFederationHealthConfig := _FederationHealthConfig{}

	err = json.Unmarshal(data, &varFederationHealthConfig)

	if err != nil {
		return err
	}

	*o = FederationHealthConfig(varFederationHealthConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "appName")
		delete(additionalProperties, "message")
		delete(additionalProperties, "ophid")
		delete(additionalProperties, "statusCode")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFederationHealthConfig struct {
	value *FederationHealthConfig
	isSet bool
}

func (v NullableFederationHealthConfig) Get() *FederationHealthConfig {
	return v.value
}

func (v *NullableFederationHealthConfig) Set(val *FederationHealthConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableFederationHealthConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableFederationHealthConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFederationHealthConfig(val *FederationHealthConfig) *NullableFederationHealthConfig {
	return &NullableFederationHealthConfig{value: val, isSet: true}
}

func (v NullableFederationHealthConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFederationHealthConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

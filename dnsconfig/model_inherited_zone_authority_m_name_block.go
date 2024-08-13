/*
DNS Configuration API

The DNS application is a BloxOne DDI service that provides cloud-based DNS configuration with on-prem host serving DNS protocol. It is part of the full-featured BloxOne DDI solution that enables customers the ability to deploy large numbers of protocol servers in the delivery of DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dnsconfig

import (
	"encoding/json"
)

// checks if the InheritedZoneAuthorityMNameBlock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InheritedZoneAuthorityMNameBlock{}

// InheritedZoneAuthorityMNameBlock Inheritance block for fields: _mname_, _protocol_mname_, _default_mname_.
type InheritedZoneAuthorityMNameBlock struct {
	// Defaults to _inherit_.
	Action *string `json:"action,omitempty"`
	// Human-readable display name for the object referred to by _source_.
	DisplayName *string `json:"display_name,omitempty"`
	// The resource identifier.
	Source *string `json:"source,omitempty"`
	// Inherited value.
	Value                *ZoneAuthorityMNameBlock `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InheritedZoneAuthorityMNameBlock InheritedZoneAuthorityMNameBlock

// NewInheritedZoneAuthorityMNameBlock instantiates a new InheritedZoneAuthorityMNameBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInheritedZoneAuthorityMNameBlock() *InheritedZoneAuthorityMNameBlock {
	this := InheritedZoneAuthorityMNameBlock{}
	return &this
}

// NewInheritedZoneAuthorityMNameBlockWithDefaults instantiates a new InheritedZoneAuthorityMNameBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInheritedZoneAuthorityMNameBlockWithDefaults() *InheritedZoneAuthorityMNameBlock {
	this := InheritedZoneAuthorityMNameBlock{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *InheritedZoneAuthorityMNameBlock) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InheritedZoneAuthorityMNameBlock) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *InheritedZoneAuthorityMNameBlock) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *InheritedZoneAuthorityMNameBlock) SetAction(v string) {
	o.Action = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *InheritedZoneAuthorityMNameBlock) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InheritedZoneAuthorityMNameBlock) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *InheritedZoneAuthorityMNameBlock) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *InheritedZoneAuthorityMNameBlock) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *InheritedZoneAuthorityMNameBlock) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InheritedZoneAuthorityMNameBlock) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *InheritedZoneAuthorityMNameBlock) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *InheritedZoneAuthorityMNameBlock) SetSource(v string) {
	o.Source = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InheritedZoneAuthorityMNameBlock) GetValue() ZoneAuthorityMNameBlock {
	if o == nil || IsNil(o.Value) {
		var ret ZoneAuthorityMNameBlock
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InheritedZoneAuthorityMNameBlock) GetValueOk() (*ZoneAuthorityMNameBlock, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InheritedZoneAuthorityMNameBlock) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given ZoneAuthorityMNameBlock and assigns it to the Value field.
func (o *InheritedZoneAuthorityMNameBlock) SetValue(v ZoneAuthorityMNameBlock) {
	o.Value = &v
}

func (o InheritedZoneAuthorityMNameBlock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InheritedZoneAuthorityMNameBlock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InheritedZoneAuthorityMNameBlock) UnmarshalJSON(data []byte) (err error) {
	varInheritedZoneAuthorityMNameBlock := _InheritedZoneAuthorityMNameBlock{}

	err = json.Unmarshal(data, &varInheritedZoneAuthorityMNameBlock)

	if err != nil {
		return err
	}

	*o = InheritedZoneAuthorityMNameBlock(varInheritedZoneAuthorityMNameBlock)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "display_name")
		delete(additionalProperties, "source")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInheritedZoneAuthorityMNameBlock struct {
	value *InheritedZoneAuthorityMNameBlock
	isSet bool
}

func (v NullableInheritedZoneAuthorityMNameBlock) Get() *InheritedZoneAuthorityMNameBlock {
	return v.value
}

func (v *NullableInheritedZoneAuthorityMNameBlock) Set(val *InheritedZoneAuthorityMNameBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableInheritedZoneAuthorityMNameBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableInheritedZoneAuthorityMNameBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInheritedZoneAuthorityMNameBlock(val *InheritedZoneAuthorityMNameBlock) *NullableInheritedZoneAuthorityMNameBlock {
	return &NullableInheritedZoneAuthorityMNameBlock{value: val, isSet: true}
}

func (v NullableInheritedZoneAuthorityMNameBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInheritedZoneAuthorityMNameBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

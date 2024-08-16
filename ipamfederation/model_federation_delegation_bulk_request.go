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

// checks if the FederationDelegationBulkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FederationDelegationBulkRequest{}

// FederationDelegationBulkRequest DelegationBulk request object.
type FederationDelegationBulkRequest struct {
	// list of action/delegation requests to execute.
	Delegations          []FederationDelegationBulkItem `json:"delegations,omitempty"`
	Fields               *ProtobufFieldMask             `json:"fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FederationDelegationBulkRequest FederationDelegationBulkRequest

// NewFederationDelegationBulkRequest instantiates a new FederationDelegationBulkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFederationDelegationBulkRequest() *FederationDelegationBulkRequest {
	this := FederationDelegationBulkRequest{}
	return &this
}

// NewFederationDelegationBulkRequestWithDefaults instantiates a new FederationDelegationBulkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFederationDelegationBulkRequestWithDefaults() *FederationDelegationBulkRequest {
	this := FederationDelegationBulkRequest{}
	return &this
}

// GetDelegations returns the Delegations field value if set, zero value otherwise.
func (o *FederationDelegationBulkRequest) GetDelegations() []FederationDelegationBulkItem {
	if o == nil || IsNil(o.Delegations) {
		var ret []FederationDelegationBulkItem
		return ret
	}
	return o.Delegations
}

// GetDelegationsOk returns a tuple with the Delegations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationDelegationBulkRequest) GetDelegationsOk() ([]FederationDelegationBulkItem, bool) {
	if o == nil || IsNil(o.Delegations) {
		return nil, false
	}
	return o.Delegations, true
}

// HasDelegations returns a boolean if a field has been set.
func (o *FederationDelegationBulkRequest) HasDelegations() bool {
	if o != nil && !IsNil(o.Delegations) {
		return true
	}

	return false
}

// SetDelegations gets a reference to the given []FederationDelegationBulkItem and assigns it to the Delegations field.
func (o *FederationDelegationBulkRequest) SetDelegations(v []FederationDelegationBulkItem) {
	o.Delegations = v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *FederationDelegationBulkRequest) GetFields() ProtobufFieldMask {
	if o == nil || IsNil(o.Fields) {
		var ret ProtobufFieldMask
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationDelegationBulkRequest) GetFieldsOk() (*ProtobufFieldMask, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *FederationDelegationBulkRequest) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given ProtobufFieldMask and assigns it to the Fields field.
func (o *FederationDelegationBulkRequest) SetFields(v ProtobufFieldMask) {
	o.Fields = &v
}

func (o FederationDelegationBulkRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FederationDelegationBulkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Delegations) {
		toSerialize["delegations"] = o.Delegations
	}
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FederationDelegationBulkRequest) UnmarshalJSON(data []byte) (err error) {
	varFederationDelegationBulkRequest := _FederationDelegationBulkRequest{}

	err = json.Unmarshal(data, &varFederationDelegationBulkRequest)

	if err != nil {
		return err
	}

	*o = FederationDelegationBulkRequest(varFederationDelegationBulkRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "delegations")
		delete(additionalProperties, "fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFederationDelegationBulkRequest struct {
	value *FederationDelegationBulkRequest
	isSet bool
}

func (v NullableFederationDelegationBulkRequest) Get() *FederationDelegationBulkRequest {
	return v.value
}

func (v *NullableFederationDelegationBulkRequest) Set(val *FederationDelegationBulkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFederationDelegationBulkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFederationDelegationBulkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFederationDelegationBulkRequest(val *FederationDelegationBulkRequest) *NullableFederationDelegationBulkRequest {
	return &NullableFederationDelegationBulkRequest{value: val, isSet: true}
}

func (v NullableFederationDelegationBulkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFederationDelegationBulkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
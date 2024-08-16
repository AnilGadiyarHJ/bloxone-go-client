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

// checks if the FederationCreateDelegationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FederationCreateDelegationResponse{}

// FederationCreateDelegationResponse The response format to create the __Delegation__ object.
type FederationCreateDelegationResponse struct {
	// The created Delegation object.
	Result               *FederationDelegation `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FederationCreateDelegationResponse FederationCreateDelegationResponse

// NewFederationCreateDelegationResponse instantiates a new FederationCreateDelegationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFederationCreateDelegationResponse() *FederationCreateDelegationResponse {
	this := FederationCreateDelegationResponse{}
	return &this
}

// NewFederationCreateDelegationResponseWithDefaults instantiates a new FederationCreateDelegationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFederationCreateDelegationResponseWithDefaults() *FederationCreateDelegationResponse {
	this := FederationCreateDelegationResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *FederationCreateDelegationResponse) GetResult() FederationDelegation {
	if o == nil || IsNil(o.Result) {
		var ret FederationDelegation
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationCreateDelegationResponse) GetResultOk() (*FederationDelegation, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *FederationCreateDelegationResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given FederationDelegation and assigns it to the Result field.
func (o *FederationCreateDelegationResponse) SetResult(v FederationDelegation) {
	o.Result = &v
}

func (o FederationCreateDelegationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FederationCreateDelegationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FederationCreateDelegationResponse) UnmarshalJSON(data []byte) (err error) {
	varFederationCreateDelegationResponse := _FederationCreateDelegationResponse{}

	err = json.Unmarshal(data, &varFederationCreateDelegationResponse)

	if err != nil {
		return err
	}

	*o = FederationCreateDelegationResponse(varFederationCreateDelegationResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFederationCreateDelegationResponse struct {
	value *FederationCreateDelegationResponse
	isSet bool
}

func (v NullableFederationCreateDelegationResponse) Get() *FederationCreateDelegationResponse {
	return v.value
}

func (v *NullableFederationCreateDelegationResponse) Set(val *FederationCreateDelegationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFederationCreateDelegationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFederationCreateDelegationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFederationCreateDelegationResponse(val *FederationCreateDelegationResponse) *NullableFederationCreateDelegationResponse {
	return &NullableFederationCreateDelegationResponse{value: val, isSet: true}
}

func (v NullableFederationCreateDelegationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFederationCreateDelegationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

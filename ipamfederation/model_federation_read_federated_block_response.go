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

// checks if the FederationReadFederatedBlockResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FederationReadFederatedBlockResponse{}

// FederationReadFederatedBlockResponse The response format to retrieve the __FederatedBlock__ object.
type FederationReadFederatedBlockResponse struct {
	// The FederatedBlock object.
	Result               *FederationFederatedBlock `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FederationReadFederatedBlockResponse FederationReadFederatedBlockResponse

// NewFederationReadFederatedBlockResponse instantiates a new FederationReadFederatedBlockResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFederationReadFederatedBlockResponse() *FederationReadFederatedBlockResponse {
	this := FederationReadFederatedBlockResponse{}
	return &this
}

// NewFederationReadFederatedBlockResponseWithDefaults instantiates a new FederationReadFederatedBlockResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFederationReadFederatedBlockResponseWithDefaults() *FederationReadFederatedBlockResponse {
	this := FederationReadFederatedBlockResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *FederationReadFederatedBlockResponse) GetResult() FederationFederatedBlock {
	if o == nil || IsNil(o.Result) {
		var ret FederationFederatedBlock
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationReadFederatedBlockResponse) GetResultOk() (*FederationFederatedBlock, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *FederationReadFederatedBlockResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given FederationFederatedBlock and assigns it to the Result field.
func (o *FederationReadFederatedBlockResponse) SetResult(v FederationFederatedBlock) {
	o.Result = &v
}

func (o FederationReadFederatedBlockResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FederationReadFederatedBlockResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FederationReadFederatedBlockResponse) UnmarshalJSON(data []byte) (err error) {
	varFederationReadFederatedBlockResponse := _FederationReadFederatedBlockResponse{}

	err = json.Unmarshal(data, &varFederationReadFederatedBlockResponse)

	if err != nil {
		return err
	}

	*o = FederationReadFederatedBlockResponse(varFederationReadFederatedBlockResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFederationReadFederatedBlockResponse struct {
	value *FederationReadFederatedBlockResponse
	isSet bool
}

func (v NullableFederationReadFederatedBlockResponse) Get() *FederationReadFederatedBlockResponse {
	return v.value
}

func (v *NullableFederationReadFederatedBlockResponse) Set(val *FederationReadFederatedBlockResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFederationReadFederatedBlockResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFederationReadFederatedBlockResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFederationReadFederatedBlockResponse(val *FederationReadFederatedBlockResponse) *NullableFederationReadFederatedBlockResponse {
	return &NullableFederationReadFederatedBlockResponse{value: val, isSet: true}
}

func (v NullableFederationReadFederatedBlockResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFederationReadFederatedBlockResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

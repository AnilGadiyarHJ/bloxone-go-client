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

// checks if the CreateNextAvailableReservedBlockResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNextAvailableReservedBlockResponse{}

// CreateNextAvailableReservedBlockResponse The response format to allocate next available __ReservedBlock__ objects.
type CreateNextAvailableReservedBlockResponse struct {
	Results              []ReservedBlock `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateNextAvailableReservedBlockResponse CreateNextAvailableReservedBlockResponse

// NewCreateNextAvailableReservedBlockResponse instantiates a new CreateNextAvailableReservedBlockResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNextAvailableReservedBlockResponse() *CreateNextAvailableReservedBlockResponse {
	this := CreateNextAvailableReservedBlockResponse{}
	return &this
}

// NewCreateNextAvailableReservedBlockResponseWithDefaults instantiates a new CreateNextAvailableReservedBlockResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNextAvailableReservedBlockResponseWithDefaults() *CreateNextAvailableReservedBlockResponse {
	this := CreateNextAvailableReservedBlockResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *CreateNextAvailableReservedBlockResponse) GetResults() []ReservedBlock {
	if o == nil || IsNil(o.Results) {
		var ret []ReservedBlock
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNextAvailableReservedBlockResponse) GetResultsOk() ([]ReservedBlock, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *CreateNextAvailableReservedBlockResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ReservedBlock and assigns it to the Results field.
func (o *CreateNextAvailableReservedBlockResponse) SetResults(v []ReservedBlock) {
	o.Results = v
}

func (o CreateNextAvailableReservedBlockResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNextAvailableReservedBlockResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateNextAvailableReservedBlockResponse) UnmarshalJSON(data []byte) (err error) {
	varCreateNextAvailableReservedBlockResponse := _CreateNextAvailableReservedBlockResponse{}

	err = json.Unmarshal(data, &varCreateNextAvailableReservedBlockResponse)

	if err != nil {
		return err
	}

	*o = CreateNextAvailableReservedBlockResponse(varCreateNextAvailableReservedBlockResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateNextAvailableReservedBlockResponse struct {
	value *CreateNextAvailableReservedBlockResponse
	isSet bool
}

func (v NullableCreateNextAvailableReservedBlockResponse) Get() *CreateNextAvailableReservedBlockResponse {
	return v.value
}

func (v *NullableCreateNextAvailableReservedBlockResponse) Set(val *CreateNextAvailableReservedBlockResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNextAvailableReservedBlockResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNextAvailableReservedBlockResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNextAvailableReservedBlockResponse(val *CreateNextAvailableReservedBlockResponse) *NullableCreateNextAvailableReservedBlockResponse {
	return &NullableCreateNextAvailableReservedBlockResponse{value: val, isSet: true}
}

func (v NullableCreateNextAvailableReservedBlockResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNextAvailableReservedBlockResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
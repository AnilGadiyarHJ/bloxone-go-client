/*
DFP API

BloxOne Cloud is a SaaS offering designed to provide protection to devices on and off-premises, including roaming, remote, and branch offices. It provides visibility into infected and compromised devices, prevents DNS-based data exfiltration, and automatically stops device communications with command-and-control servers (C&Cs) and botnets, in addition to providing recursive DNS services in the cloud. You can access the services by deploying the BloxOne Endpoint agent or the DNS forwarding proxy.  For remote office deployments or in cases where installing an endpoint agent is not desirable or possible, you can use the DNS forwarding proxy. It is a software that runs on bare-metal, VM infrastructures, or Infoblox NIOS appliances; and it embeds the client IPs in DNS queries before forwarding them to BloxOne Cloud. The communications are encrypted and client visibility is maintained. The proxy also provides DNS resolution to local DNS zones when you configure local resolvers. Once you set up a DNS forwarding proxy, it becomes the main DNS server for your remote site. It will also cache responses to speed resolution of future queries.  By implementing the DNS forwarding proxy, you can rest assured that BloxOne Cloud effectively enforces DNS client-based security policies at your remote sites. On-premises devices that send DNS queries reveal their actual client IP addresses (instead of their NAT IP address), which allows BloxOne Cloud to apply the security policies applicable to the respective endpoints and identify infected clients.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dfp

import (
	"encoding/json"
)

// checks if the DfpCreateOrUpdateDfp400Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfpCreateOrUpdateDfp400Response{}

// DfpCreateOrUpdateDfp400Response struct for DfpCreateOrUpdateDfp400Response
type DfpCreateOrUpdateDfp400Response struct {
	Error *DfpCreateOrUpdateDfp400ResponseError `json:"error,omitempty"`
}

// NewDfpCreateOrUpdateDfp400Response instantiates a new DfpCreateOrUpdateDfp400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfpCreateOrUpdateDfp400Response() *DfpCreateOrUpdateDfp400Response {
	this := DfpCreateOrUpdateDfp400Response{}
	return &this
}

// NewDfpCreateOrUpdateDfp400ResponseWithDefaults instantiates a new DfpCreateOrUpdateDfp400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfpCreateOrUpdateDfp400ResponseWithDefaults() *DfpCreateOrUpdateDfp400Response {
	this := DfpCreateOrUpdateDfp400Response{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *DfpCreateOrUpdateDfp400Response) GetError() DfpCreateOrUpdateDfp400ResponseError {
	if o == nil || IsNil(o.Error) {
		var ret DfpCreateOrUpdateDfp400ResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfpCreateOrUpdateDfp400Response) GetErrorOk() (*DfpCreateOrUpdateDfp400ResponseError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *DfpCreateOrUpdateDfp400Response) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given DfpCreateOrUpdateDfp400ResponseError and assigns it to the Error field.
func (o *DfpCreateOrUpdateDfp400Response) SetError(v DfpCreateOrUpdateDfp400ResponseError) {
	o.Error = &v
}

func (o DfpCreateOrUpdateDfp400Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfpCreateOrUpdateDfp400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableDfpCreateOrUpdateDfp400Response struct {
	value *DfpCreateOrUpdateDfp400Response
	isSet bool
}

func (v NullableDfpCreateOrUpdateDfp400Response) Get() *DfpCreateOrUpdateDfp400Response {
	return v.value
}

func (v *NullableDfpCreateOrUpdateDfp400Response) Set(val *DfpCreateOrUpdateDfp400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDfpCreateOrUpdateDfp400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDfpCreateOrUpdateDfp400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfpCreateOrUpdateDfp400Response(val *DfpCreateOrUpdateDfp400Response) *NullableDfpCreateOrUpdateDfp400Response {
	return &NullableDfpCreateOrUpdateDfp400Response{value: val, isSet: true}
}

func (v NullableDfpCreateOrUpdateDfp400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfpCreateOrUpdateDfp400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
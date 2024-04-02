/*
BloxOne Anycast API

Anycast capability enables HA (High Availability) configuration of BloxOne applications that run on equipment located on customer's premises (on-prem hosts). Anycast supports DNS, as well as DNS-forwarding services.  Anycast-enabled application setups use multiple on-premises installations for one particular application type. Multiple application instances are configured to use the same endpoint address. Anycast capability is collocated with such application instance, monitoring the local application instance and advertising to the upstream router (a customer equipment) a per-instance, local route to the common application endpoint address, as long as the local application instance is available. Depending on the type of the upstream router, the customer may configure local route advertisement via either BGP (Boarder Gateway Protocol) or OSPF (Open Shortest Path First) routing protocols. Both protocols may be enabled as well. Multiple routes to the common application service address provide redundancy without the need to reconfigure application clients.  Should an application instance become unavailable, the local route advertisements stop, resulting in withdrawal of the route (in the upstream router) to the application instance that has gone out of service and ensuring that subsequent application requests thus get routed to the remaining available application instances.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package anycast

import (
	"encoding/json"
)

// checks if the ProtoBgpConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProtoBgpConfig{}

// ProtoBgpConfig struct for ProtoBgpConfig
type ProtoBgpConfig struct {
	Asn *int64 `json:"asn,omitempty"`
	// Examples:     ASDOT        ASPLAIN     INTEGER     VALID/INVALID     0.1          1           1           Valid     1            1           1           Valid     65535        65535       65535       Valid     0.65535      65535       65535       Valid     1.0          65536       65536       Valid     1.1          65537       65537       Valid     1.65535      131071      131071      Valid     65535.0      4294901760  4294901760  Valid     65535.1      4294901761  4294901761  Valid     65535.65535  4294967295  4294967295  Valid      0.65536                              Invalid     65535.655536                         Invalid     65536.0                              Invalid     65536.65535                          Invalid                  4294967296              Invalid
	AsnText       *string            `json:"asn_text,omitempty"`
	Fields        *ProtobufFieldMask `json:"fields,omitempty"`
	HolddownSecs  *int64             `json:"holddown_secs,omitempty"`
	KeepAliveSecs *int64             `json:"keep_alive_secs,omitempty"`
	LinkDetect    *bool              `json:"link_detect,omitempty"`
	Neighbors     []ProtoBgpNeighbor `json:"neighbors,omitempty"`
	// Any predefined BGP configuration, with embedded new lines; the preamble will be prepended to the generated BGP configuration.
	Preamble *string `json:"preamble,omitempty"`
}

// NewProtoBgpConfig instantiates a new ProtoBgpConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtoBgpConfig() *ProtoBgpConfig {
	this := ProtoBgpConfig{}
	return &this
}

// NewProtoBgpConfigWithDefaults instantiates a new ProtoBgpConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtoBgpConfigWithDefaults() *ProtoBgpConfig {
	this := ProtoBgpConfig{}
	return &this
}

// GetAsn returns the Asn field value if set, zero value otherwise.
func (o *ProtoBgpConfig) GetAsn() int64 {
	if o == nil || IsNil(o.Asn) {
		var ret int64
		return ret
	}
	return *o.Asn
}

// GetAsnOk returns a tuple with the Asn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtoBgpConfig) GetAsnOk() (*int64, bool) {
	if o == nil || IsNil(o.Asn) {
		return nil, false
	}
	return o.Asn, true
}

// HasAsn returns a boolean if a field has been set.
func (o *ProtoBgpConfig) HasAsn() bool {
	if o != nil && !IsNil(o.Asn) {
		return true
	}

	return false
}

// SetAsn gets a reference to the given int64 and assigns it to the Asn field.
func (o *ProtoBgpConfig) SetAsn(v int64) {
	o.Asn = &v
}

// GetAsnText returns the AsnText field value if set, zero value otherwise.
func (o *ProtoBgpConfig) GetAsnText() string {
	if o == nil || IsNil(o.AsnText) {
		var ret string
		return ret
	}
	return *o.AsnText
}

// GetAsnTextOk returns a tuple with the AsnText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtoBgpConfig) GetAsnTextOk() (*string, bool) {
	if o == nil || IsNil(o.AsnText) {
		return nil, false
	}
	return o.AsnText, true
}

// HasAsnText returns a boolean if a field has been set.
func (o *ProtoBgpConfig) HasAsnText() bool {
	if o != nil && !IsNil(o.AsnText) {
		return true
	}

	return false
}

// SetAsnText gets a reference to the given string and assigns it to the AsnText field.
func (o *ProtoBgpConfig) SetAsnText(v string) {
	o.AsnText = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *ProtoBgpConfig) GetFields() ProtobufFieldMask {
	if o == nil || IsNil(o.Fields) {
		var ret ProtobufFieldMask
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtoBgpConfig) GetFieldsOk() (*ProtobufFieldMask, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *ProtoBgpConfig) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given ProtobufFieldMask and assigns it to the Fields field.
func (o *ProtoBgpConfig) SetFields(v ProtobufFieldMask) {
	o.Fields = &v
}

// GetHolddownSecs returns the HolddownSecs field value if set, zero value otherwise.
func (o *ProtoBgpConfig) GetHolddownSecs() int64 {
	if o == nil || IsNil(o.HolddownSecs) {
		var ret int64
		return ret
	}
	return *o.HolddownSecs
}

// GetHolddownSecsOk returns a tuple with the HolddownSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtoBgpConfig) GetHolddownSecsOk() (*int64, bool) {
	if o == nil || IsNil(o.HolddownSecs) {
		return nil, false
	}
	return o.HolddownSecs, true
}

// HasHolddownSecs returns a boolean if a field has been set.
func (o *ProtoBgpConfig) HasHolddownSecs() bool {
	if o != nil && !IsNil(o.HolddownSecs) {
		return true
	}

	return false
}

// SetHolddownSecs gets a reference to the given int64 and assigns it to the HolddownSecs field.
func (o *ProtoBgpConfig) SetHolddownSecs(v int64) {
	o.HolddownSecs = &v
}

// GetKeepAliveSecs returns the KeepAliveSecs field value if set, zero value otherwise.
func (o *ProtoBgpConfig) GetKeepAliveSecs() int64 {
	if o == nil || IsNil(o.KeepAliveSecs) {
		var ret int64
		return ret
	}
	return *o.KeepAliveSecs
}

// GetKeepAliveSecsOk returns a tuple with the KeepAliveSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtoBgpConfig) GetKeepAliveSecsOk() (*int64, bool) {
	if o == nil || IsNil(o.KeepAliveSecs) {
		return nil, false
	}
	return o.KeepAliveSecs, true
}

// HasKeepAliveSecs returns a boolean if a field has been set.
func (o *ProtoBgpConfig) HasKeepAliveSecs() bool {
	if o != nil && !IsNil(o.KeepAliveSecs) {
		return true
	}

	return false
}

// SetKeepAliveSecs gets a reference to the given int64 and assigns it to the KeepAliveSecs field.
func (o *ProtoBgpConfig) SetKeepAliveSecs(v int64) {
	o.KeepAliveSecs = &v
}

// GetLinkDetect returns the LinkDetect field value if set, zero value otherwise.
func (o *ProtoBgpConfig) GetLinkDetect() bool {
	if o == nil || IsNil(o.LinkDetect) {
		var ret bool
		return ret
	}
	return *o.LinkDetect
}

// GetLinkDetectOk returns a tuple with the LinkDetect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtoBgpConfig) GetLinkDetectOk() (*bool, bool) {
	if o == nil || IsNil(o.LinkDetect) {
		return nil, false
	}
	return o.LinkDetect, true
}

// HasLinkDetect returns a boolean if a field has been set.
func (o *ProtoBgpConfig) HasLinkDetect() bool {
	if o != nil && !IsNil(o.LinkDetect) {
		return true
	}

	return false
}

// SetLinkDetect gets a reference to the given bool and assigns it to the LinkDetect field.
func (o *ProtoBgpConfig) SetLinkDetect(v bool) {
	o.LinkDetect = &v
}

// GetNeighbors returns the Neighbors field value if set, zero value otherwise.
func (o *ProtoBgpConfig) GetNeighbors() []ProtoBgpNeighbor {
	if o == nil || IsNil(o.Neighbors) {
		var ret []ProtoBgpNeighbor
		return ret
	}
	return o.Neighbors
}

// GetNeighborsOk returns a tuple with the Neighbors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtoBgpConfig) GetNeighborsOk() ([]ProtoBgpNeighbor, bool) {
	if o == nil || IsNil(o.Neighbors) {
		return nil, false
	}
	return o.Neighbors, true
}

// HasNeighbors returns a boolean if a field has been set.
func (o *ProtoBgpConfig) HasNeighbors() bool {
	if o != nil && !IsNil(o.Neighbors) {
		return true
	}

	return false
}

// SetNeighbors gets a reference to the given []ProtoBgpNeighbor and assigns it to the Neighbors field.
func (o *ProtoBgpConfig) SetNeighbors(v []ProtoBgpNeighbor) {
	o.Neighbors = v
}

// GetPreamble returns the Preamble field value if set, zero value otherwise.
func (o *ProtoBgpConfig) GetPreamble() string {
	if o == nil || IsNil(o.Preamble) {
		var ret string
		return ret
	}
	return *o.Preamble
}

// GetPreambleOk returns a tuple with the Preamble field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtoBgpConfig) GetPreambleOk() (*string, bool) {
	if o == nil || IsNil(o.Preamble) {
		return nil, false
	}
	return o.Preamble, true
}

// HasPreamble returns a boolean if a field has been set.
func (o *ProtoBgpConfig) HasPreamble() bool {
	if o != nil && !IsNil(o.Preamble) {
		return true
	}

	return false
}

// SetPreamble gets a reference to the given string and assigns it to the Preamble field.
func (o *ProtoBgpConfig) SetPreamble(v string) {
	o.Preamble = &v
}

func (o ProtoBgpConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProtoBgpConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asn) {
		toSerialize["asn"] = o.Asn
	}
	if !IsNil(o.AsnText) {
		toSerialize["asn_text"] = o.AsnText
	}
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !IsNil(o.HolddownSecs) {
		toSerialize["holddown_secs"] = o.HolddownSecs
	}
	if !IsNil(o.KeepAliveSecs) {
		toSerialize["keep_alive_secs"] = o.KeepAliveSecs
	}
	if !IsNil(o.LinkDetect) {
		toSerialize["link_detect"] = o.LinkDetect
	}
	if !IsNil(o.Neighbors) {
		toSerialize["neighbors"] = o.Neighbors
	}
	if !IsNil(o.Preamble) {
		toSerialize["preamble"] = o.Preamble
	}
	return toSerialize, nil
}

type NullableProtoBgpConfig struct {
	value *ProtoBgpConfig
	isSet bool
}

func (v NullableProtoBgpConfig) Get() *ProtoBgpConfig {
	return v.value
}

func (v *NullableProtoBgpConfig) Set(val *ProtoBgpConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableProtoBgpConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableProtoBgpConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtoBgpConfig(val *ProtoBgpConfig) *NullableProtoBgpConfig {
	return &NullableProtoBgpConfig{value: val, isSet: true}
}

func (v NullableProtoBgpConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtoBgpConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

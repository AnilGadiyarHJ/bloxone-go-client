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

// checks if the ProtoOspfConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProtoOspfConfig{}

// ProtoOspfConfig struct for ProtoOspfConfig
type ProtoOspfConfig struct {
	// OSPF area identifier; usually in the format of an IPv4 address (although not an address itself)
	Area                *string `json:"area,omitempty"`
	AreaType            *string `json:"area_type,omitempty"`
	AuthenticationKey   *string `json:"authentication_key,omitempty"`
	AuthenticationKeyId *int64  `json:"authentication_key_id,omitempty"`
	AuthenticationType  *string `json:"authentication_type,omitempty"`
	Cost                *int64  `json:"cost,omitempty"`
	DeadInterval        *int64  `json:"dead_interval,omitempty"`
	HelloInterval       *int64  `json:"hello_interval,omitempty"`
	// Name of the interface that is configured with external IP address of the host
	Interface *string `json:"interface,omitempty"`
	// Any predefined OSPF configuration, with embedded new lines; the preamble will be prepended to the generated BGP configuration.
	Preamble           *string `json:"preamble,omitempty"`
	RetransmitInterval *int64  `json:"retransmit_interval,omitempty"`
	TransmitDelay      *int64  `json:"transmit_delay,omitempty"`
}

// NewProtoOspfConfig instantiates a new ProtoOspfConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtoOspfConfig() *ProtoOspfConfig {
	this := ProtoOspfConfig{}
	return &this
}

// NewProtoOspfConfigWithDefaults instantiates a new ProtoOspfConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtoOspfConfigWithDefaults() *ProtoOspfConfig {
	this := ProtoOspfConfig{}
	return &this
}

// GetArea returns the Area field value if set, zero value otherwise.
func (o *ProtoOspfConfig) GetArea() string {
	if o == nil || IsNil(o.Area) {
		var ret string
		return ret
	}
	return *o.Area
}

// GetAreaOk returns a tuple with the Area field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtoOspfConfig) GetAreaOk() (*string, bool) {
	if o == nil || IsNil(o.Area) {
		return nil, false
	}
	return o.Area, true
}

// HasArea returns a boolean if a field has been set.
func (o *ProtoOspfConfig) HasArea() bool {
	if o != nil && !IsNil(o.Area) {
		return true
	}

	return false
}

// SetArea gets a reference to the given string and assigns it to the Area field.
func (o *ProtoOspfConfig) SetArea(v string) {
	o.Area = &v
}

// GetAreaType returns the AreaType field value if set, zero value otherwise.
func (o *ProtoOspfConfig) GetAreaType() string {
	if o == nil || IsNil(o.AreaType) {
		var ret string
		return ret
	}
	return *o.AreaType
}

// GetAreaTypeOk returns a tuple with the AreaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtoOspfConfig) GetAreaTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AreaType) {
		return nil, false
	}
	return o.AreaType, true
}

// HasAreaType returns a boolean if a field has been set.
func (o *ProtoOspfConfig) HasAreaType() bool {
	if o != nil && !IsNil(o.AreaType) {
		return true
	}

	return false
}

// SetAreaType gets a reference to the given string and assigns it to the AreaType field.
func (o *ProtoOspfConfig) SetAreaType(v string) {
	o.AreaType = &v
}

// GetAuthenticationKey returns the AuthenticationKey field value if set, zero value otherwise.
func (o *ProtoOspfConfig) GetAuthenticationKey() string {
	if o == nil || IsNil(o.AuthenticationKey) {
		var ret string
		return ret
	}
	return *o.AuthenticationKey
}

// GetAuthenticationKeyOk returns a tuple with the AuthenticationKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtoOspfConfig) GetAuthenticationKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationKey) {
		return nil, false
	}
	return o.AuthenticationKey, true
}

// HasAuthenticationKey returns a boolean if a field has been set.
func (o *ProtoOspfConfig) HasAuthenticationKey() bool {
	if o != nil && !IsNil(o.AuthenticationKey) {
		return true
	}

	return false
}

// SetAuthenticationKey gets a reference to the given string and assigns it to the AuthenticationKey field.
func (o *ProtoOspfConfig) SetAuthenticationKey(v string) {
	o.AuthenticationKey = &v
}

// GetAuthenticationKeyId returns the AuthenticationKeyId field value if set, zero value otherwise.
func (o *ProtoOspfConfig) GetAuthenticationKeyId() int64 {
	if o == nil || IsNil(o.AuthenticationKeyId) {
		var ret int64
		return ret
	}
	return *o.AuthenticationKeyId
}

// GetAuthenticationKeyIdOk returns a tuple with the AuthenticationKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtoOspfConfig) GetAuthenticationKeyIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AuthenticationKeyId) {
		return nil, false
	}
	return o.AuthenticationKeyId, true
}

// HasAuthenticationKeyId returns a boolean if a field has been set.
func (o *ProtoOspfConfig) HasAuthenticationKeyId() bool {
	if o != nil && !IsNil(o.AuthenticationKeyId) {
		return true
	}

	return false
}

// SetAuthenticationKeyId gets a reference to the given int64 and assigns it to the AuthenticationKeyId field.
func (o *ProtoOspfConfig) SetAuthenticationKeyId(v int64) {
	o.AuthenticationKeyId = &v
}

// GetAuthenticationType returns the AuthenticationType field value if set, zero value otherwise.
func (o *ProtoOspfConfig) GetAuthenticationType() string {
	if o == nil || IsNil(o.AuthenticationType) {
		var ret string
		return ret
	}
	return *o.AuthenticationType
}

// GetAuthenticationTypeOk returns a tuple with the AuthenticationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtoOspfConfig) GetAuthenticationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationType) {
		return nil, false
	}
	return o.AuthenticationType, true
}

// HasAuthenticationType returns a boolean if a field has been set.
func (o *ProtoOspfConfig) HasAuthenticationType() bool {
	if o != nil && !IsNil(o.AuthenticationType) {
		return true
	}

	return false
}

// SetAuthenticationType gets a reference to the given string and assigns it to the AuthenticationType field.
func (o *ProtoOspfConfig) SetAuthenticationType(v string) {
	o.AuthenticationType = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *ProtoOspfConfig) GetCost() int64 {
	if o == nil || IsNil(o.Cost) {
		var ret int64
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtoOspfConfig) GetCostOk() (*int64, bool) {
	if o == nil || IsNil(o.Cost) {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *ProtoOspfConfig) HasCost() bool {
	if o != nil && !IsNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given int64 and assigns it to the Cost field.
func (o *ProtoOspfConfig) SetCost(v int64) {
	o.Cost = &v
}

// GetDeadInterval returns the DeadInterval field value if set, zero value otherwise.
func (o *ProtoOspfConfig) GetDeadInterval() int64 {
	if o == nil || IsNil(o.DeadInterval) {
		var ret int64
		return ret
	}
	return *o.DeadInterval
}

// GetDeadIntervalOk returns a tuple with the DeadInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtoOspfConfig) GetDeadIntervalOk() (*int64, bool) {
	if o == nil || IsNil(o.DeadInterval) {
		return nil, false
	}
	return o.DeadInterval, true
}

// HasDeadInterval returns a boolean if a field has been set.
func (o *ProtoOspfConfig) HasDeadInterval() bool {
	if o != nil && !IsNil(o.DeadInterval) {
		return true
	}

	return false
}

// SetDeadInterval gets a reference to the given int64 and assigns it to the DeadInterval field.
func (o *ProtoOspfConfig) SetDeadInterval(v int64) {
	o.DeadInterval = &v
}

// GetHelloInterval returns the HelloInterval field value if set, zero value otherwise.
func (o *ProtoOspfConfig) GetHelloInterval() int64 {
	if o == nil || IsNil(o.HelloInterval) {
		var ret int64
		return ret
	}
	return *o.HelloInterval
}

// GetHelloIntervalOk returns a tuple with the HelloInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtoOspfConfig) GetHelloIntervalOk() (*int64, bool) {
	if o == nil || IsNil(o.HelloInterval) {
		return nil, false
	}
	return o.HelloInterval, true
}

// HasHelloInterval returns a boolean if a field has been set.
func (o *ProtoOspfConfig) HasHelloInterval() bool {
	if o != nil && !IsNil(o.HelloInterval) {
		return true
	}

	return false
}

// SetHelloInterval gets a reference to the given int64 and assigns it to the HelloInterval field.
func (o *ProtoOspfConfig) SetHelloInterval(v int64) {
	o.HelloInterval = &v
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *ProtoOspfConfig) GetInterface() string {
	if o == nil || IsNil(o.Interface) {
		var ret string
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtoOspfConfig) GetInterfaceOk() (*string, bool) {
	if o == nil || IsNil(o.Interface) {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *ProtoOspfConfig) HasInterface() bool {
	if o != nil && !IsNil(o.Interface) {
		return true
	}

	return false
}

// SetInterface gets a reference to the given string and assigns it to the Interface field.
func (o *ProtoOspfConfig) SetInterface(v string) {
	o.Interface = &v
}

// GetPreamble returns the Preamble field value if set, zero value otherwise.
func (o *ProtoOspfConfig) GetPreamble() string {
	if o == nil || IsNil(o.Preamble) {
		var ret string
		return ret
	}
	return *o.Preamble
}

// GetPreambleOk returns a tuple with the Preamble field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtoOspfConfig) GetPreambleOk() (*string, bool) {
	if o == nil || IsNil(o.Preamble) {
		return nil, false
	}
	return o.Preamble, true
}

// HasPreamble returns a boolean if a field has been set.
func (o *ProtoOspfConfig) HasPreamble() bool {
	if o != nil && !IsNil(o.Preamble) {
		return true
	}

	return false
}

// SetPreamble gets a reference to the given string and assigns it to the Preamble field.
func (o *ProtoOspfConfig) SetPreamble(v string) {
	o.Preamble = &v
}

// GetRetransmitInterval returns the RetransmitInterval field value if set, zero value otherwise.
func (o *ProtoOspfConfig) GetRetransmitInterval() int64 {
	if o == nil || IsNil(o.RetransmitInterval) {
		var ret int64
		return ret
	}
	return *o.RetransmitInterval
}

// GetRetransmitIntervalOk returns a tuple with the RetransmitInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtoOspfConfig) GetRetransmitIntervalOk() (*int64, bool) {
	if o == nil || IsNil(o.RetransmitInterval) {
		return nil, false
	}
	return o.RetransmitInterval, true
}

// HasRetransmitInterval returns a boolean if a field has been set.
func (o *ProtoOspfConfig) HasRetransmitInterval() bool {
	if o != nil && !IsNil(o.RetransmitInterval) {
		return true
	}

	return false
}

// SetRetransmitInterval gets a reference to the given int64 and assigns it to the RetransmitInterval field.
func (o *ProtoOspfConfig) SetRetransmitInterval(v int64) {
	o.RetransmitInterval = &v
}

// GetTransmitDelay returns the TransmitDelay field value if set, zero value otherwise.
func (o *ProtoOspfConfig) GetTransmitDelay() int64 {
	if o == nil || IsNil(o.TransmitDelay) {
		var ret int64
		return ret
	}
	return *o.TransmitDelay
}

// GetTransmitDelayOk returns a tuple with the TransmitDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtoOspfConfig) GetTransmitDelayOk() (*int64, bool) {
	if o == nil || IsNil(o.TransmitDelay) {
		return nil, false
	}
	return o.TransmitDelay, true
}

// HasTransmitDelay returns a boolean if a field has been set.
func (o *ProtoOspfConfig) HasTransmitDelay() bool {
	if o != nil && !IsNil(o.TransmitDelay) {
		return true
	}

	return false
}

// SetTransmitDelay gets a reference to the given int64 and assigns it to the TransmitDelay field.
func (o *ProtoOspfConfig) SetTransmitDelay(v int64) {
	o.TransmitDelay = &v
}

func (o ProtoOspfConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProtoOspfConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Area) {
		toSerialize["area"] = o.Area
	}
	if !IsNil(o.AreaType) {
		toSerialize["area_type"] = o.AreaType
	}
	if !IsNil(o.AuthenticationKey) {
		toSerialize["authentication_key"] = o.AuthenticationKey
	}
	if !IsNil(o.AuthenticationKeyId) {
		toSerialize["authentication_key_id"] = o.AuthenticationKeyId
	}
	if !IsNil(o.AuthenticationType) {
		toSerialize["authentication_type"] = o.AuthenticationType
	}
	if !IsNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}
	if !IsNil(o.DeadInterval) {
		toSerialize["dead_interval"] = o.DeadInterval
	}
	if !IsNil(o.HelloInterval) {
		toSerialize["hello_interval"] = o.HelloInterval
	}
	if !IsNil(o.Interface) {
		toSerialize["interface"] = o.Interface
	}
	if !IsNil(o.Preamble) {
		toSerialize["preamble"] = o.Preamble
	}
	if !IsNil(o.RetransmitInterval) {
		toSerialize["retransmit_interval"] = o.RetransmitInterval
	}
	if !IsNil(o.TransmitDelay) {
		toSerialize["transmit_delay"] = o.TransmitDelay
	}
	return toSerialize, nil
}

type NullableProtoOspfConfig struct {
	value *ProtoOspfConfig
	isSet bool
}

func (v NullableProtoOspfConfig) Get() *ProtoOspfConfig {
	return v.value
}

func (v *NullableProtoOspfConfig) Set(val *ProtoOspfConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableProtoOspfConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableProtoOspfConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtoOspfConfig(val *ProtoOspfConfig) *NullableProtoOspfConfig {
	return &NullableProtoOspfConfig{value: val, isSet: true}
}

func (v NullableProtoOspfConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtoOspfConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
IPAM Federation API

The DDI IPAM Federation application enables a SaaS administrator to manage multiple IPAM systems from one central control point CSP.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipamfederation

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the ReservedBlock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReservedBlock{}

// ReservedBlock A __ReservedBlock__ object (_federation/reserved_block_) is a set of contiguous IP addresses with no gap, expressed as a CIDR block. It is explicitly associated with a Federated Realm. A __ReservedBlock__ indicates an address range for which authority is expressly forbidden. Cooperating IPAM services must not make allocations in this range.
type ReservedBlock struct {
	// The address field in form “a.b.c.d/n” where the “/n” may be omitted. In this case, the CIDR value must be defined in the _cidr_ field. When reading, the _address_ field is always in the form “a.b.c.d”.
	Address string `json:"address"`
	// The CIDR of the reserved block. This is required field, if _address_ does not specify it in its input.
	Cidr *int64 `json:"cidr,omitempty"`
	// The description for the reserved block. May contain 0 to 1024 characters. Can include UTF-8.
	Comment *string `json:"comment,omitempty"`
	// Time when the object has been created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The resource identifier.
	FederatedRealm string `json:"federated_realm"`
	// The resource identifier.
	Id *string `json:"id,omitempty"`
	// The name of the reserved block. May contain 1 to 256 characters. Can include UTF-8.
	Name *string `json:"name,omitempty"`
	// The resource identifier.
	Parent *string `json:"parent,omitempty"`
	// The type of protocol of reserved block (_ip4_ or _ip6_).
	Protocol *string `json:"protocol,omitempty"`
	// The tags for the reserved block in JSON format.
	Tags map[string]interface{} `json:"tags,omitempty"`
	// Time when the object has been updated. Equals to _created_at_ if not updated after creation.
	UpdatedAt            *time.Time `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReservedBlock ReservedBlock

// NewReservedBlock instantiates a new ReservedBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReservedBlock(address string, federatedRealm string) *ReservedBlock {
	this := ReservedBlock{}
	this.Address = address
	this.FederatedRealm = federatedRealm
	return &this
}

// NewReservedBlockWithDefaults instantiates a new ReservedBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReservedBlockWithDefaults() *ReservedBlock {
	this := ReservedBlock{}
	return &this
}

// GetAddress returns the Address field value
func (o *ReservedBlock) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *ReservedBlock) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *ReservedBlock) SetAddress(v string) {
	o.Address = v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *ReservedBlock) GetCidr() int64 {
	if o == nil || IsNil(o.Cidr) {
		var ret int64
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservedBlock) GetCidrOk() (*int64, bool) {
	if o == nil || IsNil(o.Cidr) {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *ReservedBlock) HasCidr() bool {
	if o != nil && !IsNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given int64 and assigns it to the Cidr field.
func (o *ReservedBlock) SetCidr(v int64) {
	o.Cidr = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ReservedBlock) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservedBlock) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ReservedBlock) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ReservedBlock) SetComment(v string) {
	o.Comment = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ReservedBlock) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservedBlock) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ReservedBlock) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ReservedBlock) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetFederatedRealm returns the FederatedRealm field value
func (o *ReservedBlock) GetFederatedRealm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FederatedRealm
}

// GetFederatedRealmOk returns a tuple with the FederatedRealm field value
// and a boolean to check if the value has been set.
func (o *ReservedBlock) GetFederatedRealmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FederatedRealm, true
}

// SetFederatedRealm sets field value
func (o *ReservedBlock) SetFederatedRealm(v string) {
	o.FederatedRealm = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReservedBlock) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservedBlock) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReservedBlock) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ReservedBlock) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReservedBlock) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservedBlock) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ReservedBlock) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ReservedBlock) SetName(v string) {
	o.Name = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *ReservedBlock) GetParent() string {
	if o == nil || IsNil(o.Parent) {
		var ret string
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservedBlock) GetParentOk() (*string, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *ReservedBlock) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given string and assigns it to the Parent field.
func (o *ReservedBlock) SetParent(v string) {
	o.Parent = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *ReservedBlock) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservedBlock) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *ReservedBlock) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *ReservedBlock) SetProtocol(v string) {
	o.Protocol = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ReservedBlock) GetTags() map[string]interface{} {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservedBlock) GetTagsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Tags) {
		return map[string]interface{}{}, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ReservedBlock) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *ReservedBlock) SetTags(v map[string]interface{}) {
	o.Tags = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ReservedBlock) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReservedBlock) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ReservedBlock) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ReservedBlock) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o ReservedBlock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReservedBlock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	if !IsNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	toSerialize["federated_realm"] = o.FederatedRealm
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReservedBlock) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
		"federated_realm",
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

	varReservedBlock := _ReservedBlock{}

	err = json.Unmarshal(data, &varReservedBlock)

	if err != nil {
		return err
	}

	*o = ReservedBlock(varReservedBlock)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "cidr")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "federated_realm")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "protocol")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReservedBlock struct {
	value *ReservedBlock
	isSet bool
}

func (v NullableReservedBlock) Get() *ReservedBlock {
	return v.value
}

func (v *NullableReservedBlock) Set(val *ReservedBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableReservedBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableReservedBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReservedBlock(val *ReservedBlock) *NullableReservedBlock {
	return &NullableReservedBlock{value: val, isSet: true}
}

func (v NullableReservedBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReservedBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
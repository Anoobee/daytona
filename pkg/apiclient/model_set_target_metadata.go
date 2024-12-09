/*
Daytona Server API

Daytona Server API

API version: v0.0.0-dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SetTargetMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetTargetMetadata{}

// SetTargetMetadata struct for SetTargetMetadata
type SetTargetMetadata struct {
	ProviderMetadata *string `json:"providerMetadata,omitempty"`
	Uptime           int32   `json:"uptime"`
}

type _SetTargetMetadata SetTargetMetadata

// NewSetTargetMetadata instantiates a new SetTargetMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetTargetMetadata(uptime int32) *SetTargetMetadata {
	this := SetTargetMetadata{}
	this.Uptime = uptime
	return &this
}

// NewSetTargetMetadataWithDefaults instantiates a new SetTargetMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetTargetMetadataWithDefaults() *SetTargetMetadata {
	this := SetTargetMetadata{}
	return &this
}

// GetProviderMetadata returns the ProviderMetadata field value if set, zero value otherwise.
func (o *SetTargetMetadata) GetProviderMetadata() string {
	if o == nil || IsNil(o.ProviderMetadata) {
		var ret string
		return ret
	}
	return *o.ProviderMetadata
}

// GetProviderMetadataOk returns a tuple with the ProviderMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetTargetMetadata) GetProviderMetadataOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderMetadata) {
		return nil, false
	}
	return o.ProviderMetadata, true
}

// HasProviderMetadata returns a boolean if a field has been set.
func (o *SetTargetMetadata) HasProviderMetadata() bool {
	if o != nil && !IsNil(o.ProviderMetadata) {
		return true
	}

	return false
}

// SetProviderMetadata gets a reference to the given string and assigns it to the ProviderMetadata field.
func (o *SetTargetMetadata) SetProviderMetadata(v string) {
	o.ProviderMetadata = &v
}

// GetUptime returns the Uptime field value
func (o *SetTargetMetadata) GetUptime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Uptime
}

// GetUptimeOk returns a tuple with the Uptime field value
// and a boolean to check if the value has been set.
func (o *SetTargetMetadata) GetUptimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uptime, true
}

// SetUptime sets field value
func (o *SetTargetMetadata) SetUptime(v int32) {
	o.Uptime = v
}

func (o SetTargetMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetTargetMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProviderMetadata) {
		toSerialize["providerMetadata"] = o.ProviderMetadata
	}
	toSerialize["uptime"] = o.Uptime
	return toSerialize, nil
}

func (o *SetTargetMetadata) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"uptime",
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

	varSetTargetMetadata := _SetTargetMetadata{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSetTargetMetadata)

	if err != nil {
		return err
	}

	*o = SetTargetMetadata(varSetTargetMetadata)

	return err
}

type NullableSetTargetMetadata struct {
	value *SetTargetMetadata
	isSet bool
}

func (v NullableSetTargetMetadata) Get() *SetTargetMetadata {
	return v.value
}

func (v *NullableSetTargetMetadata) Set(val *SetTargetMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableSetTargetMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableSetTargetMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetTargetMetadata(val *SetTargetMetadata) *NullableSetTargetMetadata {
	return &NullableSetTargetMetadata{value: val, isSet: true}
}

func (v NullableSetTargetMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetTargetMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

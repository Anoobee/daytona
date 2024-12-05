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

// checks if the TargetConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetConfig{}

// TargetConfig struct for TargetConfig
type TargetConfig struct {
	Deleted bool   `json:"deleted"`
	Id      string `json:"id"`
	Name    string `json:"name"`
	// JSON encoded map of options
	Options      string             `json:"options"`
	ProviderInfo TargetProviderInfo `json:"providerInfo"`
}

type _TargetConfig TargetConfig

// NewTargetConfig instantiates a new TargetConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetConfig(deleted bool, id string, name string, options string, providerInfo TargetProviderInfo) *TargetConfig {
	this := TargetConfig{}
	this.Deleted = deleted
	this.Id = id
	this.Name = name
	this.Options = options
	this.ProviderInfo = providerInfo
	return &this
}

// NewTargetConfigWithDefaults instantiates a new TargetConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetConfigWithDefaults() *TargetConfig {
	this := TargetConfig{}
	return &this
}

// GetDeleted returns the Deleted field value
func (o *TargetConfig) GetDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value
// and a boolean to check if the value has been set.
func (o *TargetConfig) GetDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deleted, true
}

// SetDeleted sets field value
func (o *TargetConfig) SetDeleted(v bool) {
	o.Deleted = v
}

// GetId returns the Id field value
func (o *TargetConfig) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TargetConfig) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TargetConfig) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TargetConfig) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TargetConfig) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TargetConfig) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value
func (o *TargetConfig) GetOptions() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *TargetConfig) GetOptionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *TargetConfig) SetOptions(v string) {
	o.Options = v
}

// GetProviderInfo returns the ProviderInfo field value
func (o *TargetConfig) GetProviderInfo() TargetProviderInfo {
	if o == nil {
		var ret TargetProviderInfo
		return ret
	}

	return o.ProviderInfo
}

// GetProviderInfoOk returns a tuple with the ProviderInfo field value
// and a boolean to check if the value has been set.
func (o *TargetConfig) GetProviderInfoOk() (*TargetProviderInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderInfo, true
}

// SetProviderInfo sets field value
func (o *TargetConfig) SetProviderInfo(v TargetProviderInfo) {
	o.ProviderInfo = v
}

func (o TargetConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TargetConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deleted"] = o.Deleted
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["options"] = o.Options
	toSerialize["providerInfo"] = o.ProviderInfo
	return toSerialize, nil
}

func (o *TargetConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"deleted",
		"id",
		"name",
		"options",
		"providerInfo",
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

	varTargetConfig := _TargetConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTargetConfig)

	if err != nil {
		return err
	}

	*o = TargetConfig(varTargetConfig)

	return err
}

type NullableTargetConfig struct {
	value *TargetConfig
	isSet bool
}

func (v NullableTargetConfig) Get() *TargetConfig {
	return v.value
}

func (v *NullableTargetConfig) Set(val *TargetConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetConfig(val *TargetConfig) *NullableTargetConfig {
	return &NullableTargetConfig{value: val, isSet: true}
}

func (v NullableTargetConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}